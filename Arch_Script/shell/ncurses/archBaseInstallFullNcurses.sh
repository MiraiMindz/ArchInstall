#!/usr/bin/env bash

### Variables
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
STEPS="19"
DEFAULT_TITLE="Arch Linux Mirai Install"
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
    dialog --title "$DEFAULT_TITLE" --msgbox "\n${WelcomeBox}" 25 90 && clear
    dialog --title "$DEFAULT_TITLE 1/$STEPS" --msgbox "\nLoading br-abnt2 keys" 7 30 && clear
    loadkeys br-abnt2
    dialog --title "$DEFAULT_TITLE 2/$STEPS" --msgbox "\nUpdating the system clock" 7 34 && clear
    timedatectl set-ntp true

    counter=0;
    pacstrap /mnt base linux linux-firmware nano &;
    pid=$!;
    trap "kill $pid 2> /dev/null" EXIT;
    while kill -0 $pid 2> /dev/null; do
        (( counter+=1 ))
        echo $counter | dialog --title "$DEFAULT_TITLE 3/$STEPS" --gauge "Installing base packages" 7 50 0;
        sleep 0.1
    done;
    trap - EXIT
    counter=100
    echo $counter | dialog --title "$DEFAULT_TITLE 4/$STEPS" --gauge "Installing base packages" 7 50 0
    clear

    dialog --title "$DEFAULT_TITLE 5/$STEPS" --msgbox "\nGenerating FSTab" 7 34 && clear
    genfstab -U /mnt >> /mnt/etc/fstab
    dialog --title "$DEFAULT_TITLE 6/$STEPS" --msgbox "\nChecking FSTab" 7 34 && clear
    if dialog --title "Is this correct?" --yesno "$(while read -r line; do echo "$line"; done </etc/fstab)" 22 94; then
        dialog --title "$DEFAULT_TITLE" --msgbox "Proceeding" 7 30
    else
        dialog --title "$DEFAULT_TITLE" --msgbox "Please fix the fstab and run this script again" 7 44
        clear
        printf "%s\n" "Please fix the fstab and run this script again"
        exit
    fi

    echo "cliInstall() {" >> archnc2.sh
    echo "    sleep 5" >> archnc2.sh
    echo "    printf \"Updating the ZoneInfo to America/Sao_Paulo\n\"" >> archnc2.sh
    echo "    ln -sf /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime" >> archnc2.sh
    echo "    " >> archnc2.sh
    echo "    printf \"Syncronizing the hardware clock to the system clock\n\"" >> archnc2.sh
    echo "    hwclock --systohc" >> archnc2.sh
    echo "    " >> archnc2.sh
    echo "    printf \"Editing /etc/locale.gen to pt_BR.UTF-8\n\"" >> archnc2.sh
    echo "    sed -i \"s/#pt_BR.UTF-8 UTF-8/pt_BR.UTF-8 UTF-8/\" /etc/locale.gen" >> archnc2.sh
    echo "    locale-gen" >> archnc2.sh
    echo "    printf \"Saving the locale in /etc/locale.conf\n\"" >> archnc2.sh
    echo "    if [ -e /etc/locale.conf ]; then" >> archnc2.sh
    echo "        echo \"LANG=pt_BR.UTF-8\" >> /etc/locale.conf" >> archnc2.sh
    echo "    else" >> archnc2.sh
    echo "        touch /etc/locale.conf" >> archnc2.sh
    echo "        echo \"LANG=pt_BR.UTF-8\" >> /etc/locale.conf" >> archnc2.sh
    echo "    fi" >> archnc2.sh
    echo "    printf \"Saving the keyboard layout in /etc/vconsole.conf\n\"" >> archnc2.sh
    echo "    if [ -e /etc/vconsole.conf ]; then" >> archnc2.sh
    echo "        echo \"KEYMAP=br-abnt2\" >> /etc/vconsole.conf" >> archnc2.sh
    echo "    else" >> archnc2.sh
    echo "        touch /etc/vconsole.conf" >> archnc2.sh
    echo "        echo \"KEYMAP=br-abnt2\" >> /etc/vconsole.conf" >> archnc2.sh
    echo "    fi" >> archnc2.sh
    echo "    " >> archnc2.sh
    echo "    printf \"Generating the hostname file\n\"" >> archnc2.sh
    echo "    read -e -p \"Enter this machine hostname: \" HSTNM" >> archnc2.sh
    echo "    if [[ -e /etc/hostname ]]; then" >> archnc2.sh
    echo "        echo \$HSTNM >> /etc/hostname" >> archnc2.sh
    echo "    else" >> archnc2.sh
    echo "        touch /etc/hostname" >> archnc2.sh
    echo "        echo \$HSTNM >> /etc/hostname" >> archnc2.sh
    echo "    fi" >> archnc2.sh
    echo "    " >> archnc2.sh
    echo "    printf \"Generating the hosts file\n\"" >> archnc2.sh
    echo "    echo \"# =====================================\" >> /etc/hosts" >> archnc2.sh
    echo "    echo \"# IPv4	Config\" >> /etc/hosts" >> archnc2.sh
    echo "    echo \"127.0.0.1	localhost\" >> /etc/hosts" >> archnc2.sh
    echo "    echo \"::1		localhost\" >> /etc/hosts" >> archnc2.sh
    echo "    echo \"127.0.1.1	\${HSTNM}.localdomain	\${HSTNM}\" >> /etc/hosts" >> archnc2.sh
    echo "    echo \"127.0.0.1	local\" > /etc/Hosts" >> archnc2.sh
    echo "    echo \"# =====================================\" >> /etc/hosts" >> archnc2.sh
    echo "    echo \"::1		ip6-localhost\" >> /etc/hosts" >> archnc2.sh
    echo "    echo \"::1		ip6-loopback\" >> /etc/hosts" >> archnc2.sh
    echo "    echo \"fe80::1%lo0 	localhost\" >> /etc/hosts" >> archnc2.sh
    echo "    echo \"ff00::0		ip6-localnet\" >> /etc/hosts" >> archnc2.sh
    echo "    echo \"ff00::0		ip6-mcastprefix\" >> /etc/hosts" >> archnc2.sh
    echo "    echo \"ff02::1		ip6-allnodes\" >> /etc/hosts" >> archnc2.sh
    echo "    echo \"ff02::2		ip6-allrouters\" >> /etc/hosts" >> archnc2.sh
    echo "    echo \"ff02::3		ip6-allhosts\" >> /etc/hosts" >> archnc2.sh
    echo "    echo \"0.0.0.0		0.0.0.0\" >> /etc/hosts" >> archnc2.sh
    echo "    " >> archnc2.sh
    echo "    echo -e -n \"Do you want to add Custom Hosts to this file too (y/n)? \"" >> archnc2.sh
    echo "    old_stty_cfg=\$(stty -g)" >> archnc2.sh
    echo "    stty raw -echo" >> archnc2.sh
    echo "    answer=\$( while ! head -c 1 | grep -i '[ny]' ;do true ;done )" >> archnc2.sh
    echo "    stty \$old_stty_cfg" >> archnc2.sh
    echo "    if echo \"\$answer\" | grep -iq \"^y\" ;then" >> archnc2.sh
    echo "        printf \"Adding custom hosts\n\"" >> archnc2.sh
    echo "        curl -fL \"https://raw.githubusercontent.com/MiraiMindz/.dotfiles/main/misc/MiraiHosts.txt\" >> /etc/hosts" >> archnc2.sh
    echo "    else" >> archnc2.sh
    echo "        printf \"Proceeding with the installation\n\"" >> archnc2.sh
    echo "    fi" >> archnc2.sh
    echo "    " >> archnc2.sh
    echo "    printf \"Setting up the root password\n\"" >> archnc2.sh
    echo "    passwd" >> archnc2.sh
    echo "    " >> archnc2.sh
    echo "    echo -e -n \"What is your processor \${DARK_BLUE}I\${NOCOLOR}ntel or \${DARK_RED}A\${NOCOLOR}MD (\${DARK_BLUE}i\${NOCOLOR}/\${DARK_RED}a\${NOCOLOR})? \"" >> archnc2.sh
    echo "    old_stty_cfg=\$(stty -g)" >> archnc2.sh
    echo "    stty raw -echo" >> archnc2.sh
    echo "    answer=\$( while ! head -c 1 | grep -i '[ai]' ;do true ;done )" >> archnc2.sh
    echo "    stty \$old_stty_cfg" >> archnc2.sh
    echo "    if echo \"\$answer\" | grep -iq \"^i\" ;then" >> archnc2.sh
    echo "        pacman -S intel-ucode" >> archnc2.sh
    echo "    else" >> archnc2.sh
    echo "        pacman -S amd-ucode" >> archnc2.sh
    echo "    fi" >> archnc2.sh
    echo "    " >> archnc2.sh
    echo "    printf \"Downloading bootloader and other packages\n\"" >> archnc2.sh
    echo "    pacman -S grub networkmanager dialog wireless_tools wpa_supplicant os-prober mtools dosfstools base-devel linux-headers iwd dhcpcd" >> archnc2.sh
    echo "    " >> archnc2.sh
    echo "    printf \"Installing bootloader\n\"" >> archnc2.sh
    echo "    printf \"Please answer the following question with the full name (/dev/sdX)\n\"" >> archnc2.sh
    echo "    read -e -p \"What is the name of your disk: \" DSKNM" >> archnc2.sh
    echo "    grub-install --target=i386-pc \$DSKNM" >> archnc2.sh
    echo "    grub-mkconfig -o /boot/grub/grub.cfg" >> archnc2.sh
    echo "    " >> archnc2.sh
    echo "    printf \"INSTRUCTIONS READ BEFORE DOING\n\"" >> archnc2.sh
    echo "    printf \"Please exit the installation media by typing: \${DARK_YELLOW}exit\${NOCOLOR}\n\"" >> archnc2.sh
    echo "    printf \"unmount the partition by typing: \${DARK_YELLOW}umount -a\${NOCOLOR}\n\"" >> archnc2.sh
    echo "    printf \"reboot your system by typing: \${DARK_YELLOW}reboot\${NOCOLOR}\n\"" >> archnc2.sh
    echo "    printf \"after the reboot activate the internet with this command:\n\"" >> archnc2.sh
    echo "    printf \"\${DARK_YELLOW}systemctl start NetworkManager\${NOCOLOR}\n\"" >> archnc2.sh
    echo "    printf \"if you are on Wi-Fi you can connect using this command: \${DARK_YELLOW}iwctl\${NOCOLOR}\n\"" >> archnc2.sh
    echo "    printf \"on the [iwd]# shell you will do the following to connect into a network:\n\"" >> archnc2.sh
    echo "    printf \"list wireless devices names with: \${DARK_YELLOW}device list\${NOCOLOR}\n\"" >> archnc2.sh
    echo "    printf \"scan for networks with: \${DARK_YELLOW}station \${device} scan\${NOCOLOR}\n\"" >> archnc2.sh
    echo "    printf \"list all available networks with: \${DARK_YELLOW}station \${device} get-networks\${NOCOLOR}\n\"" >> archnc2.sh
    echo "    printf \"to connect to a network type: \${DARK_YELLOW}station \${device} connect \${SSID}\${NOCOLOR}\n\"" >> archnc2.sh
    echo "    printf \"clone the After First Boot script with this command:\n\"" >> archnc2.sh
    echo "    printf \"\${DARK_YELLOW}curl -fLo archInstallAfter.sh \\\"https://raw.githubusercontent.com/MiraiMindz/.dotfiles/main/archInstallAfter.sh\\\"\${NOCOLOR}\n\"" >> archnc2.sh
    echo "    printf \"Run the new script with: \${DARK_YELLOW}sh archInstallAfter.sh\${NOCOLOR}\n\"" >> archnc2.sh
    echo "    printf \"\${DARK_GREEN}Good Luck\${NOCOLOR}! \n\"" >> archnc2.sh
    echo "    exit" >> archnc2.sh
    echo "    " >> archnc2.sh
    echo "}" >> archnc2.sh
    echo "" >> archnc2.sh
    echo "# Checks if dialog exists, if not go with default cli installation" >> archnc2.sh
    echo "if [[ ! -e \$(command -v dialog) || ! -f /usr/bin/dialog ]]; then" >> archnc2.sh
    echo "    printf \"%s\n\" \"dialog not found, proceeding with default cli installation\"" >> archnc2.sh
    echo "    cliInstall" >> archnc2.sh
    echo "fi" >> archnc2.sh
    echo "" >> archnc2.sh
    echo "DEFAULT_TITLE=\"Arch Linux Mirai Install\"" >> archnc2.sh
    echo "STEPS=\"19\"" >> archnc2.sh
    echo "" >> archnc2.sh
    echo "sleep 5" >> archnc2.sh
    echo "dialog --title \"\$DEFAULT_TITLE 8/\$STEPS\" --msgbox \"\nUpdating the ZoneInfo to America/Sao_Paulo\" 25 90 && clear" >> archnc2.sh
    echo "ln -sf /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime" >> archnc2.sh
    echo "" >> archnc2.sh
    echo "dialog --title \"\$DEFAULT_TITLE 9/\$STEPS\" --msgbox \"\nSyncronizing the hardware clock to the system clock\" 25 90 && clear" >> archnc2.sh
    echo "hwclock --systohc" >> archnc2.sh
    echo "" >> archnc2.sh
    echo "dialog --title \"\$DEFAULT_TITLE 10/\$STEPS\" --msgbox \"\nEditing /etc/locale.gen to pt_BR.UTF-8\" 25 90 && clear" >> archnc2.sh
    echo "sed -i \"s/#pt_BR.UTF-8 UTF-8/pt_BR.UTF-8 UTF-8/\" /etc/locale.gen" >> archnc2.sh
    echo "locale-gen" >> archnc2.sh
    echo "dialog --title \"\$DEFAULT_TITLE 11/\$STEPS\" --msgbox \"\nSaving the locale in /etc/locale.conf\" 25 90 && clear" >> archnc2.sh
    echo "if [ -e /etc/locale.conf ]; then" >> archnc2.sh
    echo "    echo \"LANG=pt_BR.UTF-8\" >> /etc/locale.conf" >> archnc2.sh
    echo "else" >> archnc2.sh
    echo "    touch /etc/locale.conf" >> archnc2.sh
    echo "    echo \"LANG=pt_BR.UTF-8\" >> /etc/locale.conf" >> archnc2.sh
    echo "fi" >> archnc2.sh
    echo "dialog --title \"\$DEFAULT_TITLE 12/\$STEPS\" --msgbox \"\nSaving the keyboard layout in /etc/vconsole.conf\" 25 90 && clear" >> archnc2.sh
    echo "if [ -e /etc/vconsole.conf ]; then" >> archnc2.sh
    echo "    echo \"KEYMAP=br-abnt2\" >> /etc/vconsole.conf" >> archnc2.sh
    echo "else" >> archnc2.sh
    echo "    touch /etc/vconsole.conf" >> archnc2.sh
    echo "    echo \"KEYMAP=br-abnt2\" >> /etc/vconsole.conf" >> archnc2.sh
    echo "fi" >> archnc2.sh
    echo "" >> archnc2.sh
    echo "HSTNM=\$(dialog --title \"\$DEFAULT_TITLE 13/\$STEPS\" --inputbox \"Enter this machine hostname: \" 25 90 3>&1- 1>&2- 2>&3-); clear" >> archnc2.sh
    echo "if [[ -e /etc/hostname ]]; then" >> archnc2.sh
    echo "    echo \$HSTNM >> /etc/hostname" >> archnc2.sh
    echo "else" >> archnc2.sh
    echo "    touch /etc/hostname" >> archnc2.sh
    echo "    echo \$HSTNM >> /etc/hostname" >> archnc2.sh
    echo "fi" >> archnc2.sh
    echo "" >> archnc2.sh
    echo "dialog --title \"\$DEFAULT_TITLE 14/\$STEPS\" --msgbox \"\nGenerating the hosts file\" 25 90 && clear" >> archnc2.sh
    echo "echo \"# =====================================\" >> /etc/hosts" >> archnc2.sh
    echo "echo \"# IPv4	Config\" >> /etc/hosts" >> archnc2.sh
    echo "echo \"127.0.0.1	localhost\" >> /etc/hosts" >> archnc2.sh
    echo "echo \"::1		localhost\" >> /etc/hosts" >> archnc2.sh
    echo "echo \"127.0.1.1	\${HSTNM}.localdomain	\${HSTNM}\" >> /etc/hosts" >> archnc2.sh
    echo "echo \"127.0.0.1	local\" > /etc/Hosts" >> archnc2.sh
    echo "echo \"# =====================================\" >> /etc/hosts" >> archnc2.sh
    echo "echo \"::1		ip6-localhost\" >> /etc/hosts" >> archnc2.sh
    echo "echo \"::1		ip6-loopback\" >> /etc/hosts" >> archnc2.sh
    echo "echo \"fe80::1%lo0 	localhost\" >> /etc/hosts" >> archnc2.sh
    echo "echo \"ff00::0		ip6-localnet\" >> /etc/hosts" >> archnc2.sh
    echo "echo \"ff00::0		ip6-mcastprefix\" >> /etc/hosts" >> archnc2.sh
    echo "echo \"ff02::1		ip6-allnodes\" >> /etc/hosts" >> archnc2.sh
    echo "echo \"ff02::2		ip6-allrouters\" >> /etc/hosts" >> archnc2.sh
    echo "echo \"ff02::3		ip6-allhosts\" >> /etc/hosts" >> archnc2.sh
    echo "echo \"0.0.0.0		0.0.0.0\" >> /etc/hosts" >> archnc2.sh
    echo "" >> archnc2.sh
    echo "if dialog --title \"\$DEFAULT_TITLE 15/\$STEPS\" --yes-label \"Yes, add\" --no-label \"No, don't add\" --yesno \"\nDo you want to add Custom Hosts to this file too?\" 7 64; then" >> archnc2.sh
    echo "    clear" >> archnc2.sh
    echo "    counter=0;" >> archnc2.sh
    echo "    curl -fL \"https://raw.githubusercontent.com/MiraiMindz/.dotfiles/main/misc/MiraiHosts.txt\" >> /etc/hosts" >> archnc2.sh
    echo "    pid=\$!;" >> archnc2.sh
    echo "    trap \"kill \$pid 2> /dev/null\" EXIT;" >> archnc2.sh
    echo "    while kill -0 \$pid 2> /dev/null; do" >> archnc2.sh
    echo "        (( counter+=1 ))" >> archnc2.sh
    echo "        echo \$counter | dialog --title \"\$DEFAULT_TITLE 15/\$STEPS\" --gauge \"Adding custom hosts\" 7 50 0;" >> archnc2.sh
    echo "        sleep 0.1" >> archnc2.sh
    echo "    done;" >> archnc2.sh
    echo "    trap - EXIT" >> archnc2.sh
    echo "    counter=100" >> archnc2.sh
    echo "    echo \$counter | dialog --title \"\$DEFAULT_TITLE 15/\$STEPS\" --gauge \"Adding custom hosts\" 7 50 0" >> archnc2.sh
    echo "    clear" >> archnc2.sh
    echo "else" >> archnc2.sh
    echo "    dialog --title \"\$DEFAULT_TITLE 15/\$STEPS\" --msgbox \"\nProceeding with the installation\" 5 12 && clear" >> archnc2.sh
    echo "fi" >> archnc2.sh
    echo "" >> archnc2.sh
    echo "RTPASS=\$(dialog --title \"\$DEFAULT_TITLE 16/\$STEPS\" --passwordbox \"Enter the root password: \" 12 50 3>&1- 1>&2- 2>&3-); clear" >> archnc2.sh
    echo "echo -e \"\$RTPASS\n\$RTPASS\" | sudo passwd -q" >> archnc2.sh
    echo "" >> archnc2.sh
    echo "if dialog --title \"\$DEFAULT_TITLE 17/\$STEPS\" --yes-label \"AMD\" --no-label \"Intel\" --yesno \"\nIs your processor AMD or Inter?\" 7 64; then" >> archnc2.sh
    echo "    pacman -S amd-ucode" >> archnc2.sh
    echo "else" >> archnc2.sh
    echo "    pacman -S intel-ucode" >> archnc2.sh
    echo "fi" >> archnc2.sh
    echo "" >> archnc2.sh
    echo "counter=0;" >> archnc2.sh
    echo "pacman -S grub networkmanager dialog wireless_tools wpa_supplicant os-prober mtools dosfstools base-devel linux-headers iwd dhcpcd" >> archnc2.sh
    echo "pid=\$!;" >> archnc2.sh
    echo "trap \"kill \$pid 2> /dev/null\" EXIT;" >> archnc2.sh
    echo "while kill -0 \$pid 2> /dev/null; do" >> archnc2.sh
    echo "    (( counter+=1 ))" >> archnc2.sh
    echo "    echo \$counter | dialog --title \"\$DEFAULT_TITLE 18/\$STEPS\" --gauge \"Downloading bootloader and other packages\" 7 50 0;" >> archnc2.sh
    echo "    sleep 0.1" >> archnc2.sh
    echo "done;" >> archnc2.sh
    echo "trap - EXIT" >> archnc2.sh
    echo "counter=100" >> archnc2.sh
    echo "echo \$counter | dialog --title \"\$DEFAULT_TITLE 18/\$STEPS\" --gauge \"Downloading bootloader and other packages\" 7 50 0" >> archnc2.sh
    echo "clear" >> archnc2.sh
    echo "" >> archnc2.sh
    echo "printf \"Installing bootloader\n\"" >> archnc2.sh
    echo "DSKNM=\$(dialog --title \"\$DEFAULT_TITLE 19/\$STEPS\" --inputbox \"What is the full name of your disk (/dev/sdX): \" 25 90 3>&1- 1>&2- 2>&3-); clear" >> archnc2.sh
    echo "grub-install --target=i386-pc \$DSKNM" >> archnc2.sh
    echo "grub-mkconfig -o /boot/grub/grub.cfg" >> archnc2.sh
    echo "" >> archnc2.sh
    echo "printf \"INSTRUCTIONS READ BEFORE DOING\n\"" >> archnc2.sh
    echo "printf \"Please exit the installation media by typing: \${DARK_YELLOW}exit\${NOCOLOR}\n\"" >> archnc2.sh
    echo "printf \"unmount the partition by typing: \${DARK_YELLOW}umount -a\${NOCOLOR}\n\"" >> archnc2.sh
    echo "printf \"reboot your system by typing: \${DARK_YELLOW}reboot\${NOCOLOR}\n\"" >> archnc2.sh
    echo "printf \"after the reboot activate the internet with this command:\n\"" >> archnc2.sh
    echo "printf \"\${DARK_YELLOW}systemctl start NetworkManager\${NOCOLOR}\n\"" >> archnc2.sh
    echo "printf \"if you are on Wi-Fi you can connect using this command: \${DARK_YELLOW}iwctl\${NOCOLOR}\n\"" >> archnc2.sh
    echo "printf \"on the [iwd]# shell you will do the following to connect into a network:\n\"" >> archnc2.sh
    echo "printf \"list wireless devices names with: \${DARK_YELLOW}device list\${NOCOLOR}\n\"" >> archnc2.sh
    echo "printf \"scan for networks with: \${DARK_YELLOW}station \${device} scan\${NOCOLOR}\n\"" >> archnc2.sh
    echo "printf \"list all available networks with: \${DARK_YELLOW}station \${device} get-networks\${NOCOLOR}\n\"" >> archnc2.sh
    echo "printf \"to connect to a network type: \${DARK_YELLOW}station \${device} connect \${SSID}\${NOCOLOR}\n\"" >> archnc2.sh
    echo "printf \"clone the After First Boot script with this command:\n\"" >> archnc2.sh
    echo "printf \"\${DARK_YELLOW}curl -fLo archInstallAfter.sh \\\"https://raw.githubusercontent.com/MiraiMindz/.dotfiles/main/archInstallAfter.sh\\\"\${NOCOLOR}\n\"" >> archnc2.sh
    echo "printf \"Run the new script with: \${DARK_YELLOW}sh archInstallAfter.sh\${NOCOLOR}\n\"" >> archnc2.sh
    echo "printf \"\${DARK_GREEN}Good Luck\${NOCOLOR}! \n\"" >> archnc2.sh
    echo "exit" >> archnc2.sh
    dialog --title "$DEFAULT_TITLE 7/$STEPS" --msgbox "\nChrooting" 7 34 && clear
    mv -v ./archnc2.sh /mnt
    arch-chroot /mnt bash -c 'sh ./archnc2.sh'
    rm -v /mnt/archnc2.sh
    exit
}

### Installation
# Checks if dialog exists, if not go with default cli installation
if [[ ! -e $(command -v dialog) || ! -f /usr/bin/dialog ]]; then
    printf "%s\n" "dialog not found, proceeding with default cli installation"
    cliInstall
fi

if dialog --title "$DEFAULT_TITLE" --yes-label "Automated" --no-label "Manual" --yesno "\nWelcome to my install script using ncurses/dialog utility\nDo you want the Automated Install or the Manual Install?" 7 64; then
    clear
    INSTALLMTD="auto"
    ncursesAutoInstall
else
    INSTALLMTD="manu"
    dialog --title "$DEFAULT_TITLE" --msgbox "\nNono" 5 7 && clear
fi

# Clear the screen after dialog script ends
sh archInstallBaseNCurses.sh
clear
exit
