package app

const speedTestCount = 30

func (serv *dnsServer) speedTestUncached() {
	// 	// skip test if doesn't respond to A or AAAA
	// 	if !serv.RespondsA && !serv.RespondsAAAA {
	// 		return
	// 	}

	// 	// uncached test (A will be used unless server only supported AAAA)
	// 	req := &dns.Msg{}
	// 	if serv.RespondsA {
	// 		req.SetQuestion(randomness.GenerateRandomComDomain(), dns.TypeA)
	// 	} else {
	// 		req.SetQuestion(randomness.GenerateRandomComDomain(), dns.TypeAAAA)
	// 	}

	// 	// startTime := time.Now()
	// 	// _, err := serv.upstream.Exchange(req)
	// 	// queryDuration := time.Since(startTime).Round(time.Millisecond / 100)

	// 	// failed
	// 	// if err != nil {
	// 	// 	serv.UncachedFails++
	// 	// 	return
	// 	// }

	// 	// prevSuccesses := serv.UncachedAttempts - serv.UncachedFails
	// 	// serv.UncachedResponse = (serv.UncachedResponse*time.Duration(prevSuccesses) + queryDuration) / time.Duration(prevSuccesses+1)
	// 	// serv.UncachedAttempts++
	// }

	// // speedTest queries all dns servers in round robin fashion to calculate their response
	// // times and also records if there are any failed queries
	// func (app *App) speedTest() {
	// 	// none of this is async to avoid possible distortion of results due to overlapping
	// 	// queries and/or network congestion
	// 	for i := 1; i <= speedTestCount; i++ {
	// 		for j := 0; j < len(app.dnsServers); j++ {
	// 			// uncached
	// 			app.dnsServers[j].speedTestUncached()

	// 		}

	// 		// cached

	// 		// root servers

	// }
}
