package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  z0 := 6 * math.Pow(10, 2)
  z := z0
    var old_z float64 = 0.0

  for math.Abs(old_z - z) > 0.0005 {
        old_z = z
    z = z - ((z*z - x) / (2 * z))
  }

  return z
}

func main() {
  fmt.Println(Sqrt(2))
}
