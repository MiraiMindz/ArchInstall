package sections

/*******************************************************************************
THIS  FILE  ONLY EXISTS  FOR TESTING, THE  CORRECT FILE IS LOCATED AT ./Startup/
THIS  WAS  DONE  BECAUSE I WOULD  NEED TO  ALL PARTS AS  STANDALONE  BINARIES TO
EXECUTE  IN THE MAIN SCRIPT  ON THE INSTALLATION PROCESS, AND  TO DO THAT I NEED
TO CREATE ANOTHER MODULE. THEY ARE IDENTICAL BTW. - MIRAI
*******************************************************************************/


import (
	"ArchInstall/helpers"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/jaypipes/ghw"
)

const sectionName string = "Startup"

func createConfigFile(filePath string) {
	file_loc := fmt.Sprintf("%s/config.json", filePath)
	if !helpers.CheckFileExists(file_loc) {
		if !helpers.CheckFileExists(filePath) {
			os.Mkdir(filePath, 0755)
		}
		err := os.WriteFile(file_loc, []byte(""), 0644)
		if err != nil {
			panic(err)
		}
	}
}

func checkRootUser() {
	currUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	if currUser.Uid != "0" {
		panic("Root user is required to run this script.")
	}
}

func checkOS() {
	if helpers.CheckFileExists("/etc/arch-release") {
		panic("This script is made to run on ArchLinux.")
	}
}

func checkPacman() {
	if helpers.CheckFileExists("/var/lib/pacman/db.lck") {
		panic("Pacman is locked.")
	}
}

func checks() {
	checkRootUser()
	checkOS()
	checkPacman()
}

func setPassword() string {
	var passwd1, passwd2 string
	helpers.ClearConsole()
	helpers.PrintHeader(sectionName, "Password")

	passwd1 = helpers.PromptReadPassword("Enter your Password")
	passwd2 = helpers.PromptReadPassword("Re-type your Password")
	if passwd1 == passwd2 {
		fmt.Println("Passwords do match.")
	} else {
		fmt.Println("Passwords do not match.")
		setPassword()
	}

	return passwd1
}

func setTimeZone() string {
	response := helpers.CurlResponse("https://ipapi.co/timezone")
	helpers.ClearConsole()
	helpers.PrintHeader(sectionName, "Time Zone")

	if helpers.YesNo(fmt.Sprintf("TimeZone detected to be %s is this correct?", response)) {
		return response
	} else {
		var userTZ string
		fmt.Print("Enter your desired timezone e.g: Europe/London: ")
		fmt.Scanf("%s", &userTZ)
		fmt.Println(userTZ)
		return userTZ
	}
}

func keyboardLayout() string {
	var returnValue string
	var keyLists []string
	for _, file := range helpers.FindFiles("/usr/share/kbd/keymaps", ".gz", ".map.gz", true) {
		keyLists = append(keyLists, file)
	}
	helpers.ClearConsole()
	helpers.PrintHeader(sectionName, "Keyboard Layout")

	_, returnValue = helpers.PromptSelect("Select your keyboard layout", keyLists)

	fmt.Printf("Selected Layout: %s\n", returnValue)
	if !helpers.YesNo("Is this correct?") {
		returnValue = keyboardLayout()
	}
	return returnValue
}

func setDiskVars() string {
	var diskList []*ghw.Disk
	var selectedDisk *ghw.Disk
	var textDiskList []string

	bInfo, err := ghw.Block()
	if err != nil {
		panic(err)
	}

	helpers.ClearConsole()
	helpers.PrintHeader(sectionName, "Disk")

	for _, disk := range bInfo.Disks {
		if disk.DriveType.String() == "HDD" || disk.DriveType.String() == "SSD" {
			convertedSize := helpers.ByteSizeConverter(disk.SizeBytes)
			diskList = append(diskList, disk)
			textDiskList = append(textDiskList, fmt.Sprintf("/dev/%s - %s", disk.Name, convertedSize))
		}
	}

	numRes, _ := helpers.PromptSelect("Select your disk", textDiskList)

	for index, dsk := range diskList {
		if numRes == index {
			selectedDisk = dsk
		} else {
			setDiskVars()
		}
	}
	fmt.Printf("%s: /dev/%s (%s) - %s\n",
				selectedDisk.DriveType.String(),
				selectedDisk.Name,
				helpers.ByteSizeConverter(selectedDisk.SizeBytes),
				selectedDisk.Model)

	if helpers.YesNo("Is this correct?") {
		fmt.Printf("Selected device: /dev/%s\n", selectedDisk.Name)
	} else {
		setDiskVars()
	}

	return selectedDisk.Name
}

func userInfo() (string, string) {
	helpers.ClearConsole()
	helpers.PrintHeader(sectionName, "User Info")

	uName := helpers.InputPrompt("Enter your username")
	userName := strings.ToLower(uName)
	hName := helpers.InputPrompt("Enter your hostname")
	hostName := strings.ToLower(hName)

	fmt.Println("Note: the username and hostname are automatically converted to lowercase.")
	fmt.Printf("USERNAME: %s\nHOSTNAME: %s\n", userName, hostName)
	if !helpers.YesNo("Is this correct?") {
		userName, hostName = userInfo()
	}


	return userName, hostName

}

func aurHelper() string {
	helpers.ClearConsole()
	helpers.PrintHeader(sectionName, "AUR Helpers")

	var aurHelpersList = []string{"aura", "nix", "pacaur", "paru", "picaur", "trizen", "yay", "none"}
	fmt.Println("select \"none\" if you don't want any or \"nix\" to use the Nix Package Manager.")
	_, answer := helpers.PromptSelect("Select your AUR Helper", aurHelpersList)
	return answer
}

func Startupp() {
	var CONFIG_DIR string = fmt.Sprintf("%s/config", helpers.GetCurrDirPath())

	if helpers.YesNo("Run checks?") {
		checks()
	}

	createConfigFile(CONFIG_DIR)
	var CONFIG_FILE string = fmt.Sprintf("%s/config.json", CONFIG_DIR)


	passwd := setPassword()
	helpers.JsonUpdater(CONFIG_FILE, "userPassword", passwd, false)

	timeZone := setTimeZone()
	helpers.JsonUpdater(CONFIG_FILE, "timeZone", timeZone, false)

	keyLayout := keyboardLayout()
	helpers.JsonUpdater(CONFIG_FILE, "keyboardLayout", keyLayout, false)

	disk := setDiskVars()
	helpers.JsonUpdater(CONFIG_FILE, "mountOptions", "defaults", false)
	helpers.JsonUpdater(CONFIG_FILE, "disk", disk, false)

	userName, hostName := userInfo()
	helpers.JsonUpdater(CONFIG_FILE, "userName", userName, false)
	helpers.JsonUpdater(CONFIG_FILE, "hostName", hostName, false)

	aurH := aurHelper()
	helpers.JsonUpdater(CONFIG_FILE, "aurHelper", aurH, false)
}
