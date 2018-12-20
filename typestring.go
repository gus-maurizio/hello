package main

import (
	"encoding/json"
	"fmt"
)

// Duplexqueue represents a single instance of the duplexqueue data structure.
type MyType struct {
	Buf   []interface{}
	Head  int
	Tail  int
	Count int
}

func main() {

	q := MyType{	Buf:  nil,
			Head: 4,
			Tail: 1,
			Count: 0,
		}
	q.Buf = []int{4,5,6}
	fmt.Println(q) 
	vjson, _ := json.MarshalIndent(q, "", "\t")
	fmt.Println(string(vjson)) 
}
