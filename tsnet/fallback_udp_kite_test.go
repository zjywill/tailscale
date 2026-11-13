package tsnet

import (
	"net/netip"
	"tailscale.com/types/nettype"
	"testing"
)

func TestUDPFallbackDeregister(t *testing.T) {
	s := new(Server)
	remove := s.RegisterFallbackUDPHandler(func(src, dst netip.AddrPort) (func(nettype.ConnPacketConn), bool) {
		return func(nettype.ConnPacketConn) {}, true
	})
	if len(s.fallbackUDPHandlers) != 1 {
		t.Fatal("handler not registered")
	}
	remove()
	remove()
	if len(s.fallbackUDPHandlers) != 0 {
		t.Fatal("handler not removed")
	}
}
