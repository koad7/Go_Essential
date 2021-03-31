package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + +i + os.Args[i] + os.Args[0]
		sep = "\n"
	}
	fmt.Println(s)
}
