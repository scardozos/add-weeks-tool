package main

import (
	"log"
	"os"

	"github.com/scardozos/add-weeks-tool/cmd/httpserver"
)

var (
	grpcServerAddr = os.Getenv("GRPC_SERVER_ADDR")
	httpServerAddr = os.Getenv("HTTP_SERVER_ADDR")
)

func main() {
	server := httpserver.NewHttpServer(grpcServerAddr, httpServerAddr)
	server.Serve()
	log.Printf("Started serving")
}
