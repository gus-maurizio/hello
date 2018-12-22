package main

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
)

func main() {

	f, _ := disk.Partitions(false)
	fmt.Printf("DISK Partitions: %+v\n\n", f)

	for _, part := range f {
		g, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("DISK Usage: %+v\n", g)
	}

}
