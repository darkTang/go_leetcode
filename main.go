package main

import "fmt"

func main() {
	res := []int{1, 2}
	fmt.Printf("%p\n", res)
	res = append(res, 1)
	fmt.Printf("%p\n", res)
}