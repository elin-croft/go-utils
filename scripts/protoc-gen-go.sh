#!/bin/bash

rm -rf ./gen-go/proto/*
for d in `find protos -name "*.proto"`;do
    protoc  --go_out=. --go-grpc_out=. \
    --proto_path=./protos \
    ./${d};
done