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
	"regexp"
	"strings"

	"github.com/jaypipes/ghw"
)

func createConfigFile(filePath string) {
	file_loc := fmt.Sprintf("%s/config.json", filePath)
	if !helpers.CheckFileExists(file_loc) {
		if !helpers.CheckFileExists(filePath) {
			os.Mkdir(filePath, 0755)
		}
		err := os.WriteFile(file_loc, []byte(""), 0644)
		helpers.Check(err)
	}
}

func checkRootUser() {
	currUser, err := user.Current()
	helpers.Check(err)
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
	helpers.PrintHeader("Startup", "Password")

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
	helpers.PrintHeader("Startup", "Time Zone")

	if helpers.YesNo(fmt.Sprintf("TimeZone detected to be %s is this correct?", response)) {
		return response
	} else {
		userTZ := helpers.InputPrompt("Enter your desired timezone e.g: Europe/London")
		fmt.Println(userTZ)
		return userTZ
	}
}

func keyboardLayout() string {
	var returnValue string
	var keyLists []string
	//for _, file := range helpers.FindFiles("/usr/share/kbd/keymaps", ".gz", ".map.gz", true) {
	//	keyLists = append(keyLists, file)
	//}
	keyLists = append(keyLists, helpers.FindFiles("/usr/share/kbd/keymaps", ".gz", ".map.gz", true)...)
	helpers.ClearConsole()
	helpers.PrintHeader("Startup", "Keyboard Layout")

	_, returnValue = helpers.PromptSelect("Select your keyboard layout", keyLists)

	fmt.Printf("Selected Layout: %s\n", returnValue)
	if !helpers.YesNo("Is this correct?") {
		returnValue = keyboardLayout()
	}
	return returnValue
}

func loadKeyboardLayout(cfgFile string) {
	keyLayout := helpers.JsonGetter(cfgFile, "keyboardLayout")
	if helpers.YesNo(fmt.Sprintf("Do you want to load the keyboard layout %s?", keyLayout)) {
		helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "loadkeys", keyLayout)
	}
}

func setDiskVars() (string, uint64) {
	var diskList []*ghw.Disk
	var selectedDisk *ghw.Disk
	var textDiskList []string

	bInfo, err := ghw.Block()
	helpers.Check(err)

	helpers.ClearConsole()
	helpers.PrintHeader("Startup", "Disk")

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

	return fmt.Sprintf("/dev/%s", selectedDisk.Name), selectedDisk.SizeBytes
}

func _readName(prompt string) string {
	uName := helpers.InputPrompt(prompt)
	usrNm := strings.ToLower(uName)
	return usrNm
}

func _askUserName() string {
	userNamePattern := "^[a-z_]([a-z0-9_-]{0,31}|[a-z0-9_-]{0,30}\\$)$"

	usrNm := _readName("Enter your username")
	userNameMatches, _ := regexp.MatchString(userNamePattern, usrNm)

	if !userNameMatches {
		helpers.ClearConsole()
		helpers.PrintHeader("Startup", "User Info")
		fmt.Println(helpers.PrintError("User name doesn't conforms the UNIX standards."))
		usrNm = _askUserName()
	}

	return usrNm
}

func _askHostName() string {
	hstNm := _readName("Enter your hostname")

	hostNameMatches := helpers.IsValidHostname(hstNm)

	if !hostNameMatches {
		helpers.ClearConsole()
		helpers.PrintHeader("Startup", "User Info")
		fmt.Println(helpers.PrintError("Host name doesn't conforms the UNIX standards."))
		hstNm = _askHostName()
	}

	return hstNm
}

func userInfo() (string, string) {
	var userName, hostName string
	helpers.ClearConsole()
	helpers.PrintHeader("Startup", "User Info")
	userName = _askUserName()
	helpers.ClearConsole()
	helpers.PrintHeader("Startup", "User Info")
	hostName = _askHostName()

	helpers.ClearConsole()
	helpers.PrintHeader("Startup", "User Info")
	fmt.Println("Note: the username and hostname are automatically converted to lowercase.")
	fmt.Printf("USERNAME: %s\nHOSTNAME: %s\n", userName, hostName)
	if !helpers.YesNo("Is this correct?") {
		userName, hostName = userInfo()
	}

	return userName, hostName
}

func rootPasswd() string {
	var passwd1, passwd2 string
	helpers.ClearConsole()
	helpers.PrintHeader("Startup", "ROOT Password")

	passwd1 = helpers.PromptReadPassword("Enter the ROOT Password")
	passwd2 = helpers.PromptReadPassword("Re-type the ROOT Password")
	if passwd1 == passwd2 {
		fmt.Println("Passwords do match.")
	} else {
		fmt.Println("Passwords do not match.")
		setPassword()
	}

	return passwd1
}

func selectInstallationType() string {
	helpers.ClearConsole()
	helpers.PrintHeader("Startup", "Installation Type")
	installTypes := []string{
		"PC",
		"Server",
		"Removable Medium",
	}

	_, selectedType := helpers.PromptSelect("Select your type of installation", installTypes)

	return selectedType
}

func aurHelper() string {
	helpers.ClearConsole()
	helpers.PrintHeader("Startup", "AUR Helpers")

	var aurHelpersList = []string{"aura", "nix", "pacaur", "paru", "picaur", "pikaur", "trizen", "yay", "none"}
	fmt.Println("select \"none\" if you don't want any or select \"nix\" to use the Nix Package Manager.")
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

	helpers.JsonUpdater(CONFIG_FILE, "installLocation", helpers.GetCurrDirPath(), false)

	selectedInstallType := selectInstallationType()
	helpers.JsonUpdater(CONFIG_FILE, "installType", selectedInstallType, false)

	keyLayout := keyboardLayout()
	helpers.JsonUpdater(CONFIG_FILE, "keyboardLayout", keyLayout, false)
	loadKeyboardLayout(CONFIG_FILE)

	timeZone := setTimeZone()
	helpers.JsonUpdater(CONFIG_FILE, "timeZone", timeZone, false)

	disk, diskSize := setDiskVars()
	helpers.JsonUpdater(CONFIG_FILE, "mountOptions", "defaults", false)
	helpers.JsonUpdater(CONFIG_FILE, "disk", disk, false)
	helpers.JsonUpdater(CONFIG_FILE, "diskSize", fmt.Sprint(diskSize), false)

	userName, hostName := userInfo()
	helpers.JsonUpdater(CONFIG_FILE, "userName", userName, false)
	helpers.JsonUpdater(CONFIG_FILE, "hostName", hostName, false)

	passwd := setPassword()
	helpers.JsonUpdater(CONFIG_FILE, "userPassword", passwd, false)

	rPasswd := rootPasswd()
	helpers.JsonUpdater(CONFIG_FILE, "rootPassword", rPasswd, false)

	aurH := aurHelper()
	helpers.JsonUpdater(CONFIG_FILE, "aurHelper", aurH, false)
}
