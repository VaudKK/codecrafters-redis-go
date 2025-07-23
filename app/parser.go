package main

import (
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

	// remove the last empty token if it exists
	if tokens[len(tokens)-1] == "" {
		tokens = tokens[:len(tokens)-1]
	}

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
		handler := Handler[strings.ToUpper(tokens[0])]
		handler(tokens, connection)
	default:
		panic("Unknown RESP type")
	}
}


func parseArray(tokens []string) []string {
	data := make([]string, 0)

	for index, token := range tokens {
		if strings.HasPrefix(token, "$") {
			continue
		}

		if(strings.HasPrefix(token, "*") && index > 0) {
			continue
		}

		data = append(data, token)
	}

	return data
}
