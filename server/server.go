package main

import (
	"context"
	"flag"
	"fmt"
	gRPC "grpcBrr/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	gRPC.UnimplementedTimeAskServiceServer
	name string
	port int
}

// to allow falgs in console go run server/server.go -port ?
var port = flag.Int("port", 8080, "server port number")

func main() {
	flag.Parse()

	server := &Server{
		name: "server1",
		port: *port,
	}

	go StartServer(server)

	for {
	}
}

func StartServer(server *Server) {
	gRPCServer := grpc.NewServer()

	listener, err := net.Listen("tcp", ":"+fmt.Sprint(server.port))

	if err != nil {
		log.Fatal("F - listener unable to start")
	}

	log.Printf("Server %s started, listening on port %d\n", server.name, server.port)

	gRPC.RegisterTimeAskServiceServer(gRPCServer, server)

	serverErr := gRPCServer.Serve(listener)
	if serverErr != nil {
		log.Fatal("F - could not register server")
	}
}

func (c *Server) GetTime(ctx context.Context, in *gRPC.AskForTimeMsg) (*gRPC.TimeMsg, error) {
	log.Printf("Client with ID %d asked for time\n", in.ClientId)

	return &gRPC.TimeMsg{
		Time:       time.Now().String(),
		ServerName: c.name,
	}, nil
}
