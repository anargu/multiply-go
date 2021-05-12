package restserver

import (
	"fmt"
	"log"
	"net/http"
)

func RunServer(addr string) {
	fmt.Printf("Server running on %s ...\n", addr)
	http.HandleFunc("/v1/multiply", MultiplyRestHandlerV1)
	log.Fatal(http.ListenAndServe(addr, nil))
}
