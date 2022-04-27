/* Exercise: Equivalent Binary Trees
 * See https://go.dev/tour/concurrency/8
 * Description: Walk binary tree and compare two trees for equivalance
 */

package main

import (
    "golang.org/x/tour/tree"
    "fmt"
)

// Print tree to standard output
func Print(t *tree.Tree) {
    ch := make(chan int)
    go Walk(t, ch)

    for i := range ch {
        fmt.Printf("%v ", i)
    }
    fmt.Print("\n")
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    var wr func(t *tree.Tree)

    defer close(ch)
    wr = func(t *tree.Tree) {
        if t.Left != nil {
            wr(t.Left)
        }
        if t != nil {
            ch <- t.Value
        } else {
            return
        }
        if t.Right != nil {
            wr(t.Right)
        }
    }
    wr(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go Walk(t1, ch1)
    go Walk(t2, ch2)

    for {
        v1,ok1 := <- ch1
        v2,ok2 := <- ch2

        if v1 != v2 || ok1 != ok2 {
            return false
        }
        if !ok1 {
            break
        }
    }

    return true
}

func main() {
    tr1 := tree.New(1)
    tr2 := tree.New(2)

    fmt.Print("Random tree 1: ")
    Print(tr1)
    fmt.Print("Random tree 2: ")
    Print(tr2)

    fmt.Println("First  comparision (tree 1 and 1): ", Same(tr1, tr1))
    fmt.Println("Second comparision (tree 1 and 2): ", Same(tr1, tr2))
}

