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
	gRPC.UnimplementedChittyChatServer
	name string
	port int
}

var (
	// to allow falgs in console go run server/server.go -port ?
	port    = flag.Int("port", 8080, "server port number")
	streams []gRPC.ChittyChat_SendMessageServer

	lastId int
	serverVClock []int32
)

func main() {
	flag.Parse()

	server := &Server{
		name: "server1",
		port: *port,
	}

	streams = make([]gRPC.ChittyChat_SendMessageServer, 0)

	go StartServer(server)

	for {
		time.Sleep(5 * time.Second)
	}
}

func StartServer(server *Server) {
	gRPCServer := grpc.NewServer()

	listener, err := net.Listen("tcp", ":"+fmt.Sprint(server.port))

	if err != nil {
		log.Fatal("F - listener unable to start")
	}

	log.Printf("Server %s started, listening on port %d\n", server.name, server.port)

	gRPC.RegisterChittyChatServer(gRPCServer, server)

	serverErr := gRPCServer.Serve(listener)
	if serverErr != nil {
		log.Fatal("F - could not register server")
	}
}

func (c *Server) SendMessage(stream gRPC.ChittyChat_SendMessageServer) error {
	streams = append(streams, stream)
	log.Printf("Length of streams slice: %d\n", cap(streams))
	for {
		clientIn, err := stream.Recv()

		if err != nil {
			log.Printf("whoops\n")
			return err
		}
		log.Printf("Server recieved %s with timestamp: %v", clientIn.Msg, clientIn.Timestamp)

		var sendErr error
		for i, s := range streams {
			deadStream := s.Send(&gRPC.ServerReponse{
				Msg:       clientIn.Msg,
				Timestamp: clientIn.Timestamp,
			})
			if deadStream != nil {
				streams = removeClientStream(streams, i)
			}
		}

		if sendErr != nil {
			return sendErr
		}

	}
}

func removeClientStream(s []gRPC.ChittyChat_SendMessageServer, i int) []gRPC.ChittyChat_SendMessageServer {
	s[i] = s[len(s)-1]
	new := s[:len(s)-1]
	return new
}

func (c *Server) JoinRoom(ctx context.Context, in *gRPC.ClientJoin) (*gRPC.ServerWelcome, error) {
	//vClock = append(vClock, 0)
	//fresh slice for new client
	VClock := make([]int32, lastId + 1)
	lastId += 1;

	for _, s := range streams {
		s.Send(&gRPC.ServerReponse{
			Timestamp: serverVClock,
			Msg: fmt.Sprintf("Client with id: %d, and name: %s", lastId, in.ClientName),
		})
	}

	return &gRPC.ServerWelcome{
		VClock: VClock,
		Id: int32(lastId-1),
	}, nil

}
