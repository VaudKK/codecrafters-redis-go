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

var pos int = 0

func readDatabaseFile(path string) []byte {
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

	return data
}


func parseRdb(fileData []byte) string {
	readHeader(fileData)

	currentByte := fileData[pos]

	for currentByte != EOF {
		switch currentByte {
		case AUX:
			value := readMetadata(fileData)
			value += readMetadata(fileData)
		case SELECTDB:
			value := readByte(fileData)
			fmt.Println("Database index: ", string(value))
		case RESIZEDB:
			hashTableSize := int(readByte(fileData))
			keysWithExpiry := int(readByte(fileData))
			fmt.Println("HashTable info: ",hashTableSize,keysWithExpiry)
		}
	}
	
	return ""
	
}

func readHeader(fileData []byte) string{
	return string(readBytesOffset(fileData,0,9))
}

func readMetadata(fileData []byte) string{
	length := int(readByte(fileData))
	return string(readBytesOffset(fileData,pos,length))
}


func readByte(fileData []byte) byte {
	value := fileData[pos]
	pos += 1
	return value
}

func readBytesOffset(fileData []byte,offset int, length int) []byte{
	destination := fileData[offset:(offset + length)]
	pos += length - 1
	return destination
}

func readBytes(fileData []byte,length int) []byte{
	value := fileData[pos:(pos + length)]
	pos += length -1
	return value
}
