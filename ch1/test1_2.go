package main

import (
	"fmt"
	"os"
)

func main() {
	for val,arg := range os.Args[1:]{
		fmt.Print(val)
		fmt.Print(arg)
		fmt.Print("\n")
	}
}
