package main

import (
	"github.com/theinpu/canon_parser/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()
	server.RegisterParserIOServer(srv, &server.ParserServer{})
	err = srv.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}
