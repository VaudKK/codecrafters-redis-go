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

	data := make([]byte, 512)

	_, err = file.Read(data)

	if err != nil {
		fmt.Println("Error reading file: ", err)
		return nil
	}

	return data
}

func parseRdb(fileData []byte) string {
	header := readHeader(fileData)
	fmt.Println("Read header: ",string(header))

	for fileData[pos] != EOF {
		fmt.Printf("%d",pos)
		switch fileData[pos] {
		case AUX:
			value := readMetadata(fileData)
			value += readMetadata(fileData)
		case SELECTDB:
			value := readByte(fileData)
			fmt.Println("Database index: ", string(value))
		case RESIZEDB:
			hashTableSize := int(readByte(fileData))
			keysWithExpiry := int(readByte(fileData))
			fmt.Println("HashTable info: ", hashTableSize, keysWithExpiry)
		// case 0x00:
		// 	key := readStringEncoding(fileData)
		// 	value := readStringEncoding(fileData)
		// 	keyValue[string(key)] = struct {
		// 		value      string
		// 		expiration int64
		// 	}{
		// 		string(value), -1,
		// 	}
		// 	fmt.Println("Read first key: ", keyValue)
		}
	}

	return ""

}

func readHeader(fileData []byte) string {
	return string(readBytesOffset(fileData, 0, 9))
}

func readMetadata(fileData []byte) string {
	length := int(readByte(fileData))
	return string(readBytesOffset(fileData, pos, length))
}

func readByte(fileData []byte) byte {
	value := fileData[pos]
	pos += 1
	return value
}

func readBytesOffset(fileData []byte, offset int, length int) []byte {
	destination := fileData[offset:(offset + length)]
	pos += length
	return destination
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
