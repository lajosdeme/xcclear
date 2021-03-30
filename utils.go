package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Masterminds/semver"
)

//CountOf: Count the number of times a boolean occurs in a boolean slice. Used to count the number of flags.
func CountOf(val bool, slice []bool) int {
	count := 0
	for _, el := range slice {
		if val == el {
			count = count + 1
		}
	}
	return count
}

//GetLatestVer: Find the folder name with the highest version in the iOS DeviceSupport folder.
func getLatestVer(dir string) (*semver.Version, error) {
	//open file
	d, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer d.Close()

	//get all file names
	names, err := d.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	//convert names to semver
	//get latest version
	var latest *semver.Version

	for _, name := range names {
		verStr := strings.Split(name, " ")[0]
		ver, err := semver.NewVersion(verStr)
		if err != nil {
			continue
		}
		if latest == nil || latest.LessThan(ver) {
			latest = ver
		}
	}

	return latest, nil
}

//Asks user for confirmation before proceeding.
func askForConfirmation() bool {
	var response string

	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}

	switch strings.ToLower(response) {
	case "y", "yes":
		return true
	case "n", "no":
		return false
	default:
		fmt.Println("I'm sorry but I didn't get what you meant, please type (y)es or (n)o and then press enter:")
		return askForConfirmation()
	}
}
