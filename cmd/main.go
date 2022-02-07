package main

import (
	"log"
	"os"

	"github.com/scardozos/add-weeks-tool/cmd/dbclient"
)

var (
	grpcServerAddr = os.Getenv("GRPC_SERVER_ADDR")
)

func main() {
	localClient := dbclient.NewLocalClient(grpcServerAddr, false)
	log.Print(localClient.GetStaticWeeks())
}
