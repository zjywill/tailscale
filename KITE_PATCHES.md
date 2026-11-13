# Kite Tailscale dependency

Base: a2263542f260e73260e821a2d1b81475c35b0766 (the existing Kite pin).
This branch intentionally avoids a broad dependency upgrade.

- Add a UDP counterpart to tsnet's fallback TCP handler API. Explicit listeners
  retain precedence; normal packet filtering is unchanged. No handler means the
  existing default refusal. Callbacks run under Server.mu, as with TCP, and must
  not call Server methods. Host address policy and resource bounds belong to the
  consuming Kite-Go-Core meshservice.HostForwarder.
- Apply IPv6 fragment GRO checksum handling and regressions from
  23de235c3accb83592e6d03438d6c61b35817de5. This prevents first fragments from being
  treated as complete transport packets before reassembly. Atomic fragments
  still receive transport checksum validation. Original copyright and
  BSD-3-Clause notices are preserved.

Validation: go test -race -count=1 ./tsnet ./wgengine/netstack/gro
./wgengine/netstack, plus Kite-Go-Core mesh/ingresstest real control/DERP tests
covering IPv4/IPv6 60,000-byte UDP, general TCP/UDP host ingress, listener priority,
ACL denial and three repeated lifecycles. See Core artifacts/G46 for exact logs.
