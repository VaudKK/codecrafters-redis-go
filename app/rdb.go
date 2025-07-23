package main

import (
	"fmt"
	"os"
)

func readDatabase(path string) map[string]string {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error opening rdb file: ", err)
		return nil
	}

	data := make([]byte, 1024)

	_, err = file.Read(data)

	if err != nil {
		fmt.Println("Error reading file: ", err)
		return nil
	}

	header := data[:0xFA]
	fmt.Println("Header", string(header))
	metadata := data[0xFA:0xFE]
	fmt.Println("Metadata", string(metadata))

	return nil
}
