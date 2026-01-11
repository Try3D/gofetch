package main

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/jaypipes/ghw"
)

func get_gpu_model() (string, error) {
	// Try ghw first (works on Linux)
	gpu, err := ghw.GPU()
	if err == nil && len(gpu.GraphicsCards) > 0 {
		card := gpu.GraphicsCards[0]
		if card.DeviceInfo != nil && card.DeviceInfo.Product != nil {
			name := card.DeviceInfo.Product.Name
			if name != "" && name != "unknown" {
				return name, nil
			}
		}
	}

	// Fallback for macOS
	if runtime.GOOS == "darwin" {
		cmd := exec.Command("sh", "-c", "system_profiler SPDisplaysDataType 2>/dev/null | awk -F': ' '/Chipset Model/{print $2}' | head -1")
		output, err := cmd.Output()
		if err == nil {
			result := strings.TrimSpace(string(output))
			if result != "" {
				return result, nil
			}
		}
	}

	// Fallback for Linux if ghw failed
	if runtime.GOOS == "linux" {
		cmd := exec.Command("sh", "-c", "lspci 2>/dev/null | grep -i 'vga\\|3d\\|display' | head -1 | sed 's/.*: //'")
		output, err := cmd.Output()
		if err == nil {
			result := strings.TrimSpace(string(output))
			if result != "" {
				return result, nil
			}
		}
	}

	return "Unknown", nil
}
