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
			filename := args[0]
			data, err := os.ReadFile(filename)
			if err != nil {
				fmt.Printf("error: unable to read file '%s'\n", filename)
				os.Exit(1)
			}
			content := string(data)
			fmt.Println(content)
		}
	}
}
