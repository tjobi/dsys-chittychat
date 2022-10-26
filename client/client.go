package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	gRPC "grpcBrr/proto"
	"io"
	"log"
	"math"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	id         int
	portNumber int
	//vClock     []int32
}

var (
	clientPort = flag.Int("clientPort", 8081, "client port number")
	id         = flag.Int("id", 1, "clientId")
	serverPort = flag.Int("serverPort", 8080, "server port number")
	name       = flag.String("clientName", "Client", "client name")
	vClock     []int32
)

func main() {
	flag.Parse()

	client := &Client{
		id:         *id,
		portNumber: *clientPort,
	}

	go startClient(client)

	for {
		time.Sleep(5 * time.Second)
	}
}

func startClient(client *Client) {
	conn, toClose := GetServerConnection()
	defer toClose.Close()

	stream, err := conn.SendMessage(context.Background())
	if err != nil {
		log.Fatal("client ded")
	}

	//
	welcome, err := conn.JoinRoom(context.Background(), &gRPC.ClientJoin{
		ClientName: *name,
	})

	if err != nil {
		log.Fatal("Welcome failed")
	}
	client.id = int(welcome.Id)
	vClock = welcome.VClock
	fmt.Println("joined, ", client.id)

	scanner := bufio.NewScanner(os.Stdin)

	didExit := false

	// goroutine that only listens
	go func() {
		for {
			in, err := stream.Recv()

			if didExit {
				log.Println("You succesfully disconnected.")
				return
			}

			if err == io.EOF {
				log.Println("Stream ended")
			}

			if err != nil {
				log.Fatal("Something went wrong, you have been disconnected.")
			}

			if vClock != nil {
				vClock[client.id] += 1
			}
			MergeVClock(in.Timestamp)

			log.Printf("%s received msg: %s, at timestamp: %d\n", *name, in.Msg, vClock)
		}
	}()

	for scanner.Scan() {
		input := scanner.Text()

		if input == "-exit" {
			vClock[client.id] += 1
			conn.LeaveRoom(context.Background(), &gRPC.ClientLeave{
				ClientId:   int32(client.id),
				Timestamp:  vClock,
				ClientName: *name,
			})
			stream.CloseSend()
			didExit = true
			return
		}

		msgToSend := &gRPC.ChatMessage{
			Msg:       input,
			ClientId:  int32(client.id),
			Timestamp: vClock,
		}

		err := stream.Send(msgToSend)

		if err != nil {
			log.Printf("Could not get time\n")
		}

		vClock[client.id] += 1
		log.Printf("%s send: %s, at timestamp: %d\n", *name, msgToSend.Msg, vClock)
	}
}

func GetServerConnection() (gRPC.ChittyChatClient, grpc.ClientConn) {
	conn, err := grpc.Dial(":"+fmt.Sprint(*serverPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("Client F - could not dial")
	}

	log.Printf("Dialed succesfully")

	return gRPC.NewChittyChatClient(conn), *conn
}

func MergeVClock(serverClock []int32) {
	if len(serverClock) > len(vClock) {
		for i := 0; i < len(serverClock)-len(vClock); i++ {
			vClock = append(vClock, 0)
		}
	}

	for i := range serverClock {
		vClock[i] = int32(math.Max(float64(vClock[i]), float64(serverClock[i])))
	}
}
