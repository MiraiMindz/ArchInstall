package main

/* Notes
Create a common config file (JSON) to store common variables for the script files.
Compile the files separated and uses a shell script to launch then accordingly.

Split the installation in multiple sections:
	[ ] - Startup
	[ ] - Pre-Setup
	[ ] - Setup
	[ ] - User Configuration
	[ ] - Post Install
	[ ] - Rice System
Startup Steps:
	[ ] - Create Config File
	[ ] - Get User Password
	[ ] - Get TimeZone
	[ ] - Get Keyboard Layout
	[ ] - Get Disk Device
	[ ] - Get Disk Size
	[ ] - Set Mount Options
	[ ] - Get User Name
	[ ] - Get Host Name
	[ ] - Get AUR Helper or NIX Package Manager
Pre-Setup Steps:
	[ ] - Get Country ISO
	[ ] - Synchronize Hardware Clock
	[ ] - Set Pacman
	[ ] - Download Essential Packages
	[ ] - Format Disk
	[ ] - Get BOOT System
	[ ] - Get SWAP Partition
	[ ] - Partition Disk
	[ ] - Mount Partitions
	[ ] - Install Base System
	[ ] - Generate File System Table (FSTAB)
	[ ] - Install GRUB
Setup Steps:
	[ ] - Set Network
	[ ] - Config Pacman (again)
	[ ] - Set MAKEPKG Config
	[ ] - Set Language
	[ ] - Set Locale
	[ ] - Set Time Zone
	[ ] - Set NOPASSWD SUDO (temporary)
	[ ] - Config Pacman (again)
	[ ] - Installing Setup Packages
	[ ] - Install Microcode
	[ ] - Install Graphics Cards (NVIDIA|AMD|INTEGRATED|INTEL)
	[ ] - Set User
	[ ] - Set Hosts file
User Configuration Steps:
	[ ] - Creates .cache
	[ ] - Sets SHELL
	[ ] - Installs AUR/NIX
	[ ] - Install Packages
Post Install Steps:
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
*/

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
		sections.PreInstalll()
		//sections.BaseInstall()
		//sections.UserConfig()
		//sections.PostInstall()
	}
}
