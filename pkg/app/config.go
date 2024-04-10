package app

import (
	"os"

	"github.com/ameshkov/dnsstamps"
	"gopkg.in/yaml.v2"
)

const dnsServersFile = "./dns_speed_mate.yaml"

type dnsServerConfig struct {
	IPAndPort string                   `yaml:"ip_and_port" json:"ip_and_port"`
	Hostname  string                   `yaml:"hostname" json:"hostname"`
	Protocol  dnsstamps.StampProtoType `yaml:"protocol" json:"protocol"`
}

// config is app's overall config
type config struct {
	DNSServers []dnsServerConfig `yaml:"dns_servers"`
}

// default config for failed to load config or no config exists
var cfgDefault = &config{
	DNSServers: []dnsServerConfig{
		{
			IPAndPort: "1.1.1.1:53",
			Hostname:  "cloudflare-dns.com",
			Protocol:  dnsstamps.StampProtoTypePlain,
		},
	},
}

// loadConfig reads the DNSSpeedMate config file and saves it to the app struct; if
// there is an issue loading the config file, a default config is loaded instead
func (app *App) loadConfig() {
	// func to load config file
	cfgLoader := func() error {
		// load config file
		configFile, err := os.ReadFile(dnsServersFile)
		if err != nil {
			return err
		}

		cfg := &config{}
		err = yaml.Unmarshal(configFile, cfg)
		if err != nil {
			return err
		}

		// set cfg
		app.cfg = cfg

		return nil
	}

	// run config loader and if err use default
	err := cfgLoader()
	if err != nil {
		app.cfg = cfgDefault
	}
}

// TODO: save config
