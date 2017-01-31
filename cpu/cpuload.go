package cpu

import (
	"fmt"
	"os"
)

// CPULoadInfo Cpu load info
type CPULoadInfo struct {
	User   int
	Nice   int
	System int
	Idle   int
	Wait   int
}

// GetCPULoadInfo get info about cpu load
func GetCPULoadInfo() CPULoadInfo {
	procFile, err := os.Open("/proc/stat")
	if err != nil {
		fmt.Println(err)
	}
	var user, nice, system, idle, iowait, irq, softirq, steal, guest, guestnice int
	var n int
	var cpuName string
	n, err = fmt.Fscanln(procFile, &cpuName, &user, &nice, &system, &idle, &iowait, &irq, &softirq, &steal, &guest, &guestnice)
	procFile.Close()
	var result CPULoadInfo

	if n == 0 || err != nil {
		fmt.Println(err)
		return result
	}

	result.User = user
	result.Nice = nice
	result.System = system
	result.Idle = idle
	result.Wait = iowait
	return result
}
