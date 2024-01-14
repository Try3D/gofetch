package main

import (
	"github.com/shirou/gopsutil/host"
	"os/user"
)

type SystemInfo struct {
	Uptime          int
	OS              string
	Distro          string
	HostName        string
	PlatformVersion string
	CurrentUser     string
}

func getSystemInfo() (SystemInfo, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return SystemInfo{}, err
	}

	currentUser, err := user.Current()
	if err != nil {
		return SystemInfo{}, err
	}

	uptime := int(hostInfo.Uptime)
	hostname := hostInfo.Hostname
	distro := hostInfo.Platform
	os := hostInfo.OS
	platformversion := hostInfo.PlatformVersion
	username := currentUser.Username

	info := SystemInfo{
		Uptime:          uptime,
		HostName:        hostname,
		Distro:          distro,
		OS:              os,
		PlatformVersion: platformversion,
		CurrentUser:     username,
	}

	return info, err
}
