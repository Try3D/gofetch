package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func main() {
	blue := "\x1b[1;34m"
	bold := "\x1b[1m"
	reset := "\x1b[0m"

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

	asciiArt := []string{
		"    .---.    ",
		"   /     \\   ",
		"   \\.@-@./   ",
		"   /`\\_/`\\   ",
		"  //  _  \\\\  ",
		" | \\     )|_ ",
		"/`\\_`>  <_/ \\",
		"\\__/'---'\\__/",
	}
	asciiWidth := 13

	infoStrings := []string{
		fmt.Sprintf("OS: %v", capitalize(SystemInfo.OS)),
		fmt.Sprintf("Distro: %v %v", capitalize(SystemInfo.Distro), SystemInfo.PlatformVersion),
		fmt.Sprintf("CPU: %v", CpuInfo),
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
	// Content: "│ " + ascii + " " + info + padding + " │"
	// We need inner width (between the │ chars)
	innerWidth := 1 + asciiWidth + 1 + maxInfoWidth + 1 // space + ascii + space + info + space

	// Header needs: "───" + "gofetch" + "───" + user@host + "───"
	headerInner := 3 + 7 + 3 + len(username) + 1 + len(hostname) + 3
	if headerInner > innerWidth {
		innerWidth = headerInner
	}

	// Print top border
	title := bold + blue + "gofetch" + reset
	prompt := bold + blue + username + reset + "@" + bold + blue + hostname + reset
	headerPadding := innerWidth - 3 - 7 - 3 - len(username) - 1 - len(hostname) - 3
	fmt.Printf("╭───%s───%s%s───╮\n", title, strings.Repeat("─", headerPadding), prompt)

	// Print content rows
	for i, info := range infoStrings {
		coloredInfo := colorizeLabel(info, bold, blue, reset)
		padding := innerWidth - 1 - asciiWidth - 1 - len(info) - 1
		fmt.Printf("│ %s %s%s │\n", asciiArt[i], coloredInfo, strings.Repeat(" ", padding))
	}

	// Print bottom border
	fmt.Printf("╰%s╯\n", strings.Repeat("─", innerWidth))
}

// colorizeLabel adds color to the label part (before the colon)
func colorizeLabel(s, bold, blue, reset string) string {
	for i, c := range s {
		if c == ':' {
			return bold + blue + s[:i+1] + reset + s[i+1:]
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
