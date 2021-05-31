package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "Welcome to Golang!!" //Generic output

	if len(os.Args) > 1 {
		var strTemp = []string{os.Args[1], "Welcome!"}
		who = strings.Join(strTemp, " ") //using input args
	}

	fmt.Println("Hello,", who)
}
