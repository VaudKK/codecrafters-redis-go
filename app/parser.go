package main

import "strings"

type RESPType int

const (
	SimpleString RESPType = iota
	Integer
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

func parse(respType RESPType,tokens []string){
	switch respType {
	case Arrays:
		parseArray(tokens)
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

func delegate()
