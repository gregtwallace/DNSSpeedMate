package app

import (
	"context"
	"time"
)

// App struct
type App struct {
	ctx context.Context
	cfg *config

	dnsServers []dnsServer
}

// New creates a new App application struct
func New() *App {
	app := &App{}

	// load app config
	app.loadConfig()

	return app
}

// Startup is to be called when the app starts.
func (app *App) Startup(ctx context.Context) {
	// save context for runtime
	app.ctx = ctx

	// initialize dns servers
	time.Sleep(2 * time.Second)
	app.initDNSServers()

	// run pre-test for rebinding protection
	// app.RebindPretest()

	// // make app
	// app := App{}

	// // initialize dns servers (connect to each and do basic testing)
	// app.initDNSServers()

	// // speed testing
	// app.speedTest()

	// // log results
	// app.output()

}
