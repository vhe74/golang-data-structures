package main

import "fmt"

type Item struct {
	data string
}

type Queue struct {
	items []Item
}

func (q *Queue) Enqueue(i Item) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() Item {
	if q.Size() != 0 {
		toRemove := q.items[0]
		q.items = q.items[1:]
		return toRemove
	} else {
		return Item{}
	}
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Display() {
	print("[")
	for _, v := range q.items {
		print(v.data)
		print(" ")
	}
	println("]")
}

func main() {
	q := Queue{}
	q.Enqueue(Item{"Hello"})
	q.Enqueue(Item{"World"})
	q.Display()
	fmt.Println("Dequeuing ", q.Dequeue())
	q.Display()
	fmt.Println("Dequeuing ", q.Dequeue())
	q.Display()
	fmt.Println("Dequeuing ", q.Dequeue())
	q.Display()
	fmt.Println("Dequeuing ", q.Dequeue())
}
