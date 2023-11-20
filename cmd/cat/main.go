package main

import (
	"bufio"
	"fmt"
	"os"
)

func help() {
	fmt.Println(`cat - concatenate file(s) to standard output

Usage: cat [OPTION]... [FILE]...

	-h : display help text
	-  : read standard input
	    (press ctrl+d to signal the end of input)

Examples:
	cat f - g  output f's contents, then standard input, then g's contents
	cat        copy standard input to standard output
 `)
	os.Exit(0)
}

func userInput() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		os.Stdout.Write([]byte(input))
	}
}

func main() {
	var i int
	args := os.Args[1:]

	if len(args) == 0 {
		userInput()
		os.Exit(0)
	}
	if args[0] == "-h" {
		help()
	}

	for i = 0; i < len(args); i++ {
		if args[i] == "-" {
			userInput()
			continue
		}
		data, _ := os.ReadFile(args[i])
		os.Stdout.Write(data)
	}
}
