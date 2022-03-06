#!/bin/bash
export GRPC_SERVER_ADDR="localhost:9001"
export HTTP_SERVER_ADDR="localhost:8002"
go run -race cmd/main.go