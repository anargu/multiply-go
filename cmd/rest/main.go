package main

import (
	"flag"
	restserver "multiply/rest"
)

func main() {
	addr := flag.String("addr", ":11000", "addr in which rest service will serve. Ex: \"0.0.0.0:11000\"")
	flag.Parse()

	restserver.RunServer(*addr)
}
