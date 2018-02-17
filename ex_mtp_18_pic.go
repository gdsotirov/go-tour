/* Exercise: Slices
 * See https://tour.golang.org/moretypes/18
 * Description: Create a slice of slices (i.e. a two-dimensional array),
 *              inialize it's values with a function and generate picture
 */

package main

import (
    //"math" /* uncomment when necessary */
    "golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
    /* create two dimensional array */    
    matrix := make([][]uint8, dy)
    for i := 0; i < dy; i++ {
        matrix[i] = make([]uint8, dx)
    }

    /* initialize values */
    for i := 0; i < dx; i++ {
        for j := 0; j < dy; j++ {
            matrix[i][j] = uint8(i * j)
                     /* or uint8((i + j)/2) or */
                     /* or uint8(math.Pow(float64(i+j), 2)) */
        }
    }

    return matrix
}

func main() {
    pic.Show(Pic)
}

