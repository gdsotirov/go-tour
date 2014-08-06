/* Exercise: Fibonacci closure
 * See http://tour.golang.org/#46
 * Description: Implement a Fibonacci numbers closure
 */

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    prev1 := 0
    prev2 := 0
    next  := 0
    return func() int {
        next = prev1 + prev2
        if prev1 == 0 {
            prev1++
        } else {
            prev1 = prev2
        }
        prev2 = next
        return next
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 42; i++ {
        fmt.Printf("n%d = %v\n", i, f())
    }
}
