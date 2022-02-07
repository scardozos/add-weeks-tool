#!/bin/bash
export GRPC_SERVER_ADDR="localhost:9001"
export HTTP_PORT=""
go run -race cmd/main.go