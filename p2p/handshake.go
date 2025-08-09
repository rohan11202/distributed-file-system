package p2p

type HandshakerFunc func(Peer) error

func NOPHandshakeFunc(Peer) error {
	return nil
}
