#!/bin/bash

echo "Compiling the apps..."
go build -o main ./cmd/main.go

echo "Run the compiled binary.."
./main
