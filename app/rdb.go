package main

import (
	"fmt"
	"os"
)

func readDatabase(path string) map[string]string {
	_, err := os.Open(path)

	if err != nil {
		fmt.Println("Error reading rdb file: ", err)
		return nil
	}

	return nil
}
