package main

import (
	"fmt"
	"os"
)

type PR struct {
	Name string
}

func main() {
	argsCount := len(os.Args)
	if argsCount == 1 {
		return
	}

	if os.Args[1] == "pr" {
		fmt.Println("the message is: ", os.Args[1])
	}
}
