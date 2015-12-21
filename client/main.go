package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/hatajoe/grpc-sample/Player"
)

type Headers []string

const CLIENT_NAME = "GoClient"
const CLIENT_ID = 2
const CLIENT_DESCRIPTION = "This is a Go Protobuf client!!"

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func (h Headers) getHeaderIndex(headername string) int {
	if len(headername) >= 2 {
		for index, s := range h {
			if s == headername {
				return index
			}
		}
	}
	return -1
}

func retrieveData() ([]byte, error) {
	fmt.Println("retrieveData")
	ProtoMessage := new(Player.Player)
	ProtoMessage.Name = proto.String(CLIENT_NAME)
	ProtoMessage.Id = proto.Int32(CLIENT_ID)

	//Populate items
	pos := new(Player.Player_Position)
	x := int32(1)
	y := int32(0)
	pos.X = &x
	pos.Y = &y

	ProtoMessage.Pos = append(ProtoMessage.Pos, pos)
	fmt.Println(ProtoMessage)
	return proto.Marshal(ProtoMessage)
}

func sendDataToDest(data []byte, dst *string) {
	fmt.Println("sendDataToDest")
	conn, err := net.Dial("tcp", *dst)
	checkError(err)
	n, err := conn.Write(data)
	checkError(err)
	fmt.Println("Sent " + strconv.Itoa(n) + " bytes")
}

func main() {
	fmt.Println("prepared")
	dest := flag.String("d", "127.0.0.1:2110", "Enter the destnation socket address")
	flag.Parse()
	data, err := retrieveData()
	checkError(err)
	sendDataToDest(data, dest)
}
