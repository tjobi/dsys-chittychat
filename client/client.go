package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	gRPC "grpcBrr/proto"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	id         int
	portNumber int
}

var (
	clientPort = flag.Int("clientPort", 8081, "client port number")
	serverPort = flag.Int("serverPort", 8080, "server port number")
)

func main() {
	flag.Parse()

	client := &Client{
		id:         1,
		portNumber: *clientPort,
	}

	go startClient(client)

	for {
		time.Sleep(5 * time.Second)
	}
}

func startClient(client *Client) {
	conn := GetServerConnection()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		log.Printf("Client: %d input %s\n", client.id, input)

		timemsg, err := conn.GetTime(context.Background(), 
									&gRPC.AskForTimeMsg{ClientId: int64(client.id)})

		if err != nil {
			log.Printf("Could not get time\n")
		}

		log.Printf("Server[%s] says that the time is %s\n",timemsg.ServerName, timemsg.Time)
	}
}

func GetServerConnection() gRPC.TimeAskServiceClient {
	conn, err := grpc.Dial(":"+fmt.Sprint(*serverPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("Client F - could not dial")
	}

	log.Printf("Dialed succesfully")

	return gRPC.NewTimeAskServiceClient(conn)
}
