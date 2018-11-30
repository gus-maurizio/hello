package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)
	type Animal struct {
		Name  string
		Order string
	}
	os.Stdout.Write(jsonBlob)
	fmt.Println()
	var animals interface{}
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("unmarshal error:", err)
	}
	fmt.Printf("%T\n", animals)
	fmt.Printf("%+v\n", animals)
        b, merr := json.Marshal(animals)
        if merr != nil {
                fmt.Println("marshal error:", merr)
        }
        os.Stdout.Write(b)
	fmt.Println()
}
