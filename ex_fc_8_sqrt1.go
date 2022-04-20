/* Exercise: Loops and Functions 
 * See https://go.dev/tour/flowcontrol/8
 * Description: Calculate square root with Newton's method in 10 iterations
 */

package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  z0 := 6 * math.Pow(10, 2)
  z := z0

  for i := 0; i < 10; i++ {
    z = z - ( (z * z - x) / (2 * z) )
  }

  return z
}

func main() {
  fmt.Println("Sqrt(2)      = ", Sqrt(2))
  fmt.Println("math.Sqrt(2) = ", math.Sqrt(2))
}

