#!/bin/bash

# check arguments
if [ $# -gt 0 ]; then
  if [[ "$1" == "build" ]]; then
    # build args exist
    echo "Compiling the apps..."
    go build -o main ./cmd/main.go
  fi
fi

echo "Run the compiled binary.."
./main
