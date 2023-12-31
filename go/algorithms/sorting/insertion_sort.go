// insertion_sort.go

package main

import "fmt"

func insertionSort(arr []int) {
    n := len(arr)

    for i := 1; i < n; i++ {
        key := arr[i]
        j := i - 1

        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j--
        }
        
        arr[j+1] = key
    }
}

func main() {
    array := []int{64, 25, 12, 22, 11}
    
    fmt.Println("Unsorted array:", array)

    insertionSort(array)

    fmt.Println("Sorted array:", array)
}
