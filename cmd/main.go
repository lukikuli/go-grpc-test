package main

import (
	"go-grpc/handler"
	"go-grpc/protobuf/period"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	//Create TCP Server on localhost:5001
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}
	//Create new gRPC server handler
	server := grpc.NewServer()

	//register gRPC UserService to gRPC server handler
	period.RegisterPeriodHandlerServer(server, &handler.PeriodHandler{})

	//Run server
	if err := server.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}
