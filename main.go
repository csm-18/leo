package main

import (
	"fmt"
	"os"
	"strings"
)

const LEO_VERSION = "leo 0.0.1"

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println(LEO_VERSION)
	} else if len(args) == 1 {
		if strings.HasSuffix(args[0], ".leo") {
			fmt.Println("compiling", args[0])
		}
	}
}
