#!/usr/bin/env bash

# TODO:
#	[ ] 01 - Clone repository if folder don't exist
#	[ ] 02 - Install i3
#	[ ] 03 - Install Polybar
#	[ ] 04 - Install Picom
#	[ ] 05 - Install Dunst
#	[ ] 06 - Install Rofi
#	[ ] 07 - Install SDDM
#	[ ] 08 - Install Bash
#	[ ] 09 - Install ZSH
#	[ ] 10 - Install Terminator
#	[ ] 11 - Install CoolRetroTerm
#	[ ] 12 - Install VIM
#	[ ] 13 - Install NVIM

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
install_steps=12
curr_step=1
xdg_method=""
usr_theme=""
dots_loc=""
install_location=""
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

random_arr() {
	local arr=("$@")
	printf '%s\n' "${arr[RANDOM % $#]}"
}

#endregion

find_dotfiles() {
	if [[ -e "$HOME/.dotfiles" ]]; then
		dots_loc="$HOME/.dotfiles"
	else
		if [[ -e "$HOME/dotfiles" ]]; then
			dots_loc="$HOME/dotfiles"
		else
			read -e -p "Enter the fullpath for the dotfiles folder " dots_locc
			dots_loc="$dots_locc"
		fi
	fi
}

default_xdg_install() {
	xdg_method="default"
}

custom_xdg_install() {
	xdg_method="custom"
}

# get_install_loc() {
# 	if [[ "$xdg_method" -eq "default" ]]; then
# 		install_location="$HOME/.config"
# 	else
# 		if [[ "$xdg_method" -eq "custom" ]]; then
# 			install_location="$HOME/src"
# 		fi
# 	fi
# }

select_theme() {
	printf "${DARK_CYAN}[ASK]${NOCOLOR}\tSelect one of the following kernels\n"
	printf "\t${DARK_RED}(1)${NOCOLOR} - Catpuccin\t\t${DARK_BLUE}[Soothing pastel theme for the high-spirited]${NOCOLOR}\n"
	printf "\t${DARK_RED}(2)${NOCOLOR} - Dracula\t\t${DARK_BLUE}[A dark theme for many editors, shells, and more]${NOCOLOR}\n"
	printf "\t${DARK_RED}(3)${NOCOLOR} - Material Ocean\t${DARK_BLUE}[The most epic theme]${NOCOLOR}\n"
	printf "\t${DARK_RED}(4)${NOCOLOR} - Nord\t\t${DARK_BLUE}[An arctic, north-bluish color palette]${NOCOLOR}\n"
	printf "\t${DARK_RED}(5)${NOCOLOR} - 8-Bits\t\t${DARK_BLUE}[Catpuccin + Animated Wallpaper and 8bit fonts]${NOCOLOR}\n"
	printf "\t${DARK_RED}(*)${NOCOLOR} - RANDOM\t\t${DARK_BLUE}[Enter any number above 5 to select a random theme]${NOCOLOR}\n"
	printf "\t" && read -e -p "Enter your numeric choice: " theme_name
	case $theme_name in
		"1" | 1)
			info_print "Using Catpuccin Theme"
			usr_theme="catpuccin"
		;;
		"2" | 2)
			info_print "Using Dracula Theme"
			usr_theme="dracula"
		;;
		"3" | 3)
			info_print "Using Material Ocean Theme"
			usr_theme="material_ocean"
		;;
		"4" | 4)
			info_print "Using Nord Theme"
			usr_theme="nord"
		;;
		"5" | 5)
			info_print "Using 8-Bits Theme"
			usr_theme="bits"
		;;
		*)
			warn_print "Option out of range, using Random Theme instead"
			usr_theme=$(random_arr "catpuccin" "dracula" "material_ocean" "nord" "bits")
			warn_print "Theme selected is: ${usr_theme}"
		;;
	esac
}

install_WM() {
	# if [[ -e "${install_location}/i3" ]]; then
	# 	rm -rvi "${install_location}/i3"
	# fi

	# ln -sfv "$dots_loc/i3" "$install_location/i3"

	printf "linking %s to %s\n" "$dots_loc/i3" "$install_location/i3"

	# if [[ -e "$dots_loc/i3/config" ]]; then
	# 	rm -vi "$dots_loc/i3/config"
	# fi

	case $usr_theme in
		"catpuccin")
			info_print "Installing i3WM Catpuccin Theme"
			printf "Moving %s to %s\n" "$dots_loc/i3/Themes/Catpuccin" "$dots_loc/i3/config"
			#ln -sfv "$dots_loc/i3/Themes/Catpuccin" "$dots_loc/i3/config"
		;;
		"dracula")
			info_print "Installing i3WM Dracula Theme"
			printf "Moving %s to %s\n" "$dots_loc/i3/Themes/Dracula" "$dots_loc/i3/config"
			#ln -sfv "$dots_loc/i3/Themes/Dracula" "$dots_loc/i3/config"
		;;
		"material_ocean")
			info_print "Installing i3WM Material Ocean Theme"
			printf "Moving %s to %s\n" "$dots_loc/i3/Themes/MaterialOcean" "$dots_loc/i3/config"
			#ln -sfv "$dots_loc/i3/Themes/MaterialOcean" "$dots_loc/i3/config"
		;;
		"nord")
			info_print "Installing i3WM Nord Theme"
			printf "Moving %s to %s\n" "$dots_loc/i3/Themes/Nord" "$dots_loc/i3/config"
			#ln -sfv "$dots_loc/i3/Themes/Nord" "$dots_loc/i3/config"
		;;
		"bits") # TODO
			info_print "Installing i3WM 8-Bits Theme"
			printf "Moving %s to %s\n"  "$dots_loc/i3/Themes/Catpuccin" "$dots_loc/i3/config"
			#ln -sfv "$dots_loc/i3/Themes/Catpuccin" "$dots_loc/i3/config"
		;;
	esac
}

install_configs() {
	case $xdg_method in
		"default")
			install_location="$HOME/.config"
		;;
		"custom")
			install_location="$HOME/src"
		;;
	esac

	info_print "Installing Configs"
	install_WM
}

clear

echo -e -n "${DARK_CYAN}[ASK]${NOCOLOR}\tDo you want to install my .dotfiles (y/n)? "
old_stty_cfg=$(stty -g)
stty raw -echo
answer=$( while ! head -c 1 | grep -i '[ny]' ;do true ;done )
stty $old_stty_cfg
if echo "$answer" | grep -iq "^y" ;then
	printf "Yes\n"
	printf "${DARK_GREEN}[OK]${NOCOLOR}\t%s\n" "Installing .dotfiles"
	echo -e -n "${DARK_CYAN}[ASK]${NOCOLOR}\tAre you using the default or the custom XDG config (d/c)? "
	old_stty_cfg=$(stty -g)
	stty raw -echo
	answer=$( while ! head -c 1 | grep -i '[cd]' ;do true ;done )
	stty $old_stty_cfg
	if echo "$answer" | grep -iq "^d" ;then
		printf "Default\n"
		printf "${DARK_GREEN}[OK]${NOCOLOR}\t%s\n" "Proceeding with Default XDG Setup"
		default_xdg_install
	else
		printf "Custom\n"
		printf "${DARK_GREEN}[OK]${NOCOLOR}\t%s\n" "Proceeding with Custom XDG Setup"
		custom_xdg_install
	fi
	select_theme
	find_dotfiles
	install_configs
else
	printf "No\n"
	error_print "Aborting"
	exit
fi
