package main

import (
	"fmt"
	"net"
)


 var Handler = map[string] func([]string,net.Conn){
	"PING": ping,
	"ECHO": echo,
	"SET":  set,
	"GET":  get,
}

var keyValue = make(map[string]string)


func echo(tokens []string, connection net.Conn) {
	response := ""
	for _, token := range tokens[1:] {
		response += fmt.Sprintf("$%d\r\n%s\r\n", len(token), token)
	}
	connection.Write([]byte(response))
}

func ping(tokens []string,connection net.Conn) {
	connection.Write([]byte("+PONG\r\n"))
}

func set(tokens []string, connection net.Conn) {
	_,ok := keyValue[tokens[1]]

	if !ok {
		keyValue[tokens[1]] = tokens[2]
	}

	connection.Write([]byte("+OK\r\n"))
}

func get(tokens []string, connection net.Conn) {
	value, ok := keyValue[tokens[1]]
	if !ok {
		connection.Write([]byte("$-1\r\n"))
		return
	}

	response := fmt.Sprintf("$%d\r\n%s\r\n", len(value), value)
	connection.Write([]byte(response))
}