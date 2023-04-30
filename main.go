package main

/*******************************************************************************
Notes:
Create a common config file to store common variables for the script files.
Compile the files separated and uses a shell script to launch then accordingly.

Planned Features:
Use the i18n and i10n pkgs to implement internationalization.
Implements a verbose/explanatory script log like the NIX installer.

THINGS THAT I *NEED* TO DO:
Implement check/search criteria in packages to prevent the user to install
packages that are not available in the Arch Linux main repos.

Split the installation in multiple sections:
	[x] - Startup
	[x] - Pre-Install
	[ ] - Base-Install
	[ ] - User Configuration
	[ ] - Post-Install
	[ ] - Rice System
Startup Steps:
	[x] - Create Config File
	[x] - Get User Password
	[x] - Get TimeZone
	[x] - Get Keyboard Layout
	[x] - Get Disk Device
	[x] - Get Disk Size
	[x] - Set Mount Options
	[x] - Get User Name
	[x] - Get Host Name
	[x] - Get AUR Helper or NIX Package Manager
Pre-Install Steps:
	[x] - Get Country ISO
	[x] - Synchronize Hardware Clock
	[x] - Set Pacman
	[x] - Download Essential Packages
	[x] - Format Disk
	[x] - Get BOOT System
	[x] - Get SWAP Partition
	[x] - Partition Disk
	[x] - Mount Partitions
	[x] - Install Base System
	[x] - Generate File System Table (FSTAB)
	[x] - Install GRUB
Base-Install Steps:
	[x] - Set Network
	[x] - Config Pacman (again)
	[x] - Set MAKEPKG Config
	[x] - Set Language
	[x] - Set Locale
	[x] - Set Time Zone
	[x] - Set NOPASSWD SUDO (temporary)
	[x] - Config Pacman (again)
	[x] - Installing Setup Packages
	[x] - Install Microcode
	[x] - Install Graphics Cards (NVIDIA|AMD|INTEGRATED|INTEL)
	[x] - Set User
	[x] - Set Hosts file
User Configuration Steps:
	[ ] - Creates .cache
	[ ] - Sets SHELL
	[ ] - Installs AUR/NIX
	[ ] - Install Packages
Post-Install Steps:
	[ ] - Customizes GRUB
	[ ] - Installs Display Manager
	[ ] - Customizes Display Manager
	[ ] - Enable Services
	[ ] - Customizing Plymounth
	[ ] - Remove NOPASSWD SUDO
	[ ] - Set Default SUDO
Rice System Steps:
	[ ] -
	[ ] -
	[ ] -
	[ ] -
*******************************************************************************/

// IMPORTS
import (
	"ArchInstall/helpers"
	"ArchInstall/sections"
	"fmt"
)

// HELPER FUNCTIONS

func welcomeScreenPrint() {

	helpers.PrintAsciiArt()

	fmt.Println(helpers.PrintYellow("Some quick notes about this script:"))
	fmt.Println(helpers.PrintHiBlack("\tIt doesn't supports Wayland."))
	fmt.Println(helpers.PrintHiBlack("\tAs of ArchISO, the default keyboard layout is en-US, so, be careful when typing if you use other layout."))
	fmt.Println(helpers.PrintHiBlack("\tThis was made for me to deploy my system easily, so there is a lot of personal taste here."))
	fmt.Println(helpers.PrintHiBlack("\tYet it's a very broad script, so you might be able to use it for bootstrapping your own system/rice."))
	fmt.Println(helpers.PrintHiBlack("\tI tried to automate/simplify everything that I could from the installation guide."))
	fmt.Println(helpers.PrintHiBlack("\tIt uses the standard 1024B measurement instead of 1000B (this means GiB instead of GB)"))
	fmt.Println(helpers.PrintHiBlack("\tThis script is divided in 5 sections:"))
	fmt.Println(helpers.PrintHiBlack("\t\t1. Startup: Only sets some basic options like keyboard layout, HDD/SDD, etc..."))
	fmt.Println(helpers.PrintHiBlack("\t\t2. Pre-Install: Setup the drive and pacstrap for installation."))
	fmt.Println(helpers.PrintHiBlack("\t\t3. Base-Install: Installs and configure the system, installs base packages and creates users."))
	fmt.Println(helpers.PrintHiBlack("\t\t4. User Config: User customizations and custom (AUR/NIX) packages."))
	fmt.Println(helpers.PrintHiBlack("\t\t5. Post-Install: Enables services, clear ups the installation."))

	fmt.Println("All of this being said. I would like to highlight that this was a fun project to work on.")
	fmt.Print("Considers following me on GitHub @MiraiMindz.\n\n")
}

// MAIN FUNC
func main() {
	welcomeScreenPrint()

	if helpers.YesNo("Can we proceed to the installation?") {
		//sections.Startupp()
		//sections.PreInstalll()
		//sections.BaseInstall()
		sections.UserConfig()
		// sections.PostInstall()
		//fmt.Println(helpers.GetEnvironmentVariables("HOME"))
		//helpers.CopyDir("/media/Arquivos/Programming/Projects/ArchInstall/a", "/media/Arquivos/Programming/Projects/ArchInstall/")
	}
}
