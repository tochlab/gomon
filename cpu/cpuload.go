package cpu

import (
    "os"
    "fmt"
)

type CpuLoadInfo struct {
    User float32 
    Nice float32 
    System float32 
    Idle float32 
    Wait float32 
}

func CpuLoad() CpuLoadInfo {
    procFile, err:=os.Open("/proc/stat")
    if err != nil {
        fmt.Println(err)
    }
    var user, nice, system, idle, iowait, irq, softirq, steal, guest, guestnice int
    var n int
    var cpuName string
    n, err = fmt.Fscanln(procFile, &cpuName, &user, &nice, &system, &idle, &iowait, &irq, &softirq, &steal, &guest, &guestnice)
    procFile.Close();
    var result CpuLoadInfo

    if( n == 0 || err !=nil) {
        fmt.Println(err)
        return result
    }

    sumtotal := user + nice + system + idle + iowait + irq +softirq + steal + guest + guestnice
    
    var sumfloat =  float32(sumtotal)
    result.User = float32(user) / sumfloat * 100
    result.Nice = float32(nice) / sumfloat * 100
    result.System = float32(system) / sumfloat * 100
    result.Idle = float32(idle) / sumfloat * 100
    result.Wait = float32(iowait) / sumfloat * 100
    return result
}