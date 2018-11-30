package main

import (
	"fmt"
	"net"
)

func main() {
	ifaces, _ := net.Interfaces()
	for idx, i := range ifaces {
		if i.Flags&net.FlagLoopback     != 0 { continue }
		if i.Flags&net.FlagPointToPoint != 0 { continue }
		addrs, _ := i.Addrs()
		if len(addrs) == 0 { continue }
		fmt.Printf("idx %d %s %d %v ", idx, i.Name, i.MTU, i.HardwareAddr)
		for  _, addr := range addrs {
			var ip net.IP
	                switch v := addr.(type) {
        	        case *net.IPNet:
                	        ip = v.IP
                	case *net.IPAddr:
                        	ip = v.IP
                	}
			fmt.Printf("%#v ",ip.String())
		}
		fmt.Printf("\n")

	}
}
