package main

import (
	"os"
	"log"
	"fmt"
	"strings"
)

func addSpace(a string) (b string){
	for _, zz := range a {
		z := string(zz)
		r := strings.ContainsAny(z, "+*-/^=")
		fmt.Println(r)
		if r {
			b = b + " "
		}
		b = b + z
		if r {
                        b = b + " "
                }
	}
	return
}

func verif(a string)  (b []op, ok bool) {
	a = strings.Join(strings.Fields(addSpace(a))," ")
	g := strings.Split(a, " ")
	fmt.Println(g)
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
