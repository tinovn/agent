package main

import (
	"fmt"

	"github.com/vitodeploy/agent/pkg/cpu"
	"github.com/vitodeploy/agent/pkg/disk"
	"github.com/vitodeploy/agent/pkg/memory"
)

func main() {
	// cfg := config.GetConfig()
	cpuInfo := cpu.GetCPUInfo()
	diskInfo := disk.GetDiskInfo()
	memoryInfo := memory.GetMemoryInfo()
	fmt.Println(cpuInfo)
	fmt.Println(diskInfo)
	fmt.Println(memoryInfo)
}
