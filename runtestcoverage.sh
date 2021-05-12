#!/bin/sh

go test ./... -coverprofile coverage
go tool cover -func=coverage
go tool cover -html=coverage -o coverage.html