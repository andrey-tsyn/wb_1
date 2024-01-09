package main

import "fmt"

type Drive interface {
	DriveInfo() string
}

type HDD struct{}

func (h *HDD) DriveInfo() string {
	return "I'm HDD"
}

type FloppyDisk struct{}

func (f *FloppyDisk) FloppyDiskInformation() string {
	return "I'm floppy disk"
}

type FloppyDiskToDriveAdapter struct {
	floppyDisk *FloppyDisk
}

func (f *FloppyDiskToDriveAdapter) DriveInfo() string {
	return f.floppyDisk.FloppyDiskInformation()
}

func main() {
	var hdd Drive = &HDD{}
	fmt.Printf("HDD drive info: %s\n", hdd.DriveInfo())

	floppyDisk := FloppyDisk{}
	var adapter Drive = &FloppyDiskToDriveAdapter{&floppyDisk}
	fmt.Printf("Floppy disk drive info: %s\n", adapter.DriveInfo())
}
