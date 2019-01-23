package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func Encode(msg string) string {
	str := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Printf("%q\n", str)
	return str
}

func Decode(msg string) string {
	data, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	fmt.Printf("%q\n", data)
	return string(data)
}

func main() {
	// msg := "Hello, 世界"
	encoded := Encode(os.Args[1]) 
	fmt.Println(encoded)
	decoded := Decode(encoded)
	fmt.Println(string(decoded))
}
