package main

import (
	"os"
)

func getTerminalEmulator() string {
	term := os.Getenv("TERM")
	return term
}

func getShell() string {
	shell := os.Getenv("SHELL")
	return shell
}
