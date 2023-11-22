package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

func help() {
	fmt.Println(`basename - strip directory and suffix from filenames

Usage: basename OPTION... NAME...

	-h : display help text
	-s : remove a trailing suffix

Examples:
	basename /usr/bin/dumplings
		-> "dumplings"

	basename -s=.h include/stdio.h
		-> "stdio"

	basename any/str1 any/str2
		-> "str1"
		   "str2"
 `)
	os.Exit(0)
}

func main() {
	s := flag.String("s", "", "trailing suffix")
	h := flag.Bool("h", false, "help text")

	flag.Parse()
	names := flag.Args()

	for _, dir := range names {
		stripName := strings.Split(dir, "/")
		stripName = slices.DeleteFunc(stripName, func(st string) bool {
			return st == ""
		})
		nosuffixName := strings.TrimSuffix(stripName[len(stripName)-1], *s)
		fmt.Println(nosuffixName)
	}

	if *h == true {
		help()
	}
}
