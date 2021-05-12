package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	grpcsrv "multiply/grpc"

	"google.golang.org/grpc"
)

func main() {
	port := flag.String("port", "12000", "port in which grpc server serves")
	host := flag.String("host", "", "host in which grpc server serves")
	x := flag.Float64("x", 0.0, "first number")
	y := flag.Float64("y", 0.0, "second number")
	flag.Parse()
	var err error

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", *host, *port), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := grpcsrv.NewClient(conn)
	result, err := grpcsrv.ExecuteOperation(float32(*x), float32(*y), client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stdout, "%v\n", *result)
}
