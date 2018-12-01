package main

import (
	"fmt"
	"time"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/net"
)

func main() {

	cpu.Percent(0, true)
	c, _ := cpu.Info()
	f, _ := disk.Partitions(true)

	t0 := time.Now().UnixNano()
	v, _ := mem.VirtualMemory()
	d, _ := cpu.Times(true)
	e, _ := cpu.Percent(0, true)
	for _, part := range f {
		g, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("DISK Usage: %+v\n", g)
	}

	t1 := time.Now().UnixNano()

	// almost every return value is a struct
	fmt.Printf("CPU info: %+v\n\n", c)
	fmt.Printf("%dns Total: %v, Free:%v, UsedPercent:%f%%\n\n", t1-t0, v.Total, v.Free, v.UsedPercent)
	fmt.Printf("CPU Times info: %+v\n\n", d)
	fmt.Printf("CPU Percent info: %+v\n\n", e)
	fmt.Printf("DISK Partitions: %+v\n\n", f)
	// convert to JSON. String() is also implemented

	fmt.Println(v)
	fmt.Println(d)
	fmt.Println(e)
}
