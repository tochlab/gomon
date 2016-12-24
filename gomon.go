package main

import "./cpu"
import "fmt"

func main() {
	cpuload := cpu.GetCPULoadInfo()
	fmt.Println(cpuload)
}
