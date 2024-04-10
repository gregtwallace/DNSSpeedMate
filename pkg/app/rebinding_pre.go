package app

import (
	"dnsspeedmate/pkg/safemap"
	"log"
	"strings"
	"sync"

	"github.com/AdguardTeam/dnsproxy/upstream"
	"github.com/ameshkov/dnsstamps"
	"github.com/miekg/dns"
)

// rebindPretestServer has the config for the dns server used for the
// rebinding pre test
var rebindPretestServerStamp = &dnsstamps.ServerStamp{
	ServerAddrStr: "1.1.1.1:853",
	ProviderName:  "cloudflare-dns.com",
	Proto:         dnsstamps.StampProtoTypeTLS,
}

// PretestInternet
func (app *App) PretestInternet() {

}

// rebindPretest confirms the rebinding records actually exist via a
// resolver that won't block them; this is done to ensure there aren't
// false "Protected" rebinding results in the event the records are
// ever deleted
func (app *App) RebindPretest() {
	results := safemap.NewSafeMap[rebindNetID, bool]()

	// make dns server for pretest
	serv, err := upstream.AddressToUpstream(rebindPretestServerStamp.String(), nil)
	if err != nil {
		// bad config, pretest fails and rebinding cant be tested
		log.Println(err)
		serv = nil
	}

	// Use a wg to make this faster (esp. in cases of failure)
	var wg sync.WaitGroup

	// confirm test for each net is working
	for i := range rebindNets {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// if pretest server failed, test failed
			if serv == nil {
				results.Add(rebindNets[i].ID, false)
				return
			}

			// make request
			rType := dns.TypeA
			if rebindNets[i].IPv6 {
				rType = dns.TypeAAAA
			}

			req := (&dns.Msg{}).SetQuestion(rebindNets[i].TestDomain, rType)

			// do request
			resp, err := serv.Exchange(req)
			if err != nil {
				log.Println(err)
			}

			// true if correct result, false anything else
			thisResult := false
			if resp != nil &&
				len(resp.Answer) == 1 &&
				strings.HasSuffix(resp.Answer[0].String(), rebindNets[i].TestExpectedResult) {

				thisResult = true
			}

			results.Add(rebindNets[i].ID, thisResult)
		}()
	}

	wg.Wait()

	// app.rebindPretestResults = results
}
