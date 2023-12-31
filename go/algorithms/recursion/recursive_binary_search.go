// binary_search.go

package main

import "fmt"

func binarySearch(arr []int, target int, low, high int) int {
    if low > high {
        return -1
    }

    mid := (low + high) / 2
    midValue := arr[mid]

    if midValue == target {
        return mid
    } else if midValue > target {
        return binarySearch(arr, target, low, mid-1)
    } else {
        return binarySearch(arr, target, mid+1, high)
    }
}

func main() {
    sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    target := 7

    result := binarySearch(sortedArray, target, 0, len(sortedArray)-1)

    if result != -1 {
        fmt.Printf("Element %d found at index %d\n", target, result)
    } else {
        fmt.Printf("Element %d not found in the array\n", target)
    }
}
