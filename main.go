package main

import "fmt"

func main() {
	var res []int
	fmt.Printf("%p\n", res)
	res = append(res,1)
	fmt.Printf("%p\n", res)
	fmt.Println(res)
}