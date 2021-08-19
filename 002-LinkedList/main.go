package main

import "fmt"

type node struct {
	data string
	next *node
}

type linkedList struct {
	length int
	head   *node
	tail   *node
}

func (l linkedList) Len() int {
	return l.length
}

func (l linkedList) Display() {
	for l.head != nil {
		fmt.Printf("%v -> ", l.head.data)
		l.head = l.head.next
	}
	fmt.Println()
}

func (l *linkedList) PushBack(n *node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {
		l.tail.next = n
		l.tail = n
		l.length++
	}
}

func main() {
	list := linkedList{}

	n := &node{data: "Vincent"}
	list.PushBack(n)
	list.Display()

	n = &node{data: "Valentine"}
	list.PushBack(n)
	list.Display()
}
