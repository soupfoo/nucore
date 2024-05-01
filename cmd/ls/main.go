package main

import (
	"fmt"
	"os"
	"strings"
)

func help() {
	fmt.Println(`ls - list contents of a directory

Usage: cat [OPTION]... [Directory]

	-h : display help text
	-a : do not ignore hidden entries
	-l : use long listing format
 `)
	os.Exit(0)
}

var (
	oneKB = 1024.0
	oneMB = oneKB * 1024
	oneGB = oneMB * 1024
	oneTB = oneGB * 1024
)

func fileSize(s float64) (float64, string) {
	switch {
	case s >= oneKB && s < oneMB:
		return s / oneKB, "K"
	case s >= oneMB && s < oneGB:
		return s / oneMB, "M"
	case s >= oneGB && s < oneTB:
		return s / oneGB, "G"
	case s >= oneTB:
		return s / oneTB, "T"
	}
	return s, "B"
}

func evaluateFlags(args []string, a, l *bool) {
	for i := 0; i < len(args); i++ {
		x := args[i]
		if strings.HasPrefix(x, "-") {
			for j := 1; j < len(x); j++ {
				switch string(x[j]) {
				case "a":
					*a = true
				case "l":
					*l = true
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
		dir         = "."
		showAll     bool
		showDetails bool
	)

	args := os.Args[1:]
	evaluateFlags(args, &showAll, &showDetails)

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			continue
		}
		if arg != "" {
			dir = arg
		}
	}
	fmt.Println("[", dir, "]:")
	list(dir, showDetails, showAll)
}

func list(dir string, showDetails, showAll bool) {
	var (
		blue   = "\033[34;1m"
		purple = "\033[35;1m"
		reset  = "\033[0m"
		icon   string
	)

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("directory", dir, "doesn't exist")
		fmt.Println("pass -h flag to see the help text")
	}

	for _, f := range files {
		name := f.Name()
		if string(name[0]) == "." && !showAll {
			continue
		}
		fileinfo, _ := f.Info()
		perms := fileinfo.Mode()
		bytes := fileinfo.Size()
		size, unit := fileSize(float64(bytes))

		if showDetails {
			fmt.Print(perms, "  ")
			fmt.Printf("%s%.1f %s%s\t", purple, size, unit, reset)
		}

		if f.IsDir() {
			icon = "\U0000f115 "
			fmt.Println(blue + icon + name + reset)
		} else {
			icon = "\U0000f016 "
			fmt.Println(icon + name)
		}
	}
	fmt.Println()
}
