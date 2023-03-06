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
	//"os"
	//"os/exec"
	"strconv"
	"strings"
)

func setCountryISO() string {
	iso := helpers.CurlResponse("https://ifconfig.co/country-iso")
	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Getting country ISO")
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
	//cmd := helpers.RunShellCommand("timedatectl", "set-ntp", "true")
	//fmt.Println(cmd)

	fmt.Printf("%s %s %s\n", "timedatectl", "set-ntp", "true")
}

func setupPacman(reflectorCountryISO string) {
	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Setting up PacMan")
	//cmd := helpers.RunShellCommand("pacman", "-S", "--noconfirm", "archlinux-keyring")
	//fmt.Print(cmd)
	//cmd = helpers.RunShellCommand("pacman", "-S", "--noconfirm", "--needed", "pacman-contrib")
	//fmt.Println(cmd)
	//helpers.ReplaceFileLine("/etc/pacman.conf", "#ParallelDownloads", "ParallelDownloads")
	//cmd = helpers.RunShellCommand("pacman", "-S", "--noconfirm", "--needed", "reflector", "grub")
	//fmt.Println(cmd)

	//helpers.CopyFile("/etc/pacman.d/mirrorlist", "/etc/pacman.d/mirrorlist.backup")
	//reflectorArgs := []string{	"-a", "48", "-c", reflectorCountryISO, "-f", "5",
	//							"-l", "20", "--sort", "rate", "--save", "/etc/pacman.d/mirrorlist"}
	//cmd = helpers.RunShellCommand("reflector", reflectorArgs...)
	//fmt.Println(cmd)
	//cmd = helpers.RunShellCommand("mkdir", "/mnt")
	//fmt.Println(cmd)

	//cmd = helpers.RunShellCommand("pacman", "-S", "--noconfirm", "--needed", "gptfdisk")
	//fmt.Println(cmd)

	fmt.Printf("%s %s %s %s\n", "pacman", "-S", "--noconfirm", "archlinux-keyring")
	fmt.Printf("%s %s %s %s %s\n", "pacman", "-S", "--noconfirm", "--needed", "pacman-contrib")
	fmt.Printf("%s %s %s\n", "/etc/pacman.conf", "#ParallelDownloads", "ParallelDownloads")
	fmt.Printf("%s %s %s %s %s %s\n", "pacman", "-S", "--noconfirm", "--needed", "reflector", "grub")
	fmt.Printf("%s %s\n", "/etc/pacman.d/mirrorlist", "/etc/pacman.d/mirrorlist.backup")
	fmt.Printf("%s %s %s %s %s %s %s %s %s %s %s %s %s\n", "reflector", "-a", "48", "-c", reflectorCountryISO, "-f", "5", "-l", "20", "--sort", "rate", "--save", "/etc/pacman.d/mirrorlist")
	fmt.Printf("%s %s\n", "mkdir", "/mnt")
	fmt.Printf("%s %s %s %s %s\n", "pacman", "-S", "--noconfirm", "--needed", "gptfdisk")

}

func formatDisk(cfgFile string) {
	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Formatting Disk")

	disk := helpers.JsonGetter(cfgFile, "disk")
	//helpers.RunShellCommand("umount", "-A", "--recursive", "/mnt")
	//helpers.RunShellCommand("sgdisk", "-Z", disk)
	//helpers.RunShellCommand("sgdisk", "-a", "2048", "-o", disk)
	fmt.Printf("%s %s %s %s\n", "umount", "-A", "--recursive", "/mnt")
	fmt.Printf("%s %s %s\n", "sgdisk", "-Z", disk)
	fmt.Printf("%s %s %s %s %s\n", "sgdisk", "-a", "2048", "-o", disk)
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



//region disk
//// I think it's cleaner now

func _getRecommendedSwapSize(totalRam int) (int, string) {
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
	return swapRecommendedSize, fmt.Sprintf("0:0:+%dG", swapRecommendedSize)
}

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
	userSize := helpers.InputPrompt(fmt.Sprintf("Enter the size of the \"%s\" partition", partName))
	if strings.Contains(userSize, ".") {
		helpers.ClearConsole()
		fmt.Println(helpers.PrintRed("Can't use decimal places"))
		_showRemainingDisk(restSize)
		retSize, sgStr = _askSizeOfDisk(restSize, diskFullSize, partName)
	} else {
		userByteSize := helpers.ConvertToByte(userSize)
		if (diskFullSize - userByteSize) <= 0 {
			helpers.ClearConsole()
			fmt.Println(helpers.PrintRed("Can't exceed drive physical space"))
			_showRemainingDisk(restSize)
			retSize, sgStr = _askSizeOfDisk(restSize, diskFullSize, partName)
		} else {
			retSize = restSize - userByteSize
			sgStr = fmt.Sprintf("0:0:+%s", helpers.ParseSizeString(userSize))
		}
	}

	return retSize, sgStr
}

func _askPartitionFileSystem(partName string) string {
	prompt := fmt.Sprintf("Select the file system for the partition %s", partName)
	fileSystems := []string{
		"ext4",
		"btrfs",
	}
	_, selectedFileSystem := helpers.PromptSelect(prompt, fileSystems)

	return selectedFileSystem
}

func _askPartitionName(prompt string) string {
	partName := helpers.InputPrompt(prompt)
	if len(partName) <= 0 {
		fmt.Println(helpers.PrintRed("Partition name can't be blank"))
		partName = _askPartitionName(prompt)
	}
	return partName
}


func _createPartitions(partitionName, sgString, partitionCode, disk string) {
	tCode := fmt.Sprintf("--typecode=0:%s", partitionCode)
	pName := fmt.Sprintf("--change-name=0:'%s'", partitionName)
	fmt.Printf("%s %s %s %s %s %s\n", "sgdisk", "-n", sgString, tCode, pName, disk)

	//cmd := helpers.RunShellCommand("sgdisk", "-n", sgString, tCode, pName, disk)
	//fmt.Println(cmd)


	// sgdisk -n 3::-0 --typecode=3:8300 --change-name=3:'ROOT' ${DISK}
	// mkfs.btrfs -L ROOT ${partition3} -f
	// mkfs.ext4 -L ROOT ${partition3}
}

func _createFileSystems(partitionName, partDisk, filesystem string) {
	if strings.ToLower(filesystem) == "btrfs" {
		//cmd = helpers.RunShellCommand("mkfs.btrfs", "-L", partitionName, "-f")
		//fmt.Println(cmd)
		fmt.Printf("%s %s %s %s %s\n","mkfs.btrfs", "-L", partitionName, partDisk, "-f")
	} else if strings.ToLower(filesystem) == "ext4" {
		//cmd = helpers.RunShellCommand("mkfs.ext4", "-L", partitionName)
		//fmt.Println(cmd)

		fmt.Printf("%s %s %s %s\n","mkfs.ext4", "-L", partitionName, partDisk)

	} else if strings.ToLower(filesystem) == "swp" {
		fmt.Printf("%s %s\n","mkswap", partDisk)
		fmt.Printf("%s %s\n","swapon", partDisk)
	}
}

func _setRootPart(isHomeSet, isFileSet, isCustomSet, isX64 bool, percentSize, restSize, diskFullSize int) (int, string, string, string, string) {
	var rootSize, returnSize int
	var sgString, selectedFS, partitionCode string
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

	selectedFS = _askPartitionFileSystem(partitionName)
	if isX64 {
		partitionCode = "8303"
	} else {
		partitionCode = "8304"
	}

	return returnSize, partitionName, sgString, selectedFS, partitionCode
}

func _setHomePart(isFileSet, isCustomSet bool, percentSize, restSize, diskFullSize int) (int, string, string, string, string) {
	var homeSize, returnSize int
	var sgString, selectedFS string
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

	selectedFS = _askPartitionFileSystem(partitionName)
	return returnSize, partitionName, sgString, selectedFS, "8302"
}

func _setFilePart(restSize int) (int, string, string, string, string) {
	partitionName := "FILES"
	selectedFS := _askPartitionFileSystem(partitionName)
	return restSize, partitionName, "0:0:0", selectedFS, "8300"
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
	var sgString, selectedFS string

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
			partName := _askPartitionName(fmt.Sprintf("What is the name of the partition %d?", (x+1)))
			helpers.ClearConsole()

			_partHint(isFileSet)
			_showRemainingDisk(returnSize)
			returnSize, sgString = _askSizeOfDisk(returnSize, diskFullSize, partName)
			selectedFS = _askPartitionFileSystem(partName)

			partInfos := []string{
				partName,
				sgString,
				selectedFS,
				"8300",
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
			partName := _askPartitionName(fmt.Sprintf("What is the name of the partition %d?", (x+1)))

			helpers.ClearConsole()
			_partHint(isFileSet)
			_showRemainingDisk(returnSize)
			returnSize, sgString = _askSizeOfDisk(returnSize, diskFullSize, partName)
			selectedFS = _askPartitionFileSystem(partName)

			partInfos := []string{
				partName,
				sgString,
				selectedFS,
				"8300",
			}

			partCommands = append(partCommands, partInfos)
		}
		helpers.ClearConsole()
		_partHint(isFileSet)
		_showRemainingDisk(returnSize)

		partName := _askPartitionName("What is the name of the last partition?")
		selectedFS = _askPartitionFileSystem(partName)
		partInfos := []string{
			partName,
			"0:0:0",
			selectedFS,
			"8300",
		}
		partCommands = append(partCommands, partInfos)
	}

	return partCommands

}

func setDiskPartVars(cfgFile, disk string, isSWAPSet bool, swapSize int) [][]string {
	helpers.ClearConsole()
	helpers.PrintHeader("Startup", "Partitioning")
	var (
		homePart, filePart, morePart, isX64 bool = false, false, false, false
		rootLabel, homeLabel, filesLabel,
		sgRoot, sgHome, sgFiles,
		rootFS, homeFS, filesFS,
		rootPartCode, homePartCode, filesPartCode string
		dSizePercent, restSize int
		commandsOrder, partLists [][]string
	)


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

	if helpers.YesNo("Are you running an 64-bit system?") {
		isX64 = true
	}

	if dSize >= (500 * helpers.GB) {
		dSizePercent = int(float64(dSize/helpers.GB) * 0.20)
	} else {
		dSizePercent = int(float64(dSize/helpers.GB) * 0.40)
	}

	_restSize := dSize - dSizePercent
	if isSWAPSet {
		restSize = _restSize - swapSize
	} else {
		restSize = _restSize
	}

	if homePart {
		restSize, rootLabel, sgRoot, rootFS, rootPartCode = _setRootPart(homePart, filePart, morePart, isX64, dSizePercent, restSize, dSize)
		restSize, homeLabel, sgHome, homeFS, homePartCode = _setHomePart(filePart, morePart, dSizePercent, restSize, dSize)
		if filePart {
			restSize, filesLabel, sgFiles, filesFS, filesPartCode = _setFilePart(restSize)
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
			restSize, rootLabel, sgRoot, rootFS, rootPartCode = _setRootPart(homePart, filePart, morePart, isX64, dSizePercent, restSize, dSize)
			restSize, filesLabel, sgFiles, filesFS, filesPartCode = _setFilePart(restSize)
			if morePart {
				partLists = _setMorePart(filePart, restSize, dSize)
			}
		} else {
			if morePart {
				restSize, rootLabel, sgRoot, rootFS, rootPartCode = _setRootPart(homePart, filePart, morePart, isX64, dSizePercent, restSize, dSize)
				partLists = _setMorePart(filePart, restSize, dSize)
			} else {
				restSize, rootLabel, sgRoot, rootFS, rootPartCode = _setRootPart(homePart, filePart, morePart, isX64, dSizePercent, restSize, dSize)
			}
		}
	}

	rootPartData := []string{
		rootLabel,
		sgRoot,
		rootFS,
		rootPartCode,
	}

	homePartData := []string{
		homeLabel,
		sgHome,
		homeFS,
		homePartCode,
	}

	filePartData := []string{
		filesLabel,
		sgFiles,
		filesFS,
		filesPartCode,
	}

	fmt.Println(rootPartData)
	fmt.Println(homePartData)
	fmt.Println(filePartData)
	fmt.Println(partLists)

	if len(rootPartData) > 0 {
		commandsOrder = append(commandsOrder, rootPartData)
	}

	if homePart && len(homePartData) > 0 {
		commandsOrder = append(commandsOrder, homePartData)
	}

	if morePart && len(partLists) > 0 {
		for _, v := range partLists {
			if len(v) > 0 {
				commandsOrder = append(commandsOrder, v)
			}
		}
	}

	if filePart && len(filePartData) > 0 {
		commandsOrder = append(commandsOrder, filePartData)
	}


	fmt.Println(commandsOrder)

	return commandsOrder
}


func partitionDisk(cfgFile string) {
	var commandsOrder [][]string
	isSWAPSet := false
	swpSize := 0
	bSys := helpers.JsonGetter(cfgFile, "bootSystem")
	disk := helpers.JsonGetter(cfgFile, "disk")
	if bSys == "BIOS" {
		commandsOrder = append(commandsOrder, []string{"BIOSBOOT", "0:0:+1M", "none", "ef02"})
	} else {
		commandsOrder = append(commandsOrder, []string{"EFIBOOT", "0:0:+300M", "none", "ef00"})
	}


	memInfo := helpers.GetLine("/proc/meminfo", "MemTotal")
	totalMem := helpers.ExtractNumbers(memInfo.(string))[0] * helpers.KiB

	if totalMem < (8 * helpers.GiB) {
		rSize, swapSize := _getRecommendedSwapSize(totalMem)
		if rSize > 0 {
			swpSize = rSize * helpers.GiB
			fmt.Println("The script detected that you have less than 8GiB of RAM")
			fmt.Printf("Based on your RAM size of %s, the script recommends a SWAP size of %s\n",
			helpers.ByteSizeConverter(uint64(totalMem)), helpers.ByteSizeConverter(uint64(rSize * helpers.GiB)))
			if helpers.YesNo("Would you like to create a SWAP partition?") {
				commandsOrder = append(commandsOrder, []string{"SWAP", swapSize, "swp", "8200"})
				isSWAPSet = true
			}
		}
	}

	fmt.Println(swpSize, isSWAPSet)
	diskSize, err := strconv.Atoi(helpers.JsonGetter(cfgFile, "diskSize"))
	helpers.Check(err)
	if diskSize > (128 * helpers.GB) {
		cOrd := setDiskPartVars(cfgFile, disk, isSWAPSet, swpSize)
		commandsOrder = append(commandsOrder, cOrd...)
	}




	//cmd := helpers.RunShellCommand("partprobe", disk)
	//fmt.Println(cmd)




	if len(commandsOrder) > 0 {
		for _, v := range commandsOrder {
			if len(v) > 0 {
				_createPartitions(v[0], v[1], v[3], disk)
			}
		}

		for i, v := range commandsOrder {
			x := i
			if len(v) > 0 {
				_createFileSystems(v[0], fmt.Sprintf("%s%d",disk,x+1), v[2])
			}
		}
	}


}

//endregion


func PreInstall() {
	var CONFIG_DIR string = fmt.Sprintf("%s/config", helpers.GetCurrDirPath())
	var CONFIG_FILE string = fmt.Sprintf("%s/config.json", CONFIG_DIR)
	//fmt.Println(CONFIG_FILE)
	fmt.Println("PreInstall")

	cISO := setCountryISO()
	helpers.JsonUpdater(CONFIG_FILE, "countryISO", cISO, false)

	setTimeDate()
	formatDisk(CONFIG_FILE)

	bootSys := bootSystem()
	helpers.JsonUpdater(CONFIG_FILE, "bootSystem", bootSys, false)

	partitionDisk(CONFIG_FILE)

}
