/* Exercise: Advanced Exercise: Complex cube roots
 * See http://tour.golang.org/#50
 * Description: Calcualte cube root by Newton's method
 */

package main

import (
    "fmt"
    "math"
    "math/cmplx"
)

func Cbrt(x complex128) complex128 {
    z0 := complex(6 * math.Pow(10, 3), 0)
    z  := z0
    var old_z complex128 = 0 + 0i

    for cmplx.Abs(old_z - z) > 0.0005 {
        old_z = z
        z = z - ((cmplx.Pow(z, 3) - x) / (3 * cmplx.Pow(z, 2)))
    }

    return z
}

func main() {
    fmt.Println(Cbrt(2))
}
