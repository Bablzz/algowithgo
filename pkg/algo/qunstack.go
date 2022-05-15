package algo

import (
	"fmt"
)

type Stack struct {
	elem []int
}

func NewStack() Stack {
	return Stack{elem: make([]int, 0)}
}

func (s *Stack) Staked(e int) {
	s.elem = append(s.elem, e)
}

func (s *Stack) UnStaked() int {
	temp := s.elem[len(s.elem)-1]
	s.elem = s.elem[:len(s.elem)-1]
	return temp
}

func (s *Stack) Display() {
	for _, v := range s.elem {
		fmt.Printf("%d <- ", v)
	}
}

type Queue struct {
	elem []int
}

func NewQueue() Queue {
	return Queue{elem: make([]int, 0)}
}

func (q *Queue) Queued(e int) {
	q.elem = append(q.elem, e)
}

func (q *Queue) DeQueued() int {
	temp := q.elem[0]
	q.elem = q.elem[1:]
	return temp
}

func (q *Queue) Display() {
	for _, v := range q.elem {
		fmt.Printf("%d -> ", v)
	}
}
