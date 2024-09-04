package main

import (
	"dist-file-system-golang/p2p"
	"fmt"
)

type FileServerOpts struct {
	StorageRoot string
	PathTransformFunc PathTransformFunc
	Transport p2p.Transport
}

type FileServer struct {
	FileServerOpts

	store *Store
	quitch chan struct{}
}

func NewFileServer(opts FileServerOpts) *FileServer {
	storeOpts := StoreOpts {
		Root : opts.StorageRoot,
		PathTransformFunc : opts.PathTransformFunc,
	}
	return &FileServer{
		FileServerOpts: opts,
		store: NewStore(storeOpts),
		quitch: make(chan struct{}),
	}
}

func (s *FileServer) Stop() {
	
}

func (s *FileServer) loop() {
	for {
		select {
		case msg := <- s.Transport.Consume():
			fmt.Println(msg)
		case <- s.quitch:
			return
		}
	}
}

func (s *FileServer) Start() error {
	if err := s.Transport.ListenAndAccept(); err != nil {
		return err
	}

	return nil
}