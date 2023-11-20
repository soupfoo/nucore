package main

import (
	"fmt"
	"os"
	"strings"
)

func help() {
	fmt.Println(`wc - print newline and byte counts for each file

Usage: wc [OPTION]... [FILE]...

	-h : display help text
	-l : line count
	-b : byte count
 `)
	os.Exit(0)
}

func countNewLines(x []byte) int {
	count := 0
	for _, val := range x {
		if string(val) == "\n" {
			count++
		}
	}
	return count
}

func countBytes(x []byte) int {
	return len(x)
}

func evaluateFlags(args []string, l, b *bool) {
	for i := 0; i < len(args); i++ {
		x := args[i]
		if strings.HasPrefix(x, "-") {
			for j := 1; j < len(x); j++ {
				switch string(x[j]) {
				case "l":
					*l = true
				case "b":
					*b = true
				case "h":
					help()
				default:
					fmt.Println("Invalid flag")
				}
			}
		}
	}
}

func main() {
	var (
		l, b           bool
		countL, countB int
	)
	args := os.Args[1:]
	evaluateFlags(args, &l, &b)

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			continue
		}
		data, _ := os.ReadFile(arg)
		if l {
			countL = countNewLines(data)
		}
		if b {
			countB = countBytes(data)
		}
		fmt.Println(countL, countB, "\t\033[34;1m\U0000f016", arg, "\033[0m")

	}
}
