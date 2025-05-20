#!/bin/bash

echo "Testing db package with coverage:"
cd internal/db
go test -covermode=count -coverprofile=coverage.out && go tool cover -func=coverage.out

echo -e "\nTesting wifi package with coverage:"
cd internal/wifi
go test -covermode=count -coverprofile=coverage.out && go tool cover -func=coverage.out