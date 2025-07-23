package main

import (
	"fmt"
	"os"
)

const EOF = 0xFF
const AUX = 0xFA
const RESIZEDB = 0xFB
const SELECTDB = 0xFE
const EXPIRETIME = 0xFD
const EXPIRETIMEMS = 0xFC

func readDatabaseFile(path string) []byte {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error opening rdb file: ", err)
		return nil
	}

	data := make([]byte, 0)

	_, err = file.Read(data)

	if err != nil {
		fmt.Println("Error reading file: ", err)
		return nil
	}

	fmt.Println(string(data[:9]))

	return nil
}
