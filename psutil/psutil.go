package psutil

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

func Demo() {
	// cpuInfo()
	// memoryVirtualMemory()
	// processProcesses()
}

// cpu
func cpuInfo() {
	c, _ := cpu.Info()
	fmt.Println(c)
}

// memory
func memoryVirtualMemory() {
	m, _ := mem.VirtualMemory()
	fmt.Println(m)
}

// process
func processProcesses() {
	p, _ := process.Processes()
	fmt.Println(p)
}
