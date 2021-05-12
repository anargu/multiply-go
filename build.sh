#!/bin/sh

go build -o grpc-client cmd/grpc-client/main.go
go build -o grpc-server cmd/grpc-server/main.go
go build -o rest-server cmd/rest/main.go
