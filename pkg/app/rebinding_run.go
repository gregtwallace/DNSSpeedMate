package app

import (
	"dnsspeedmate/pkg/safemap"
	"errors"
)

const rebindDomainRoot = "rebindtest.dnsspeedmate.com."

// rebindNetID is a unique identifier for each network that will be
// tested for rebinding protection
type rebindNetID int

const (
	rebindNet10 rebindNetID = iota
	rebindNet10v6
	rebindNet127
	rebindNet127v6
	rebindNet169
	rebindNet169v6
	rebindNet172
	rebindNet172v6
	rebindNet192
	rebindNet192v6

	rebindNetv6fc00
	rebindNetv6fe80
	rebindNetv6Loopback
)

type rebindNet struct {
	ID                 rebindNetID
	IPv6               bool
	CIDR               string
	TestDomain         string
	TestExpectedResult string
}

var rebindNets = []rebindNet{
	{
		ID:                 rebindNet10,
		IPv6:               false,
		CIDR:               "10.0.0.0/8",
		TestDomain:         "net-10." + rebindDomainRoot,
		TestExpectedResult: "10.10.10.1",
	},
	{
		ID:                 rebindNet10v6,
		IPv6:               true,
		CIDR:               "::ffff:10.0.0.0/8 (::ffff:a00:0/104)",
		TestDomain:         "net-10." + rebindDomainRoot,
		TestExpectedResult: "::ffff:10.10.10.1",
	},
	{
		ID:                 rebindNet127,
		IPv6:               false,
		CIDR:               "127.0.0.0/8",
		TestDomain:         "net-127." + rebindDomainRoot,
		TestExpectedResult: "127.0.0.1",
	},
	{
		ID:                 rebindNet127v6,
		IPv6:               true,
		CIDR:               "::ffff:127.0.0.0/8 (::ffff:7f00:0/104)",
		TestDomain:         "net-127." + rebindDomainRoot,
		TestExpectedResult: "::ffff:127.0.0.1",
	},
	{
		ID:                 rebindNet169,
		CIDR:               "169.254.0.0/16",
		IPv6:               false,
		TestDomain:         "net-169." + rebindDomainRoot,
		TestExpectedResult: "169.254.1.1",
	},
	{
		ID:                 rebindNet169v6,
		CIDR:               "::ffff:169.254.0.0/16 (::ffff:c0a8:0/112)",
		IPv6:               true,
		TestDomain:         "net-169." + rebindDomainRoot,
		TestExpectedResult: "::ffff:169.254.1.1",
	},
	{
		ID:                 rebindNet172,
		CIDR:               "172.16.0.0/12",
		IPv6:               false,
		TestDomain:         "net-172." + rebindDomainRoot,
		TestExpectedResult: "172.16.0.1",
	},
	{
		ID:                 rebindNet172v6,
		CIDR:               "::ffff:172.16.0.0/12 (::ffff:ac10:0/108)",
		IPv6:               true,
		TestDomain:         "net-172." + rebindDomainRoot,
		TestExpectedResult: "::ffff:172.16.0.1",
	},
	{
		ID:                 rebindNet192,
		CIDR:               "192.168.0.0/16",
		IPv6:               false,
		TestDomain:         "net-192." + rebindDomainRoot,
		TestExpectedResult: "192.168.1.1",
	},
	{
		ID:                 rebindNet192v6,
		CIDR:               "::ffff:192.168.0.0/16 (::ffff:a9fe:0/112)",
		IPv6:               true,
		TestDomain:         "net-192." + rebindDomainRoot,
		TestExpectedResult: "::ffff:192.168.1.1",
	},
	{
		ID:                 rebindNetv6fc00,
		CIDR:               "fc00::/8",
		IPv6:               true,
		TestDomain:         "net-v6-fc00." + rebindDomainRoot,
		TestExpectedResult: "fc00::1",
	},
	{
		ID:                 rebindNetv6fe80,
		CIDR:               "fe80::/10",
		IPv6:               true,
		TestDomain:         "net-v6-fe80." + rebindDomainRoot,
		TestExpectedResult: "fe80::1",
	},
	{
		ID:                 rebindNetv6Loopback,
		CIDR:               "::1/128",
		IPv6:               true,
		TestDomain:         "net-v6-loopback." + rebindDomainRoot,
		TestExpectedResult: "::1",
	},
}

// MarshalText for rebindNetID to cover map keys also
func (rnID rebindNetID) MarshalText() ([]byte, error) {
	for i := range rebindNets {
		if rebindNets[i].ID == rnID {
			return []byte(rebindNets[i].CIDR), nil
		}
	}

	return nil, errors.New("invalid rebindNetID")
}

type rebindStatus int

const (
	rebindStatusProtected rebindStatus = iota
	rebindStatusVulnerable
	rebindStatusTestFailed
	rebindStatusNotTested
)

func (rs rebindStatus) String() string {
	friendlyName := ""

	switch rs {
	case rebindStatusProtected:
		friendlyName = "Protected"
	case rebindStatusVulnerable:
		friendlyName = "Vulnerable"
	case rebindStatusTestFailed:
		friendlyName = "Failed to Run"
	case rebindStatusNotTested:
		friendlyName = "Not Tested"
	}

	return friendlyName
}

// MarshalText for rebindStatus to yield friendly text
func (rs rebindStatus) MarshalText() ([]byte, error) {
	return []byte(rs.String()), nil
}

// RunRebindingTest tests for dns rebinding protection on the server
func (serv *dnsServer) RunRebindingTest(pretestResults *safemap.SafeMap[rebindNetID, bool]) {
	// make new results map
	// serv.RebindResults = safemap.NewSafeMap[rebindNetID, rebindStatus]()

	// // check each rebindNet
	// var wg sync.WaitGroup

	// for i := range rebindNets {
	// 	wg.Add(1)

	// 	go func() {
	// 		defer wg.Done()

	// 		// if pretest is missing don't run test
	// 		if pretestResults == nil {
	// 			_, _ = serv.RebindResults.Add(rebindNets[i].ID, rebindStatusNotTested)
	// 			return
	// 		}

	// 		// if pretest failed or isn't in list for some reason, don't run test
	// 		preResult, err := pretestResults.Read(rebindNets[i].ID)
	// 		if !preResult || err != nil {
	// 			_, _ = serv.RebindResults.Add(rebindNets[i].ID, rebindStatusNotTested)
	// 			return
	// 		}

	// 		// if test is a record type not supported by server, don't run it
	// 		if (rebindNets[i].IPv6 && !serv.RespondsAAAA) ||
	// 			(!rebindNets[i].IPv6 && !serv.RespondsA) {

	// 			_, _ = serv.RebindResults.Add(rebindNets[i].ID, rebindStatusNotTested)
	// 			return
	// 		}

	// 		// run test
	// 		recordType := dns.TypeA
	// 		if rebindNets[i].IPv6 {
	// 			recordType = dns.TypeAAAA
	// 		}

	// 		req := (&dns.Msg{}).SetQuestion(rebindNets[i].TestDomain, recordType)

	// 		resp, err := serv.upstream.Exchange(req)
	// 		// failed to run test
	// 		if err != nil {
	// 			_, _ = serv.RebindResults.Add(rebindNets[i].ID, rebindStatusTestFailed)
	// 			return
	// 		}

	// 		// check result (Refused, NXDOMAIN, and NULL IP are considered acceptable rebind protection)
	// 		answerIsNullIP := false
	// 		if len(resp.Answer) == 1 {
	// 			// ipv4 A
	// 			if !rebindNets[i].IPv6 {
	// 				aAns, ok := resp.Answer[0].(*dns.A)
	// 				if ok && aAns.A.Equal(net.IPv4zero) {
	// 					answerIsNullIP = true
	// 				}
	// 			} else {
	// 				// ipv6 AAAA
	// 				aaaaAns, ok := resp.Answer[0].(*dns.AAAA)
	// 				if ok && aaaaAns.AAAA.Equal(net.IPv6zero) {
	// 					answerIsNullIP = true
	// 				}
	// 			}
	// 		}

	// 		if resp.MsgHdr.Rcode == dns.RcodeRefused ||
	// 			resp.MsgHdr.Rcode == dns.RcodeNameError ||
	// 			answerIsNullIP {

	// 			_, _ = serv.RebindResults.Add(rebindNets[i].ID, rebindStatusProtected)
	// 			return
	// 		}

	// 		// didn't match a protected condition = vulnerable
	// 		_, _ = serv.RebindResults.Add(rebindNets[i].ID, rebindStatusVulnerable)
	// 	}()
	// }

	// wg.Wait()
}
