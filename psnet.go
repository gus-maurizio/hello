package main

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
	"time"
)

var PluginData,
    PluginDataPrev	map[string]interface{}
var TScurrent,
    TSprevious 		int64


func main() {

	TSprevious = time.Now().UnixNano()
	PluginData     = make(map[string]interface{}, 20)
	PluginDataPrev = make(map[string]interface{}, 20)

	k, _ := net.Interfaces()
	fmt.Printf("NET Interfaces %+v \n", k)

	netio, _ := net.IOCounters(true)
	for netidx := range netio {
		PluginDataPrev[netio[netidx].Name] = netio[netidx]
	}

	for {
		time.Sleep(500 * time.Millisecond)
	
		TScurrent = time.Now().UnixNano()
       	 netio, _ = net.IOCounters(true)
       	 for netidx := range netio { PluginData[netio[netidx].Name] = netio[netidx] }
	
		Δts := TScurrent - TSprevious	// nanoseconds!

		NETS:
		for netid, _ := range PluginData {
			_, present := PluginDataPrev[netid]
			if !present {continue NETS}
			inc_precv		:= PluginData[netid].(net.IOCountersStat).PacketsRecv  - PluginDataPrev[netid].(net.IOCountersStat).PacketsRecv
			inc_psent		:= PluginData[netid].(net.IOCountersStat).PacketsSent  - PluginDataPrev[netid].(net.IOCountersStat).PacketsSent
			inc_brecv		:= PluginData[netid].(net.IOCountersStat).BytesRecv    - PluginDataPrev[netid].(net.IOCountersStat).BytesRecv
			inc_bsent		:= PluginData[netid].(net.IOCountersStat).BytesSent    - PluginDataPrev[netid].(net.IOCountersStat).BytesSent
	
			ppsrecv := float64(inc_precv) * 1e9 / float64(Δts)
			ppssent := float64(inc_psent) * 1e9 / float64(Δts) 
	
			bpsrecv := float64(8 * inc_brecv) * 1e9 / (float64(Δts) * 1024.0 * 1024.0)
			bpssent := float64(8 * inc_bsent) * 1e9 / (float64(Δts) * 1024.0 * 1024.0)
			if netid == "en0" {
				fmt.Printf("--en0--> pps %7.3f/%7.3f   mbps  %7.3f/%7.3f  B %v \n", ppsrecv, ppssent, bpsrecv, bpssent, PluginData[netid].(net.IOCountersStat).BytesRecv)
			}
		}
		// save current values as previous
		for netid, _ := range PluginData {
			_, present := PluginDataPrev[netid]
			if present { PluginDataPrev[netid] = PluginData[netid] }
		}
		TSprevious = TScurrent
	}

}
