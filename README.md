# gRPC-Golang-demoproject

**This project is to get an understanding of how RPC API works using gRPC in Golang.**
_I have implemented the four types,_
1. Unary
2. Server Streaming
3. Client Streaming
4. Bi-directional Streaming


## Command to run the greet.proto file.
### Note : Make sure to install protoc compiler to execute the greet.proto file
protoc --go_out=. --go-grpc_out=. proto/greet.proto
