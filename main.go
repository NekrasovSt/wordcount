package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("0")
	}
	fmt.Println(len(strings.Fields(os.Args[1])))
}
