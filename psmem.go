package main

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

func main() {


	v, _ := mem.VirtualMemory()


	fmt.Printf("mem info: %+v\n\n", v)

	fmt.Println(v)
}
