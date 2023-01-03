package main

import (
	"alghorithms/src/linkedList"
	"fmt"
)

func main() {
	list := linkedList.New[int]()

	list.Prepend(2)
	list.Prepend(5)
	list.Append(3)

	fmt.Println(list)
	fmt.Println(list.StringInverse())

	val, _ := list.Pop()
	fmt.Println(val)
	fmt.Println(list)
	fmt.Println(list.StringInverse())

	fmt.Println(list.Contains(5))
	fmt.Println(list.Contains(3))
}
