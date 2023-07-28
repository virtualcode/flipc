package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing input file")
	}
	fmt.Println("Flipcase of " + os.Args[1])
	oldName := os.Args[1]
	newName := strings.ToLower((oldName))

	os.Rename(oldName, newName)
}
