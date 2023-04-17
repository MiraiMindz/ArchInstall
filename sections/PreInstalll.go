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

type Country struct {
	Name   string
	Code   string
	Number int
}

func setCountryISO() string {
	iso := helpers.CurlResponse("https://ifconfig.co/country-iso")
	reflectorCountries := helpers.RunShellCommand(!helpers.COMMANDS_TEST_MODE, true, "reflector", "--list-countries")
	cLines := strings.Split(reflectorCountries, "\n")
	countries := []Country{}
	countriesOpts := []string{}
	var countriesMaps = make(map[string]Country)

	// Parse each line
	for i, line := range cLines {
		if i > 1 {
			fields := strings.Fields(line)
			if len(fields) >= 3 {
				// Extract country code from the second-to-last field
				code := fields[len(fields)-2]

				// Extract country name by joining all fields except the last two
				name := strings.Join(fields[:len(fields)-2], " ")

				country := Country{
					Name: name,
					Code: code,
				}

				// Parse integer from the last field
				var number int
				if _, err := fmt.Sscanf(fields[len(fields)-1], "%d", &number); err != nil {
					fmt.Printf("Failed to parse number from line: %s\n", line)
				} else {
					country.Number = number
					countries = append(countries, country)
				}
			}
		}
	}

	for _, v := range countries {
		x := fmt.Sprintf("[%s] - %s (%d)", v.Code, v.Name, v.Number)
		countriesOpts = append(countriesOpts, x)
		countriesMaps[x] = v
	}

	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Getting country ISO")
	prompt := fmt.Sprintf("Country ISO detected to be %s, is this correct?", strings.ReplaceAll(iso, "\n", ""))
	if helpers.YesNo(prompt) {
		return strings.ReplaceAll(iso, "\n", "")
	} else {
		fmt.Println(helpers.PrintHiBlack("[CODE] - NAME (NUMBER OF MIRRORS)"))
		_, _inputISO := helpers.PromptSelect("Select your country ISO", countriesOpts)
		inputISO := countriesMaps[_inputISO]
		return inputISO.Code
	}
}

func setTimeDate() {
	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Synchronizing hardware clock")
	helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "timedatectl", "set-ntp", "true")
}

func setupPacman(reflectorCountryISO, cfgFile string) {
	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Setting up PacMan")
	helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "pacman", "-S", "--noconfirm", "archlinux-keyring")
	helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "pacman", "-S", "--noconfirm", "--needed", "pacman-contrib")

	// parallelDowns := helpers.InputDefaultPrompt("Enter your desired number of Parallel Downloads", "5")

	//helpers.JsonUpdater(cfgFile, "pacmanParallelDownloads", parallelDowns)

	// helpers.ReplaceFileLine("/etc/pacman.conf", "#ParallelDownloads", fmt.Sprintf("ParallelDownloads = %s", parallelDowns))
	// helpers.ReplaceFileLine("/etc/pacman.conf", "#Color", "Color")
	// helpers.ReplaceFileLine("/etc/pacman.conf", "#CheckSpace", "CheckSpace")
	// helpers.ReplaceFileLine("/etc/pacman.conf", "#VerbosePkgLists", "VerbosePkgLists")

	// // helpers.ReplaceFileLine("/media/Arquivos/Programming/Projects/ArchInstall/pacman.conf", "#ParallelDownloads", fmt.Sprintf("ParallelDownloads = %s", parallelDowns))
	// // helpers.ReplaceFileLine("/media/Arquivos/Programming/Projects/ArchInstall/pacman.conf", "#Color", "Color")
	// // helpers.ReplaceFileLine("/media/Arquivos/Programming/Projects/ArchInstall/pacman.conf", "#CheckSpace", "CheckSpace")
	// // helpers.ReplaceFileLine("/media/Arquivos/Programming/Projects/ArchInstall/pacman.conf", "#VerbosePkgLists", "VerbosePkgLists")

	helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "pacman", "-S", "--noconfirm", "--needed", "reflector", "grub")

	//helpers.CopyFile("/etc/pacman.d/mirrorlist", "/etc/pacman.d/mirrorlist.backup")
	reflectorCountries := helpers.RunShellCommand(!helpers.COMMANDS_TEST_MODE, true, "reflector", "--list-countries")
	cLines := strings.Split(reflectorCountries, "\n")
	countries := []Country{}
	countriesOpts := []string{}
	selectedCountriesCodes := []string{}
	var countriesMaps = make(map[string]Country)

	// Parse each line
	for i, line := range cLines {
		if i > 1 {
			fields := strings.Fields(line)
			if len(fields) >= 3 {
				// Extract country code from the second-to-last field
				code := fields[len(fields)-2]

				// Extract country name by joining all fields except the last two
				name := strings.Join(fields[:len(fields)-2], " ")

				country := Country{
					Name: name,
					Code: code,
				}

				// Parse integer from the last field
				var number int
				if _, err := fmt.Sscanf(fields[len(fields)-1], "%d", &number); err != nil {
					fmt.Printf("Failed to parse number from line: %s\n", line)
				} else {
					country.Number = number
					countries = append(countries, country)
				}
			}
		}
	}

	for _, v := range countries {
		x := fmt.Sprintf("[%s] - %s (%d)", v.Code, v.Name, v.Number)
		countriesOpts = append(countriesOpts, x)
		countriesMaps[x] = v
	}

	for _, v := range countries {
		if v.Code == reflectorCountryISO {
			if helpers.YesNo(fmt.Sprintf("Your country has %d mirrors in total, do you want to select another country?", v.Number)) {
				fmt.Println(helpers.PrintHiBlack("[CODE] - NAME (NUMBER OF MIRRORS)"))
				selectedCountries := helpers.PromptMultiSelect("Select your desired countries", countriesOpts)
				for _, v := range selectedCountries {
					selectedCountriesCodes = append(selectedCountriesCodes, countriesMaps[v].Code)
				}
			}
		}
	}

	helpers.YesNo("")

	reflectorArgs := []string{"-a", "48", "-c", strings.Join(selectedCountriesCodes, ","), "-f", "5", "-l", "20", "--sort", "rate", "--save", "/etc/pacman.d/mirrorlist"}
	helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "reflector", reflectorArgs...)
	helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "pacman", "-S", "--noconfirm", "--needed", "gptfdisk")

	helpers.YesNo("")
}

func formatDisk(cfgFile string) {
	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Formatting Disk")

	disk := helpers.JsonGetter(cfgFile, "disk")
	helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "mkdir", "-pv", "/mnt")
	helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "umount", "-A", "--recursive", "/mnt")
	helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "sgdisk", "-Z", disk)
	helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "sgdisk", "-a", "2048", "-o", disk)
	helpers.YesNo("")
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
	fmt.Println(helpers.PrintHiBlack(remainDiskSpace))
}

func _askSizeOfDisk(restSize, diskFullSize int, partName string) (int, string) {
	var retSize int
	var sgStr string

	fmt.Println(helpers.PrintHiBlack("use 'G' for Giga, 'M' for Mega and 'K' for Kilo, example: 400M"))
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
	helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "sgdisk", "-n", sgString, tCode, pName, disk)
}

func _createFileSystems(partitionName, partDisk, filesystem string) {
	switch strings.ToLower(filesystem) {
	case "btrfs":
		helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "mkfs.btrfs", "-L", partitionName, "-f")
		break
	case "ext4":
		helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "mkfs.ext4", "-L", partitionName)
		break
	case "swp":
		helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "mkswap", partDisk)
		helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "swapon", partDisk)
		break
	case "fat":
		helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "mkfs.fat", "-F32", "-n", partitionName, partDisk)
		break
	case "vfat":
		helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "mkfs.vfat", "-F32", "-n", partitionName, partDisk)
		break
	}
}

func _setRootPart(isHomeSet, isFileSet, isCustomSet, isX64 bool, percentSize, restSize, diskFullSize int) (int, string, string, string, string) {
	var rootSize, returnSize int
	var sgString, selectedFS, partitionCode string
	partitionName := "ROOT"

	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Partitioning")
	if isHomeSet && (isFileSet || isCustomSet) {
		rootSize, _ = _calcRootHomeSize(percentSize)
		rootSizeBytes := helpers.ByteSizeConverter(uint64(rootSize))
		helpers.ClearConsole()
		helpers.PrintHeader("Pre-Install", "Partitioning")
		promptString := fmt.Sprintf("Based on you drive size, the scripts recommends a size of %s to %s, would you like to use it?", rootSizeBytes, partitionName)
		if helpers.YesNo(promptString) {
			sgString = fmt.Sprintf("0:0:+%s", helpers.ParseSizeString(rootSizeBytes))
			returnSize = restSize - rootSize
		} else {
			helpers.ClearConsole()
			helpers.PrintHeader("Pre-Install", "Partitioning")
			_showRemainingDisk(restSize)
			returnSize, sgString = _askSizeOfDisk(restSize, diskFullSize, partitionName)
		}
	} else if isHomeSet && !isFileSet && !isCustomSet {
		rootSizeBytes := helpers.ByteSizeConverter(uint64(percentSize * helpers.GiB))
		helpers.ClearConsole()
		helpers.PrintHeader("Pre-Install", "Partitioning")
		promptString := fmt.Sprintf("Based on you drive size, the scripts recommends a size of %s to %s, would you like to use it?", rootSizeBytes, partitionName)
		if helpers.YesNo(promptString) {
			sgString = fmt.Sprintf("0:0:+%s", helpers.ParseSizeString(rootSizeBytes))
			returnSize = restSize - (percentSize * helpers.GiB)
		} else {
			helpers.ClearConsole()
			helpers.PrintHeader("Pre-Install", "Partitioning")
			_showRemainingDisk(restSize)
			returnSize, sgString = _askSizeOfDisk(restSize, diskFullSize, partitionName)
		}
	} else if !isHomeSet && (isFileSet || isCustomSet) {
		rootSizeBytes := helpers.ByteSizeConverter(uint64(percentSize * helpers.GiB))
		helpers.ClearConsole()
		helpers.PrintHeader("Pre-Install", "Partitioning")
		promptString := fmt.Sprintf("Based on you drive size, the scripts recommends a size of %s to %s, would you like to use it?", rootSizeBytes, partitionName)
		if helpers.YesNo(promptString) {
			sgString = fmt.Sprintf("0:0:+%s", helpers.ParseSizeString(rootSizeBytes))
			returnSize = restSize - (percentSize * helpers.GiB)
		} else {
			helpers.ClearConsole()
			helpers.PrintHeader("Pre-Install", "Partitioning")
			_showRemainingDisk(restSize)
			returnSize, sgString = _askSizeOfDisk(restSize, diskFullSize, partitionName)
		}
	} else {
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

func _setHomePart(isFileSet, isCustomSet bool, percentSize, restSize, diskFullSize int) (int, string, string, string, string, string) {
	var homeSize, returnSize int
	var sgString, selectedFS string
	partitionName := "HOME"
	mountPoint := "home"

	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Partitioning")

	if !isFileSet && !isCustomSet {
		returnSize = restSize
		sgString = "0:0:0"
	} else if isFileSet || isCustomSet {
		_, homeSize = _calcRootHomeSize(percentSize)
		bytesHomeSize := helpers.ByteSizeConverter(uint64(homeSize))
		helpers.ClearConsole()
		helpers.PrintHeader("Pre-Install", "Partitioning")
		promptString := fmt.Sprintf("Based on you drive size, the scripts recommends a size of %s to %s, would you like to use it?", bytesHomeSize, partitionName)
		if helpers.YesNo(promptString) {
			sgString = fmt.Sprintf("0:0:+%s", helpers.ParseSizeString(bytesHomeSize))
			returnSize = restSize - (percentSize * helpers.GiB)
		} else {
			helpers.ClearConsole()
			helpers.PrintHeader("Pre-Install", "Partitioning")
			_showRemainingDisk(restSize)
			returnSize, sgString = _askSizeOfDisk(restSize, diskFullSize, partitionName)
		}
	}

	selectedFS = _askPartitionFileSystem(partitionName)
	return returnSize, partitionName, sgString, selectedFS, "8302", mountPoint
}

func _setFilePart(restSize int) (int, string, string, string, string, string) {
	partitionName := "FILES"
	mountPoint := "media/Files"
	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Partitioning")
	selectedFS := _askPartitionFileSystem(partitionName)
	return restSize, partitionName, "0:0:0", selectedFS, "8300", mountPoint
}

func _partHint(isFileSet bool) {
	if isFileSet {
		a := helpers.PrintYellow("Reminder:")
		b := helpers.PrintHiBlack("You had selected to use a files partition")
		c := helpers.PrintHiBlack("that partition will be used to fill the empty space in your disk.")
		d := helpers.PrintHiBlack("So don't worry on filling up the remaining space with these partitions.")
		fmt.Printf("%s %s %s\n%s\n", a, b, c, d)
	} else {
		z := helpers.PrintYellow("Note:")
		x := helpers.PrintHiBlack("The last partition will be used to fill the remaining space in your disk.")
		fmt.Printf("%s %s\n", z, x)
	}
}

func _setMorePart(isFileSet bool, restSize, diskFullSize int) [][]string {
	var partCommands [][]string
	var returnSize int = restSize
	var sgString, selectedFS string
	var baseMountPoint string = "media"

	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Partitioning")
	_partHint(isFileSet)
	answer, err := strconv.Atoi(helpers.InputPrompt("How many partitions do you want to create?"))
	helpers.Check(err)

	if isFileSet {
		for i := 0; i < answer; i++ {
			x := i
			helpers.ClearConsole()
			helpers.PrintHeader("Pre-Install", "Partitioning")
			_partHint(isFileSet)
			_showRemainingDisk(returnSize)
			partName := _askPartitionName(fmt.Sprintf("What is the name of the partition %d?", (x + 1)))
			helpers.ClearConsole()
			helpers.PrintHeader("Pre-Install", "Partitioning")

			_partHint(isFileSet)
			_showRemainingDisk(returnSize)
			returnSize, sgString = _askSizeOfDisk(returnSize, diskFullSize, partName)
			selectedFS = _askPartitionFileSystem(partName)
			partMountPoint := fmt.Sprintf("%s/%s", baseMountPoint, partName)

			partInfos := []string{
				partName,
				sgString,
				selectedFS,
				"8300",
				partMountPoint,
			}

			partCommands = append(partCommands, partInfos)
		}
	} else {
		for i := 0; i < (answer - 1); i++ {
			x := i
			helpers.ClearConsole()
			helpers.PrintHeader("Pre-Install", "Partitioning")
			_partHint(isFileSet)
			_showRemainingDisk(restSize)

			helpers.ClearConsole()
			helpers.PrintHeader("Pre-Install", "Partitioning")
			_partHint(isFileSet)
			_showRemainingDisk(returnSize)
			partName := _askPartitionName(fmt.Sprintf("What is the name of the partition %d?", (x + 1)))

			helpers.ClearConsole()
			helpers.PrintHeader("Pre-Install", "Partitioning")
			_partHint(isFileSet)
			_showRemainingDisk(returnSize)
			returnSize, sgString = _askSizeOfDisk(returnSize, diskFullSize, partName)
			selectedFS = _askPartitionFileSystem(partName)
			partMountPoint := fmt.Sprintf("%s/%s", baseMountPoint, partName)

			partInfos := []string{
				partName,
				sgString,
				selectedFS,
				"8300",
				partMountPoint,
			}

			partCommands = append(partCommands, partInfos)
		}
		helpers.ClearConsole()
		helpers.PrintHeader("Pre-Install", "Partitioning")
		_partHint(isFileSet)
		_showRemainingDisk(returnSize)

		partName := _askPartitionName("What is the name of the last partition?")
		selectedFS = _askPartitionFileSystem(partName)
		partMountPoint := fmt.Sprintf("%s/%s", baseMountPoint, partName)

		partInfos := []string{
			partName,
			"0:0:0",
			selectedFS,
			"8300",
			partMountPoint,
		}
		partCommands = append(partCommands, partInfos)
	}

	return partCommands

}

func setDiskPartVars(cfgFile, disk string, isSWAPSet bool, swapSize int) [][]string {
	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Partitioning")
	var (
		homePart, filePart, morePart, isX64 bool = false, false, false, false
		rootLabel, homeLabel, filesLabel,
		sgRoot, sgHome, sgFiles,
		rootFS, homeFS, filesFS,
		rootPartCode, homePartCode, filesPartCode,
		homeMountPoint, filesMountPoint string
		dSizePercent, restSize   int
		commandsOrder, partLists [][]string
	)

	dSize, err := strconv.Atoi(helpers.JsonGetter(cfgFile, "diskSize"))
	helpers.Check(err)

	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Partitioning")
	if helpers.YesNo("Would you like to create a partition for the /home folder?") {
		homePart = true
	}

	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Partitioning")
	fmt.Println(helpers.PrintHiBlack("Note: this partition will be the final one, taking the rest of the disk available space."))
	if helpers.YesNo("Would you like to create a partition for files under /media?") {
		filePart = true
	}

	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Partitioning")
	if helpers.YesNo("Would you like to create more partitions?") {
		morePart = true
	}

	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Partitioning")
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
		restSize, homeLabel, sgHome, homeFS, homePartCode, homeMountPoint = _setHomePart(filePart, morePart, dSizePercent, restSize, dSize)
		if filePart {
			restSize, filesLabel, sgFiles, filesFS, filesPartCode, filesMountPoint = _setFilePart(restSize)
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
			restSize, filesLabel, sgFiles, filesFS, filesPartCode, homeMountPoint = _setFilePart(restSize)
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
		homeMountPoint,
	}

	filePartData := []string{
		filesLabel,
		sgFiles,
		filesFS,
		filesPartCode,
		filesMountPoint,
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
		commandsOrder = append(commandsOrder, []string{"EFIBOOT", "0:0:+300M", "vfat", "ef00"})
	}

	memInfo := helpers.GetLine("/proc/meminfo", "MemTotal")
	totalMem := helpers.ExtractNumbers(memInfo.(string))[0] * helpers.KiB

	if totalMem < (8 * helpers.GiB) {
		rSize, swapSize := _getRecommendedSwapSize(totalMem)
		if rSize > 0 {
			swpSize = rSize * helpers.GiB
			helpers.ClearConsole()
			helpers.PrintHeader("Pre-Install", "Partitioning")
			fmt.Println("The script detected that you have less than 8GiB of RAM")
			fmt.Printf("Based on your RAM size of %s, the script recommends a SWAP size of %s\n",
				helpers.ByteSizeConverter(uint64(totalMem)), helpers.ByteSizeConverter(uint64(rSize*helpers.GiB)))
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

	helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "partprobe", disk)

	fmt.Println(commandsOrder)
	helpers.YesNo("")

	if len(commandsOrder) > 0 {
		for _, v := range commandsOrder {
			if len(v) > 0 {
				_createPartitions(v[0], v[1], v[3], disk)
			}
		}

		for i, v := range commandsOrder {
			x := i
			if len(v) > 0 {
				_createFileSystems(v[0], fmt.Sprintf("%s%d", disk, x+1), v[2])
			}
		}

		for _, v := range commandsOrder {
			if len(v) > 0 {
				if (v[0] != "ROOT") && (v[0] != "BIOSBOOT") && (v[0] != "EFIBOOT") {
					helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "mkdir", "-pv", fmt.Sprintf("/mnt/%s", v[4]))
				}
				if v[0] == "EFIBOOT" {
					helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "mkdir", "-pv", "/mnt/boot/efi")
				}
			}
		}
	}

	for _, v := range commandsOrder {
		//x := i
		if len(v) > 0 {
			if (v[0] != "ROOT") && (v[0] != "BIOSBOOT") && (v[0] != "EFIBOOT") { // fmt.Sprintf("/mnt/%s", v[4])
				helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "mount", "-v", "-L", v[0], fmt.Sprintf("/mnt/%s", v[4]))
			}
			if v[0] == "EFIBOOT" {
				helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "mount", "-v", "-t", v[2], "-L", "EFIBOOT", "/mnt/boot/")
			}
		}
	}

	if helpers.GetLine("/proc/mounts", "/mnt") == nil {
		helpers.ClearConsole()
		helpers.PrintHeader("Pre-Install", "Partitioning")
		fmt.Println(helpers.PrintError("Driver is not mounted, can't continue"))
		helpers.CountDown(10, "Rebooting")
		helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "reboot", "now")
	}

}

//endregion

func installArch() {
	kernelsOpts := []helpers.ItemInfo{
		{Item: "stable", Info: "Vanilla Linux kernel and modules, with a few patches applied."},
		{Item: "hardened", Info: "A security-focused Linux kernel applying a set of hardening patches. It also enables more upstream hardening features than the default kernel."},
		{Item: "long-term", Info: "Long-term support (LTS) Linux kernel and modules."},
		{Item: "real-time", Info: "This patch allows nearly all of the kernel to be preempted, with the exception of a few very small regions of code."},
		{Item: "real-time long-term", Info: "The same as the Real-Time kernel but with long-term support."},
		{Item: "zen", Info: "Result of a collaborative effort of kernel hackers to provide the best Linux kernel possible for everyday systems."},
	}
	textEditorOpts := []helpers.ItemInfo{
		{Item: "nano", Info: "Console text editor based on pico with on-screen key bindings help."},
		{Item: "vim", Info: "Advanced text editor that seeks to provide the power of the de-facto Unix editor 'vi', with a more complete feature set."},
		{Item: "emacs", Info: "The extensible, customizable, self-documenting real-time display editor by GNU."},
	}

	pkgs := []string{
		"base",
		"base-devel",
		"linux-firmware",
		"archlinux-keyring",
		"wget",
		"sudo",
	}

	command := []string{
		"/mnt",
	}

	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Installing Base System")
	_, kernelChoice := helpers.PromptSelectInfo("Select your desired kernel", kernelsOpts)
	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Installing Base System")
	_, textEditorChoice := helpers.PromptSelectInfo("Select your desired TUI Text Editor", textEditorOpts)

	switch kernelChoice {
	case "stable":
		pkgs = append(pkgs, "linux")
	case "hardened":
		pkgs = append(pkgs, "linux-hardened")
	case "long-term":
		pkgs = append(pkgs, "linux-lts")
	case "real-time":
		pkgs = append(pkgs, "linux-rt")
	case "real-time long-term":
		pkgs = append(pkgs, "linux-rt-lts")
	case "zen":
		pkgs = append(pkgs, "linux-zen")
	default:
		pkgs = append(pkgs, "linux")
	}

	switch textEditorChoice {
	case "nano":
		pkgs = append(pkgs, "nano")
	case "vim":
		pkgs = append(pkgs, "vim")
	case "emacs":
		pkgs = append(pkgs, "emacs")
	default:
		pkgs = append(pkgs, "nano")
	}

	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Installing Base System")
	fmt.Println(helpers.PrintHiBlack(fmt.Sprintf("Current Packages are %s.", pkgs)))
	if helpers.YesNo("Would you like to add more packages?") {
		helpers.ClearConsole()
		helpers.PrintHeader("Pre-Install", "Installing Base System")
		fmt.Println(helpers.PrintHiBlack(fmt.Sprintf("Current Packages are %s.", pkgs)))
		usrPkgs := helpers.InputPrompt("Enter the desired packages")
		usrPkgsSplit := strings.Fields(usrPkgs)
		pkgs = append(pkgs, usrPkgsSplit...)
	}

	command = append(command, pkgs...)
	command = append(command, "--noconfirm")
	command = append(command, "--needed")

	helpers.ClearConsole()
	helpers.PrintHeader("Pre-Install", "Installing Base System")
	fmt.Println(helpers.PrintHiBlack(fmt.Sprintf("Current Packages are %s.", pkgs)))
	fmt.Println(helpers.PrintHiBlack("If there was any typo in the addition of custom packages, say \"NO\" to re-run this part."))
	if helpers.YesNo("Would you like to proceed with the installation?") {
		helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "pacstrap", command...)
	} else {
		installArch()
	}
}

func copyNecessaryFiles(cfgFile string) {
	installLoc := helpers.JsonGetter(cfgFile, "installLocation")
	helpers.CopyFile("/etc/pacman.d/mirrorlist", "/mnt/etc/pacman.d/mirrorlist")
	helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "cp", "-Rv", installLoc, "/mnt/root/ArchInstall") // Too lazy to implement a CopyFolder function.

}

func generateFileSystemTable() {
	cmdOut := helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "genfstab", "-L", "/mnt")
	println(cmdOut)
	if helpers.YesNo("Is this correct?") {
		helpers.WriteToFile("/mnt/etc/fstab", cmdOut, 0644)
	} else {
		helpers.CountDown(5, "Rebooting")
		helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "reboot", "now")
	}

}

func installBootLoader(cfgFile string) {
	bootType := helpers.JsonGetter(cfgFile, "bootSystem")
	disk := helpers.JsonGetter(cfgFile, "disk")
	if strings.ToLower(bootType) == "bios" {
		helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "grub-install", "--target=i386-pc", "--boot-directory=/mnt/boot", disk)
	} else {
		helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "pacstrap", "/mnt", "efibootmgr", "--noconfirm", "--needed")
	}
}

func PreInstalll() {
	var CONFIG_DIR string = fmt.Sprintf("%s/config", helpers.GetCurrDirPath())
	var CONFIG_FILE string = fmt.Sprintf("%s/config.json", CONFIG_DIR)

	cISO := setCountryISO()
	helpers.JsonUpdater(CONFIG_FILE, "countryISO", cISO, false)

	setTimeDate()
	formatDisk(CONFIG_FILE)

	setupPacman(cISO, CONFIG_FILE)

	bootSys := bootSystem()
	helpers.JsonUpdater(CONFIG_FILE, "bootSystem", bootSys, false)

	formatDisk(CONFIG_FILE)
	partitionDisk(CONFIG_FILE)
	installArch()
	copyNecessaryFiles(CONFIG_FILE)
}
