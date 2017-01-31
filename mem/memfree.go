package mem

import "syscall"

// MemInfo ...
type SysInfo struct {
	Total int64
	Free  int64
	Buff  int64
	Unit  int64
}

// GetMemInfo ...
func GetMemInfo() SysInfo {
	si := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(si)
	if err != nil {
		panic("Commander, we have a problem. syscall.Sysinfo:" + err.Error())
	}
	var result SysInfo
	result.Unit = int64(si.Unit)
	result.Free = int64(si.Freeram) * result.Unit
	result.Buff = int64(si.Bufferram) * result.Unit
	result.Total = int64(si.Totalram) * result.Unit
	return result
}
