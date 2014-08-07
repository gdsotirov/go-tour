/* Exercise: Rot13 Reader
 * See http://tour.golang.org/#63
 * Description: Implement ROT13 reader on top of io.Reader
 */

package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r13 rot13Reader) Read(p []byte) (n int, err error) {
    s, err := r13.r.Read(p)
    if err != nil {
        return s, err
    }
    var cnt int = 0

    for i := 0; i < len(p); i++ {
        p[i] = rot13(p[i])
        cnt++
    }

    return cnt, nil
}

func rot13(b byte) byte {
    var sub byte
    /* apply conversion only on alphabetical characters */
    if b >= 64 && b <= 90 {
        sub = b + 13
        if sub > 90 {
            sub -= 26 /* big English letters in the ASCII table */
        }
    } else {
        if b >= 97 && b <= 122 {
            sub = b + 13
            if sub > 122 {
                sub -= 26 /* big English letters in the ASCII table */
            }
        } else {
            sub = b
        }
    }

    return sub
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
