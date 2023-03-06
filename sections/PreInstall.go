package sections

/*
	https://zetcode.com/golang/exec-command/

	[ ] -
	[ ] -
	[ ] -
	[ ] -
	[ ] -
	[ ] -
	[ ] -
	[ ] -

Run Shell commands:
	cmd := exec.Command("cmd", "arg1")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	Check(err)

	s := string(out)



8085160 totalMem
500107862016 HDD Size

if /home, SWAP, /file & /custom are set:
    [x] 6.4% goes to ROOT (1/3 of 20%)
    [x] 13.6% goes to HOME (2/3 of 20%)
    [x] subtract the SWAP
    [ ] asks for user partitioning
    [x] the rest goes to FILE
if /home, SWAP & /file are set but /custom not:
    [x] 6.4% goes to ROOT (1/3 of 20%)
    [x] 13.6% goes to HOME (2/3 of 20%)
    [x] subtract the SWAP
    [x] the rest goes to FILE
if /home, SWAP & /custom are set but /file not:
    [x] 6.4% goes to ROOT (1/3 of 20%)
    [x] 13.6% goes to HOME (2/3 of 20%)
    [x] subtract the SWAP
    [x] the rest goes to user partitioning
if /home & SWAP are set, but /file and /custom not:
    [x] 20% goes to ROOT
    [x] subtract the SWAP
    [x] the rest goes to HOME
if /file & SWAP are set but /home and /custom not:
    [x] 20% goes to ROOT
    [x] subtract the SWAP
    [x] the rest goes to FILE
if /file, SWAP & /custom are set but /home not:
    [x] 20% goes to ROOT
    [x] subtract the SWAP
    [x] asks for user partitioning
    [x] the rest goes to FILE
if /custom & SWAP are set but /home and /file not:
    [x] 20% goes to ROOT
    [x] subtract the SWAP
    [x] asks for user partitioning
if none is set:
    [x] subtract the SWAP
    [x] the rest goes to ROOT



*/

import (
	"ArchInstall/helpers"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func setCountryISO() string {
	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Getting country ISO")
	iso := helpers.CurlResponse("https://ifconfig.co/country-iso")
	if helpers.YesNo(fmt.Sprintf("Country ISO detected to be %s, is this correct?", iso)) {
		return iso
	} else {
		inputISO := helpers.InputPrompt("Enter your country ISO e.g: USA")
		return inputISO
	}
}

func setTimeDate() {
	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Synchronizing hardware clock")
	cmd := exec.Command("timedatectl", "set-ntp", "true")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	helpers.Check(err)
	fmt.Println(string(out))
}

func setupPacman(reflectorCountryISO string) {
	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Setting up PacMan")
	cmd := helpers.RunShellCommand("pacman", "-S", "--noconfirm", "archlinux-keyring")
	fmt.Print(cmd)
	cmd = helpers.RunShellCommand("pacman", "-S", "--noconfirm", "--needed", "pacman-contrib")
	fmt.Println(cmd)
	helpers.ReplaceFileLine("/etc/pacman.conf", "#ParallelDownloads", "ParallelDownloads")
	cmd = helpers.RunShellCommand("pacman", "-S", "--noconfirm", "--needed", "reflector", "grub")
	fmt.Println(cmd)

	helpers.CopyFile("/etc/pacman.d/mirrorlist", "/etc/pacman.d/mirrorlist.backup")
	reflectorArgs := []string{	"-a", "48", "-c", reflectorCountryISO, "-f", "5",
								"-l", "20", "--sort", "rate", "--save", "/etc/pacman.d/mirrorlist"}
	cmd = helpers.RunShellCommand("reflector", reflectorArgs...)
	fmt.Println(cmd)
	cmd = helpers.RunShellCommand("mkdir", "/mnt")
	fmt.Println(cmd)

	cmd = helpers.RunShellCommand("pacman", "-S", "--noconfirm", "--needed", "gptfdisk")
	fmt.Println(cmd)
}

func formatDisk(cfgFile string) {
	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Formatting Disk")

	disk := helpers.JsonGetter(cfgFile, "disk")
	helpers.RunShellCommand("umount", "-A", "--recursive", "/mnt")
	helpers.RunShellCommand("sgdisk", "-Z", disk)
	helpers.RunShellCommand("sgdisk", "-a", "2048", "-o", disk)
}

func bootSystem() string {
	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Boot System")
	bootSys := []string{"BIOS", "UEFI"}
	var retBootSys string

	if helpers.CheckFileExists("/sys/firmware/efi") {
		retBootSys = "UEFI"
	} else {
		retBootSys = "BIOS"
	}

	if helpers.YesNo(fmt.Sprintf("Script detected that this machine has a \"%s\" boot system, is this correct?", retBootSys)) {
		return retBootSys
	} else {
		_, selectedBootSystem := helpers.PromptSelect("What is your boot system?", bootSys)
		return selectedBootSystem
	}
}

func getRecommendedSwapSize(totalRam int) (int, string) {
	//var recommendedSize uint64
	var swapRecommendedSize int
	var ramSizes = []int{
		8 * helpers.GiB,
		7 * helpers.GiB,
		6 * helpers.GiB,
		5 * helpers.GiB,
		4 * helpers.GiB,
		3 * helpers.GiB,
		2 * helpers.GiB,
		1 * helpers.GiB,
	}
	for i, v := range ramSizes {
		if totalRam < v {
			swapRecommendedSize = i
		}
	}
	return swapRecommendedSize, fmt.Sprintf("0:0:+%dGiB", swapRecommendedSize)
}

func partitionDisk(cfgFile string) (bool, int) {
	isSWAPSet := false
	rrSize := 0
	bSys := helpers.JsonGetter(cfgFile, "bootSystem")
	disk := helpers.JsonGetter(cfgFile, "disk")
	if bSys == "BIOS" {
		cmd := helpers.RunShellCommand("sgdisk", "-n", "0:0:+1M", "--typecode=0:ef02", "--change-name=0:'BIOSBOOT'", disk)
		fmt.Println(cmd)
	} else {
		cmd := helpers.RunShellCommand("sgdisk", "-n", "0:0:+300M", "--typecode=0:ef00", "--change-name=0:'EFIBOOT'", disk)
		fmt.Println(cmd)
	}


	memInfo := helpers.GetLine("/proc/meminfo", "MemTotal")
	totalMem := helpers.ExtractNumbers(memInfo.(string))[0] * helpers.KiB

	if totalMem < (8 * helpers.GiB) {
		rSize, swapSize := getRecommendedSwapSize(totalMem)
		if rSize > 0 {
			rrSize = rSize * helpers.GiB
			fmt.Println("The script detected that you have less than 8GiB of RAM")
			fmt.Printf("Based on your RAM size of %s, the script recommends a SWAP size of %s\n",
			helpers.ByteSizeConverter(uint64(totalMem)), helpers.ByteSizeConverter(uint64(rSize * helpers.GiB)))

			if helpers.YesNo("Would you like to create a SWAP partition?") {
				cmd := helpers.RunShellCommand("sgdisk", "-n", swapSize, "--typecode=0:8200", "--change-name=0:'SWAP'", disk)
				fmt.Println(cmd)
				isSWAPSet = true
			}
		}
	}

	//cmd := helpers.RunShellCommand("sgdisk", "-n", "0:0:0", "--typecode=0:8300", "--change-name=0:'ROOT'", disk)
	//fmt.Println(cmd)

	// SET PARTITIONING

	// sgdisk -n 0:0:+4GiB -t 0:8200 -c 0:swap

	cmd := helpers.RunShellCommand("partprobe", disk)
	fmt.Println(cmd)

	return isSWAPSet, rrSize
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//// I think it's cleaner now
//

func _calcRootHomeSize(dSizePercent int) (int, int) {
	dRSize := int(helpers.RoundMultiple((float64(dSizePercent) / 3.0), 8))
	dHSize := dSizePercent - dRSize

	return (dRSize * helpers.GiB), (dHSize * helpers.GiB)
}


func _showRemainingDisk(restSize int) {
	remainDiskSpace := fmt.Sprintf("Remaining Disk Space: %s (%d bytes)", helpers.ByteSizeConverter(uint64(restSize)), restSize)
	fmt.Println(helpers.PrintBlack(remainDiskSpace))
}

func _askSizeOfDisk(restSize, diskFullSize int, partName string) (int, string) {
	var retSize int
	var sgStr string


	fmt.Println(helpers.PrintBlack("use 'G' for Giga, 'M' for Mega and 'K' for Kilo, example: 400M"))
	fmt.Println(helpers.PrintBlack("It doesn't supports decimal places, so if you want 2.5G write as 2500M"))
	userSize := helpers.InputPrompt(fmt.Sprintf("Enter the size of the %s partition", partName))
	userByteSize := helpers.ConvertToByte(userSize)

	if (diskFullSize - userByteSize) <= 0 {

		fmt.Println(helpers.PrintRed("Can't exceed drive physical space"))
		retSize, sgStr = _askSizeOfDisk(restSize, diskFullSize, partName)
	} else {
		retSize = restSize - userByteSize
		sgStr = fmt.Sprintf("0:0:+%s", helpers.ParseSizeString(userSize))
	}

	return retSize, sgStr
}

func _setRootPart(isHomeSet, isFileSet, isCustomSet bool, percentSize, restSize, diskFullSize int) (int, string, string) {
	var rootSize, returnSize int
	var sgString string
	partitionName := "ROOT"

	if isHomeSet && (isFileSet || isCustomSet) {
		rootSize, _ = _calcRootHomeSize(percentSize)
		rootSizeBytes := helpers.ByteSizeConverter(uint64(rootSize))
		helpers.ClearConsole()
		fmt.Println(isHomeSet)
		promptString := fmt.Sprintf("Based on you drive size, the scripts recommends a size of %s to %s, would you like to use it?", rootSizeBytes, partitionName)
		if helpers.YesNo(promptString) {
			sgString = fmt.Sprintf("0:0:+%s", helpers.ParseSizeString(rootSizeBytes))
			returnSize = restSize - rootSize
		} else {
			helpers.ClearConsole()
			_showRemainingDisk(restSize)
			returnSize, sgString = _askSizeOfDisk(restSize, diskFullSize, partitionName)
		}
	} else if isHomeSet && !isFileSet && !isCustomSet {
		rootSizeBytes := helpers.ByteSizeConverter(uint64(percentSize * helpers.GiB))
		helpers.ClearConsole()
		promptString := fmt.Sprintf("Based on you drive size, the scripts recommends a size of %s to %s, would you like to use it?", rootSizeBytes, partitionName)
		if helpers.YesNo(promptString) {
			sgString = fmt.Sprintf("0:0:+%s", helpers.ParseSizeString(rootSizeBytes))
			returnSize = restSize - (percentSize * helpers.GiB)
		} else {
			helpers.ClearConsole()
			_showRemainingDisk(restSize)
			returnSize, sgString = _askSizeOfDisk(restSize, diskFullSize, partitionName)
		}
	} else if !isHomeSet && (isFileSet || isCustomSet) {
		rootSizeBytes := helpers.ByteSizeConverter(uint64(percentSize * helpers.GiB))
		helpers.ClearConsole()
		promptString := fmt.Sprintf("Based on you drive size, the scripts recommends a size of %s to %s, would you like to use it?", rootSizeBytes, partitionName)
		if helpers.YesNo(promptString) {
			sgString = fmt.Sprintf("0:0:+%s", helpers.ParseSizeString(rootSizeBytes))
			returnSize = restSize - (percentSize * helpers.GiB)
		} else {
			helpers.ClearConsole()
			_showRemainingDisk(restSize)
			returnSize, sgString = _askSizeOfDisk(restSize, diskFullSize, partitionName)
		}
	}  else {
		returnSize = restSize
		sgString = "0:0:0"
	}

	return returnSize, partitionName, sgString
}

func _setHomePart(isFileSet, isCustomSet bool, percentSize, restSize, diskFullSize int) (int, string, string) {
	var homeSize, returnSize int
	var sgString string
	partitionName := "HOME"

	if !isFileSet && !isCustomSet {
		returnSize = restSize
		sgString = "0:0:0"
	} else if isFileSet || isCustomSet {
		_, homeSize = _calcRootHomeSize(percentSize)
		bytesHomeSize := helpers.ByteSizeConverter(uint64(homeSize))
		helpers.ClearConsole()
		promptString := fmt.Sprintf("Based on you drive size, the scripts recommends a size of %s to %s, would you like to use it?", bytesHomeSize, partitionName)
		if helpers.YesNo(promptString) {
			sgString = fmt.Sprintf("0:0:+%s", helpers.ParseSizeString(bytesHomeSize))
			returnSize = restSize - (percentSize * helpers.GiB)
		} else {
			helpers.ClearConsole()
			_showRemainingDisk(restSize)
			returnSize, sgString = _askSizeOfDisk(restSize, diskFullSize, partitionName)
		}
	}

	return returnSize, partitionName, sgString
}

func _setFilePart(restSize int) (int, string, string) {
	partitionName := "FILES"
	return restSize, partitionName, "0:0:0"
}

func _partHint(isFileSet bool) {
	if isFileSet {
		a := helpers.PrintYellow("Reminder:")
		b := helpers.PrintBlack("You had selected to use a files partition")
		c := helpers.PrintBlack("that partition will be used to fill the empty space in your disk.")
		d := helpers.PrintBlack("So don't worry on filling up the remaining space with these partitions.")
		fmt.Printf("%s %s %s\n%s\n", a, b, c, d)
	} else {
		z := helpers.PrintYellow("Note:")
		x := helpers.PrintBlack("The last partition will be used to fill the remaining space in your disk.")
		fmt.Printf("%s %s\n", z, x)
	}
}

func _setMorePart(isFileSet bool, restSize, diskFullSize int) [][]string {
	var partCommands [][]string
	var returnSize int = restSize
	var sgString string

	helpers.ClearConsole()
	_partHint(isFileSet)
	answer, err := strconv.Atoi(helpers.InputPrompt("How many partitions do you want to create?"))
	helpers.Check(err)

	if isFileSet {
		for i := 0; i < answer; i++ {
			x := i
			helpers.ClearConsole()
			_partHint(isFileSet)
			_showRemainingDisk(returnSize)
			partName := helpers.InputPrompt(fmt.Sprintf("What is the name of the partition %d?", (x+1)))
			helpers.ClearConsole()

			_partHint(isFileSet)
			_showRemainingDisk(returnSize)
			returnSize, sgString = _askSizeOfDisk(returnSize, diskFullSize, partName)

			partInfos := []string{
				partName,
				sgString,
			}

			partCommands = append(partCommands, partInfos)
		}
	} else {
		for i := 0; i < (answer-1); i++ {
			x := i
			helpers.ClearConsole()
			_partHint(isFileSet)
			_showRemainingDisk(restSize)


			helpers.ClearConsole()
			_partHint(isFileSet)
			_showRemainingDisk(returnSize)
			partName := helpers.InputPrompt(fmt.Sprintf("What is the name of the partition %d?", (x+1)))

			helpers.ClearConsole()
			_partHint(isFileSet)
			_showRemainingDisk(returnSize)
			returnSize, sgString = _askSizeOfDisk(returnSize, diskFullSize, partName)

			partInfos := []string{
				partName,
				sgString,
			}

			partCommands = append(partCommands, partInfos)
		}
		helpers.ClearConsole()
		_partHint(isFileSet)
		_showRemainingDisk(returnSize)

		partName := helpers.InputPrompt("What is the name of the last partition?")
		partInfos := []string{
			partName,
			"0:0:0",
		}
		partCommands = append(partCommands, partInfos)
	}

	return partCommands

}


func setDiskPartVars(cfgFile string, isSWAPSet bool, swapSize int) {
	helpers.ClearConsole()
	helpers.PrintHeader("Startup", "Partitioning")
	var homePart, filePart, morePart bool = false, false, false
	var dSizePercent, restSize int
	var rootLabel, homeLabel, filesLabel, sgRoot, sgHome, sgFiles string
	var partLists [][]string

	dSize, err := strconv.Atoi(helpers.JsonGetter(cfgFile, "diskSize"))
	helpers.Check(err)

	if helpers.YesNo("Would you like to create a partition for the /home folder?") {
		homePart = true
	}

	fmt.Println(helpers.PrintBlack("Note: this partition will be the final one, taking the rest of the disk available space."))
	if helpers.YesNo("Would you like to create a partition for files under /media?") {
		filePart = true
	}

	if helpers.YesNo("Would you like to create more partitions?") {
		morePart = true
	}

	if dSize >= (500 * helpers.GB) {
		dSizePercent = int(float64(dSize/helpers.GB) * 0.20)
	} else {
		dSizePercent = int(float64(dSize/helpers.GB) * 0.40)
	}

	_restSize := dSize - dSizePercent
	if isSWAPSet {
		restSize = _restSize - swapSize
	}


	if homePart {
		restSize, rootLabel, sgRoot = _setRootPart(homePart, filePart, morePart, dSizePercent, restSize, dSize)
		restSize, homeLabel, sgHome = _setHomePart(filePart, morePart, dSizePercent, restSize, dSize)
		if filePart {
			restSize, filesLabel, sgFiles = _setFilePart(restSize)
			if morePart {
				partLists = _setMorePart(filePart, restSize, dSize)
			}

		} else {
			if morePart {
				partLists = _setMorePart(filePart, restSize, dSize)
			}
		}
	} else {
		if filePart {
			restSize, rootLabel, sgRoot = _setRootPart(homePart, filePart, morePart, dSizePercent, restSize, dSize)
			restSize, filesLabel, sgFiles = _setFilePart(restSize)
			if morePart {
				partLists = _setMorePart(filePart, restSize, dSize)
			}
		} else {
			if morePart {
				restSize, rootLabel, sgRoot = _setRootPart(homePart, filePart, morePart, dSizePercent, restSize, dSize)
				partLists = _setMorePart(filePart, restSize, dSize)
			} else {
				restSize, rootLabel, sgRoot = _setRootPart(homePart, filePart, morePart, dSizePercent, restSize, dSize)
			}
		}
	}

	fmt.Println(rootLabel, sgRoot)
	fmt.Println(homeLabel, sgHome)
	fmt.Println(filesLabel, sgFiles)
	fmt.Println(partLists)

/*
DRIVES ORDER:
0 - BOOT
1 - SWAP
2 - ROOT
3 - HOME
{...} - USER DRIVES
LAST - FILES
*/

}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func PreInstall() {
	var CONFIG_DIR string = fmt.Sprintf("%s/config", helpers.GetCurrDirPath())
	var CONFIG_FILE string = fmt.Sprintf("%s/config.json", CONFIG_DIR)
	//fmt.Println(CONFIG_FILE)
	fmt.Println("PreInstall")

	//cISO := setCountryISO()
	//helpers.JsonUpdater(CONFIG_FILE, "countryISO", cISO, false)

	// setTimeDate()
	//formatDisk(CONFIG_FILE)

	//bootSys := bootSystem()
	//helpers.JsonUpdater(CONFIG_FILE, "bootSystem", bootSys, false)

	//isSwapSet, swapSize := partitionDisk(CONFIG_FILE)

	diskSize, err := strconv.Atoi(helpers.JsonGetter(CONFIG_FILE, "diskSize"))
	helpers.Check(err)
	fmt.Println(diskSize)
	if diskSize > (128 * helpers.GB) {
		//setDiskPartVars(CONFIG_FILE, isSwapSet, swapSize)
		setDiskPartVars(CONFIG_FILE, true, (1 * helpers.GiB))
	}

}
