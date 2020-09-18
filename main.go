package main

import (
	"fmt"
	"github.com/bablzz/algowithgo/pkg/algo"
)

func main() {
	var linkedList algo.LinkedList
	linkedList = algo.LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)
	fmt.Println(linkedList.NodeWithValue(5).Proprety)
	linkedList.IterateList()

	fmt.Println("\nCheck tuples")
	fmt.Println(algo.A(algo.B(12, 5)))
}
