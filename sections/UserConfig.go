package sections

import (
	"ArchInstall/helpers"
	"encoding/json"
	"fmt"
	"io/ioutil"

	//"os"
	"strings"
)

/* This file will guide the user through the installation and configuration
Of custom installed packages and their configuration, or it's ricing.

This part of the script will read a CUSTOM_CUSTOMIZATION_FILE if all of the instructions
to install any custom rices with a single file.

This file must be inside user .dotfiles repo, and this repo should follow the folder specification structure.*/

/* FOLDER SPECIFICATION STRUCTURE:

.dotfiles/
├─┬── Catppuccin/
│ ├──┬── Terminals/
│ │  ├────── Terminator/
│ │  │       ├────── config
│ │  │       ├────── additional_packages.json
│ │  │       └────── install.json
│ │  └────── Alacritty/
│ │          ├────── config
│ │          ├────── additional_packages.json
│ │          └────── install.json
│ ├──┬── CodeEditors/
│ │  ├────── VisualStudioCode/
│ │  │       ├────── extensions.json
│ │  │       ├────── profiles.json
│ │  │       ├────── settings.json
│ │  │       └────── install.json
│ │  └────── NeoVim/
│ │          ├────── configs/
│ │          ├────── extensions.json
│ │          ├────── settings.json
│ │          └────── install.json
│ ├──┬── Shells/
│ │  ├────── BASH/
│ │  │       ├────── Aliases.json
│ │  │       ├────── Functions.json
│ │  │       ├────── Prompts.json
│ │  │       └────── install.json
│ │  └────── ZSH/
│ │          ├────── Aliases.json
│ │          ├────── Functions.json
│ │          ├────── Prompts.json
│ │          └────── install.json
│ ├──┬── CliTools/
│ │  ├────── Bat/
│ │  │       ├────── config
│ │  │       └────── install.json
│ │  ├────── CAVA/
│ │  │       ├────── config
│ │  │       └────── install.json
│ │  └────── ls/
│ │          ├────── config
│ │          └────── install.json
│ ├──┬── SystemPackages/
│ │  ├────── DesktopEnvironment/
│ │  │       ├────── KDE/
│ │  │       │       ├────── configs/
│ │  │       │       └────── install.json
│ │  │       └────── GNOME/
│ │  │               ├────── configs/
│ │  │               └────── install.json
│ │  ├────── WindowManager/
│ │  │       ├────── i3/
│ │  │       │       ├────── config
│ │  │       │       └────── install.json
│ │  │       └────── BSPWM/
│ │  │               ├────── config
│ │  │               └────── install.json
│ │  ├────── DisplayManager/
│ │  │       ├────── SDDM/
│ │  │       │       ├────── configs/
│ │  │       │       └────── install.json
│ │  │       └────── LXDM/
│ │  │               ├────── configs/
│ │  │               └────── install.json
│ │  └────── Compositor/
│ │          ├────── Picom/
│ │          │       ├────── config
│ │          │       └────── install.json
│ │          └────── XCOMP/
│ │                  ├────── config
│ │                  └────── install.json
│ ├──┬── UserApps/
│ │  └────── Discord/
│ │          ├────── CustomStyles.css
│ │          └────── install.json
│ ├──┬── Assets/
│ │  ├────── Wallpapers/
│ │  │       ├────── LockScreen/
│ │  │       │       ├────── Wallpaper01.png
│ │  │       │       └────── Wallpaper02.png
│ │  │       ├────── WindowManager/
│ │  │       │       ├────── Wallpaper01.png
│ │  │       │       └────── Wallpaper02.png
│ │  │       └────── DesktopManager/
│ │  │               ├────── Wallpaper01.png
│ │  │               └────── Wallpaper02.png
│ │  └────── ProfilePictures/
│ │          ├────── ProfilePicture01.png
│ │          └────── ProfilePicture02.png
│ └───── customization.json
│
└──── Nord/ ...

*/

/* install.json structure
{
    "Instructions": [
        {
            "Command": "MOVE",
            "from": "sorce",
            "to": "source"
        },
        {
            "Command": "COPY",
            "from": "sorce",
            "to": "source"
        },
        {
            "Command": "RENAME",
            "Source": "Name",
            "NewName": "NewName"
        },
        {
            "Command": "EXECUTE",
            "UseSudo": false,
            "Source": "CommandName",
            "Args": [
                "1",
                "2",
                "3"
            ]
        }
    ]
}
*/

/* customization.json structure:

{
    "Themes": [
        {
            "Name": "Catppuccin",
			"Location": "/Catppuccin",
            "Packages": {
                "Terminal": "/Terminals/Terminator"
                "Shell": "/Shells/Terminator"
                "CodeEditors": {
                    "neovim": "/CodeEditors/NeoVim",
                    "vscode": "null"
                },
                "CliTools": {
                    "bat": "/CliTools/Bat"
                },
                "SystemPackages": {
                    "DesktopEnvironment": null,
                    "WindowManager": "DesktopEnvironment/WindowManager/i3",
                    "DisplayManager": "DisplayManager/SDDM",
					"Compositor": null,
                }
            },
			"UserApps": {
				"Discord": "/UserApps/Discord"
			},
            "Assets": {
                "Wallpaper": "Assets/Wallpapers/WindowManager/Wallpaper01.png",
                "LockscreenWallpaper": Assets/Wallpapers/LockScreen/Wallpaper01.png",
                "UserProfilePicture": Assets/ProfilePictures/ProfilePicture01.png",
                "SystemFont": null
            }
        }
    ]
}

*/

func cloneDotfiles() {
	fmt.Println("Clone")
}

func installCatppuccin() {
	helpers.ClearConsole()
	helpers.PrintHeader("User Config", "Installing Catppuccin Theme")
	parseCustomizationJSON(fmt.Sprintf("%s/Themes/customization.json", helpers.GetCurrDirPath()))

	// b, err := json.Marshal(user)
	// if err != nil {
	//     fmt.Println(err)
	//     return
	// }
	// fmt.Println(string(b))
}

func installNord() {
	helpers.ClearConsole()
	helpers.PrintHeader("User Config", "Installing Nord Theme")
}

func extractPathsFromPackages(packages map[string]interface{}, location string, category string, paths *[]string) {
	if categoryPackages, ok := packages[category].(map[string]interface{}); ok {
		for _, packagePath := range categoryPackages {
			if packagePath != nil {
				*paths = append(*paths, fmt.Sprintf("%s/%s/install.json", location, packagePath.(string)))
			}
		}
	}
}

func extractPathsFromAssets(assets map[string]interface{}, location string, paths *[]string) {
	for _, assetPath := range assets {
		if assetPath != nil {
			*paths = append(*paths, fmt.Sprintf("%s/%s", location, assetPath.(string)))
		}
	}
}

func parseCustomizationJSON(customizationLocation string) {
	var obj map[string]interface{}
	content, err := ioutil.ReadFile(customizationLocation)
	helpers.Check(err)
	if err := json.Unmarshal(content, &obj); err != nil {
		panic(err)
	}

	var packagesPaths []string
	var assetsPaths []string
	themes := obj["Themes"].([]interface{})
	for _, theme := range themes {
		themeMap := theme.(map[string]interface{})
		packages := themeMap["Packages"].(map[string]interface{})
		assets := themeMap["Assets"].(map[string]interface{})
		location := themeMap["Location"].(string)
		extractPathsFromPackages(packages, location, "SystemPackages", &packagesPaths)
		extractPathsFromPackages(packages, location, "CodeEditors", &packagesPaths)
		extractPathsFromPackages(packages, location, "CliTools", &packagesPaths)
		extractPathsFromPackages(packages, location, "UserApps", &packagesPaths)
		extractPathsFromPackages(packages, location, "OtherApps", &packagesPaths)
		extractPathsFromAssets(assets, location, &assetsPaths)
	}

	filteredPackagesPaths := make([]string, 0)
	filteredAssetsPaths := make([]string, 0)

	for _, path := range packagesPaths {
		if !strings.Contains(path, "null/") {
			filteredPackagesPaths = append(filteredPackagesPaths, path)
		}
	}

	for _, path := range filteredAssetsPaths {
		if !strings.Contains(path, "null/") {
			filteredAssetsPaths = append(filteredAssetsPaths, path)
		}
	}

	resAssets := make([]string, 0)
	resPackages := make([]string, 0)

	if len(filteredPackagesPaths) != 0 {
		resPackages = filteredPackagesPaths
	} else {
		resPackages = packagesPaths
	}

	if len(filteredAssetsPaths) != 0 {
		resAssets = filteredAssetsPaths
	} else {
		resAssets = assetsPaths
	}

	fmt.Println("SystemPackages:", resPackages)
	fmt.Println("Assets:", resAssets)

	for _, v := range resPackages {
		parserInstallJSON(v)
	}
}

func parserInstallJSON(instructionsLocation string) {
	var instructions helpers.Instructions

	fmt.Println(instructionsLocation)

	content, err := ioutil.ReadFile(instructionsLocation)
	helpers.Check(err)
	err = json.Unmarshal(content, &instructions)
	helpers.Check(err)

	for _, instruction := range instructions.Instructions {
		switch instruction.Command {
		case "MOVE":
			//err := os.Rename(instruction.From, instruction.To)
			//helpers.Check(err)
			fmt.Println("mv", instruction.From, instruction.To)
		case "COPYFILE":
			//helpers.CopyFile(instruction.From, instruction.To)
			fmt.Println("cp", instruction.From, instruction.To)
		case "COPYDIR":
			//helpers.CopyDir(instruction.From, instruction.To)
			fmt.Println("cpdir", instruction.From, instruction.To)
		case "RENAME":
			//err := os.Rename(instruction.Source, instruction.NewName)
			//helpers.Check(err)
			fmt.Println("rename", instruction.From, instruction.To)
		case "EXECUTE":
			x := append([]string{instruction.Source}, instruction.Args...)
			fmt.Println("exec", instruction.UseSudo, x)
			//if instruction.UseSudo {
			//	helpers.SudoExecuteDir(helpers.COMMANDS_TEST_MODE, false, instruction.CommandDirectory,instruction.Source, instruction.Args...)
			//} else {
			//	helpers.RunShellCommandDir(helpers.COMMANDS_TEST_MODE, false, instruction.CommandDirectory, instruction.Source, instruction.Args...)
			//}
		case "REPLACELINE":
			helpers.ReplaceFileLine(instruction.FileReplace, instruction.SourceLine, instruction.EditedLine)
		case "GITCLONE":
			gitArgs := []string{"clone", instruction.GitURL, instruction.GitDestination}
			helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "git", gitArgs...)
		default:
			fmt.Printf("Unknown command: %s\n", instruction.Command)
		}
	}
}

func createDotfiles(cfgFile string) {
	helpers.ClearConsole()
	helpers.PrintHeader("User Config", "Creating .dotfiles")
	if helpers.YesNo("Would you like to create a Git repository in your $HOME with the created configurations?") {
		helpers.JsonUpdater(cfgFile, "createDotfilesGit", true, false)
		repoName := helpers.InputPrompt("Enter the name of the repository")
		helpers.JsonUpdater(cfgFile, "dotfilesRepoName", repoName, false)
	} else {
		helpers.JsonUpdater(cfgFile, "createDotfilesGit", false, false)
	}

	helpers.ClearConsole()
	helpers.PrintHeader("User Config", "Creating .dotfiles")
	themeOptions := []helpers.ItemInfo{
		{Item: "Catppuccin", Info: "A soothing pastel theme for the high-spirited! Catppuccin is a community-driven pastel theme that aims to be the middle ground between low and high contrast themes."},
		{Item: "Nord", Info: "An arctic, north-bluish color palette, with low-contrast colors."},
	}
	_, selectTheme := helpers.PromptSelectInfo("Select your desired theme", themeOptions)
	switch strings.ToLower(selectTheme) {
	case "catppuccin":
		installCatppuccin()
		helpers.JsonUpdater(cfgFile, "selectedTheme", "catppuccin", false)
	case "nord":
		installNord()
		helpers.JsonUpdater(cfgFile, "selectedTheme", "nord", false)
	default:
		installCatppuccin()
		helpers.JsonUpdater(cfgFile, "selectedTheme", "catppuccin", false)
	}
}

func selectCloneOrCreate(cfgFile string) {
	helpers.ClearConsole()
	helpers.PrintHeader("User Config", "Cloning .dotfiles")
	if helpers.IsCommandAvailable("git") {
		fmt.Println(helpers.PrintHiYellow("NOTE: This repo needs to follow the folder structure specified in the script repository."))
		if helpers.YesNo("Would you like to clone a existing dotfiles repository and use it's files?") {
			cloneDotfiles()
		} else {
			createDotfiles(cfgFile)
		}
	} else {
		if helpers.YesNo(fmt.Sprintf("%s is not installed, would you like to install it?", helpers.PrintHiYellow("git"))) {
			//helpers.PacmanInstallPackages("git")
			selectCloneOrCreate(cfgFile)
		} else {
			createDotfiles(cfgFile)
		}
	}

}

func UserConfig() {
	var CONFIG_DIR string = fmt.Sprintf("%s/config", helpers.GetCurrDirPath())
	var CONFIG_FILE string = fmt.Sprintf("%s/config.json", CONFIG_DIR)
	helpers.ClearConsole()
	helpers.PrintHeader("User Config", "Setting Configs")
	fmt.Printf("%s %s %s\n", helpers.PrintHiYellow("Selecting"), helpers.PrintHiRed("\"NO\""), helpers.PrintHiYellow("will skip this step and leave you in a default Arch Linux installation with the proper services enabled for you to customize in your own."))
	fmt.Printf("%s %s %s\n", helpers.PrintHiYellow("But selecting"), helpers.PrintHiGreen("\"YES\""), helpers.PrintHiYellow("will ask for a Git repository, and if you don't have one, will guide you through the customization of a pre-defined theme."))
	if helpers.YesNo("Would you like to automatic rice using the script?") {
		selectCloneOrCreate(CONFIG_FILE)
	} else {
		return
	}
}
