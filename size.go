package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
)

//DirSize: Returns the size of the directory passed in as the path parameter or an error.
func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()

		}
		return err
	})
	return size, err
}

//SizeInMB: Converts bytes to megabytes.
func SizeInMB(byteSize int64) float64 {
	if byteSize < 0 {
		return 0
	}
	return float64(byteSize) / 1000.0 / 1000.0
}

//SizeInGB: Converts bytes to gigabytes.
func SizeInGB(byteSize int64) float64 {
	if byteSize < 0 {
		return 0
	}
	return float64(byteSize) / 1000.0 / 1000.0 / 1000.0
}

//SizeMBToGB: Converts megabytes to gigabytes.
func SizeMBToGB(mbSize float64) float64 {
	if mbSize < 0 {
		return 0
	}
	return mbSize / 1000.0
}

//RoundGB: Rounds gigabyte size to two digits.
func RoundGB(gbSize float64) float64 {
	if gbSize < 0 {
		return 0
	}
	return math.Round(gbSize*100) / 100
}

//SizeMsg: Outputs message describing the size of space cleaned.
func SizeMsg(mbSize float64) {
	if mbSize < 0 {
		log.Printf("Successfully freed up %d MB.", 0)
		return
	}
	if mbSize >= 1000.0 {
		gbSize := RoundGB(SizeMBToGB(mbSize))
		log.Printf("Successfully freed up %.2f GB.", gbSize)
	} else {
		log.Printf("Successfully freed up %.1f MB.", mbSize)
	}
}

//Gets the size of the latest iOS version in the DeviceSupport directory
func DeviceSupportLatestVerDirSize() int64 {
	d, err := os.Open(DeviceSupport())
	if err != nil {
		return 0
	}
	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		return 0
	}

	latest, err := getLatestVer(DeviceSupport())
	if err != nil {
		return 0
	}

	for _, name := range names {
		if latest != nil && strings.Split(name, " ")[0] == latest.Original() {
			latestDir := fmt.Sprintf("%s/%s", DeviceSupport(), name)
			latestSize, err := DirSize(latestDir)
			if err != nil {
				return 0
			}
			return latestSize
		}
	}
	return 0
}
