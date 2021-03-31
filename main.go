package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jessevdk/go-flags"
)

//Options
var opts struct {
	Version  bool   `short:"v" long:"version" description:"Current version of xcclear."`
	Diagnose bool   `short:"d" long:"diagnose" description:"Diagnose the total space occupied by derived data, caches and other XCode related files."`
	Purge    bool   `short:"p" long:"purge" description:"Clean all caches/unnecessary files. \nThis will not delete your 'Archives' folder and the last installed iOS version in the 'iOS DeviceSupport' folder."`
	Clear    string `short:"c" long:"clear" description:"Clean only the specified directories. \n Options can be:\n -'derived': The 'Derived Data' folder. \n -'device': The 'iOS DeviceSupport' folder. (Will keep the latest iOS version.) \n -'watch': 'watchOS DeviceSupport' folder. \n -'simulator': 'CoreSimulator' folder. \n -'cache': XCode caches. \n -'archives': 'Archives' folder." choice:"derived" choice:"device" choice:"watch" choice:"simulator" choice:"cache" choice:"archives"`
}

//Colors used for logging messages
var blue = color.New(color.FgBlue).Add(color.Bold).SprintFunc()
var red = color.New(color.FgRed).Add(color.Bold).SprintFunc()
var green = color.New(color.FgGreen).Add(color.Bold).SprintFunc()

func main() {
	fmt.Println(SizeInMB(892145714))
	_, err := flags.Parse(&opts)

	if err != nil {
		panic(err)
	}

	//No flags provided error
	if !opts.Version && !opts.Diagnose && !opts.Purge && opts.Clear == "" {
		panic(errors.New("No flags provided. See -h or --help for possible options."))
	}

	//More than one flags error
	setFlags := []bool{opts.Version, opts.Diagnose, opts.Purge, opts.Clear != ""}
	if CountOf(true, setFlags) > 1 {
		panic(errors.New("More than one flag. Please pass in only one option at a time. See -h or --help for possible options."))
	}

	//Version
	if opts.Version {
		fmt.Println("1.0.0")
		os.Exit(0)
	}

	//Diagnosing XCode space
	if opts.Diagnose {
		DiagnoseAll()
		defer os.Exit(0)
	}

	//Clean everything
	if opts.Purge {
		Purge()
		defer os.Exit(0)
	}

	//clean specified folder
	switch opts.Clear {
	case "derived":
		size, _ := DirSize(DerivedData())
		sizeMb := SizeInMB(size)
		defer SizeMsg(sizeMb)
		CleanDirContents(DerivedData())
	case "device":
		size, _ := DirSize(DeviceSupport())
		sizeMb := SizeInMB(size)
		defer SizeMsg(sizeMb)
		CleanDirContents(DeviceSupport())
	case "watch":
		size, _ := DirSize(WatchOSSupport())
		sizeMb := SizeInMB(size)
		defer SizeMsg(sizeMb)
		CleanDirContents(WatchOSSupport())
	case "simulator":
		size, _ := DirSize(CoreSimulator())
		sizeMb := SizeInMB(size)
		defer SizeMsg(sizeMb)
		CleanDirContents(CoreSimulator())
	case "cache":
		size, _ := DirSize(XcCache())
		sizeMb := SizeInMB(size)
		defer SizeMsg(sizeMb)
		CleanDirContents(XcCache())
	case "archives":
		fmt.Println("This will delete all files in your Archives. Please make sure you don't need any of those.\nDo you wish to continue? (yes/no)")
		confirm := askForConfirmation()
		if !confirm {
			os.Exit(0)
		}
		size, _ := DirSize(Archives())
		sizeMb := SizeInMB(size)
		defer SizeMsg(sizeMb)

		CleanDirContents(Archives())
	}
}
