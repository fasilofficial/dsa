// heap_sort.go

package main

import "fmt"

func heapSort(arr []int) {
    n := len(arr)

    buildHeap(arr, n)

    for i := n - 1; i > 0; i-- {
        arr[0], arr[i] = arr[i], arr[0]

        heapify(arr, i, 0)
    }
}

func buildHeap(arr []int, n int) {
    for i := n/2 - 1; i >= 0; i-- {
        heapify(arr, n, i)
    }
}

func heapify(arr []int, n, i int) {
    largest := i
    left := 2*i + 1
    right := 2*i + 2

    if left < n && arr[left] > arr[largest] {
        largest = left
    }

    if right < n && arr[right] > arr[largest] {
        largest = right
    }

    if largest != i {
        arr[i], arr[largest] = arr[largest], arr[i]
        heapify(arr, n, largest)
    }
}

func main() {
    array := []int{64, 25, 12, 22, 11}

    fmt.Println("Unsorted array:", array)

    heapSort(array)

    fmt.Println("Sorted array:", array)
}