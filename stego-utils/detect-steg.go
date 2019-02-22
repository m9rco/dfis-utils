package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func detect() {
	zipsig := []byte{'\x4b', '\x03', '\x04'}
	var flag bool

	file, err := os.Open("./noise.png")
	handerr(err)
	defer file.Close()

	bfio := bufio.NewReader(file)

	fstat, err := file.Stat()
	handerr(err)

	for i := int64(0); i < fstat.Size(); i++ {
		thisbyte, err := bfio.ReadByte()
		handerr(err)

		if thisbyte == '\x50' {
			byteslice := make([]byte, 3)
			byteslice, err := bfio.Peek(3)
			handerr(err)

			if bytes.Equal(byteslice, zipsig) {
				fmt.Println("Found Zip Signature at: ", i)
				flag = true
				break
			}
		}
	}

	if !flag {
		fmt.Println("Stegnography NOT detected")
	}
}