/* Exercise: Readers
 * See https://go.dev/tour/methods/22
 * Description: Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
 */

package main

import "golang.org/x/tour/reader"

type MyReader struct {
    c   byte    // an ASCII character
}

func (r *MyReader) Read(b []byte) (n int, err error) {
    for i := range b {
        b[i] = r.c
    }

    return len(b), nil // never return io.EOF
}

func main() {
    reader.Validate(&MyReader{'A'})
    reader.Validate(&MyReader{'B'}) // got byte 42 at offset 0, want 'A'
}

