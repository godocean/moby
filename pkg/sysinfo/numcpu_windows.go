package sysinfo // import "github.com/docker/docker/pkg/sysinfo"

import (
	"golang.org/x/sys/windows"
	"runtime"
)

var (
	kernel32                = windows.NewLazySystemDLL("kernel32.dll")
	getActiveProcessorCount = kernel32.NewProc("GetActiveProcessorCount")
)

const ALL_PROCESSOR_GROUPS = 0xFFFF

func numCPU() int {
	if amount, _, _ := getActiveProcessorCount.Call(uintptr(ALL_PROCESSOR_GROUPS)); amount != 0 {
		return int(amount)
	}
	return runtime.NumCPU()
}
