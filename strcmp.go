package main

import (
	"flag"
	"fmt"
	"os"
)

var help = `Usage of ./strcmp:
  strcmp STRING1 STRING2
`

func run() int {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, help)
	}
	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
		return 2
	}
	var result = (flag.Args()[0] == flag.Args()[1])
	var color = 31
	var answer = "False"
	if result {
		color = 32
		answer = "True"
	}
	fmt.Printf("\033[%d;1m%s\033[0m\n", color, answer)
	return 0
}

func main() {
	os.Exit(run())
}
