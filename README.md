# dsys-chittychat

## Running our chitty-chat

1. From the root folder run ```go run server/server.go```
   1. The server will default to listen on port 8080.
   2. A different port can be selected by using the ```-port``` flag followed by a port number.
2. In a new different terminal start the clients by running ```go run client/client.go```.
   1. Clients will default to port 8081. When running multiple clients, make sure to give each client a unique port using the ```-clientPort``` flag.
   2. If the serverport was set to something other than the default *8080*, remember to use the ```-serverPort``` to pass along the new port.
   3. Optional. Clients can be given a name using the ```-clientName``` flag.
3. The chitty chat is now ready for use!
4. When it is time for a client to *exit* the chat room, the client should type in ```-exit```.

## Copy and paste to start
   
      1. go run server/server.go
      
      2. go run client/client.go
      
      3. go run client/client.go -clientPort 8085
      
      4. go run client/client.go -clientPort 8090 

