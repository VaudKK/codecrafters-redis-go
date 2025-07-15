package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {

		connection, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(connection)

		fmt.Println("Accepted new connection from", connection.RemoteAddr())
	}

}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	data := make([]byte, 2048)
	_, err := connection.Read(data)

	if err != nil {
		if err != io.EOF {
			fmt.Println("Error reading from connection: ", err.Error())
			return
		}
	}

	handle(data, connection)
}

func handle(line []byte, connection net.Conn) {
	firstByte := []int32{'+', '-', ':', '*'}

	for _, prefix := range firstByte {
		if line[0] == byte(prefix) {
			handleCommand(strings.TrimRight(string(line), "\x00"), connection)
			return
		}
	}
}

func handleCommand(line string, connection net.Conn) {
	dataType, tokens := getRESPType(line)
	parse(dataType, tokens, connection)
}
