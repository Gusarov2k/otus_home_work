// Реализовать тип IntStack , который
// содержит стэк целых чисел. У него
// должны быть методы Push(i int) и
// Pop() int
package main

import (
	"fmt"
)

type Stack struct {
	Ids      []int
	CountPop int
}

func (s *Stack) Push(val int) {
	s.Ids = append(s.Ids, val)
	s.CountPop++
}

func (s *Stack) Pop() int {
	result := s.Ids[s.CountPop-1]
	s.CountPop--
	return result
}

func main() {
	s := Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Printf("expected 30, got %d\n", s.Pop())
	fmt.Printf("expected 20, got %d\n", s.Pop())
	fmt.Printf("expected 10, got %d\n", s.Pop())
}
