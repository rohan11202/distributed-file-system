package main

import (
	"fmt"
	"log"

	"distFile/p2p"
)

func OnPeer(p p2p.Peer) error {
	fmt.Printf("Doing something with Peer outside of TCPTransport")
	p.Close()
	return nil
}

func main() {

	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:     ":3000",
		HandshakerFunc: p2p.NOPHandshakeFunc,
		Decoder:        p2p.DefaultDecoder{},
		OnPeer:         OnPeer,
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("Received message: %+v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
