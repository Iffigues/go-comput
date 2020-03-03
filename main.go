package main

import (
	"os"
	"log"
	"strings"
)

func verif(a string)  (b []op, ok bool) {
	print(a)
	return
}

func sublime(a string) (b string) {
	return strings.Join(strings.Fields(a), " ")
}

func beg(a []op) {
}

func main() {
	var a []op
	var ok bool
	if len(os.Args) > 1 {
		if a , ok = verif(sublime(strings.Join(os.Args[1:], " "))); !ok {
			log.Fatal("error occure")
		}
	} else {
		bd:
		for  {
			if a, ok = verif(sublime(args())); ok {
				break bd
			}
		}
	}
	beg(a)
}
