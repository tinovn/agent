package disk

import (
	"fmt"
	"syscall"
)

type DiskInfo struct {
	Total uint64
	Free  uint64
	Used  uint64
}

func GetDiskInfo() DiskInfo {
	// Get disk usage of the current working directory
	info, err := diskUsage("/")
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	return info
}

// DiskUsage returns the disk usage of the path in bytes
func diskUsage(path string) (usage DiskInfo, err error) {
	fs := &syscall.Statfs_t{}
	err = syscall.Statfs(path, fs)
	if err != nil {
		return
	}
	usage.Total = (fs.Blocks * uint64(fs.Bsize)) / (1024 * 1024)
	usage.Free = (fs.Bfree * uint64(fs.Bsize)) / (1024 * 1024)
	usage.Used = usage.Total - usage.Free
	return
}
