package main

import (
	"github.com/jaypipes/ghw"
)

func get_cpu_model() (string, error) {
	cpu, err := ghw.CPU()
	if err != nil {
		return "Unknown", nil
	}

	if len(cpu.Processors) > 0 {
		return cpu.Processors[0].Model, nil
	}

	return "Unknown", nil
}
