package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
)

func get_cpu_model() (string, error) {
	cpuInfo, err := cpu.Info()
	if err != nil {
		return "", err
	}

	if len(cpuInfo) > 0 {
		return cpuInfo[0].ModelName, nil
	}

	return "", fmt.Errorf("No CPU information available")
}
