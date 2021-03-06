package main

import (
	"log"
	"time"

	"./cpu"
	"./mem"
)

func proc(x, y int) float32 {
	return (float32(x) / float32(y)) * 100.0
}

func main() {
	prevInfo := cpu.GetCPULoadInfo()
	var n uint32
	n = 0
	for {
		time.Sleep(1000000000)
		cpuload := cpu.GetCPULoadInfo()
		meminfo := mem.GetMemInfo()
		var deltaUser = cpuload.User - prevInfo.User
		var deltaNice = cpuload.Nice - prevInfo.Nice
		var deltaSyst = cpuload.System - prevInfo.System
		var deltaIdle = cpuload.Idle - prevInfo.Idle
		var deltaWait = cpuload.Wait - prevInfo.Wait
		var sumDelta = deltaUser + deltaNice + deltaSyst + deltaIdle + deltaWait
		if n%40 == 0 {
			log.Println("  User    Syst    Idle    Wait     Nice    Free")
		}
		log.Printf("%7.3f %7.3f %7.3f %7.3f %7.3f %7dMb\n",
			proc(deltaUser, sumDelta),
			proc(deltaSyst, sumDelta),
			proc(deltaIdle, sumDelta),
			proc(deltaWait, sumDelta),
			proc(deltaNice, sumDelta),
			meminfo.Free/(1024*1024),
		)
		prevInfo = cpuload
		n++
	}
}
