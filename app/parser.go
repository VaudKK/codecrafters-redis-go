package main

import (
	"fmt"
	"net"
	"strings"
)

type RESPType int

const (
	SimpleString RESPType = iota
	Integer
	BulkString
	Arrays
)

func getRESPType(line string) (RESPType, []string) {
	tokens := strings.Split(line, "\r\n")

	if strings.HasPrefix(tokens[0], "*") {
		return Arrays, tokens
	} else if strings.HasPrefix(tokens[0], ":") {
		return Integer, tokens
	}

	return SimpleString, tokens
}

func parse(respType RESPType,tokens []string, connection net.Conn) {
	switch respType {
	case Arrays:
		tokens := parseArray(tokens)
		fmt.Println("Parsed Array:", tokens)
	default:
		panic("Unknown RESP type")
	}
}


func parseArray(tokens []string) []string {
	data := make([]string, 0)

	for _, token := range tokens {
		if strings.HasPrefix(token, "$") {
			continue
		}

		if strings.HasPrefix(token, "*") {
			continue
		}

		data = append(data, token)
	}

	return data

}
