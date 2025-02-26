demo_api:
	go build -ldflags="-s -w" -p 4 -o ./bin/demo ./demo/main/main.go

gen_code: protoc thrift

protoc:
	sh scripts/protoc-gen-go.sh

thrift:
	thrift -r --strict --gen go thrift_files/ew_go_thrift/ew_go.thrift