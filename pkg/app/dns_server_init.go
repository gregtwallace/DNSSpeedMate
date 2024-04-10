package app

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/AdguardTeam/dnsproxy/upstream"
	"github.com/ameshkov/dnsstamps"
	"github.com/miekg/dns"
	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

const domainInitialTest = "google.com."

// dnsServer is config and information about a dns server
type dnsServer struct {
	// config
	dnsServerConfig

	// usable server
	upstream upstream.Upstream `json:"-"`

	// information about the server
	RespondsA    *bool `json:"responds_a"`
	RespondsAAAA *bool `json:"responds_aaaa"`
}

// initUpstream establishes the initial dns connection and does some basic connectivity
// testing for A and AAAA records
func (serv *dnsServer) initUpstream() {
	// make dns stamp
	stamp := &dnsstamps.ServerStamp{
		ServerAddrStr: serv.IPAndPort,
		ProviderName:  serv.Hostname,
		Proto:         dnsstamps.StampProtoType(serv.Protocol),
		Path:          "/dns-query",
	}

	// initial config
	var err error
	serv.upstream, err = upstream.AddressToUpstream(stamp.String(), nil)
	if err != nil {
		// bad config
		return
	}

	// make sure things are working
	var wg sync.WaitGroup

	// A
	wg.Add(1)
	go func() {
		defer wg.Done()

		req := (&dns.Msg{}).SetQuestion(domainInitialTest, dns.TypeA)
		_, err = serv.upstream.Exchange(req)
		if err != nil {
			b := false
			serv.RespondsA = &b
		} else {
			b := true
			serv.RespondsA = &b
		}
	}()

	// AAAA
	wg.Add(1)
	go func() {
		defer wg.Done()

		req := (&dns.Msg{}).SetQuestion(domainInitialTest, dns.TypeAAAA)
		_, err = serv.upstream.Exchange(req)
		if err != nil {
			b := false
			serv.RespondsAAAA = &b
		} else {
			b := true
			serv.RespondsAAAA = &b
		}
	}()

	wg.Wait()

}

// initDNSServersEvent is the struct for data to send in the
// initDNSServers Event
type initDNSServersEvent struct {
	DNSServers *[]dnsServer `json:"dns_servers"`
}

// emit emits the initDNSServersEvent
func (idse *initDNSServersEvent) emit(appCtx context.Context) {
	event, err := json.Marshal(idse)
	if err != nil {
		// should never happen
		panic(err)
	}

	log.Println(string(event))
	rt.EventsEmit(appCtx, "init_dns_servers", string(event))
}

// initDNSServers establishes the upstream connection to each DNS server
// and performs basic testing
func (app *App) initDNSServers() {
	// make servers array on app
	app.dnsServers = make([]dnsServer, 0)

	// make response struct
	resStr := &initDNSServersEvent{
		DNSServers: &app.dnsServers,
	}

	for i := range app.cfg.DNSServers {
		app.dnsServers = append(app.dnsServers, dnsServer{dnsServerConfig: app.cfg.DNSServers[i]})
	}
	// initial emit w/ server list
	resStr.emit(app.ctx)

	// async initializing each
	var wg sync.WaitGroup

	for i := range app.dnsServers {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// init upstream
			app.dnsServers[i].initUpstream()

			// emit current status
			resStr.emit(app.ctx)
		}()
	}

	wg.Wait()
}
