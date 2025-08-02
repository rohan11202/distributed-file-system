package p2p


type Peer interface {
}

//Handles the communication betwen the nodes
type Transport interface {
	ListenAndAccept() error

}