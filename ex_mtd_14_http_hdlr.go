/* Exercise: HTTP Handlers
 * See http://web.archive.org/web/20140807003324/http://tour.golang.org/#60
 * Description: Web server handling specific paths
 */

package main

import (
    "fmt"
    "net/http"
)

type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

func (str String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, str)
}

func (sct Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, sct.Greeting, sct.Punct, sct.Who)
}

func main() {
    http.Handle("/string", String("I'm a frayed knot."))
    http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
    http.ListenAndServe("localhost:4000", nil)
}

