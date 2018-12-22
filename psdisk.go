package main

import (
	"fmt"
	"time"
	"github.com/shirou/gopsutil/disk"
)

func main() {

	for {
		f, _ := disk.IOCounters()
		fmt.Printf("DISK disk1 IO ops R/W %v %v BYTES R/W %v %v\n", f["disk1"].ReadCount, f["disk1"].WriteCount, f["disk1"].ReadBytes, f["disk1"].WriteBytes)
		//for ioname, counter := range f {
		//	fmt.Printf("DISK %s IO ops R/W %v %v BYTES R/W %v %v\n", ioname, counter)
		//}
		time.Sleep(500* time.Millisecond)
	}
}
