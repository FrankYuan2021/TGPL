package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		// range : for i in len(DataList);return i, DataList[i]
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
