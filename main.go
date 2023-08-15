package main

import "fmt"

func main() {
	var a func()
	var b = 12
	a = func() {
		fmt.Println(b)
	}
	a()
}