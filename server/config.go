package server

import (
	"net"
	"time"

	"github.com/kelseyhightower/envconfig"
	dhcp "github.com/krolaw/dhcp4"
)

type Config struct {
	Server_IP_Addr string `required:"true"`
	Options        map[dhcp.OptionCode]string
	Start_IP_Addr  string        `required:"true"`
	Lease_Range    int           `required:"true"`
	Lease_Duration time.Duration `default:"1h"`
}

func NewConfig() (*Config, error) {
	c := Config{
		Options: make(map[dhcp.OptionCode]string),
	}
	if err := envconfig.Process("dhcp", &c); err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *Config) Server() *Server {
	options := make(dhcp.Options)
	for code, val := range c.Options {
		options[code] = []byte(val)
	}
	return &Server{
		Handler: &Handler{
			ServerIPAddr: net.ParseIP(c.Server_IP_Addr),
			Options:      options,
			Leases: Leases{
				StartIPAddr: net.ParseIP(c.Start_IP_Addr),
				Range:       c.Lease_Range,
				Duration:    c.Lease_Duration,
			},
		},
	}
}
