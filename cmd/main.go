package main

import (
	"fmt"
	"os"
)

func main() {
	argsCount := len(os.Args)
	if argsCount == 1 {
		return
	}

	if os.Args[1] == "req" {
		fmt.Println("the message is: ", os.Args[1])
	}
}
