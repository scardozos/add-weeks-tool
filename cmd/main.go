package main

import (
	"log"
	"os"
	"time"

	"github.com/scardozos/add-weeks-tool/cmd/dbclient"
)

var (
	grpcServerAddr = os.Getenv("GRPC_SERVER_ADDR")
)

func main() {
	localClient := dbclient.NewLocalClient(grpcServerAddr, false)
	now := time.Now()
	log.Printf("Got value %v in %v", localClient.GetStaticWeeks(), time.Since(now))
}
