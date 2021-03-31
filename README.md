# xcclear

Say hello to a few extra gigabytes of space on your Mac with xcclear, a simple and easy to use command line application written in Go for cleaning unnecessary XCode files.

## Installation

    brew install xcclear
    
## Screenshots
Run ```xcclear -d``` to diagnose your storage:  

![diagnose](https://user-images.githubusercontent.com/44027725/113197954-9d19d480-9265-11eb-90ee-e360553e1e0f.gif)


Run ```xcclear -p``` to purge all unwanted files from your storage:  

![purge](https://user-images.githubusercontent.com/44027725/113198953-d30b8880-9266-11eb-9c6e-f7cfc9b91a82.gif)


## Options

flag                       |   type    | description
-------------------------- | --------- | ------------------------------------------
`-d/--diagnose`            | `bool`    | Diagnose the total space occupied by derived data, caches and other XCode related files.<br>
`-p/--purge`               | `bool`    | Clean all caches/unnecessary files.  
`-c/--clear`               | `string`  | Clean only the specified directories.<br>Options can be:<br> - `derived`: Derived Data. - `device`: iOS DeviceSupport (The latest iOS version will be kept.)<br> - `watch`: watchOS DeviceSupport<br> - `simulator`: 'CoreSimulator'<br> - `cache`: Xcode caches located at ```~/Library/Caches/com.apple.dt.Xcode```<br> - `archives`: Archives
`-v/--version`             | `bool`    | Get the current version.

## Development

### Compiling from source
    go get github.com/lajosdeme/xcclear
If ```$GOPATH/bin``` is on your ```$PATH``` run ```xcclear -h``` to see usage options. 

### Contributing
If you have any suggestions or questions feel free to raise an issue or reach out to me directly via <a href="mailto:lajosd@protonmail.ch">lajosd@protonmail.ch</a>.  
  
If this tool made your life easier you can thank me by buying me a coffee.  

<a href="https://www.buymeacoffee.com/edgz29w" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>

