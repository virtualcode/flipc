package main

import (
	"flipc"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing input file")
	}

	oldfilename := os.Args[1]
	flipc.Flipcase(oldfilename)
}
