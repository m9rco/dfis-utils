package main

import (
	"fmt"
	"log"
)

// Global Constants
const (
	IP      = "13.85.77.179" // mini client chat
	MINPORT = 1
	MAXPORT = 65535
)

func main() {
	fmt.Println("Scanning ports ....")
	scan()

	fmt.Println("Grabbing banners ....")
	grab()
}

func handerr(err error) {
	if err != nil {
		log.Println(err)
	}
}