package main

import (
	"os/exec"
	"runtime"
	"strings"
)

func get_cpu_model() (string, error) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("sysctl", "-n", "machdep.cpu.brand_string")
	case "linux":
		cmd = exec.Command("sh", "-c", "grep -m1 'model name' /proc/cpuinfo | cut -d: -f2")
	default:
		cmd = exec.Command("sh", "-c", "echo 'Unknown CPU'")
	}

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}
