#!/bin/bash

GENPATH=gen-go/protos
if [ -e $GENPATH ]; then
    rm -rf ./gen-go/proto/*
else
    #echo "Path: $GENPATH does not exist"
    mkdir -p $GENPATH
fi

for d in `find protos -name "*.proto"`; do
    protoc --go_out=paths=source_relative:gen-go/protos \
    --go-grpc_out=paths=source_relative:gen-go/protos \
    --proto_path=./protos \
    ./${d};
done
