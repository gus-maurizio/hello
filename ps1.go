package main

import _ "net/http/pprof"

import (
        "encoding/json"
        "fmt"

        "github.com/shirou/gopsutil/mem"
        "github.com/shirou/gopsutil/cpu"


)

var myDynamicInfo map[string]interface{}

func getInfo() {
        // Get all the static information about this instance

        if myDynamicInfo == nil {
                myDynamicInfo = make(map[string]interface{},20)
        }

        myDynamicInfo["mem"]           , _ = mem.VirtualMemory()
        myDynamicInfo["cputimes"]      , _ = cpu.Times(false)
        myDynamicInfo["cputimes_i"]    , _ = cpu.Times(true)
        myDynamicInfo["cpupercent"]    , _ = cpu.Percent(0, false)
        myDynamicInfo["cpupercent_i"]  , _ = cpu.Percent(0, true)
	fmt.Println("%#v", myDynamicInfo)

}

func main() {

	fmt.Printf("%+v\n",myDynamicInfo)	

	getInfo()

        infoAnswer, ierr := json.MarshalIndent(myDynamicInfo, "", "\t")
        if ierr != nil { fmt.Println("Cannot json marshal info. Err %s", ierr) }
       	fmt.Printf("\n\n\n%s\n", infoAnswer)

}


