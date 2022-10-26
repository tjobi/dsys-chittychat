package main

import (
	"context"
	"flag"
	"fmt"
	gRPC "grpcBrr/proto"
	"log"
	"math"
	"net"
	"os"
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
	port         = flag.Int("port", 8080, "server port number")
	streams      []gRPC.ChittyChat_SendMessageServer
	streamActive map[int32]bool

	lastId         int
	serverVClock   []int32
	serverPosition int
)

func main() {
	flag.Parse()

	//https://stackoverflow.com/a/19966217
	f, err := os.OpenFile("output2.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	server := &Server{
		name: "server1",
		port: *port,
	}

	streamActive = make(map[int32]bool)

	streams = make([]gRPC.ChittyChat_SendMessageServer, 0)
	serverVClock = make([]int32, 1)
	serverPosition = 0
	lastId = 0

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

	for {
		clientIn, err := stream.Recv()

		if err != nil {
			log.Printf("Client disconnected\n")
			return err
		}

		MergeVClock(clientIn.Timestamp)

		serverVClock[serverPosition]++

		//	log.Printf("Server recieved %s with timestamp: %v", clientIn.Msg, clientIn.Timestamp)
		log.Printf("ClientClock: %d, Client %d PUBLISH: \"%s\"", clientIn.Timestamp, clientIn.ClientId, clientIn.Msg)

		var sendErr error
		for i, s := range streams {
			if clientIn.ClientId != int32(i+1) {
				if streamActive[int32(i+1)] {
					s.Send(&gRPC.ServerReponse{
						Msg:       clientIn.Msg,
						Timestamp: serverVClock,
					})
				}
			}
		}

		serverVClock[serverPosition]++

		log.Printf("ServerClock after publish: %v", serverVClock)

		if sendErr != nil {
			return sendErr
		}

	}
}

/*func removeClientStream(s []gRPC.ChittyChat_SendMessageServer, i int) []gRPC.ChittyChat_SendMessageServer {
	s[i] = s[len(s)-1]
	new := s[:len(s)-1]
	return new
} */

func (c *Server) JoinRoom(ctx context.Context, in *gRPC.ClientJoin) (*gRPC.ServerWelcome, error) {

	serverVClock = append(serverVClock, 0)
	serverVClock[serverPosition] += 1
	VClock := serverVClock

	lastId += 1

	streamActive[int32(lastId)] = true

	for i, s := range streams {
		if streamActive[int32(i+1)] {
			s.Send(&gRPC.ServerReponse{
				Timestamp: serverVClock,
				Msg:       fmt.Sprintf("JOINED: Client with id: %d, and name: %s\n", lastId, in.ClientName),
			})
		}
	}

	log.Printf("ClientClock: %d, Client %d JOINED", VClock, lastId)

	return &gRPC.ServerWelcome{
		VClock: VClock,
		Id:     int32(lastId),
	}, nil

}

func (c *Server) LeaveRoom(ctx context.Context, in *gRPC.ClientLeave) (*gRPC.ServerBye, error) {
	serverVClock[serverPosition] += 1
	MergeVClock(in.Timestamp)

	streamActive[int32(in.ClientId)] = false

	for i, s := range streams {
		if streamActive[int32(i+1)] {
			s.Send(&gRPC.ServerReponse{
				Timestamp: serverVClock,
				Msg:       fmt.Sprintf("LEFT: Client with id: %d, and name: %s\n", lastId, in.ClientName),
			})
		}
	}

	log.Printf("ClientClock: %d, Client %d LEFT", in.Timestamp, in.ClientId)

	serverVClock[serverPosition] += 1

	return &gRPC.ServerBye{
		Msg: "Subscribe to my onlyfans ...",
	}, nil
}

func MergeVClock(clientClock []int32) {
	if len(clientClock) > len(serverVClock) {
		for i := 0; i < len(clientClock)-len(serverVClock); i++ {
			serverVClock = append(serverVClock, 0)
		}
	}

	for i := range clientClock {
		serverVClock[i] = int32(math.Max(float64(serverVClock[i]), float64(clientClock[i])))
	}
}
