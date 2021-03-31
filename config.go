package main

import "os"

func home() string {
	return os.Getenv("HOME")
}
func devDir() string {
	return home() + "/Library/Developer/"
}
func xcodeDir() string {
	return devDir() + "Xcode/"
}

func DerivedData() string {
	return xcodeDir() + "DerivedData"
}

func Archives() string {
	return xcodeDir() + "Archives"
}

func DeviceSupport() string {
	return xcodeDir() + "iOS DeviceSupport"
}

func WatchOSSupport() string {
	return xcodeDir() + "watchOS DeviceSupport"
}

func CoreSimulator() string {
	return devDir() + "CoreSimulator"
}

func XcCache() string {
	return home() + "/Library/Caches/com.apple.dt.Xcode"
}

func AllDirectories() []string {
	return []string{DerivedData(), Archives(), DeviceSupport(), WatchOSSupport(), CoreSimulator(), XcCache()}
}
