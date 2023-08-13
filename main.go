package main

import (
	"container/list"
	"fmt"
)

func main() {
	st := list.New()
	st.PushBack(34)
	fmt.Println(st.Front().Value)
}
