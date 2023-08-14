package main

import "fmt"

func main() {
	str := "strings"
	for _,s := range str {
		fmt.Printf("%c\n",s)
	}
}
