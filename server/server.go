package server

import (
	dhcp "github.com/krolaw/dhcp4"
)

type Server struct {
	Handler *Handler
}

func New() (*Server, error) {
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}
	return config.Server(), nil
}

func (s *Server) Listen() error {
	return dhcp.ListenAndServe(s.Handler)
}
