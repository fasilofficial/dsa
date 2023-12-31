// quick_sort.go

package main

import "fmt"

func quickSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    pivot := arr[0]
    var left, right []int

    for _, num := range arr[1:] {
        if num <= pivot {
            left = append(left, num)
        } else {
            right = append(right, num)
        }
    }

    left = quickSort(left)
    right = quickSort(right)

    return append(append(left, pivot), right...)
}

func main() {
    array := []int{64, 25, 12, 22, 11}
    
    fmt.Println("Unsorted array:", array)
	
    array = quickSort(array)

    fmt.Println("Sorted array:", array)
}
