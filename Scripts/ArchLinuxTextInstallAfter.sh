#!/usr/bin/env sh

#############################################################################################
# [ ] TUI Guide		[ ] TUI Auto	[ ] CLI Auto											#
# [ ]				[ ]				[ ]	01 - Enable NetworkManager | systemctl				#
# [ ]				[ ]				[ ]	02 - Create User & User Password					#
# [ ]				[ ]				[ ]	03 - Add user to Wheel Group						#
# [ ]				[ ]				[ ]	04 - Set up Pacman to parallel download				#
# [ ]				[ ]				[ ]	05 - Install Graphics								#
# [ ]				[ ]				[ ]	06 - Install Display Server							#
# [ ]				[ ]				[ ]	07 - Install PulseAudio-PavuControl					#
# [ ]				[ ]				[ ]	08 - Install XDG-User-Dirs | Maybe configure it		#
# [ ]				[ ]				[ ]	09 - Install Display Manager						#
# [ ]				[ ]				[ ]	10 - Activate Display Manager						#
# [ ]				[ ]				[ ]	11 - Install Desktop Environment / Window Manager	#
# [ ]				[ ]				[ ]	12 - Install Yay AUR Helper							#
# [ ]				[ ]				[ ]	13 - Install The Custom Packages					#
# [ ]				[ ]				[ ]	14 - Apply the Rice									#
#############################################################################################
#region FORMAT_VARIABLES
DARK_BLACK='\033[30m'
DARK_RED='\033[31m'
DARK_GREEN='\033[32m'
DARK_YELLOW='\033[33m'
DARK_BLUE='\033[34m'
DARK_PURPLE='\033[35m'
DARK_CYAN='\033[36m'
DARK_WHITE='\033[37m'
LIGHT_BLACK='\033[90m'
LIGHT_RED='\033[91m'
LIGHT_GREEN='\033[92m'
LIGHT_YELLOW='\033[93m'
LIGHT_BLUE='\033[94m'
LIGHT_PURPLE='\033[95m'
LIGHT_CYAN='\033[96m'
LIGHT_WHITE='\033[97m'
TEXTBOLD_ON='\033[1m'
TEXTFAINT_ON='\033[2m'
TEXTITALIC_ON='\033[3m'
TEXTUNDERLN_ON='\033[4m'
TEXTBLINK_ON='\033[5m'
TEXTHIGHLT_ON='\033[7m'
TEXTHIDDEN_ON='\033[8m'
TEXTSTRIKE_ON='\033[9m'
TEXTBOLD_OFF='\033[21m'
TEXTFAINT_OFF='\033[22m'
TEXTITALIC_OFF='\033[23m'
TEXTUNDERLN_OFF='\033[24m'
TEXTBLINK_OFF='\033[25m'
TEXTHIGHLT_OFF='\033[27m'
TEXTHIDDEN_OFF='\033[28m'
TEXTSTRIKE_OFF='\033[29m'
NOCOLOR='\033[39m'
TEXTRESETALL='\033[m'
#endregion

#region SCRIPT_VARIABLES
sleep_time=0.5
install_steps=4
curr_step=1
#endregion

#region UTILITY_FUNCTIONS
step_print() {
	printf "${DARK_PURPLE}[%02d/%d]${NOCOLOR}\t%s\n" "$((curr_step++))" "${install_steps}" "$1" && sleep $sleep_time
}

error_print() {
	printf "${DARK_RED}[ERROR]${NOCOLOR}\t%s\n" "$1"
}

info_print() {
	printf "${DARK_BLUE}[INFO]${NOCOLOR}\t%s\n" "$1"
}

check_print() {
	if [[ $? -eq 0 ]]; then
		printf "${DARK_GREEN}[DONE]${NOCOLOR}\t%s\n" "$1" && sleep $sleep_time
	else
		printf "${DARK_RED}[ERROR]${NOCOLOR}\t%s\n" "$1" && sleep $sleep_time
	fi
}

warn_print() {
	printf "${DARK_YELLOW}[WARN]${NOCOLOR}\t%s\n" "$1"
}

done_print() {
	printf "${DARK_GREEN}[DONE]${NOCOLOR}\t%s\n" "$1"
}

#endregion

enable_network() {
	step_print "Activating Network"
	info_print "Enabling NetworkManager with SystemD"
	#systemctl enable NetworkManager
	check_print "Enabling NetworkManager with SystemD"
	info_print "Enabling IWD with SystemD"
	#systemctl enable iwd
	check_print "Enabling IWD with SystemD"
	info_print "Enabling DHCPCD with SystemD"
	#systemctl enable dhcpcd
	check_print "Enabling DHCPCD with SystemD"
}

set_user() {
	step_print "Setting up the User"
	info_print "Creating new user"
	printf "${DARK_CYAN}[ASK]${NOCOLOR}\t"
	read -e -p "Please enter the username: " usrnm
	#useradd -m -G wheel $usrnm
	check_print "Creating new user"

	info_print "Creating the password for the user ${usrnm}"
	#passwd $usrnm
	check_print "Creating the password for the user ${usrnm}"

	info_print "Adding user ${usrnm} to WHEEL group in sudoers file"
	info_print "For security reasons this need to be done by the visudo command"
	info_print "Please uncomment the following line in the sudoers file:"
	printf "${DARK_BLUE}[INFO]${NOCOLOR}\t${DARK_YELLOW}# %%wheel ALL=(ALL) ALL${NOCOLOR}\n"
	info_print "use CTRL+O to save the file and CTRL+X to exit the editor"
	sleep 3
	#EDITOR=nano visudo
	check_print "Adding user ${usrnm} to WHEEL group in sudoers file"
}

set_pacman() {
	step_print "Configuring PacMan"
	echo -e -n "${DARK_CYAN}[ASK]${NOCOLOR}\tDo you want to update pacman mirrorlist (y/n)? "
	old_stty_cfg=$(stty -g)
	stty raw -echo
	answer=$( while ! head -c 1 | grep -i '[ny]' ;do true ;done )
	stty $old_stty_cfg
	if echo "$answer" | grep -iq "^y" ;then
		info_print "Updating pacman mirrorlist"
		if [[ -e /etc/pacman.d/mirrorlist ]]; then
			rm /etc/pacman.d/mirrorlist
			while IFS="" read -r p || [ -n "$p" ]; do
				printf "%s\n" "${p##\#}" >> /etc/pacman.d/mirrorlist
			done <<< $(curl https://archlinux.org/mirrorlist/all/)
		else
			while IFS="" read -r p || [ -n "$p" ]; do
				printf "%s\n" "${p##\#}" >> /etc/pacman.d/mirrorlist
			done <<< $(curl https://archlinux.org/mirrorlist/all/)
		fi
		check_print "Updating pacman mirrorlist"
	else
		info_print "Proceeding"
	fi

	info_print "Setting up parallel download, verbosity and colors"
	sed -i "s/#ParallelDownloads = 5/ParallelDownloads = 5/" /etc/pacman.conf
	sed -i "s/#VerbosePkgLists/VerbosePkgLists/" /etc/pacman.conf
	sed -i "s/#Color/Color/" /etc/pacman.conf
	check_print "Setting up parallel download, verbosity and colors"

	info_print "Setting up pacman to use custom repositories"
	info_print "Please uncomment the repos that you want to use"
	nano /etc/pacman.conf
	pacman -Syyy
	check_print "Setting up pacman to use custom repositories"
}

install_usr_packages() {
	step_print "Installing Packages"
	pacman -S xf86-video-intel xorg pipewire lib32-pipewire pipewire-pulse pavucontrol xdg-user-dirs sddm i3-gaps terminator dunst rofi feh git discord gnome-keyring docker opera doas ccat font-manager github-cli grub-customizer lxappearance ncurses neovim vim pacman-contrib pacman-mirrorlist pacutils thunar thunar-archive-plugin thunar-media-tags-plugin thunar-volman tumbler zsh zsh-completions zsh-syntax-highlighting openssh openssl gvfs gvfs-mtp spectacle perl neofetch btop android-file-transfer android-tools android-udev pragha pkgfile shellcheck
	check_print "Installing Packages"

	sleep 2
	info_print "Enabling Pipewire and SDDM"
	systemctl enable pipewire-pulse.service
	systemctl enable sddm.service
	check_print "Enabling Pipewire and SDDM"

	info_print "Configuring DOAS"
	if [[ -e /etc/doas.conf ]]; then
		echo "permit :wheel" >> /etc/doas.conf
	else
		touch /etc/doas.conf
		echo "permit :wheel" >> /etc/doas.conf
	fi
	check_print "Configuring DOAS"

	info_print "Updating PKGFILE"
	pkgfile --update
	check_print "Updating PKGFILE"
}

config_xdg_dirs() {
	echo -e -n "${DARK_CYAN}[ASK]${NOCOLOR}\tDo you want to use custom XDG directories (y/n)? "
	old_stty_cfg=$(stty -g)
	stty raw -echo
	answer=$( while ! head -c 1 | grep -i '[ny]' ;do true ;done )
	stty $old_stty_cfg
	if echo "$answer" | grep -iq "^y" ;then
		printf "Yes\n"
		printf "${DARK_GREEN}[OK]${NOCOLOR}\t%s\n" "Setting up Custom Folders"

		info_print "Configuring XDG Directories"
		files_part=$(ls /media)

		info_print "Creating Folders"
		[[ ! -d $HOME/lib ]] && mkdir -pv $HOME/lib
		[[ ! -d $HOME/src ]] && mkdir -pv $HOME/src
		[[ ! -d $HOME/dev ]] && mkdir -pv $HOME/dev
		[[ ! -d $HOME/run ]] && mkdir -pv $HOME/run
		[[ ! -d $HOME/data ]] && mkdir -pv $HOME/data
		[[ ! -d $HOME/state ]] && mkdir -pv $HOME/state
		check_print "Creating Folders"

		if [[ -f $HOME/.config/user-dirs.dirs ]]; then
			printf "" > $HOME/.config/user-dirs.dirs
		else
			touch $HOME/.config/user-dirs.dirs
		fi

		### LIB DIR
		if [[ -d /media/$files_part ]]; then
			[[ ! -d /media/$files_part/Downloads ]] && mkdir -pv /media/$files_part/Downloads
			[[ ! -d /media/$files_part/Images ]] && mkdir -pv /media/$files_part/Images
			[[ ! -d /media/$files_part/Musica ]] && mkdir -pv /media/$files_part/Musica
			[[ ! -d /media/$files_part/Videos ]] && mkdir -pv /media/$files_part/Videos
			[[ ! -d /media/$files_part/Documents ]] && mkdir -pv /media/$files_part/Documents
			[[ ! -d /media/$files_part/Torrents ]] && mkdir -pv /media/$files_part/Torrents
			ln -sfv /media/$files_part/Downloads $HOME/lib/Downloads
			ln -sfv /media/$files_part/Images $HOME/lib/Images
			ln -sfv /media/$files_part/Musica $HOME/lib/Musica
			ln -sfv /media/$files_part/Videos $HOME/lib/Videos
			ln -sfv /media/$files_part/Documents $HOME/lib/Documents
			ln -sfv /media/$files_part/Documents $HOME/lib/Torrents
		else
			[[ ! -d $HOME/lib/Downloads ]] && mkdir -pv $HOME/lib/Downloads
			[[ ! -d $HOME/lib/Images ]] && mkdir -pv $HOME/lib/Images
			[[ ! -d $HOME/lib/Musica ]] && mkdir -pv $HOME/lib/Musica
			[[ ! -d $HOME/lib/Videos ]] && mkdir -pv $HOME/lib/Videos
			[[ ! -d $HOME/lib/Documents ]] && mkdir -pv $HOME/lib/Documents
			[[ ! -d $HOME/lib/Torrents ]] && mkdir -pv $HOME/lib/Torrents
		fi

		### SRC DIR
		export XDG_CONFIG_HOME="$HOME/src"
		if [[ -d $HOME/.dotfiles ]]; then
			ln -sfv $HOME/.dotfiles $HOME/dotfiles
		fi

		### DEV DIR
		if [[ -d /media/$files_part ]]; then
			[[ ! -d /media/$files_part/Programming ]] && mkdir -pv /media/$files_part/Programming
			[[ ! -d /media/$files_part/MusicProduction ]] && mkdir -pv /media/$files_part/MusicProduction
			ln -sfv /media/$files_part/Programming $HOME/dev/code
			ln -sfv /media/$files_part/MusicProduction $HOME/dev/songs
			mkdir -pv $HOME/dev/code/LangFolders
		else
			mkdir -pv $HOME/dev/code/Courses
			mkdir -pv $HOME/dev/code/Lessions
			mkdir -pv $HOME/dev/code/Libraries
			mkdir -pv $HOME/dev/code/Pratices
			mkdir -pv $HOME/dev/code/Projects
			mkdir -pv $HOME/dev/code/LangFolders
			mkdir -pv $HOME/dev/songs/MIDIs
			mkdir -pv $HOME/dev/songs/Projects
			mkdir -pv $HOME/dev/songs/SamplePacks
		fi

		# OTHER DIRS
		export XDG_CACHE_HOME="/tmp"
		export XDG_STATE_HOME="$HOME/state"
		export XDG_DATA_HOME="$HOME/data"
		export XDG_RUNTIME_DIR="$HOME/run"

		echo "XDG_DOWNLOAD_DIR=\"\$HOME/lib/Downloads\"" >> $HOME/.config/user-dirs.dirs
		echo "XDG_PICTURES_DIR=\"\$HOME/lib/Images\"" >> $HOME/.config/user-dirs.dirs
		echo "XDG_MUSIC_DIR=\"\$HOME/lib/Musica\"" >> $HOME/.config/user-dirs.dirs
		echo "XDG_VIDEOS_DIR=\"\$HOME/lib/Videos\"" >> $HOME/.config/user-dirs.dirs
		echo "XDG_DOCUMENTS_DIR=\"\$HOME/lib/Documents\"" >> $HOME/.config/user-dirs.dirs
		echo "XDG_DESKTOP_DIR=\"\"" >> $HOME/.config/user-dirs.dirs
		echo "XDG_TEMPLATES_DIR=\"\"" >> $HOME/.config/user-dirs.dirs
		echo "XDG_PUBLICSHARE_DIR=\"\"" >> $HOME/.config/user-dirs.dirs

		done_print "Finished"
	else
		printf "No\n"
		printf "${DARK_GREEN}[OK]${NOCOLOR}\t%s\n" "Proceeding"
	fi

}

clear
printf "[=====================================================================]\n"
printf "             Welcome to my Arch Linux Installation Script              \n"
printf "        This script was made to install my system automatically        \n"
printf "                                 Enjoy                                 \n"
printf "[=====================================================================]\n"

echo -e -n "${DARK_CYAN}[ASK]${NOCOLOR}\tDo you have any internet connection on this machine (y/n)? "
old_stty_cfg=$(stty -g)
stty raw -echo
answer=$( while ! head -c 1 | grep -i '[ny]' ;do true ;done )
stty $old_stty_cfg
if echo "$answer" | grep -iq "^y" ;then
	printf "Yes\n"
	printf "${DARK_GREEN}[OK]${NOCOLOR}\t%s\n" "Initializing System Installation"
	enable_network
	set_user
	set_pacman
	install_usr_packages
	config_xdg_dirs
else
	printf "No\n"
	printf "${DARK_GREEN}[OK]${NOCOLOR}\t%s\n" "Please Ensure Any Internet Connection"
	exit
fi
