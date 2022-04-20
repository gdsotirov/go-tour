/* Exercise: Errors
 * See https://go.dev/tour/methods/20
 * Description: Make Sqrt function from previous exercise return error when given a negative number argument
 */

package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    msg := fmt.Sprintf("Error: cannot calcualte square root of a negative number: %v\n", float64(e))
    return msg
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, ErrNegativeSqrt(x)
    }
    z0 := 6 * math.Pow(10, 2)
    z := z0
    var old_z float64 = 0.0

    for math.Abs(old_z-z) > 0.0005 {
        old_z = z
        z = z - ((z*z - x) / (2 * z))
    }

    return z, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
