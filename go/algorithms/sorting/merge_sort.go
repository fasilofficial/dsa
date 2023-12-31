// merge_sort.go

package main

import "fmt"

func merge(left, right []int) []int {
    result := make([]int, 0, len(left)+len(right))

    for len(left) > 0 || len(right) > 0 {
        if len(left) == 0 {
            return append(result, right...)
        }
        if len(right) == 0 {
            return append(result, left...)
        }

        if left[0] <= right[0] {
            result = append(result, left[0])
            left = left[1:]
        } else {
            result = append(result, right[0])
            right = right[1:]
        }
    }

    return result
}

func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    mid := len(arr) / 2
    left := mergeSort(arr[:mid])
    right := mergeSort(arr[mid:])

    return merge(left, right)
}

func main() {
    array := []int{64, 25, 12, 22, 11}
    
    fmt.Println("Unsorted array:", array)

    array = mergeSort(array)

    fmt.Println("Sorted array:", array)
}
