package main

import "code.google.com/p/go-tour/pic"

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
        }
    }

    return matrix
}

func main() {
    pic.Show(Pic)
}