package main

import (
	"log"

	pb "github.com/Minz10/gRPC-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Akhil", "Alice", "Bob"},
	}

	//<-----UNARY API----->
	//callSayHello(client)
	
	//<-----Server Streaming----->
	//callSayHelloServerStream(client, names)

	//<-----Client Streaming---->
	//callSayHelloClientStream(client, names)

	//<-----Bi-Directional Streaming----->
	callHelloBidirectionalStream(client, names)
}