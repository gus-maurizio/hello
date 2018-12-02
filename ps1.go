package main

import _ "net/http/pprof"

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

var myDynamicInfo map[string]interface{}

type MeasureConfig struct {
	CPUblue		float64
	CPUgreen	float64
	CPUred		float64
}

type MeasureState struct {
	CPUcount	int
	CPUcountf	float64

	SYSLOADalert	bool
	SYSLOADgrow	bool
	SYSLOADtrend	bool
	SYSLOADlast	load.AvgStat
	SYSLOAD		load.AvgStat

}

var MyMeasureConfig MeasureConfig
var MyMeasureState  MeasureState

func getInfo() {
	// Get all the static information about this instance

	if myDynamicInfo == nil {
		myDynamicInfo = make(map[string]interface{}, 20)
		MyMeasureState.SYSLOADlast.Load1  = 1.0
		MyMeasureState.SYSLOADlast.Load5  = 1.0
		MyMeasureState.SYSLOADlast.Load15 = 1.0
	}

	myDynamicInfo["mem"], _ 		= mem.VirtualMemory()
	myDynamicInfo["cputimes"], _ 		= cpu.Times(false)
	myDynamicInfo["cpu%"], _ 		= cpu.Percent(0, false)
	myDynamicInfo["core%"], _ 		= cpu.Percent(0, true)
	myDynamicInfo["systemloadavg"], _ 	= load.Avg()
	myDynamicInfo["miscstats"], _ 		= load.Misc()

	MyMeasureState.SYSLOAD.Load1		= myDynamicInfo["systemloadavg"].(*load.AvgStat).Load1  / MyMeasureState.CPUcountf
	MyMeasureState.SYSLOAD.Load5		= myDynamicInfo["systemloadavg"].(*load.AvgStat).Load5  / MyMeasureState.CPUcountf
	MyMeasureState.SYSLOAD.Load15		= myDynamicInfo["systemloadavg"].(*load.AvgStat).Load15 / MyMeasureState.CPUcountf
	myDynamicInfo["sysload"]		= MyMeasureState.SYSLOAD

        // What exactly is good?
        // systemloadavg (load1 load5 and load15) should be divided by #CPUs
        // and will give an idea of the trend. Your mileage might vary BUT,
        // we can say that any value higher than 1.0 is indication of
        // SATURATION!!! While your system might work fine at 1.5, values
        // should not exceed 2.0
        //
        // Utilization metrics
        // per-CPU and per-Process utilization. Utilization metrics are useful
        // for workload characterization.
        //
        // Saturation metrics
        // per-thread run queue, CPU run queue latency, CPU run queue length.
        // Saturation metrics are useful for identifying a performance problem.
        //
        // In 1993, a Linux engineer found a nonintuitive case with load averages,
        // and with a three-line patch changed them forever from "CPU load averages"
        // to what one might call "system load averages."
        // His change included tasks in the uninterruptible state, so that
        // load averages reflected demand for disk resources and not just CPUs.
        // These system load averages count the number of threads working and
        // waiting to work, and are summarized as a triplet of exponentially-damped
        // moving sum averages that use 1, 5, and 15 minutes as constants in an equation.
        // This triplet of numbers lets you see if load is increasing or decreasing,
        // and their greatest value may be for relative comparisons with themselves.
        //
        // The use of the uninterruptible state has since grown in the Linux kernel,
        // and nowadays includes uninterruptible lock primitives.
        // If the load average is a measure of demand in terms of running and
        // waiting threads (and not strictly threads wanting hardware resources),
        // then they are still working the way we want them to.
        //
	if MyMeasureState.SYSLOAD.Load1 >  MyMeasureState.SYSLOAD.Load5 && 
	   MyMeasureState.SYSLOAD.Load5 >= MyMeasureState.SYSLOAD.Load15 {
		MyMeasureState.SYSLOADgrow = true
	} else {MyMeasureState.SYSLOADgrow = false}

        if MyMeasureState.SYSLOAD.Load1  >= MyMeasureState.SYSLOADlast.Load1 &&
           MyMeasureState.SYSLOAD.Load5  >= MyMeasureState.SYSLOADlast.Load5 &&
           MyMeasureState.SYSLOAD.Load15 >= MyMeasureState.SYSLOADlast.Load15 {
		MyMeasureState.SYSLOADtrend = true
	} else {MyMeasureState.SYSLOADtrend = false}

	if MyMeasureState.SYSLOADgrow && MyMeasureState.SYSLOADtrend {
		MyMeasureState.SYSLOADalert = true
	} else {MyMeasureState.SYSLOADalert = false}

        fmt.Printf("%+v\n", myDynamicInfo)
        fmt.Printf("%+v\n", MyMeasureState)

	MyMeasureState.SYSLOADlast = MyMeasureState.SYSLOAD
}

func main() {

	MyMeasureConfig.CPUblue		= 5.0/100.0 
	MyMeasureConfig.CPUgreen	= 50.0/100.0 
	MyMeasureConfig.CPUred		= 80.0/100.0 

	MyMeasureState.CPUcount , _ 	= cpu.Counts(true)
	MyMeasureState.CPUcountf 	= float64(MyMeasureState.CPUcount)

	for i := 0; i < 3; i++ {
		time.Sleep(1000 * time.Millisecond)
		getInfo()
	}
	// What exactly is good?
	// systemloadavg (load1 load5 and load15) should be divided by #CPUs 
	// and will give an idea of the trend. Your mileage might vary BUT,
	// we can say that any value higher than 1.0 is indication of
	// SATURATION!!! While your system might work fine at 1.5, values
	// should not exceed 2.0
	//
	// Utilization metrics
	// per-CPU and per-Process utilization. Utilization metrics are useful 
	// for workload characterization.
	//
	// Saturation metrics
	// per-thread run queue, CPU run queue latency, CPU run queue length.
	// Saturation metrics are useful for identifying a performance problem.
	//
	// In 1993, a Linux engineer found a nonintuitive case with load averages, 
	// and with a three-line patch changed them forever from "CPU load averages"
	// to what one might call "system load averages." 
	// His change included tasks in the uninterruptible state, so that 
	// load averages reflected demand for disk resources and not just CPUs. 
	// These system load averages count the number of threads working and 
	// waiting to work, and are summarized as a triplet of exponentially-damped 
	// moving sum averages that use 1, 5, and 15 minutes as constants in an equation. 
	// This triplet of numbers lets you see if load is increasing or decreasing, 
	// and their greatest value may be for relative comparisons with themselves.
	// 
	// The use of the uninterruptible state has since grown in the Linux kernel, 
	// and nowadays includes uninterruptible lock primitives. 
	// If the load average is a measure of demand in terms of running and 
	// waiting threads (and not strictly threads wanting hardware resources), 
	// then they are still working the way we want them to.
	// 

	
	infoAnswer, ierr := json.MarshalIndent(myDynamicInfo, "", "\t")
	if ierr != nil {
		fmt.Println("Cannot json marshal info. Err %s", ierr)
	}
	fmt.Printf("\n%+v\n\n%s\n", MyMeasureConfig, infoAnswer)

}
