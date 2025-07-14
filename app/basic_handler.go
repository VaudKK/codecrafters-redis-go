package main

import "net"

func pingBasic(connection net.Conn) {
	connection.Write([]byte("+PONG\r\n"))
}