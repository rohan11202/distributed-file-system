package p2p

type Peer interface {
	Close() error
}

// Handles the communication betwen the nodes
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
