// queue.go

package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Display() {
	fmt.Println("Queue:", q.items)
}

func main() {
	queue := &Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println("Queue Size:", queue.Size())

	item, err := queue.Dequeue()
	if err == nil {
		fmt.Println("Dequeued Item:", item)
	} else {
		fmt.Println(err)
	}

	fmt.Println("Is Queue Empty:", queue.IsEmpty())
	queue.Display()
}
