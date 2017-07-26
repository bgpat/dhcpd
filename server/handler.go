package server

import (
	"net"
	"time"

	dhcp "github.com/krolaw/dhcp4"
)

type Handler struct {
	ServerIPAddr net.IP
	Options      dhcp.Options
	Leases       Leases
}

func (h *Handler) ServeDHCP(p dhcp.Packet, msgType dhcp.MessageType, options dhcp.Options) dhcp.Packet {
	switch msgType {
	case dhcp.Discover:
		lease := h.Leases.Get(p.CHAddr())
		if lease == nil {
			return nil
		}
		return dhcp.ReplyPacket(
			p,
			dhcp.Offer,
			h.ServerIPAddr,
			lease.IPAddr,
			h.Leases.Duration,
			h.Options.SelectOrderOrAll(options[dhcp.OptionParameterRequestList]),
		)
	case dhcp.Request:
		if addr, ok := options[dhcp.OptionServerIdentifier]; ok && !net.IP(addr).Equal(h.ServerIPAddr) {
			return nil
		}
		req := net.IP(options[dhcp.OptionRequestedIPAddress])
		if req == nil {
			req = net.IP(p.CIAddr())
		}
		if len(req) != 4 || req.Equal(net.IPv4zero) {
			return dhcp.ReplyPacket(p, dhcp.NAK, h.ServerIPAddr, nil, 0, nil)
		}
		i := dhcp.IPRange(h.Leases.StartIPAddr, req) - 1
		if i < 0 || h.Leases.Range <= i {
			return dhcp.ReplyPacket(p, dhcp.NAK, h.ServerIPAddr, nil, 0, nil)
		}
		lease := h.Leases.Table[i]
		if lease == nil {
			return dhcp.ReplyPacket(p, dhcp.NAK, h.ServerIPAddr, nil, 0, nil)
		}
		lease.Expiry = time.Now().Add(h.Leases.Duration)
		return dhcp.ReplyPacket(
			p,
			dhcp.ACK,
			h.ServerIPAddr,
			req,
			h.Leases.Duration,
			h.Options.SelectOrderOrAll(options[dhcp.OptionParameterRequestList]),
		)
	case dhcp.Release, dhcp.Decline:
		h.Leases.Delete(p.CHAddr())
	}
	return nil
}
