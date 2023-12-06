package main

import (
	"log"
	"os"
)

func main() {
	if err := os.WriteFile("file_write.txt", []byte("Hello World!!"), 0666); err != nil {
		log.Fatal(err)
	}
}
