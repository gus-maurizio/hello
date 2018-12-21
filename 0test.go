package main

import (
	"fmt"
	"github.com/gus-maurizio/structures/duplexqueue"
)

func main() {
	var qi duplexqueue.Duplexqueue
	qi.Init(10,0)
        qi.Do( func(m interface{}) { fmt.Printf("%v, ",m) })
	fmt.Printf("%v\n",qi)
	fmt.Printf("\n") 
	for i := 0; i < 25; i++ {
		old := qi.PushPop(10+i)
		fmt.Printf("iter %3d old: %3v --> ",i,old) 
        	qi.Do( func(m interface{}) { fmt.Printf("%2v, ",m) })
		//fmt.Printf("\n") 
		fmt.Printf(" 1: %3v 5: %3v 10: %3v buff:  %v\n",qi.Index(1-1), qi.Index(5-1), qi.Index(10-1), qi)
	}

        fmt.Printf("-- DoFor ->")
        qi.DoFor( 0, 4, func(m interface{}) {
                fmt.Printf(" %v",m)
                })
        fmt.Printf(" \n")

	fmt.Printf("%v\n",qi)
	fmt.Printf("%v\n",qi.Slice())
	fmt.Printf("%v\n",qi.Index(3))
	fmt.Printf("%v\n",qi.Index(-2))
	fmt.Printf("%v\n",qi.Index(13))
}
