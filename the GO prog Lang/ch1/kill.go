package main

import (
	"fmt"
	"os"
)

func main() {
	<-sigs
	fmt.Println("Got an interrupt")
	cleanup()
	os.Exit(2)
}
