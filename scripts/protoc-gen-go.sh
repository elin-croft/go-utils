#!/bin/bash

rm -rf ./gen-go/proto/*
for d in `find protos -name "*.proto"`; do
    protoc --go_out=paths=source_relative:gen-go/protos \
    --go-grpc_out=paths=source_relative:gen-go/protos \
    --proto_path=./protos \
    ./${d};
done
