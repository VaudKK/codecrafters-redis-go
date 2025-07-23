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

	data := make([]byte,1024)

	_,err = file.Read(data)

	if err != nil {
		fmt.Println("Error reading file: ",err)
		return nil
	}

	header := data[:9]

	fmt.Println(header)

	return nil
}
