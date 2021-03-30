package main

import (
	"fmt"
	"strings"
)

var totalSize int64 = 0

//DiagnoseAll: Get the size of each directory and print it on the command line.
func DiagnoseAll() {
	for _, dir := range AllDirectories() {
		dirSlice := strings.Split(dir, "/")
		dirName := dirSlice[len(dirSlice)-1]
		if dir == XcCache() {
			dirName = "XCode Cache"
		}
		diagnoseDir(dir, dirName)
	}

	mbSizeTotal := SizeInMB(totalSize)
	if mbSizeTotal >= 1000.0 {
		gbSize := RoundGB(SizeMBToGB(mbSizeTotal))
		gbStr := fmt.Sprintf("%.2f", gbSize)
		fmt.Printf("\n%s: %s %s \n", blue("Total:"), addPadding("Total:"), red(gbStr+" GB"))
	} else {
		mbStr := fmt.Sprintf("%.1f", mbSizeTotal)
		fmt.Printf("%s: %s %s \n", blue("Total:"), addPadding("Total:"), green(mbStr+" MB"))
	}
	fmt.Printf("\nType %s or %s to clear everything. Archives and the last installed iOS version in iOS DeviceSupport won't be deleted.", red("xcclear -p"), red("xcclear --purge"))
}

//diagnoseDir: Get the size of the directory passed in as the path parameter and print description message.
func diagnoseDir(path string, dirName string) {
	sizeOfDir, err := DirSize(path)
	totalSize = totalSize + sizeOfDir

	if err != nil {
		fmt.Printf("%s: - directory doesn't exist.\n", dirName)
		return
	}

	mbSize := SizeInMB(sizeOfDir)

	if mbSize >= 1000.0 {
		gbSize := RoundGB(SizeMBToGB(mbSize))
		gbStr := fmt.Sprintf("%.2f", gbSize)
		fmt.Printf("%s: %s %s \n", dirName, addPadding(dirName), red(gbStr+" GB"))
	} else {
		mbStr := fmt.Sprintf("%.1f", mbSize)
		fmt.Printf("%s: %s %s \n", dirName, addPadding(dirName), green(mbStr+" MB"))
	}
}

//addPadding: Add a padding between the directory name and the size of the directory.
func addPadding(title string) string {
	padding := ""
	for i := 0; i < (40 - len(title)); i++ {
		padding = padding + "-"
	}
	return padding
}
