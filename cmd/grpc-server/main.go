package main

import (
	"flag"
	"fmt"
	"log"
	grpcsvc "multiply/grpc"
	"multiply/grpc/proto"
	"net"

	"google.golang.org/grpc"
)

func main() {
	port := flag.String("port", "12000", "port in which grpc server serves")
	host := flag.String("host", "", "host in which grpc server serves")
	flag.Parse()
	RunServer(*host, *port)
}

func RunServer(host string, port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalln(err)
	}
	srv := grpc.NewServer()
	proto.RegisterMultiplyServiceServer(srv, grpcsvc.MultiplyGrpcSrv)
	// reflection.Register(srv)

	log.Fatal(srv.Serve(listener))
}
