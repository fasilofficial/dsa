// linear_search.go

package main

import "fmt"

func linearSearch(arr []int, target int) int {
    for i, value := range arr {
        if value == target {
            return i
        }
    }

    return -1
}

func main() {
    array := []int{10, 4, 6, 8, 2, 7, 1}
    target := 6

    result := linearSearch(array, target)

    if result != -1 {
        fmt.Printf("Element %d found at index %d\n", target, result)
    } else {
        fmt.Printf("Element %d not found in the array\n", target)
    }
}
