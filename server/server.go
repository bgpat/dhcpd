package server

import (
	dhcp "github.com/krolaw/dhcp4"
)

type Server struct {
	Handler   *Handler
	Interface *string
}

func New() (*Server, error) {
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}
	return config.Server(), nil
}

func (s *Server) Listen() error {
	if s.Interface == nil {
		return dhcp.ListenAndServe(s.Handler)
	}
	return dhcp.ListenAndServeIf(*s.Interface, s.Handler)
}
