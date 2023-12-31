// min_heap.go

package main

import "fmt"

type MinHeap struct {
    heap []int
}

func NewMinHeap() *MinHeap {
    return &MinHeap{}
}

func (h *MinHeap) Insert(value int) {
    h.heap = append(h.heap, value)
    h.heapifyUp(len(h.heap) - 1)
}

func (h *MinHeap) ExtractMin() (int, error) {
    if len(h.heap) == 0 {
        return 0, fmt.Errorf("heap is empty")
    }

    min := h.heap[0]
    lastIdx := len(h.heap) - 1

    h.heap[0], h.heap[lastIdx] = h.heap[lastIdx], h.heap[0]
    h.heap = h.heap[:lastIdx]

    h.heapifyDown(0)

    return min, nil
}

func (h *MinHeap) heapifyUp(index int) {
    parent := (index - 1) / 2

    for index > 0 && h.heap[index] < h.heap[parent] {
        h.heap[index], h.heap[parent] = h.heap[parent], h.heap[index]
        index = parent
        parent = (index - 1) / 2
    }
}

func (h *MinHeap) heapifyDown(index int) {
    for {
        leftChild := 2*index + 1
        rightChild := 2*index + 2
        smallest := index

        if leftChild < len(h.heap) && h.heap[leftChild] < h.heap[smallest] {
            smallest = leftChild
        }

        if rightChild < len(h.heap) && h.heap[rightChild] < h.heap[smallest] {
            smallest = rightChild
        }

        if smallest != index {
            h.heap[index], h.heap[smallest] = h.heap[smallest], h.heap[index]
            index = smallest
        } else {
            break
        }
    }
}

func main() {
    heap := NewMinHeap()

    heap.Insert(4)
    heap.Insert(10)
    heap.Insert(3)
    heap.Insert(5)
    heap.Insert(1)

    fmt.Println("Min Heap:", heap.heap)

    min, err := heap.ExtractMin()
    if err == nil {
        fmt.Printf("Extracted Min Element: %d\n", min)
        fmt.Println("Updated Min Heap:", heap.heap)
    } else {
        fmt.Println(err)
    }
}
