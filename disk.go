package main

import (
	"github.com/shirou/gopsutil/v3/disk"
)

type DiskInfo struct {
	Used    int
	Total   int
	Percent int
}

func getDiskInfo() (DiskInfo, error) {
	diskUsage, err := disk.Usage("/")
	if err != nil {
		return DiskInfo{}, err
	}

	total := diskUsage.Total / 1024 / 1024 / 1024
	used := diskUsage.Used / 1024 / 1024 / 1024
	usedPercent := diskUsage.UsedPercent

	info := DiskInfo{
		Total:   int(total),
		Used:    int(used),
		Percent: int(usedPercent),
	}
	return info, nil
}
