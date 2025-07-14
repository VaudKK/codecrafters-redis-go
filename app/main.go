package main

import (
	"bufio"
	"fmt"
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
	}

}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	scanner := bufio.NewScanner(connection)

	for scanner.Scan() {
		line := scanner.Text()
		handle(line, connection)
	}

	fmt.Println("Connection closed by client")

}

func handle(line string, connection net.Conn) {
	firstByte := []string{"+", "-", ":", "*"}

	fmt.Println("Received line:", line)

	for _, prefix := range firstByte {
		if strings.HasPrefix(line, prefix) {
			handleEncodedCommand(line, connection)
			return
		}
	}

	handleBasicCommand(line, connection)
}

func handleBasicCommand(line string, connection net.Conn) {
	command := strings.Split(line, " ")

	switch strings.ToUpper(command[0]) {
	case "PING":
		pingBasic(connection)
	case "ECHO":
		echo(command[1:], connection)
	default:
		connection.Write([]byte("-ERR unknown command '" + command[0] + "'\r\n"))
	}
}

func handleEncodedCommand(line string, connection net.Conn) {
	dataType, tokens := getRESPType(line)
	parse(dataType, tokens)
}
