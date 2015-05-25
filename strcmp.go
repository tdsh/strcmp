package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Specify 2 strings you want to compare.")
		os.Exit(1)
	}
	var result = (os.Args[1] == os.Args[2])
	var color = 31
	var answer = "False"
	if result {
		color = 32
		answer = "True"
	}
	fmt.Printf("\033[%d;1m%s\033[0m\n", color, answer)
	os.Exit(0)
}
