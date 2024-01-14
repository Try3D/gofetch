package main

import (
	"fmt"
	"math"
	"unicode"

	"github.com/mattn/go-runewidth"
)

func main() {
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

	title := "gofetch"
	prompt := fmt.Sprintf("%v@%v", SystemInfo.CurrentUser, SystemInfo.HostName)

	asciiLinux := []string{
		"    .---.    ",
		"   /     \\   ",
		"   \\.@-@./   ",
		"   /`\\_/`\\   ",
		"  //  _  \\\\  ",
		" | \\     )|_ ",
		"/`\\_`>  <_/ \\",
		"\\__/'---'\\__/",
	}

	formattedStrings := []string{
		fmt.Sprintf("OS: %v", capitalize(SystemInfo.OS)),
		fmt.Sprintf("Distro: %v %v", capitalize(SystemInfo.Distro), SystemInfo.PlatformVersion),
		fmt.Sprintf("CPU: %v", CpuInfo),
		fmt.Sprintf("Terminal: %v", Terminal),
		fmt.Sprintf("Shell: %v", Shell),
		fmt.Sprintf("Disk (/): %vG / %vG (%v%%)", DiskInfo.Used, DiskInfo.Total, DiskInfo.Percent),
		fmt.Sprintf("Battery: %.f%% %v", BatteryInfo.Charge, BatteryInfo.State),
		fmt.Sprintf("Time to full: %.f h %.f min", math.Floor(BatteryInfo.TimeToFull), 60*(BatteryInfo.TimeToFull-math.Floor(BatteryInfo.TimeToFull))),
	}

	maxWidth := 0

	for _, s := range formattedStrings {
		if maxWidth < runewidth.StringWidth(s) {
			maxWidth = runewidth.StringWidth(s)
		}
	}

	maxWidth += len(asciiLinux[0])

	fmt.Printf("╭───%v", title)
	printChar('─', maxWidth - runewidth.StringWidth(prompt) - runewidth.StringWidth(title) - 3)
	fmt.Printf("%v───╮\n", prompt)

	for i, s := range formattedStrings {
		fmt.Printf("│ ")
		fmt.Printf("%v ", asciiLinux[i])
		fmt.Printf("%v", s)
		printChar(' ', maxWidth-runewidth.StringWidth(s) - len(asciiLinux[0]) + 1)
		fmt.Printf("│\n")
	}

	fmt.Printf("╰")
	printChar('─', maxWidth + 3)
	fmt.Printf("╯\n")
}

func capitalize(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func printChar(s rune, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", s)
	}
}
