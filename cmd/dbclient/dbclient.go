package dbclient

import (
	weeks_pb "github.com/scardozos/esplai-weeks-db/api/weeksdb"
	"google.golang.org/grpc"
)

type DbClientContext struct {
	Client weeks_pb.WeeksDatabaseClient
	Conn   *grpc.ClientConn
}

type DbClient struct {
	*DbClientContext
}

func NewDbClient(host string, secure bool) *DbClient {
	clientContext := DbClientContext{}
	clientContext.CreateConn(host, secure)
	clientContext.Client = weeks_pb.NewWeeksDatabaseClient(clientContext.Conn)
	return &DbClient{&clientContext}
}
