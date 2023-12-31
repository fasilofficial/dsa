// selection_sort.go

package main

import "fmt"

func selectionSort(arr []int) {
    n := len(arr)

    for i := 0; i < n-1; i++ {
        minIndex := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIndex] {
                minIndex = j
            }
        }

        arr[i], arr[minIndex] = arr[minIndex], arr[i]
    }
}

func main() {
    array := []int{64, 25, 12, 22, 11}
    
    fmt.Println("Unsorted array:", array)

    selectionSort(array)

    fmt.Println("Sorted array:", array)
}
