// array.go

package main

import "fmt"

func main() {
    array := []int{1, 2, 3, 4, 5}

    fmt.Println("Original Array:", array)

    array = append(array, 6)
    fmt.Println("After Insertion 1 (at the end):", array)

    position := 2
    value := 10
    array = insertAtPosition(array, position, value)
    fmt.Printf("After Insertion 2 (at position %d): %v\n", position, array)

    value = 100
    array = insertAtBeginning(array, value)
    fmt.Println("After Insertion 3 (at the beginning):", array)

    array = deleteFromEnd(array)
    fmt.Println("After Deletion 1 (from the end):", array)

    position = 1
    array = deleteFromPosition(array, position)
    fmt.Printf("After Deletion 2 (from position %d): %v\n", position, array)

    array = deleteFromBeginning(array)
    fmt.Println("After Deletion 3 (from the beginning):", array)
}

func insertAtPosition(arr []int, position, value int) []int {
    if position < 0 || position > len(arr) {
        return arr
    }

    arr = append(arr[:position], append([]int{value}, arr[position:]...)...)
    return arr
}

func insertAtBeginning(arr []int, value int) []int {
    arr = append([]int{value}, arr...)
    return arr
}

func deleteFromEnd(arr []int) []int {
    if len(arr) == 0 {
        return arr
    }

    arr = arr[:len(arr)-1]
    return arr
}

func deleteFromPosition(arr []int, position int) []int {
    if position < 0 || position >= len(arr) {
        return arr
    }

    arr = append(arr[:position], arr[position+1:]...)
    return arr
}

func deleteFromBeginning(arr []int) []int {
    if len(arr) == 0 {
        return arr
    }

    arr = arr[1:]
    return arr
}
