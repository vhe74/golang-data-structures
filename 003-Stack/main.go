package main

import "fmt"

type Item struct {
	data string
}

type Stack struct {
	items []Item
}

// Push will add a value at the end (top)
func (s *Stack) Push(i Item) {
	s.items = append(s.items, i)
}

// Pop will remove the value at the end (top)
func (s *Stack) Pop() Item {
	if s.Size() != 0 {
		lastIndex := len(s.items) - 1
		toRemove := s.items[lastIndex]
		s.items = s.items[:lastIndex]
		return toRemove
	} else {
		return Item{}
	}
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Display() {
	print("[")
	for _, v := range s.items {
		print(v.data)
		print(" ")
	}
	println("]")
}

func main() {
	s := Stack{}
	s.Push(Item{"Hello"})
	s.Push(Item{"World"})
	s.Push(Item{"!"})
	s.Display()
	fmt.Println("Stack size is ", s.Size())
	fmt.Println("Pop ", s.Pop())
	s.Display()
	fmt.Println("Pop ", s.Pop())
	s.Display()
	fmt.Println("Pop ", s.Pop())
	s.Display()
	fmt.Println("Pop ", s.Pop())
	s.Display()
	fmt.Println("Pop ", s.Pop())
	s.Display()
}
