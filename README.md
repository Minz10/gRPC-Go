gRPC-Golang-demoproject


#command to run the greet.proto
protoc --go_out=. --go-grpc_out=. proto/greet.proto