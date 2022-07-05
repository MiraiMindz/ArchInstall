#!/usr/bin/env sh

#################################################################################
# [ ] TUI Guide		[ ] TUI Auto	[x] CLI Auto								#
# [ ]				[ ]				[x]	01	- Set the HDD			(SCRIPT 1)	#
# [ ]				[ ]				[x]	02	- Load keys				(SCRIPT 1)	#
# [ ]				[ ]				[x]	03	- Update system clock	(SCRIPT 1)	#
# [ ]				[ ]				[x]	04	- Generate FSTab		(SCRIPT 1)	#
# [ ]				[ ]				[x]	05	- Check FSTab			(SCRIPT 1)	#
# [ ]				[ ]				[x]	06	- Install base packages	(SCRIPT 1)	#
# [ ]				[ ]				[x]	07	- Chroot				(SCRIPT 1)	#
# [ ]				[ ]				[x]	08	- Update ZoneInfo					#
# [ ]				[ ]				[x]	09	- Sync hardware clock				#
# [ ]				[ ]				[x]	10	- Generate locales					#
# [ ]				[ ]				[x]	11	- Save Locales						#
# [ ]				[ ]				[x]	12	- Save keyboard layout				#
# [ ]				[ ]				[x]	13	- Set hostname						#
# [ ]				[ ]				[x]	14	- Hosts file and custom hosts		#
# [ ]				[ ]				[x]	15	- Root password						#
# [ ]				[ ]				[x]	16	- Processor micro-code				#
# [ ]				[ ]				[x]	17	- Download bootloader and more		#
# [ ]				[ ]				[x]	18	- Install bootloader				#
# [ ]				[ ]				[x]	19	- Create Dotfiles Install Script	#
#################################################################################

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
install_steps=18
curr_step=1
second_script_name=TEMP_SECOND_SCRIPT.sh
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

#region CONSOLE iNSTALL FUNCTIONS

# generate_second_script() {

# }

console_install_partition() {
	step_print "Setting up the System Partitions"
	#lsblk
	#read -e -p "enter your disk name (sdX): " disk_name

	info_print "Creating GPT Partition Table" && sleep $sleep_time
	#sgdisk -o /dev/$disk_name
	check_print "Creating GPT Partition Table" && sleep $sleep_time

	info_print "Creating ROOT Partition " && sleep $sleep_time
	#sgdisk -n 1::+50G --typecode 2:8300 /dev/$disk_name
	check_print "Creating ROOT Partition " && sleep $sleep_time

	info_print "Creating HOME Partition " && sleep $sleep_time
	#sgdisk -n 2::+30G --typecode 2:8300 /dev/$disk_name
	check_print "Creating HOME Partition " && sleep $sleep_time

	info_print "Creating FILES Partition" && sleep $sleep_time
	#sgdisk -n 3::0 --typecode 2:8300 /dev/$disk_name
	check_print "Creating FILES Partition" && sleep $sleep_time

	info_print "Mounting ROOT Partition " && sleep $sleep_time
	#mount "/dev/${disk_name}1" /mnt
	check_print "Mounting ROOT Partition " && sleep $sleep_time

	info_print "Mounting HOME Partition " && sleep $sleep_time
	#mkdir -p /mnt/home
	#mount "/dev/${disk_name}2"
	check_print "Mounting HOME Partition " && sleep $sleep_time

	info_print "Mounting FILES Partition" && sleep $sleep_time
	#read -e -p "enter the name for your files partition: " files_part_name
	#mkdir -p /media/$files_part_name
	#mount "/dev/${disk_name}3"
	check_print "Mounting FILES Partition" && sleep $sleep_time

	info_print "Checking Table Integrity" && sleep $sleep_time
	#sgdisk -v /dev/$disk_name
	check_print "Checking Table Integrity" && sleep $sleep_time
}

console_install_loadkeys() {
	step_print "Loading BR-ABNT2 Keys"
	#loadkeys br-abnt2
	check_print "Loading BR-ABNT2 Keys" && sleep $sleep_time
}

console_install_update_sysclock() {
	step_print "Updating System Clock"
	#timedatectl set-ntp true
	check_print "Updating System Clock" && sleep $sleep_time
}

console_install_fstab() {
	step_print "Generating File System Table"
	#genfstab -U /mnt >> /mnt/etc/fstab
	check_print "Generating File System Table" && sleep $sleep_time

	step_print "Checking File System Table"
	#cat /mnt/etc/fstab
	echo -e -n "${DARK_CYAN}[ASK]${NOCOLOR}\tIs this correct (y/n)? "
	old_stty_cfg=$(stty -g)
	stty raw -echo
	answer=$( while ! head -c 1 | grep -i '[ny]' ;do true ;done )
	stty $old_stty_cfg
	if echo "$answer" | grep -iq "^y" ;then
		printf "Yes\n"
		done_print "Proceeding"
	else
		printf "No\n"
		error_print "Aborting"
		warn_print "Please Fix the partitioning and re-run this script"
		exit
	fi
}

console_install_packages() {
	printf "${DARK_CYAN}[ASK]${NOCOLOR}\tSelect one of the following kernels\n"
	printf "\t${DARK_RED}(1)${NOCOLOR} - Stable\t${DARK_BLUE}[Default]${NOCOLOR}\n"
	printf "\t${DARK_RED}(2)${NOCOLOR} - Hardened\t${DARK_BLUE}[Security Focused]${NOCOLOR}\n"
	printf "\t${DARK_RED}(3)${NOCOLOR} - LTS\t${DARK_BLUE}[Long Term Support]${NOCOLOR}\n"
	printf "\t${DARK_RED}(4)${NOCOLOR} - Zen\t${DARK_BLUE}[Stability Focused for daily-drive systems]${NOCOLOR}\n"
	printf "\t" && read -e -p "Enter your numeric choice: " kern_name
	case $kern_name in
		"1" | 1)
			info_print "Using STABLE Kernel"
			user_kernel="linux"
		;;
		"2" | 2)
			info_print "Using HARDENED Kernel"
			user_kernel="linux-hardened"
		;;
		"3" | 3)
			info_print "Using LTS Kernel"
			user_kernel="linux-lts"
		;;
		"4" | 4)
			info_print "Using ZEN Kernel"
			user_kernel="linux-zen"
		;;
		*)
			warn_print "Option out of range, using STABLE kernel instead"
			user_kernel="linux"
		;;
	esac
	step_print "Installing Base Packages"
	#pacstrap /mnt base $user_kernel linux-firmware nano
	check_print "Installing Base Packages" && sleep $sleep_time
}

console_install_chroot() {
	#generate_second_script
	#mv -v ./TEMP_SECOND_SCRIPT.sh /mnt
	step_print "Changing Root Directory"
	#arch-chroot /mnt bash -c 'sh ./TEMP_SECOND_SCRIPT.sh'
	#sh ./TEMP_SECOND_SCRIPT.sh ### REMOVE ===================================================================================================
	#exit
}

console_install_update_zoneinfo() {
	sleep 5
	step_print "Updating Zone Info"
	#ln -sf /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime
	check_print "Installing Base Packages" && sleep $sleep_time
}

console_install_sync_hardclock() {
	step_print "Syncronizing Hardware Clock"
	#hwclock --systohc
	check_print "Syncronizing Hardware Clock" && sleep $sleep_time
}

console_install_locales_keyboard() {
	step_print "Generating Locales to pt_BR.UTF-8"
	#sed -i "s/#pt_BR.UTF-8 UTF-8/pt_BR.UTF-8 UTF-8/" /etc/locale.gen
    #locale-gen
	check_print "Generating Locales to pt_BR.UTF-8" && sleep $sleep_time

	step_print "Saving Locales"
	# if [[ -e /etc/locale.conf ]]; then
	# 	#echo "LANG=pt_BR.UTF-8" >> /etc/locale.conf
	# else
	# 	#touch /etc/locale.conf
	# 	#echo "LANG=pt_BR.UTF-8" >> /etc/locale.conf
	# fi
	check_print "Saving Locales" && sleep $sleep_time

	step_print "Saving Keyboard Layout"
	# if [[ -e /etc/vconsole.conf ]]; then
	# 	#echo "KEYMAP=br-abnt2" >> /etc/vconsole.conf
	# else
	# 	#touch /etc/vconsole.conf
	# 	#echo "KEYMAP=br-abnt2" >> /etc/vconsole.conf
	# fi
	check_print "Saving Keyboard Layout" && sleep $sleep_time
}

console_install_hostname_hosts() {
	step_print "Creating Hostname"
	printf "${DARK_CYAN}[ASK]${NOCOLOR}\t"
	read -e -p "Enter this machine hostname: " hstnm
	# if [[ -e /etc/hostname ]]; then
	# 	#echo $hstnm >> /etc/hostname
	# else
	# 	#touch /etc/hostname
	# 	#echo $hstnm >> /etc/hostname
	# fi
	check_print "Creating Hostname" && sleep $sleep_time

	step_print "Generating the hosts file"
	#echo "# =====================================" >> /etc/hosts
    #echo "# IPv4	Config" >> /etc/hosts
    #echo "127.0.0.1	localhost" >> /etc/hosts
    #echo "::1		localhost" >> /etc/hosts
    #echo "127.0.1.1	${hstnm}.localdomain	${hstnm}" >> /etc/hosts
    #echo "127.0.0.1	local" > /etc/Hosts
    #echo "# =====================================" >> /etc/hosts
    #echo "::1		ip6-localhost" >> /etc/hosts
    #echo "::1		ip6-loopback" >> /etc/hosts
    #echo "fe80::1%lo0 	localhost" >> /etc/hosts
    #echo "ff00::0		ip6-localnet" >> /etc/hosts
    #echo "ff00::0		ip6-mcastprefix" >> /etc/hosts
    #echo "ff02::1		ip6-allnodes" >> /etc/hosts
    #echo "ff02::2		ip6-allrouters" >> /etc/hosts
    #echo "ff02::3		ip6-allhosts" >> /etc/hosts
    #echo "0.0.0.0		0.0.0.0" >> /etc/hosts
	check_print "Generating the hosts file" && sleep $sleep_time

	info_print "Getting Custom Hosts"
	#curl -fL "https://raw.githubusercontent.com/MiraiMindz/.dotfiles/main/misc/MiraiHosts.txt" >> /etc/hosts
	check_print "Getting Custom Hosts" && sleep $sleep_time
}

console_install_root_passwd() {
	step_print "Setting Root Password"
	#passwd
	check_print "Setting Root Password" && sleep $sleep_time
}

console_install_processor_microcode() {
	step_print "Installing Processor Micro Code"
	echo -e -n "${DARK_CYAN}[ASK]${NOCOLOR}\tWhat is your processor ${DARK_BLUE}I${NOCOLOR}ntel or ${DARK_RED}A${NOCOLOR}MD (${DARK_BLUE}I${NOCOLOR}/${DARK_RED}A${NOCOLOR})? "
	old_stty_cfg=$(stty -g)
	stty raw -echo
	answer=$( while ! head -c 1 | grep -i '[ai]' ;do true ;done )
	stty $old_stty_cfg
	if echo "$answer" | grep -iq "^i" ;then
		printf "INTEL\n"
		info_print "Installing Intel Processor Microcode"
		#pacman -S intel-ucode
	else
		printf "AMD\n"
		info_print "Installing AMD Processor Microcode"
		#pacman -S amd-ucode
	fi
	check_print "Installing Processor Micro Code" && sleep $sleep_time
}

console_install_bootloader() {
	step_print "Downloading Bootloader and Other Packages"
	#pacman -S grub networkmanager dialog wireless_tools wpa_supplicant os-prober mtools dosfstools base-devel linux-headers iwd dhcpcd
	check_print "Downloading Bootloader and Other Packages" && sleep $sleep_time

	step_print "Installing Bootloader"
	printf "${DARK_CYAN}[ASK]${NOCOLOR}\t"
    read -e -p "What is the path to your disk (/dev/sdX): " dsknm
    #grub-install --target=i386-pc $dsknm
    #grub-mkconfig -o /boot/grub/grub.cfg
	check_print "Installing Bootloader" && sleep $sleep_time
}

#endregion

# error_print "Random Error"
# warn_print "Random Warning"
# done_print "Random Success"
# #check_print "Random Info Step"
# info_print "Random Info"


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
	console_install_partition
	console_install_loadkeys
	console_install_update_sysclock
	console_install_fstab
	console_install_packages
	console_install_chroot

	console_install_update_zoneinfo
	console_install_sync_hardclock
	console_install_locales_keyboard
	console_install_hostname_hosts
	console_install_root_passwd
	console_install_processor_microcode
	console_install_bootloader
else
	printf "No\n"
	printf "${DARK_GREEN}[OK]${NOCOLOR}\t%s\n" "Please Ensure Any Internet Connection"
	exit
fi
