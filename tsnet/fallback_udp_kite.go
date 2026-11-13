// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause
// Kite G46: UDP counterpart of upstream RegisterFallbackTCPHandler.
package tsnet

import (
	"net/netip"
	"tailscale.com/types/nettype"
)

// FallbackUDPHandler handles a connected UDP flow only when no explicit
// tsnet listener matches. It runs after normal tailnet packet filtering.
// A nil handler with intercept=true rejects the flow. Callbacks must not call
// back into Server: as with TCP fallback registration, Server.mu is held.
type FallbackUDPHandler func(src, dst netip.AddrPort) (handler func(nettype.ConnPacketConn), intercept bool)

func (s *Server) RegisterFallbackUDPHandler(cb FallbackUDPHandler) func() {
	s.mu.Lock()
	defer s.mu.Unlock()
	hnd := s.fallbackUDPHandlers.Add(cb)
	return func() { s.mu.Lock(); defer s.mu.Unlock(); delete(s.fallbackUDPHandlers, hnd) }
}
