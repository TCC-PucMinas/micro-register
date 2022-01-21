package main

import (
	"fmt"
	"net"

	"github.com/TCC-PucMinas/micro-register/communicate"
	"google.golang.org/grpc"
)

func main() {

	// port := os.Getenv("PORT")
	port := 4000
	host := fmt.Sprintf("0.0.0.0:%v", port)

	listener, err := net.Listen("tcp", host)

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	communicate.RegisterAuthenticateCommunicateServer(grpcServer, &communicate.Server{})

	fmt.Println("[x] - Server register listen...")

	if err := grpcServer.Serve(listener); err != nil {
		panic(err.Error())
	}

}
