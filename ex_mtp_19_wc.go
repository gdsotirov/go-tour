/* Exercise: Maps
 * See http://tour.golang.org/moretypes/19
 * Description: Implement a word count function with a map
 */

package main

import (
    "golang.org/x/tour/wc"
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
