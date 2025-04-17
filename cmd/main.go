package main

import (
	"orders/handlers"

	"fmt"
	"net"
	"orders/db"

	pb "cloud_commons/order"

	"google.golang.org/grpc"
)

var (
	order_handler *handlers.OrderHandler
)

func main() {

	defer db.CloseConnection()

	lis, err := net.Listen("tcp", ":50052")

	if err != nil {
		panic("Failed to listen tcp on 50052")
	}

	server := grpc.NewServer()
	order_handler = handlers.NewOrderHandler()

	pb.RegisterOrderServiceServer(server, order_handler)

	fmt.Println("Starting inventory grpc on 50052")

	err = server.Serve(lis)

	if err != nil {
		panic("Failed to serve the inventory: " + err.Error())
	}
}
