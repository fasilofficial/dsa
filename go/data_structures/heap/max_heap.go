// max_heap.go

package main

import "fmt"

type MaxHeap struct {
    heap []int
}

func NewMaxHeap() *MaxHeap {
    return &MaxHeap{}
}

func (h *MaxHeap) Insert(value int) {
    h.heap = append(h.heap, value)
    h.heapifyUp(len(h.heap) - 1)
}

func (h *MaxHeap) ExtractMax() (int, error) {
    if len(h.heap) == 0 {
        return 0, fmt.Errorf("heap is empty")
    }

    max := h.heap[0]
    lastIdx := len(h.heap) - 1

    h.heap[0], h.heap[lastIdx] = h.heap[lastIdx], h.heap[0]
    h.heap = h.heap[:lastIdx]

    h.heapifyDown(0)

    return max, nil
}

func (h *MaxHeap) heapifyUp(index int) {
    parent := (index - 1) / 2

    for index > 0 && h.heap[index] > h.heap[parent] {
        h.heap[index], h.heap[parent] = h.heap[parent], h.heap[index]
        index = parent
        parent = (index - 1) / 2
    }
}

func (h *MaxHeap) heapifyDown(index int) {
    for {
        leftChild := 2*index + 1
        rightChild := 2*index + 2
        largest := index

        if leftChild < len(h.heap) && h.heap[leftChild] > h.heap[largest] {
            largest = leftChild
        }

        if rightChild < len(h.heap) && h.heap[rightChild] > h.heap[largest] {
            largest = rightChild
        }

        if largest != index {
            h.heap[index], h.heap[largest] = h.heap[largest], h.heap[index]
            index = largest
        } else {
            break
        }
    }
}

func main() {
    heap := NewMaxHeap()

    heap.Insert(4)
    heap.Insert(10)
    heap.Insert(3)
    heap.Insert(5)
    heap.Insert(1)

    fmt.Println("Max Heap:", heap.heap)

    max, err := heap.ExtractMax()
    if err == nil {
        fmt.Printf("Extracted Max Element: %d\n", max)
        fmt.Println("Updated Max Heap:", heap.heap)
    } else {
        fmt.Println(err)
    }
}
