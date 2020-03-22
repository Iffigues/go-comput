package main

import (
	"os"
)

func main() {
	var a []string
	if len(os.Args) == 1 {
		a = append(a, args())
	} else {
		a = os.Args[1:]
	}
	newObj(a)
}
