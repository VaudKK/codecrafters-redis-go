package main

import (
	"fmt"
	"os"
)

const EOF = 0xff
const AUX = 0xfa
const RESIZEDB = 0xfb
const SELECTDB = 0xfe
const EXPIRETIME = 0xfd
const EXPIRETIMEMS = 0xfc

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
	//Read the first 9 bytes
	value := readHeader(fileData)
	fmt.Println("Header:",string(value))
	
	for  {
		switch readByte(fileData) {
		case AUX:
			value := readMetadata(fileData)
			fmt.Println("Read:",value)
			fmt.Println("Next metadata will start from:",pos)
			value += readMetadata(fileData)
			fmt.Println("Aux:", string(value))
		case SELECTDB:
			value := readByte(fileData)
			fmt.Println("Database index:", string(value))
		case RESIZEDB:
			hashTableSize := int(readByte(fileData))
			keysWithExpiry := int(readByte(fileData))
			fmt.Println("HashTable info:", hashTableSize, keysWithExpiry)
		case EXPIRETIME:
		case EXPIRETIMEMS:
		case EOF:
			return ""
		default:
			return ""
		}
	}
}

func readHeader(fileData []byte) string {
	return string(readBytesOffset(fileData, 0, 9))
}

func readMetadata(fileData []byte) string {
	length := int(readByte(fileData))
	fmt.Println("Length to be read:", length)
	return string(readBytesOffset(fileData, pos, length))
}

func readByte(fileData []byte) byte {
	value := fileData[pos]

	//Advance by one to the next byte to be read
	pos += 1
	return value
}

func readBytesOffset(fileData []byte, offset int, length int) []byte {
	data := make([]byte,length)

	j := 0
	for i := offset; i < length; i++ {
		data[j] = readByte(fileData)
		j += 1
	}

	return data
}

// func readBytes(fileData []byte,length int) []byte{
// 	value := fileData[pos:(pos + length)]
// 	pos += length -1
// 	return value
// }

func readStringEncoding(fileData []byte) []byte {
	length := int(readByte(fileData))
	return readBytesOffset(fileData, pos, length)
}
