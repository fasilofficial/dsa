// factorial.go

package main

import "fmt"

func factorial(n int) int {
    result := 1

    for i := 1; i <= n; i++ {
        result *= i
    }

    return result
}

func main() {
    n := 5
    result := factorial(n)
    fmt.Printf("Factorial of %d is: %d\n", n, result)
}