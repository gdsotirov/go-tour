/* Exercise: Generic types
 * See https://go.dev/tour/generics/2
 * Description: Add some functionality to a singly-linked list holding any type of value
 */

package main

import "fmt"

type ListItem[T any] struct {
    val  T
    next *ListItem[T]
}

type List[T any] struct {
    count uint64
    head *ListItem[T]
}

func (l *List[T]) Add(item T) {
    if l.head == nil {
        l.head = &ListItem[T]{item, nil}
    } else {
        // Move to the end of list
        ptr := l.head
        for ; ptr.next != nil; ptr = ptr.next {
        }
        ptr.next = &ListItem[T]{item, nil}
    }
    l.count++
}

func (l *List[T]) Print() {
    for i := l.head; i != nil; i = i.next {
        fmt.Println(i.val)
    }
}

func main() {
    var nums List[int]
    var names List[string]

    // Populate list of numbers
    for i := 1; i < 10; i++ {
        nums.Add(i)
    }

    // Populate list of names
    names.Add("Georgi")
    names.Add("Dimitar")
    names.Add("Ivan")

    fmt.Printf("List with %d numbers:\n", nums.count)
    nums.Print()

    fmt.Printf("List with %d names:\n", names.count)
    names.Print()
}

