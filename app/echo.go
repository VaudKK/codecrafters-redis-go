package main

import (
	"net"
	"strings"
)

func echo(tokens []string, connection net.Conn) {
	response := strings.Join(tokens, " ")

	bulkString := "$" + string(len(response)) + "\r\n" + response + "\r\n"
	connection.Write([]byte(bulkString))
}