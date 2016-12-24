package main

import "./cpu"
import "fmt"

func main() {
    cpuload := cpu.CpuLoad()
    fmt.Println(cpuload)
}