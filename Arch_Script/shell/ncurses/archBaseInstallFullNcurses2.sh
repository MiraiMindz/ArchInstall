#!/usr/bin/env bash

######################################################################################################
# ┌──────────────┬──────────────┬──────────────────────────────────────────────────────────────────┐ #
# │ [x] CLI Auto │ [x] TUI Auto │ [x] TUI Guided │                       STEPS          ┌──────────┤ #
# │     [x]      │      [x]     │       [x]      │  01 - Load keys                      │ SCRIPT 1 │ #
# │     [x]      │      [x]     │       [x]      │  02 - Update system clock            │ SCRIPT 1 │ #
# │     [x]      │      [x]     │       [x]      │  03 - Install base packages          │ SCRIPT 1 │ #
# │     [x]      │      [x]     │       [x]      │  04 - Generate FSTab                 │ SCRIPT 1 │ #
# │     [x]      │      [x]     │       [x]      │  05 - Check FSTab                    │ SCRIPT 1 │ #
# │     [x]      │      [x]     │       [x]      │  06 - Chroot                         │ SCRIPT 1 │ #
# │     [x]      │      [x]     │       [x]      │  07 - Update ZoneInfo                └──────────┤ #
# │     [x]      │      [x]     │       [x]      │  08 - Sync hardware clock                       │ #
# │     [x]      │      [x]     │       [x]      │  09 - Generate locales                          │ #
# │     [x]      │      [x]     │       [x]      │  10 - Save Locales                              │ #
# │     [x]      │      [x]     │       [x]      │  11 - Save keyboard layout                      │ #
# │     [x]      │      [x]     │       [x]      │  12 - Set hostname                              │ #
# │     [x]      │      [x]     │       [x]      │  13 - Hosts file and custom hosts               │ #
# │     [x]      │      [x]     │       [x]      │  14 - Root password                             │ #
# │     [x]      │      [x]     │       [x]      │  15 - Processor micro-code                      │ #
# │     [x]      │      [x]     │       [x]      │  16 - Download bootloader and more              │ #
# │     [x]      │      [x]     │       [x]      │  17 - Install bootloader                        │ #
# └────────────────────────────────────────────────────────────────────────────────────────────────┘ #
######################################################################################################

### Variables
DARK_BLACK='\033[30m'a
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
steps="17"
default_title="Arch Linux Mirai Install"
WelcomeBox="┌────────────────────────────────────────────────────────────────────────────────────┐
│                    Welcome to Mirai's Arch Linux Install Script                    │
│ This is the BEFORE FIRST BOOT Script                                               │
│ This script was made to make my life easy when installing Arch Linux               │
│ It will install base packages for MY USE CASE, in this part                        │
│ On the After First Boot Script, that script will install my rice and my packages   │
│ Before using this one you will need to set some stuff, here is the list:           │
│    - Network Connection                                                            │
│    - Partitions                                                                    │
│ Here is some reminders of my personal pre-config                                   │
│ PARTITIONS:                                                                        │
│ NAME      SIZE            TYPE    MOUNTPOINT          FILESYSTEM                   │
│ sdX       FULL-DISK       disk                                                     │
│ ├─sdXY    4GB~12GB        part    [SWAP]              LINUX-SWAP                   │
│ ├─sdXY    30GB            part    /home/\${username}   EXT4                         │
│ ├─sdXY    50GB            part    /                   EXT4                         │
│ └─sdXY    REST-OF-DISK    part    /media/Arquivos     EXT4                         │
│                                                                                    │
│                   So without further ado, let's start this script                  │
└────────────────────────────────────────────────────────────────────────────────────┘"
installmtd=""

### Functions
cliInstall() {
    printf "┌──────────────────────────────────────────────────────────────────────────────────────────┐\n"
    printf "│                       Welcome to Mirai's Arch Linux Install Script                       │\n"
    printf "│ This is the ${DARK_YELLOW}BEFORE FIRST BOOT${NOCOLOR} Script\t\t\t\t\t\t\t   │\n"
    printf "│ This script was made to make my life easy when installing Arch Linux\t\t\t   │\n"
    printf "│ It will install base packages for MY USE CASE, in this part\t\t\t\t   │\n"
    printf "│ On the After First Boot Script, that script will install my rice and my packages\t   │\n"
    printf "│ Before using this one you will need to set some stuff, here is the list:\t\t   │\n"
    printf "│\t- Network Connection\t\t\t\t\t\t\t\t   │\n"
    printf "│\t- Partitions\t\t\t\t\t\t\t\t\t   │\n"
    printf "│ Here is some reminders of my personal pre-config\t\t\t\t\t   │\n"
    printf "│ PARTITIONS:\t\t\t\t\t\t\t\t\t\t   │\n"
    printf "│ NAME\t\t\tSIZE\t\tTYPE\tMOUNTPOINT\t\t\tFILESYSTEM │\n"
    printf "│ sdX\t\t\tFULL-DISK\tdisk\t\t\t\t\t\t   │\n"
    printf "│ ├─sdXY\t\t4GB~12GB\tpart\t[SWAP]\t\t\t\tLINUX-SWAP │\n"
    printf "│ ├─sdXY\t\t30GB\t\tpart\t/home/\${username}\t\tEXT4\t   │\n"
    printf "│ ├─sdXY\t\t50GB\t\tpart\t/\t\t\t\tEXT4\t   │\n"
    printf "│ └─sdXY\t\tREST-OF-DISK\tpart\t/media/Arquivos\t\t\tEXT4\t   │\n"
    printf "│                     So without further ado, let's start this script                      │\n"
    printf "└──────────────────────────────────────────────────────────────────────────────────────────┘\n"

    printf "Loading keyboard layout\n"
    loadkeys br-abnt2

    printf "Updating the system clock\n"
    timedatectl set-ntp true

    printf "Installing base packages\n"
    pacstrap /mnt base linux linux-firmware nano

    printf "Generating FSTab\n"
    genfstab -U /mnt >> /mnt/etc/fstab
    printf "Checking FSTab\n"
    cat /mnt/etc/fstab
    echo -e -n "Is this correct (y/n)? "
    old_stty_cfg=$(stty -g)
    stty raw -echo
    answer=$( while ! head -c 1 | grep -i '[ny]' ;do true ;done )
    stty $old_stty_cfg
    if echo "$answer" | grep -iq "^y" ;then
        printf "Proceeding\n"
    else
        printf "Aborting\n"
        printf "Please Fix the partitioning and re-run this script"
        exit
    fi

    echo "sleep 5" >> archBaseInstall2.sh
    echo "printf \"Updating the ZoneInfo to America/Sao_Paulo\n\"" >> archBaseInstall2.sh
    echo "ln -sf /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime" >> archBaseInstall2.sh
    echo "" >> archBaseInstall2.sh
    echo "printf \"Syncronizing the hardware clock to the system clock\n\"" >> archBaseInstall2.sh
    echo "hwclock --systohc" >> archBaseInstall2.sh
    echo "" >> archBaseInstall2.sh
    echo "printf \"Editing /etc/locale.gen to pt_BR.UTF-8\n\"" >> archBaseInstall2.sh
    echo "sed -i \"s/#pt_BR.UTF-8 UTF-8/pt_BR.UTF-8 UTF-8/\" /etc/locale.gen" >> archBaseInstall2.sh
    echo "locale-gen" >> archBaseInstall2.sh
    echo "printf \"Saving the locale in /etc/locale.conf\n\"" >> archBaseInstall2.sh
    echo "if [ -e /etc/locale.conf ]; then" >> archBaseInstall2.sh
    echo "    echo \"LANG=pt_BR.UTF-8\" >> /etc/locale.conf" >> archBaseInstall2.sh
    echo "else" >> archBaseInstall2.sh
    echo "    touch /etc/locale.conf" >> archBaseInstall2.sh
    echo "    echo \"LANG=pt_BR.UTF-8\" >> /etc/locale.conf" >> archBaseInstall2.sh
    echo "fi" >> archBaseInstall2.sh
    echo "printf \"Saving the keyboard layout in /etc/vconsole.conf\n\"" >> archBaseInstall2.sh
    echo "if [ -e /etc/vconsole.conf ]; then" >> archBaseInstall2.sh
    echo "    echo \"KEYMAP=br-abnt2\" >> /etc/vconsole.conf" >> archBaseInstall2.sh
    echo "else" >> archBaseInstall2.sh
    echo "    touch /etc/vconsole.conf" >> archBaseInstall2.sh
    echo "    echo \"KEYMAP=br-abnt2\" >> /etc/vconsole.conf" >> archBaseInstall2.sh
    echo "fi" >> archBaseInstall2.sh
    echo "" >> archBaseInstall2.sh
    echo "printf \"Generating the hostname file\n\"" >> archBaseInstall2.sh
    echo "read -e -p \"Enter this machine hostname: \" HSTNM" >> archBaseInstall2.sh
    echo "if [[ -e /etc/hostname ]]; then" >> archBaseInstall2.sh
    echo "    echo \$HSTNM >> /etc/hostname" >> archBaseInstall2.sh
    echo "else" >> archBaseInstall2.sh
    echo "    touch /etc/hostname" >> archBaseInstall2.sh
    echo "    echo \$HSTNM >> /etc/hostname" >> archBaseInstall2.sh
    echo "fi" >> archBaseInstall2.sh
    echo "" >> archBaseInstall2.sh
    echo "printf \"Generating the hosts file\n\"" >> archBaseInstall2.sh
    echo "echo \"# =====================================\" >> /etc/hosts" >> archBaseInstall2.sh
    echo "echo \"# IPv4	Config\" >> /etc/hosts" >> archBaseInstall2.sh
    echo "echo \"127.0.0.1	localhost\" >> /etc/hosts" >> archBaseInstall2.sh
    echo "echo \"::1		localhost\" >> /etc/hosts" >> archBaseInstall2.sh
    echo "echo \"127.0.1.1	\${HSTNM}.localdomain	\${HSTNM}\" >> /etc/hosts" >> archBaseInstall2.sh
    echo "echo \"127.0.0.1	local\" > /etc/Hosts" >> archBaseInstall2.sh
    echo "echo \"# =====================================\" >> /etc/hosts" >> archBaseInstall2.sh
    echo "echo \"::1		ip6-localhost\" >> /etc/hosts" >> archBaseInstall2.sh
    echo "echo \"::1		ip6-loopback\" >> /etc/hosts" >> archBaseInstall2.sh
    echo "echo \"fe80::1%lo0 	localhost\" >> /etc/hosts" >> archBaseInstall2.sh
    echo "echo \"ff00::0		ip6-localnet\" >> /etc/hosts" >> archBaseInstall2.sh
    echo "echo \"ff00::0		ip6-mcastprefix\" >> /etc/hosts" >> archBaseInstall2.sh
    echo "echo \"ff02::1		ip6-allnodes\" >> /etc/hosts" >> archBaseInstall2.sh
    echo "echo \"ff02::2		ip6-allrouters\" >> /etc/hosts" >> archBaseInstall2.sh
    echo "echo \"ff02::3		ip6-allhosts\" >> /etc/hosts" >> archBaseInstall2.sh
    echo "echo \"0.0.0.0		0.0.0.0\" >> /etc/hosts" >> archBaseInstall2.sh
    echo "" >> archBaseInstall2.sh
    echo "echo -e -n \"Do you want to add Custom Hosts to this file too (y/n)? \"" >> archBaseInstall2.sh
    echo "old_stty_cfg=\$(stty -g)" >> archBaseInstall2.sh
    echo "stty raw -echo" >> archBaseInstall2.sh
    echo "answer=\$( while ! head -c 1 | grep -i '[ny]' ;do true ;done )" >> archBaseInstall2.sh
    echo "stty \$old_stty_cfg" >> archBaseInstall2.sh
    echo "if echo \"\$answer\" | grep -iq \"^y\" ;then" >> archBaseInstall2.sh
    echo "    printf \"Adding custom hosts\n\"" >> archBaseInstall2.sh
    echo "    curl -fL \"https://raw.githubusercontent.com/MiraiMindz/.dotfiles/main/misc/MiraiHosts.txt\" >> /etc/hosts" >> archBaseInstall2.sh
    echo "else" >> archBaseInstall2.sh
    echo "    printf \"Proceeding with the installation\n\"" >> archBaseInstall2.sh
    echo "fi" >> archBaseInstall2.sh
    echo "" >> archBaseInstall2.sh
    echo "printf \"Setting up the root password\n\"" >> archBaseInstall2.sh
    echo "passwd" >> archBaseInstall2.sh
    echo "" >> archBaseInstall2.sh
    echo "echo -e -n \"What is your processor \${DARK_BLUE}I\${NOCOLOR}ntel or \${DARK_RED}A\${NOCOLOR}MD (\${DARK_BLUE}i\${NOCOLOR}/\${DARK_RED}a\${NOCOLOR})? \"" >> archBaseInstall2.sh
    echo "old_stty_cfg=\$(stty -g)" >> archBaseInstall2.sh
    echo "stty raw -echo" >> archBaseInstall2.sh
    echo "answer=\$( while ! head -c 1 | grep -i '[ai]' ;do true ;done )" >> archBaseInstall2.sh
    echo "stty \$old_stty_cfg" >> archBaseInstall2.sh
    echo "if echo \"\$answer\" | grep -iq \"^i\" ;then" >> archBaseInstall2.sh
    echo "    pacman -S intel-ucode" >> archBaseInstall2.sh
    echo "else" >> archBaseInstall2.sh
    echo "    pacman -S amd-ucode" >> archBaseInstall2.sh
    echo "fi" >> archBaseInstall2.sh
    echo "" >> archBaseInstall2.sh
    echo "printf \"Downloading bootloader and other packages\n\"" >> archBaseInstall2.sh
    echo "pacman -S grub networkmanager dialog wireless_tools wpa_supplicant os-prober mtools dosfstools base-devel linux-headers iwd dhcpcd" >> archBaseInstall2.sh
    echo "" >> archBaseInstall2.sh
    echo "printf \"Installing bootloader\n\"" >> archBaseInstall2.sh
    echo "printf \"Please answer the following question with the full name (/dev/sdX)\n\"" >> archBaseInstall2.sh
    echo "read -e -p \"What is the name of your disk: \" DSKNM" >> archBaseInstall2.sh
    echo "grub-install --target=i386-pc \$DSKNM" >> archBaseInstall2.sh
    echo "grub-mkconfig -o /boot/grub/grub.cfg" >> archBaseInstall2.sh
    echo "" >> archBaseInstall2.sh
    echo "printf \"INSTRUCTIONS READ BEFORE DOING\n\"" >> archBaseInstall2.sh
    echo "printf \"Please exit the installation media by typing: \${DARK_YELLOW}exit\${NOCOLOR}\n\"" >> archBaseInstall2.sh
    echo "printf \"unmount the partition by typing: \${DARK_YELLOW}umount -a\${NOCOLOR}\n\"" >> archBaseInstall2.sh
    echo "printf \"reboot your system by typing: \${DARK_YELLOW}reboot\${NOCOLOR}\n\"" >> archBaseInstall2.sh
    echo "printf \"after the reboot activate the internet with this command:\n\"" >> archBaseInstall2.sh
    echo "printf \"\${DARK_YELLOW}systemctl start NetworkManager\${NOCOLOR}\n\"" >> archBaseInstall2.sh
    echo "printf \"if you are on Wi-Fi you can connect using this command: \${DARK_YELLOW}iwctl\${NOCOLOR}\n\"" >> archBaseInstall2.sh
    echo "printf \"on the [iwd]# shell you will do the following to connect into a network:\n\"" >> archBaseInstall2.sh
    echo "printf \"list wireless devices names with: \${DARK_YELLOW}device list\${NOCOLOR}\n\"" >> archBaseInstall2.sh
    echo "printf \"scan for networks with: \${DARK_YELLOW}station \${device} scan\${NOCOLOR}\n\"" >> archBaseInstall2.sh
    echo "printf \"list all available networks with: \${DARK_YELLOW}station \${device} get-networks\${NOCOLOR}\n\"" >> archBaseInstall2.sh
    echo "printf \"to connect to a network type: \${DARK_YELLOW}station \${device} connect \${SSID}\${NOCOLOR}\n\"" >> archBaseInstall2.sh
    echo "printf \"clone the After First Boot script with this command:\n\"" >> archBaseInstall2.sh
    echo "printf \"\${DARK_YELLOW}curl -fLo archInstallAfter.sh \\\"https://raw.githubusercontent.com/MiraiMindz/.dotfiles/main/archInstallAfter.sh\\\"\${NOCOLOR}\n\"" >> archBaseInstall2.sh
    echo "printf \"Run the new script with: \${DARK_YELLOW}sh archInstallAfter.sh\${NOCOLOR}\n\"" >> archBaseInstall2.sh
    echo "printf \"\${DARK_GREEN}Good Luck\${NOCOLOR}! \n\"" >> archBaseInstall2.sh
    echo "exit" >> archBaseInstall2.sh
    echo "" >> archBaseInstall2.sh

    mv -v ./archBaseInstall2.sh /mnt

    printf "CHRooting to installation\n"
    arch-chroot /mnt bash -c 'sh ./archBaseInstall2.sh'
    rm -v /mnt/archBaseInstall2.sh
    exit
}

ncursesAutoInstall() {
    dialog --title "$default_title" --msgbox "\n${WelcomeBox}" 25 90 && clear
    dialog --title "$default_title 1/$steps" --msgbox "\nLoading br-abnt2 keys" 7 30 && clear
    loadkeys br-abnt2
    dialog --title "$default_title 2/$steps" --msgbox "\nUpdating the system clock" 7 34 && clear
    timedatectl set-ntp true

    dialog --title "$default_title 3/$steps" --msgbox "\nGenerating FSTab" 7 34 && clear
    genfstab -U /mnt >> /mnt/etc/fstab
    dialog --title "$default_title 4/$steps" --msgbox "\nChecking FSTab" 7 34 && clear
    if dialog --title "Is this correct?" --yesno "$(while read -r line; do echo "$line"; done </mnt/etc/fstab)" 22 94; then
        dialog --title "$default_title" --msgbox "Proceeding" 7 30
    else
        dialog --title "$default_title" --msgbox "Please fix the fstab and run this script again" 7 44
        clear
        printf "%s\n" "Please fix the fstab and run this script again"
        exit
    fi

    dialog --title "$default_title 5/$steps" --msgbox "\nInstalling base packages" 7 34 && clear
    pacstrap /mnt base linux linux-firmware nano

    dialog --title "$default_title 6/$steps" --msgbox "\nChrooting" 7 34 && clear

}

ncursesManualInstall() {
    usrkymp=$(dialog --title "$default_title 1/$steps" --output-fd 1 --inputbox "Type 'MENU' to use a menu selection\nYour desired keymap: \nExample: br-abnt2" 25 50); clear
    lwcs_usrkymp=$(echo "$usrkymp" | tr '[:upper:]' '[:lower:]')
    if [[ "$lwcs_usrkymp" == "menu" ]]; then # big menu
        usrkymp=$(dialog --title "$default_title 1/$steps" --no-tags --output-fd 1 --menu "Select your desired keymap: " 7 60 0 "amiga-de" "amiga-de" "amiga-us" "amiga-us" "atari-de" "atari-de" "atari-se" "atari-se" "atari-uk-falcon" "atari-uk-falcon" "atari-us" "atari-us" "mac-euro" "mac-euro" "mac-euro2" "mac-euro2" "apple-a1048-sv" "apple-a1048-sv" "apple-a1243-sv-fn-reverse" "apple-a1243-sv-fn-reverse" "apple-a1243-sv" "apple-a1243-sv" "apple-internal-0x0253-sv-fn-reverse" "apple-internal-0x0253-sv-fn-reverse" "apple-internal-0x0253-sv" "apple-internal-0x0253-sv" "mac-be" "mac-be" "mac-de-latin1-nodeadkeys" "mac-de-latin1-nodeadkeys" "mac-de-latin1" "mac-de-latin1" "mac-de_CH" "mac-de_CH" "mac-dk-latin1" "mac-dk-latin1" "mac-dvorak" "mac-dvorak" "mac-es" "mac-es" "mac-fi-latin1" "mac-fi-latin1" "mac-fr" "mac-fr" "mac-fr_CH-latin1" "mac-fr_CH-latin1" "mac-it" "mac-it" "mac-no-latin1" "mac-no-latin1" "mac-pl" "mac-pl" "mac-pt-latin1" "mac-pt-latin1" "mac-se" "mac-se" "mac-template" "mac-template" "mac-uk" "mac-uk" "mac-us" "mac-us" "sun-pl-altgraph" "sun-pl-altgraph" "sun-pl" "sun-pl" "sundvorak" "sundvorak" "sunkeymap" "sunkeymap" "sunt4-es" "sunt4-es" "sunt4-fi-latin1" "sunt4-fi-latin1" "sunt4-no-latin1" "sunt4-no-latin1" "sunt5-cz-us" "sunt5-cz-us" "sunt5-de-latin1" "sunt5-de-latin1" "sunt5-es" "sunt5-es" "sunt5-fi-latin1" "sunt5-fi-latin1" "sunt5-fr-latin1" "sunt5-fr-latin1" "sunt5-ru" "sunt5-ru" "sunt5-uk" "sunt5-uk" "sunt5-us-cz" "sunt5-us-cz" "sunt6-uk" "sunt6-uk" "sv-latin1" "sv-latin1" "tj_alt-UTF8" "tj_alt-UTF8" "tr_q-latin5" "tr_q-latin5" "tralt" "tralt" "trf" "trf" "trq" "trq" "ttwin_alt-UTF-8" "ttwin_alt-UTF-8" "ttwin_cplk-UTF-8" "ttwin_cplk-UTF-8" "ttwin_ct_sh-UTF-8" "ttwin_ct_sh-UTF-8" "ttwin_ctrl-UTF-8" "ttwin_ctrl-UTF-8" "ua-cp1251" "ua-cp1251" "ua-utf-ws" "ua-utf-ws" "ua-utf" "ua-utf" "ua-ws" "ua-ws" "ua" "ua" "uk" "uk" "us-acentos" "us-acentos" "us" "us" "us1" "us1" "mk-cp1251" "mk-cp1251" "mk-utf" "mk-utf" "mk" "mk" "mk0" "mk0" "nl" "nl" "nl2" "nl2" "no-latin1" "no-latin1" "no" "no" "pc110" "pc110" "pl" "pl" "pl1" "pl1" "pl2" "pl2" "pl3" "pl3" "pl4" "pl4" "pt-latin1" "pt-latin1" "pt-latin9" "pt-latin9" "ro" "ro" "ro_std" "ro_std" "ro_win" "ro_win" "ru-cp1251" "ru-cp1251" "ru-ms" "ru-ms" "ru-yawerty" "ru-yawerty" "ru" "ru" "ru1" "ru1" "ru2" "ru2" "ru3" "ru3" "ru4" "ru4" "ru_win" "ru_win" "ruwin_alt-CP1251" "ruwin_alt-CP1251" "ruwin_alt-KOI8-R" "ruwin_alt-KOI8-R" "ruwin_alt-UTF-8" "ruwin_alt-UTF-8" "ruwin_alt_sh-UTF-8" "ruwin_alt_sh-UTF-8" "ruwin_cplk-CP1251" "ruwin_cplk-CP1251" "ruwin_cplk-KOI8-R" "ruwin_cplk-KOI8-R" "ruwin_cplk-UTF-8" "ruwin_cplk-UTF-8" "ruwin_ct_sh-CP1251" "ruwin_ct_sh-CP1251" "ruwin_ct_sh-KOI8-R" "ruwin_ct_sh-KOI8-R" "ruwin_ct_sh-UTF-8" "ruwin_ct_sh-UTF-8" "ruwin_ctrl-CP1251" "ruwin_ctrl-CP1251" "ruwin_ctrl-KOI8-R" "ruwin_ctrl-KOI8-R" "ruwin_ctrl-UTF-8" "ruwin_ctrl-UTF-8" "se-fi-ir209" "se-fi-ir209" "se-fi-lat6" "se-fi-lat6" "se-ir209" "se-ir209" "se-lat6" "se-lat6" "sg-latin1-lk450" "sg-latin1-lk450" "sg-latin1" "sg-latin1" "sg" "sg" "sk-prog-qwerty" "sk-prog-qwerty" "sk-prog-qwertz" "sk-prog-qwertz" "sk-qwerty" "sk-qwerty" "sk-qwertz" "sk-qwertz" "slovene" "slovene" "sr-cy" "sr-cy" "sr-latin" "sr-latin" "bashkir" "bashkir" "bg-cp1251" "bg-cp1251" "bg-cp855" "bg-cp855" "bg_bds-cp1251" "bg_bds-cp1251" "bg_bds-utf8" "bg_bds-utf8" "bg_pho-cp1251" "bg_pho-cp1251" "bg_pho-utf8" "bg_pho-utf8" "br-abnt" "br-abnt" "br-abnt2" "br-abnt2" "br-latin1-abnt2" "br-latin1-abnt2" "br-latin1-us" "br-latin1-us" "by-cp1251" "by-cp1251" "by" "by" "bywin-cp1251" "bywin-cp1251" "ca" "ca" "cf" "cf" "croat" "croat" "cz-cp1250" "cz-cp1250" "cz-lat2-prog" "cz-lat2-prog" "cz-lat2" "cz-lat2" "cz-qwertz" "cz-qwertz" "cz-us-qwertz" "cz-us-qwertz" "cz" "cz" "de-latin1-nodeadkeys" "de-latin1-nodeadkeys" "de-latin1" "de-latin1" "de-mobii" "de-mobii" "de" "de" "de_alt_UTF-8" "de_alt_UTF-8" "de_CH-latin1" "de_CH-latin1" "defkeymap" "defkeymap" "defkeymap_V1" "defkeymap_V1" "dk-latin1" "dk-latin1" "dk" "dk" "emacs" "emacs" "emacs2" "emacs2" "es-cp850" "es-cp850" "es" "es" "et-nodeadkeys" "et-nodeadkeys" "et" "et" "fa" "fa" "fi" "fi" "fr_CH-latin1" "fr_CH-latin1" "fr_CH" "fr_CH" "gr-pc" "gr-pc" "gr" "gr" "hu" "hu" "hu101" "hu101" "il-heb" "il-heb" "il-phonetic" "il-phonetic" "il" "il" "is-latin1-us" "is-latin1-us" "is-latin1" "is-latin1" "it-ibm" "it-ibm" "it" "it" "it2" "it2" "jp106" "jp106" "kazakh" "kazakh" "ky_alt_sh-UTF-8" "ky_alt_sh-UTF-8" "kyrgyz" "kyrgyz" "la-latin1" "la-latin1" "lt" "lt" "lt" "lt" "lt" "lt" "lv-tilde" "lv-tilde" "lv" "lv" "applkey" "applkey" "unicode" "unicode" "windowkeys" "windowkeys" "backspace" "backspace" "ctrl" "ctrl" "euro" "euro" "euro1" "euro1" "euro2" "euro2" "keypad" "keypad" "wangbe" "wangbe" "wangbe2" "wangbe2" "azerty" "azerty" "be-latin1" "be-latin1" "fr-latin1" "fr-latin1" "fr-latin9" "fr-latin9" "fr-pc" "fr-pc" "fr" "fr" "adnw" "adnw" "neo" "neo" "neoqwertz" "neoqwertz" "bone" "bone" "koy" "koy" "pt-olpc" "pt-olpc" "es-olpc" "es-olpc" "fr-bepo-latin9" "fr-bepo-latin9" "fr-bepo" "fr-bepo" "ANSI-dvorak" "ANSI-dvorak" "dvorak-ca-fr" "dvorak-ca-fr" "dvorak-es" "dvorak-es" "dvorak-fr" "dvorak-fr" "dvorak-l" "dvorak-l" "dvorak-la" "dvorak-la" "dvorak-no" "dvorak-no" "dvorak-programmer" "dvorak-programmer" "dvorak-r" "dvorak-r" "dvorak-ru" "dvorak-ru" "dvorak-sv-a1" "dvorak-sv-a1" "dvorak-sv-a5" "dvorak-sv-a5" "dvorak-uk" "dvorak-uk" "dvorak-ukp" "dvorak-ukp" "dvorak" "dvorak" "colemak" "colemak" "carpalx-full" "carpalx-full" "carpalx" "carpalx" "tr_f-latin5" "tr_f-latin5" "trf-fgGIod" "trf-fgGIod"); clear
        loadkeys $usrkymp
    else
        loadkeys $usrkymp
    fi

    dialog --title "$default_title 2/$steps" --infobox "Updating system clock" 20 40
    timedatectl set-ntp true

    bspkgs=$(dialog --title "$default_title 3/$steps" --output-fd 1 --inputbox "Type \"AUTO\" to install the packages from the auto installation\nAUTO=base linux linux-firmware nano\nEnter the packages that you want to install:" 25 75); clear
    lwcs_bspkgs=$(echo "$bspkgs" | tr '[:upper:]' '[:lower:]')
    if [[ "$lwcs_bspkgs" == "auto" ]]; then
        dialog --title "$default_title 3/$steps" --infobox "Installing base packages" 7 50 0;
        pacstrap /mnt base linux linux-firmware nano
    else
        dialog --title "$default_title 3/$steps" --infobox "Installing packages" 7 50 0;
        pacstrap /mnt "$bspkgs"
    fi
    dialog --title "$default_title 4/$steps" --infobox "\nGenerating FSTab" 7 34 && clear
    dialog --title "$default_title 5/$steps" --infobox "\nChecking FSTab" 7 34 && clear
    if dialog --title "Is this correct?" --yesno "$(while read -r line; do echo "$line"; done </etc/fstab)" 22 94; then
        dialog --title "$default_title" --infobox "Proceeding" 7 30
    else
        dialog --title "$default_title" --infobox "Please fix the fstab and run this script again" 7 44
        clear
        printf "%s\n" "Please fix the fstab and run this script again"
        exit
    fi
    dialog --title "$default_title 6/$steps" --msgbox "\nChrooting" 7 34 && clear
}

### Installation
# Checks if dialog exists, if not go with default cli installation
if [[ ! -e $(command -v dialog) || ! -f /usr/bin/dialog || ! -f /bin/dialog ]]; then
    printf "%s\n" "dialog not found, proceeding with default cli installation"
    cliInstall
fi

if dialog --title "$default_title" --yes-label "Automated" --no-label "Manual" --yesno "\nWelcome to my install script using ncurses/dialog utility\nDo you want the Automated Install or the Manual Install?" 7 64; then
    installmtd="auto"
    clear
    ncursesAutoInstall
else
    installmtd="manu"
    clear
    ncursesManualInstall
fi
