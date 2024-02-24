# Notes

Create a common config file to store common variables for the script files.
Compile the files separated and uses a shell script to launch then accordingly.

Split the screen in 2 sections side-by-side, on the left put the script
mechanic, on the right put the explanation of things about the current step
- which commands are being run
- the objective of the step
- provide more details

prompt for the user to select the language of the script on the beginning

on the import of configs/user files, search for an install.sh script
and scan throught it searching for harmful commands like "sudo rm -rf /"
on the step where this happens, on the right-side (info side) provide
a list of harmful commands.
if a command is found, warn the user about it and ask if they are sure
about executing the install.sh script. if no, skip; if yes, execute.

create a top header of each step and a bottom footer too.
header contains:
- number of current step/remaining steps
- name of step
- other things

footer contains:
- exit message
- help message
- other things

Reminders:
Each script will be a bubbletea app, with multiple views representing the
accordingly sections.

Planned Features:
Use the i18n and i10n pkgs to implement internationalization.
Implements a verbose/explanatory script log like the NIX installer.

THINGS THAT I *NEED* TO DO:
Implement check/search criteria in packages to prevent the user to install
packages that are not available in the Arch Linux main repos.

Split the installation in multiple sections:
- [ ] Startup
- [ ] Pre-Install
- [ ] Base-Install
- [ ] User Configuration
- [ ] Post-Install
- [ ] Rice System
Startup Steps:
- [ ] Create Config File
- [ ] Get User Password
- [ ] Get TimeZone
- [ ] Get Keyboard Layout
- [ ] Get Disk Device
- [ ] Get Disk Size
- [ ] Set Mount Options
- [ ] Get User Name
- [ ] Get Host Name
- [ ] Get AUR Helper or NIX Package Manager
Pre-Install Steps:
- [ ] Get Country ISO
- [ ] Synchronize Hardware Clock
- [ ] Set Pacman
- [ ] Download Essential Packages
- [ ] Format Disk
- [ ] Get BOOT System
- [ ] Get SWAP Partition
- [ ] Partition Disk
- [ ] Mount Partitions
- [ ] Install Base System
- [ ] Generate File System Table (FSTAB)
- [ ] Install GRUB
Base-Install Steps:
- [ ] Set Network
- [ ] Config Pacman (again)
- [ ] Set MAKEPKG Config
- [ ] Set Language
- [ ] Set Locale
- [ ] Set Time Zone
- [ ] Set NOPASSWD SUDO (temporary)
- [ ] Config Pacman (again)
- [ ] Installing Setup Packages
- [ ] Install Microcode
- [ ] Install Graphics Cards (NVIDIA|AMD|INTEGRATED|INTEL)
- [ ] Set User
- [ ] Set Hosts file
User Configuration Steps:
- [ ] Creates .cache
- [ ] Sets SHELL
- [ ] Installs AUR/NIX
- [ ] Install Packages
Post-Install Steps:
- [ ] Customizes GRUB
- [ ] Installs Display Manager
- [ ] Customizes Display Manager
- [ ] Enable Services
- [ ] Customizing Plymounth
- [ ] Remove NOPASSWD SUDO
- [ ] Set Default SUDO
Rice System Steps:
- [ ] import and install user dotfiles (version systems like git/svn)
- [ ] import and install user files from hard-drive

## TODOs

- [ ] dual-screen
    - [x] split screen into 2 sides
    - [x] switch between screens
    - [ ] add behaviour so they are independent from each other

need to make the list function on the left screen and the pager on the right
