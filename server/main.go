package main

import (
	"fmt"
	"net"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/hatajoe/grpc-sample/Player"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func writeValuesTofile(datatowrite *Player.Player) {
	fmt.Println(datatowrite)
}

func handleProtoClient(conn net.Conn, c chan *Player.Player) {
	fmt.Println("Connection established")
	//Close the connection when the function exits
	defer conn.Close()
	//Create a data buffer of type byte slice with capacity of 4096
	data := make([]byte, 4096)
	//Read the data waiting on the connection and put it in the data buffer
	n, err := conn.Read(data)
	checkError(err)
	fmt.Println("Decoding Protobuf message")
	//Create an struct pointer of type *Player.Player struct
	protodata := new(Player.Player)
	//Convert all the data retrieved into the Player.Player struct type
	err = proto.Unmarshal(data[0:n], protodata)
	checkError(err)
	//Push the protobuf message into a channel
	c <- protodata
}

func main() {
	fmt.Printf("Started ProtoBuf Server")
	c := make(chan *Player.Player)
	go func() {
		for {
			message := <-c
			writeValuesTofile(message)
		}
	}()
	//Listen to the TCP port
	listener, err := net.Listen("tcp", "127.0.0.1:2110")
	checkError(err)
	for {
		if conn, err := listener.Accept(); err == nil {
			//If err is nil then that means that data is available for us so we take up this data and pass it to a new goroutine
			go handleProtoClient(conn, c)
		} else {
			continue
		}
	}
}
