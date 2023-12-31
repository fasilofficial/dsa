// stack.go

package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Display() {
	fmt.Println("Stack:", s.items)
}

func main() {
	stack := &Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Stack Size:", stack.Size())

	item, err := stack.Pop()
	if err == nil {
		fmt.Println("Popped Item:", item)
	} else {
		fmt.Println(err)
	}

	fmt.Println("Is Stack Empty:", stack.IsEmpty())
	stack.Display()
}