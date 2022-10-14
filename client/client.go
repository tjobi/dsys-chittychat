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

	scanner := bufio.NewScanner(os.Stdin)

	// goroutine that only listens
	go func() {
		for {
			in, err := stream.Recv()

			if err == io.EOF {
				log.Println("Stream ended")
			}

			if err != nil {
				log.Fatal("ded")
			}

			if vClock == nil {
				vClock[client.id] += 1
			}
			MergeVClock(in.Timestamp)

			log.Printf("Client received msg: %s, at timestamp: %v\n", in.Msg, in.Timestamp)
		}
	}()

	for scanner.Scan() {
		input := scanner.Text()

		log.Printf("Client: %d input %s\n", client.id, input)

		vClock[client.id] += 1

		msgToSend := &gRPC.ChatMessage{
			Msg:       input,
			ClientId:  int32(client.id),
			Timestamp: vClock,
		}

		err := stream.Send(msgToSend)

		//timemsg, err := conn.SendMessage(context.Background())

		if err != nil {
			log.Printf("Could not get time\n")
		}

		log.Printf("Client send: %s\n", msgToSend.Msg)

		//log.Printf("Server[%s] says that the time is %s\n", timemsg.ServerName, timemsg.Time)
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

	for i, _ := range serverClock {
		vClock[i] = int32(math.Max(float64(vClock[i]), float64(serverClock[i])))
	}
}
