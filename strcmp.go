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
	fmt.Printf("%t\n", os.Args[1] == os.Args[2])
	os.Exit(0)
}
