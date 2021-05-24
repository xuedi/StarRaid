package server

import (
	"fmt"
	"os"

	ini "gopkg.in/ini.v1"
)

type Config struct {
	dsn       string
	targetTps int
	debug     bool
	host      string
	port      int
}

func (c *Config) load(configFile string) {

	// load ini file
	cfg, err := ini.Load(configFile)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	c.debug, err = cfg.Section("server").Key("debug").Bool()
	if err != nil {
		fmt.Printf("config missing 'debug': %v", err)
		os.Exit(1)
	}

	c.targetTps, err = cfg.Section("server").Key("targetTps").Int()
	if err != nil || c.targetTps <= 0 {
		fmt.Printf("config has an invalid 'targetTps' setting: %v", c.targetTps)
		os.Exit(1)
	}

	c.port, err = cfg.Section("server").Key("port").Int()
	if err != nil || c.port <= 0 {
		fmt.Printf("config has an invalid 'port' setting: %v", c.port)
		os.Exit(1)
	}

	c.host = cfg.Section("server").Key("host").String()

	c.dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		cfg.Section("database").Key("user").String(),
		cfg.Section("database").Key("pass").String(),
		cfg.Section("database").Key("host").String(),
		cfg.Section("database").Key("port").String(),
		cfg.Section("database").Key("name").String(),
	)

}
