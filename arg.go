package main

import (
	"bufio"
	"fmt"
	"os"
)

func args() (b string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	b = text[:len(text)-1]
	return
}
