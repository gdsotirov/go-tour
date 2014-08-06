/* Exercise: Maps
 * See http://tour.golang.org/#43
 * Description: Implement a word count function with a map
 */

package main

import (
    "code.google.com/p/go-tour/wc"
    "fmt"
    "strings"
)

func WordCount(s string) map[string]int {
    wc := make(map[string]int)
    words := strings.Fields(s)

    for i := 0; i < len(words); i++ {
      wc[words[i]]++
    }

    fmt.Println("Map is ", wc);
    return wc
}

func main() {
    wc.Test(WordCount)
}
