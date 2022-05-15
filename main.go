package main

import (
	"net/http"
	"net/url"

	"github.com/bablzz/algowithgo/pkg/algo"
)

type Opts struct {
	URL         string
	RawResponse string
	Method      string
	Status      int
	Mux         *http.ServeMux
	QueryParams url.Values
}

type Foo struct {
	aa  bool
	ccc bool
	bb  int32
}

func main() {
	// var linkedList algo.LinkedList
	// linkedList = algo.LinkedList{}
	// linkedList.AddToHead(5)
	// linkedList.AddToEnd(1)
	// linkedList.AddToEnd(3)
	// linkedList.AddToEnd(7)
	// linkedList.AddToEnd(3)
	// linkedList.IterateList()
	// exercise.EvenOdd(&linkedList)
	// fmt.Print("\n")
	// linkedList.IterateList()

	// type Key string
	//
	// key := "Bond"
	// ctx := context.WithValue(context.Background(), key, "James")
	// ctx = context.WithValue(ctx, key, "007")
	// fmt.Println(ctx.Value(key))
	// nums := []int{2,3,4,-3,-4,5,8,15,10,6,7}
	// target := 10
	// fmt.Print(algo.SumOfThree(nums, target))
	// fmt.Print(algo.MissNumber([]int{1, 2, 4, 5, 6}))

	// s := algo.NewStack()
	// s.Staked(5)
	// s.Staked(7)
	// s.UnStaked()
	// s.Staked(8)
	// s.Display()

	q := algo.NewQueue()
	q.Queued(5)
	q.Queued(7)
	q.DeQueued()
	q.Display()
}
