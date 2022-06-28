package main

import (
	"github.com/kalogs-c/zip_landia/unzip"
	"log"
	"os"
)

func main() {
	args := os.Args
	err := ziplandia.UnzipSource(args[1], args[2])
	if err != nil {
		log.Fatal(err)
	}
}
