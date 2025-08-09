package p2p

import "net"

// RPC holds any arbuitory data that is being
// sent over the p2p network
type RPC struct {
	From    net.Addr
	Payload []byte
}
