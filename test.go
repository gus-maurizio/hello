package main

import (
	"fmt"
	"github.com/gus-maurizio/structures/duplexqueue"
)

func main() {
	var q duplexqueue.Duplexqueue
	var qi duplexqueue.Duplexqueue
	q.PushBack("foo")
	q.PushBack("bar")
	q.PushBack("baz")

	fmt.Println(q.Len())   // Prints: 3
	fmt.Println(q.Front()) // Prints: foo
	fmt.Println(q.Back())  // Prints: baz

	q.PopFront() // remove "foo"
	q.PopBack()  // remove "baz"

	q.PushFront("hello")
	q.PushBack("world")

	// Consume duplexqueue and print elements.
	for q.Len() != 0 {
		fmt.Println(q.PopFront())
	}

	qi.PushBack(1)
	qi.PushBack(2)
	qi.PushBack(3)
	qi.PushBack(4)
	qi.PushBack(5)
	qi.Do( func(m interface{}) {
		fmt.Printf("%v\n",m)
		})

	fmt.Printf("------ %v\n",qi.Index(0))

        qi.DoIndex( -2, func(m interface{}) {
                fmt.Printf("%v\n",m)
                })

	qi.Init(10,[2]int{3,4})
        qi.Do( func(m interface{}) { fmt.Printf("%2d, ",m) })

	fmt.Printf("\n") 
	for i := 0; i < 25; i++ {
		old := qi.PushPop(10+i)
		fmt.Printf("iter %3d old: %v --> ",i,old) 
        	qi.Do( func(m interface{}) { fmt.Printf("%2d, ",m) })
		fmt.Printf("\n") 
	}

}
