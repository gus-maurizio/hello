package main

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
)

func main() {


	k, _ := net.Interfaces()
	n, _ := net.Connections("all") 

	var TCPStatuses = map[string]string{
	    "01": "ESTABLISHED",
	    "02": "SYN_SENT",
	    "03": "SYN_RECV",
	    "04": "FIN_WAIT1",
	    "05": "FIN_WAIT2",
	    "06": "TIME_WAIT",
	    "07": "CLOSE",
	    "08": "CLOSE_WAIT",
	    "09": "LAST_ACK",
	    "0A": "LISTEN",
	    "0B": "CLOSING",
	}
	
	connCounter := make(map[string]int,len(TCPStatuses))
	for _, status := range TCPStatuses { connCounter[status] = 0 }
	for _, connStat := range n         { connCounter[connStat.Status] += 1 }
	// almost every return value is a struct
	fmt.Printf("NET Interfaces %+v \n\n%v \n", k, connCounter)
	fmt.Printf("NET Conn %+v \n", n)
	for _, c := range n {
		fmt.Printf("--> <%s> %+v\n",c.Status,c)
	}
	for status, count   := range connCounter { 
		if status == "" { continue }
		fmt.Printf("status %s: %d\n",status,count)
	}

}
