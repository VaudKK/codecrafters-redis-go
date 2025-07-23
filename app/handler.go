package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

var Handler = map[string]func([]string, net.Conn){
	"PING":   ping,
	"ECHO":   echo,
	"SET":    set,
	"GET":    get,
	"CONFIG": config,
	"KEYS":   readKeys,
}

var keyValue = make(map[string]struct {
	value      string
	expiration int64
})

func echo(tokens []string, connection net.Conn) {
	response := ""
	for _, token := range tokens[1:] {
		response += fmt.Sprintf("$%d\r\n%s\r\n", len(token), token)
	}
	connection.Write([]byte(response))
}

func ping(tokens []string, connection net.Conn) {
	connection.Write([]byte("+PONG\r\n"))
}

func config(tokens []string, connection net.Conn) {
	if tokens[1] == "GET" {
		switch tokens[2] {
		case "dir":
			data := map[string]int{"dir": len("dir"), redisConfig.Dir: len(redisConfig.Dir)}
			connection.Write([]byte(writeArray(data)))
		case "dbfilename":
			data := map[string]int{"dbfilename": len("dbfilename"), redisConfig.DBFileName: len(redisConfig.DBFileName)}
			connection.Write([]byte(writeArray(data)))
		}
	}
}

func set(tokens []string, connection net.Conn) {
	_, ok := keyValue[tokens[1]]

	if !ok {
		keyValue[tokens[1]] = struct {
			value      string
			expiration int64
		}{
			value:      tokens[2],
			expiration: -1,
		}
	}

	// set the expirtation time if provided
	if len(tokens) > 3 {
		px, err := strconv.Atoi(tokens[4])
		if err == nil {
			setPx(tokens[1], int64(px))
		}
	}

	connection.Write([]byte("+OK\r\n"))
}

func setPx(key string, px int64) {
	value := keyValue[key]
	value.expiration = time.Now().UnixMilli() + px
	keyValue[key] = value
}

func get(tokens []string, connection net.Conn) {
	value, ok := keyValue[tokens[1]]
	if !ok || (value.expiration > -1 && value.expiration < time.Now().UnixMilli()) {
		writeNoKeyFound(connection)
		return
	}

	response := fmt.Sprintf("$%d\r\n%s\r\n", len(value.value), value.value)
	connection.Write([]byte(response))
}

func readKeys(tokens []string, connection net.Conn) {
	readDatabaseFile(fmt.Sprintf("%s%c%s", redisConfig.Dir, os.PathSeparator, redisConfig.DBFileName))
}

func writeArray(contents map[string]int) string {
	length := len(contents)
	value := fmt.Sprintf("*%d\r\n", length)

	for key, val := range contents {
		value += fmt.Sprintf("$%d\r\n%s\r\n", val, key)
	}

	return value
}

func writeNoKeyFound(connection net.Conn) {
	connection.Write([]byte("$-1\r\n"))
}
