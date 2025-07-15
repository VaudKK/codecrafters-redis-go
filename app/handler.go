package main

import (
	"fmt"
	"net"
)


 var Handler = map[string] func([]string,net.Conn){
	"PING": ping,
	"ECHO": echo,
}


func echo(tokens []string, connection net.Conn) {
	response := ""
	for _, token := range tokens[1:] {
		response += fmt.Sprintf("$%d\r\n%s\r\n", len(token), token)
	}
	connection.Write([]byte(response))
}

func ping(tokens []string,connection net.Conn) {
	fmt.Println("Received PING command")
	connection.Write([]byte("+PONG\r\n"))
}