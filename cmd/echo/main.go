package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		if string(arg[0]) == "$" {
			enVar := os.Getenv(arg)
			fmt.Print(enVar, " ")
		} else {
			fmt.Print(arg, " ")
		}
	}
}
