package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//Purge: Delete every file from every directory.
func Purge() {
	log.Printf("%s", red("Purging started..."))
	log.Println("Archives and the last installed iOS version in iOS DeviceSupport won't be deleted.")

	//close simulators
	cmd := exec.Command("killall", "Simulator")
	cmd.Start()

	var dirSizes int64 = 0
	for _, dir := range AllDirectories() {
		if dir == Archives() {
			continue
		}

		size, _ := DirSize(dir)
		dirSizes = dirSizes + size

		CleanDirContents(dir)
	}
	dirSizes = dirSizes - DeviceSupportLatestVerDirSize()
	totalMB := SizeInMB(dirSizes)
	SizeMsg(totalMB)
}

//CleanDirContents: Delete every file from the directory passed in as dir parameter.
func CleanDirContents(dir string) error {
	//open directory
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()

	//get all file names
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}

	dirSlice := strings.Split(dir, "/")
	dirName := dirSlice[len(dirSlice)-1]
	log.Printf("%s  %s\n", red("Purging"), blue(dirName))

	//loop through dir
	for _, name := range names {
		//if device support dir, we keep the latest os version
		if dir == DeviceSupport() {
			latest, err := getLatestVer(dir)
			if err != nil {
				continue
			}

			if latest != nil && strings.Split(name, " ")[0] == latest.Original() {
				continue
			}
		}

		//close simulators
		if dir == CoreSimulator() {
			cmd := exec.Command("killall", "Simulator")
			cmd.Start()
		}

		//removing all files
		err := os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
