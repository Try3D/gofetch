package main

import (
	"fmt"
	"math"
	"unicode"
)

func main() {
  blue := "\x1b[1;34m";
  bold := "\x1b[1m"
  reset := "\x1b[0m";

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

	title := bold + blue + "gofetch" + reset
	prompt := fmt.Sprintf("%v%v%v%v@%v%v%v%v", bold, blue, SystemInfo.CurrentUser, reset, bold, blue, SystemInfo.HostName, reset)

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
    fmt.Sprintf("%v%vOS:%v %v", bold, blue, reset, capitalize(SystemInfo.OS)),
		fmt.Sprintf("%v%vDistro:%v %v %v", bold, blue, reset, capitalize(SystemInfo.Distro), SystemInfo.PlatformVersion),
		fmt.Sprintf("%v%vCPU:%v %v", bold, blue, reset, CpuInfo),
		fmt.Sprintf("%v%vTerminal:%v %v", bold, blue, reset, Terminal),
		fmt.Sprintf("%v%vShell:%v %v", bold, blue, reset, Shell),
		fmt.Sprintf("%v%vDisk (/):%v %vG / %vG (%v%%)", bold, blue, reset, DiskInfo.Used, DiskInfo.Total, DiskInfo.Percent),
		fmt.Sprintf("%v%vBattery:%v %.f%% %v", bold, blue, reset, BatteryInfo.Charge, BatteryInfo.State),
		fmt.Sprintf("%v%vTime to full:%v %.f h %.f min", bold, blue, reset, math.Floor(BatteryInfo.TimeToFull), 60*(BatteryInfo.TimeToFull-math.Floor(BatteryInfo.TimeToFull))),
	}

	maxWidth := 0

	for _, s := range formattedStrings {
		if maxWidth < len(s) {
			maxWidth = len(s)
		}
	}

	maxWidth += len(asciiLinux[0])

	fmt.Printf("╭───%v", title)
	printChar('─', maxWidth - len(prompt) - len(title) + 27)
	fmt.Printf("%v───╮\n", prompt)

	for i, s := range formattedStrings {
		fmt.Printf("│ ")
		fmt.Printf("%v ", asciiLinux[i])
		fmt.Printf("%v", s)
		printChar(' ', maxWidth-len(s) - len(asciiLinux[0]) + 1)
		fmt.Printf("│\n")
	}

	fmt.Printf("╰")
	printChar('─', maxWidth - 12)
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
