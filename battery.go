package main

import (
	"fmt"
	"github.com/distatus/battery"
)

type BatteryInfo struct {
	State      string
	Charge     float64
	TimeToFull float64
}

func get_battery() (BatteryInfo, error) {
	batteries, err := battery.GetAll()
	if err != nil {
		fmt.Println("Could not get battery info!")
		return BatteryInfo{}, err
	}

	battery := batteries[0]

	state := battery.State.String()
	charge := battery.Current * 100 / battery.Full
	time_to_full := (battery.Full - battery.Current) / battery.ChargeRate

	info := BatteryInfo{
		State:      state,
		Charge:     charge,
		TimeToFull: time_to_full,
	}

	return info, nil
}
