package main

import (
	"fmt"
	"math"
	"runtime"
	"strings"
	"unicode"
)

func main() {
	blue := "\x1b[1;34m"
	green := "\x1b[1;32m"
	bold := "\x1b[1m"
	reset := "\x1b[0m"

	// Select color based on OS
	accentColor := blue
	if runtime.GOOS == "darwin" {
		accentColor = blue
	} else if runtime.GOOS == "linux" {
		accentColor = green
	}

	BatteryInfo, err := get_battery()
	if err != nil {
		fmt.Println(err)
		return
	}

	CpuInfo, err := get_cpu_model()
	if err != nil {
		fmt.Println(err)
		return
	}

	GpuInfo, err := get_gpu_model()
	if err != nil {
		fmt.Println(err)
		return
	}

	SystemInfo, err := getSystemInfo()
	if err != nil {
		fmt.Println(err)
		return
	}

	Terminal := getTerminalEmulator()
	Shell := getShell()

	DiskInfo, err := getDiskInfo()
	if err != nil {
		fmt.Println(err)
		return
	}

	username := SystemInfo.CurrentUser
	hostname := SystemInfo.HostName

	// Select ASCII art based on OS
	var asciiArt []string
	if runtime.GOOS == "darwin" {
		// Apple logo
		asciiArt = []string{
			"         _      ",
			"        (/      ",
			"   .---__--.    ",
			"  /         \\   ",
			" |         /    ",
			" |         \\_   ",
			"  \\         /   ",
			"   `._.-._.'    ",
			"                ",
		}
	} else {
		// Tux (Linux penguin)
		asciiArt = []string{
			"    .---.       ",
			"   /     \\      ",
			"   \\.@-@./      ",
			"   /`\\_/`\\      ",
			"  //  _  \\\\     ",
			" | \\     )|_    ",
			"/`\\_`>  <_/ \\   ",
			"\\__/'---'\\__/   ",
			"                ",
		}
	}
	asciiWidth := 16

	infoStrings := []string{
		fmt.Sprintf("OS: %v", capitalize(SystemInfo.OS)),
		fmt.Sprintf("Distro: %v %v", capitalize(SystemInfo.Distro), SystemInfo.PlatformVersion),
		fmt.Sprintf("CPU: %v", CpuInfo),
		fmt.Sprintf("GPU: %v", GpuInfo),
		fmt.Sprintf("Terminal: %v", Terminal),
		fmt.Sprintf("Shell: %v", Shell),
		fmt.Sprintf("Disk (/): %vG / %vG (%v%%)", DiskInfo.Used, DiskInfo.Total, DiskInfo.Percent),
		fmt.Sprintf("Battery: %.f%% %v", BatteryInfo.Charge, BatteryInfo.State),
		fmt.Sprintf("Time to full: %.f h %.f min", math.Floor(BatteryInfo.TimeToFull), 60*(BatteryInfo.TimeToFull-math.Floor(BatteryInfo.TimeToFull))),
	}

	// Find max width of info strings
	maxInfoWidth := 0
	for _, s := range infoStrings {
		if len(s) > maxInfoWidth {
			maxInfoWidth = len(s)
		}
	}

	// Calculate widths
	innerWidth := 1 + asciiWidth + 1 + maxInfoWidth + 1

	// Header needs: "───" + "gofetch" + "───" + user@host + "───"
	headerInner := 3 + 7 + 3 + len(username) + 1 + len(hostname) + 3
	if headerInner > innerWidth {
		innerWidth = headerInner
	}

	// Print top border
	title := bold + accentColor + "gofetch" + reset
	prompt := bold + accentColor + username + reset + "@" + bold + accentColor + hostname + reset
	headerPadding := innerWidth - 3 - 7 - 3 - len(username) - 1 - len(hostname) - 3
	fmt.Printf("╭───%s───%s%s───╮\n", title, strings.Repeat("─", headerPadding), prompt)

	// Print content rows
	for i, info := range infoStrings {
		coloredInfo := colorizeLabel(info, bold, accentColor, reset)
		padding := innerWidth - 1 - asciiWidth - 1 - len(info) - 1
		fmt.Printf("│ %s %s%s │\n", asciiArt[i], coloredInfo, strings.Repeat(" ", padding))
	}

	// Print bottom border
	fmt.Printf("╰%s╯\n", strings.Repeat("─", innerWidth))
}

// colorizeLabel adds color to the label part (before the colon)
func colorizeLabel(s, bold, color, reset string) string {
	for i, c := range s {
		if c == ':' {
			return bold + color + s[:i+1] + reset + s[i+1:]
		}
	}
	return s
}

func capitalize(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
