package memory

import (
	"fmt"
	"os"
	"strings"
)

type MemoryInfo struct {
	Total string
	Free  string
	Used  string
}

func GetMemoryInfo() MemoryInfo {
	var memoryInfo MemoryInfo
	// Read the contents of /proc/meminfo
	memInfo, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		fmt.Println("Error reading /proc/meminfo:", err)
		return memoryInfo
	}

	// Parse the contents to get total, used, and free memory
	lines := strings.Split(string(memInfo), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		key := strings.TrimRight(fields[0], ":")
		value := fields[1]
		switch key {
		case "MemTotal":
			memoryInfo.Total = value
		case "MemFree":
			memoryInfo.Free = value
		case "MemAvailable":
			memoryInfo.Used = value
		}
	}
	return memoryInfo
}
