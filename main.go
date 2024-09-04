package main

import (
	"dist-file-system-golang/p2p"
	"log"
)

func main() {
	tcptransportOpts := p2p.TCPTransportOpts {
		ListenAddr: ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder: p2p.DefaultDecoder{},
		// TODO: onPeer func
	}
	tcpTransport := p2p.NewTCPTransport(tcptransportOpts)

	fileServerOpts := FileServerOpts{
		StorageRoot: "3000_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport: tcpTransport,
	}

	s := NewFileServer(fileServerOpts)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

	select {}
}