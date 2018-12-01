package main

import (
	"fmt"
	"time"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
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

	h, _ := host.Info()
	j, _ := host.Users()

	k, _ := net.Interfaces()
	l, _ := net.IOCounters(true) 
	m, _ := net.IOCounters(false) 
	n, _ := net.Connections("all") 
	
	p, _ := process.Processes() 
        for _, proc := range p {
                q, _ := proc.Connections()
		if len(q) == 0 {continue}
		pcmd, _ :=  proc.Cmdline()
                fmt.Printf("Process %v %v connections %+v\n", proc, pcmd, q)
        }


	
	t1 := time.Now().UnixNano()

	// almost every return value is a struct
	fmt.Printf("CPU info: %+v\n\n", c)
	fmt.Printf("%dns Total: %v, Free:%v, UsedPercent:%f%%\n\n", t1-t0, v.Total, v.Free, v.UsedPercent)
	fmt.Printf("CPU Times info: %+v\n\n", d)
	fmt.Printf("CPU Percent info: %+v\n\n", e)
	fmt.Printf("DISK Partitions: %+v\n\n", f)
	fmt.Printf("Host Info %+v %+v\n\n", h, j)
	fmt.Printf("NET Interfaces %+v \n\n\n%+v\n\n%+v\n\n\n\n", k, l, m)
	fmt.Printf("NET Connections %+v\n\n\n", n)
	fmt.Printf("PROCESSes: %+v\n\n\n", len(p))
	// convert to JSON. String() is also implemented

	fmt.Println(v)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(h)
}
