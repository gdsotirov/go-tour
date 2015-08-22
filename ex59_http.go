/* Exercise: Web servers
 * See http://web.archive.org/web/20140807003324/http://tour.golang.org/#59
 * Description: Simple web server that prints a greeting
 */

package main

import (
    "fmt"
    "net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello!")
}

func main() {
    var h Hello
    http.ListenAndServe("localhost:4000", h)
}

