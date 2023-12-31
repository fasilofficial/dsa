// factorial.go

package main

import "fmt"

func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    n := 5
    result := factorial(n)

    fmt.Printf("Factorial of %d is: %d\n", n, result)
}
