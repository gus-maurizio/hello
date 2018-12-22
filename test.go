package main

import (
	"fmt"
	"github.com/gus-maurizio/structures/duplexqueue"
)

func main() {
	var qi duplexqueue.Duplexqueue
	qi.Init(11,"----")
        qi.Do( func(m interface{}) { fmt.Printf("%v, ",m) })
	fmt.Printf("%v\n",qi)
	fmt.Printf("\n") 
	for i := 0; i < 25; i++ {
		old := qi.PushPop(fmt.Sprintf("<%d, %v>",i,10+i))
		fmt.Printf("iter %3d old: %3v --> ",i,old) 
        	qi.Do( func(m interface{}) { fmt.Printf("%v, ",m) })
		fmt.Printf("\n") 
		// fmt.Printf("%v\n",qi)
	}

        fmt.Printf("-- DoFor ->")
        qi.DoFor( 0, 4, func(m interface{}) {
                fmt.Printf(" %v",m)
                })
        fmt.Printf(" \n")

	fmt.Printf("%v\n",qi)
	fmt.Printf("%v\n",qi.Slice())
}
