package flipc

import (
	"fmt"
	"os"
	"strings"
)

func Flipcase(oldName string) {

	fmt.Println("Flipcase of " + os.Args[1])

	newName := strings.ToLower((oldName))

	os.Rename(oldName, newName)
}
