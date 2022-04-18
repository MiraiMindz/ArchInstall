cliInstall() {
    sleep 5
    printf "Updating the ZoneInfo to America/Sao_Paulo\n"
    ln -sf /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime

    printf "Syncronizing the hardware clock to the system clock\n"
    hwclock --systohc

    printf "Editing /etc/locale.gen to pt_BR.UTF-8\n"
    sed -i "s/#pt_BR.UTF-8 UTF-8/pt_BR.UTF-8 UTF-8/" /etc/locale.gen
    locale-gen
    printf "Saving the locale in /etc/locale.conf\n"
    if [ -e /etc/locale.conf ]; then
        echo "LANG=pt_BR.UTF-8" >> /etc/locale.conf
    else
        touch /etc/locale.conf
        echo "LANG=pt_BR.UTF-8" >> /etc/locale.confa
    fi
    printf "Saving the keyboard layout in /etc/vconsole.conf\n"
    if [ -e /etc/vconsole.conf ]; then
        echo "KEYMAP=br-abnt2" >> /etc/vconsole.conf
    else
        touch /etc/vconsole.conf
        echo "KEYMAP=br-abnt2" >> /etc/vconsole.conf
    fi

    printf "Generating the hostname file\n"
    read -e -p "Enter this machine hostname: " HSTNM
    if [[ -e /etc/hostname ]]; then
        echo $HSTNM >> /etc/hostname
    else
        touch /etc/hostname
        echo $HSTNM >> /etc/hostname
    fi

    printf "Generating the hosts file\n"
    echo "# =====================================" >> /etc/hosts
    echo "# IPv4	Config" >> /etc/hosts
    echo "127.0.0.1	localhost" >> /etc/hosts
    echo "::1		localhost" >> /etc/hosts
    echo "127.0.1.1	${HSTNM}.localdomain	${HSTNM}" >> /etc/hosts
    echo "127.0.0.1	local" > /etc/Hosts
    echo "# =====================================" >> /etc/hosts
    echo "::1		ip6-localhost" >> /etc/hosts
    echo "::1		ip6-loopback" >> /etc/hosts
    echo "fe80::1%lo0 	localhost" >> /etc/hosts
    echo "ff00::0		ip6-localnet" >> /etc/hosts
    echo "ff00::0		ip6-mcastprefix" >> /etc/hosts
    echo "ff02::1		ip6-allnodes" >> /etc/hosts
    echo "ff02::2		ip6-allrouters" >> /etc/hosts
    echo "ff02::3		ip6-allhosts" >> /etc/hosts
    echo "0.0.0.0		0.0.0.0" >> /etc/hosts

    echo -e -n "Do you want to add Custom Hosts to this file too (y/n)? "
    old_stty_cfg=$(stty -g)
    stty raw -echo
    answer=$( while ! head -c 1 | grep -i '[ny]' ;do true ;done )
    stty $old_stty_cfg
    if echo "$answer" | grep -iq "^y" ;then
        printf "Adding custom hosts\n"
        curl -fL "https://raw.githubusercontent.com/MiraiMindz/.dotfiles/main/misc/MiraiHosts.txt" >> /etc/hosts
    else
        printf "Proceeding with the installation\n"
    fi

    printf "Setting up the root password\n"
    passwd

    echo -e -n "What is your processor ${DARK_BLUE}I${NOCOLOR}ntel or ${DARK_RED}A${NOCOLOR}MD (${DARK_BLUE}i${NOCOLOR}/${DARK_RED}a${NOCOLOR})? "
    old_stty_cfg=$(stty -g)
    stty raw -echo
    answer=$( while ! head -c 1 | grep -i '[ai]' ;do true ;done )
    stty $old_stty_cfg
    if echo "$answer" | grep -iq "^i" ;then
        pacman -S intel-ucode
    else
        pacman -S amd-ucode
    fi

    printf "Downloading bootloader and other packages\n"
    pacman -S grub networkmanager dialog wireless_tools wpa_supplicant os-prober mtools dosfstools base-devel linux-headers iwd dhcpcd

    printf "Installing bootloader\n"
    printf "Please answer the following question with the full name (/dev/sdX)\n"
    read -e -p "What is the name of your disk: " DSKNM
    grub-install --target=i386-pc $DSKNM
    grub-mkconfig -o /boot/grub/grub.cfg

    printf "INSTRUCTIONS READ BEFORE DOING\n"
    printf "Please exit the installation media by typing: ${DARK_YELLOW}exit${NOCOLOR}\n"
    printf "unmount the partition by typing: ${DARK_YELLOW}umount -a${NOCOLOR}\n"
    printf "reboot your system by typing: ${DARK_YELLOW}reboot${NOCOLOR}\n"
    printf "after the reboot activate the internet with this command:\n"
    printf "${DARK_YELLOW}systemctl start NetworkManager${NOCOLOR}\n"
    printf "if you are on Wi-Fi you can connect using this command: ${DARK_YELLOW}iwctl${NOCOLOR}\n"
    printf "on the [iwd]# shell you will do the following to connect into a network:\n"
    printf "list wireless devices names with: ${DARK_YELLOW}device list${NOCOLOR}\n"
    printf "scan for networks with: ${DARK_YELLOW}station ${device} scan${NOCOLOR}\n"
    printf "list all available networks with: ${DARK_YELLOW}station ${device} get-networks${NOCOLOR}\n"
    printf "to connect to a network type: ${DARK_YELLOW}station ${device} connect ${SSID}${NOCOLOR}\n"
    printf "clone the After First Boot script with this command:\n"
    printf "${DARK_YELLOW}curl -fLo archInstallAfter.sh \"https://raw.githubusercontent.com/MiraiMindz/.dotfiles/main/archInstallAfter.sh\"${NOCOLOR}\n"
    printf "Run the new script with: ${DARK_YELLOW}sh archInstallAfter.sh${NOCOLOR}\n"
    printf "${DARK_GREEN}Good Luck${NOCOLOR}! \n"
    exit
}

# Checks if dialog exists, if not go with default cli installation
if [[ ! -e $(command -v dialog) || ! -f /usr/bin/dialog ]]; then
    printf "%s\n" "dialog not found, proceeding with default cli installation"
    cliInstall
fi

DEFAULT_TITLE="Arch Linux Mirai Install"
STEPS="19"

sleep 5
dialog --title "$DEFAULT_TITLE 8/$STEPS" --msgbox "\nUpdating the ZoneInfo to America/Sao_Paulo" 25 90 && clear
#ln -sf /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime

dialog --title "$DEFAULT_TITLE 9/$STEPS" --msgbox "\nSyncronizing the hardware clock to the system clock" 25 90 && clear
#hwclock --systohc

dialog --title "$DEFAULT_TITLE 10/$STEPS" --msgbox "\nEditing /etc/locale.gen to pt_BR.UTF-8" 25 90 && clear
#sed -i "s/#pt_BR.UTF-8 UTF-8/pt_BR.UTF-8 UTF-8/" /etc/locale.gen
#locale-gen
dialog --title "$DEFAULT_TITLE 11/$STEPS" --msgbox "\nSaving the locale in /etc/locale.conf" 25 90 && clear
#if [ -e /etc/locale.conf ]; then
#    echo "LANG=pt_BR.UTF-8" >> /etc/locale.conf
#else
#    touch /etc/locale.conf
#    echo "LANG=pt_BR.UTF-8" >> /etc/locale.conf
#fi
dialog --title "$DEFAULT_TITLE 12/$STEPS" --msgbox "\nSaving the keyboard layout in /etc/vconsole.conf" 25 90 && clear
#if [ -e /etc/vconsole.conf ]; then
#    echo "KEYMAP=br-abnt2" >> /etc/vconsole.conf
#else
#    touch /etc/vconsole.conf
#    echo "KEYMAP=br-abnt2" >> /etc/vconsole.conf
#fi

HSTNM=$(dialog --title "$DEFAULT_TITLE 13/$STEPS" --inputbox "Enter this machine hostname: " 25 90 3>&1- 1>&2- 2>&3-); clear
#if [[ -e /etc/hostname ]]; then
#    echo $HSTNM >> /etc/hostname
#else
#    touch /etc/hostname
#    echo $HSTNM >> /etc/hostname
#fi

dialog --title "$DEFAULT_TITLE 14/$STEPS" --msgbox "\nGenerating the hosts file" 25 90 && clear
#echo "# =====================================" >> /etc/hosts
#echo "# IPv4	Config" >> /etc/hosts
#echo "127.0.0.1	localhost" >> /etc/hosts
#echo "::1		localhost" >> /etc/hosts
#echo "127.0.1.1	${HSTNM}.localdomain	${HSTNM}" >> /etc/hosts
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

if dialog --title "$DEFAULT_TITLE 15/$STEPS" --yes-label "Yes, add" --no-label "No, don't add" --yesno "\nDo you want to add Custom Hosts to this file too?" 7 64; then
    clear
    counter=0;
    #curl -fL "https://raw.githubusercontent.com/MiraiMindz/.dotfiles/main/misc/MiraiHosts.txt" >> /etc/hosts
    pid=$!;
    trap "kill $pid 2> /dev/null" EXIT;
    while kill -0 $pid 2> /dev/null; do
        (( counter+=1 ))
        echo $counter | dialog --title "$DEFAULT_TITLE 15/$STEPS" --gauge "Adding custom hosts" 7 50 0;
        sleep 0.1
    done;
    trap - EXIT
    counter=100
    echo $counter | dialog --title "$DEFAULT_TITLE 15/$STEPS" --gauge "Adding custom hosts" 7 50 0
    clear
else
    dialog --title "$DEFAULT_TITLE 15/$STEPS" --msgbox "\nProceeding with the installation" 5 12 && clear
fi

RTPASS=$(dialog --title "$DEFAULT_TITLE 16/$STEPS" --passwordbox "Enter the root password: " 12 50 3>&1- 1>&2- 2>&3-); clear
#echo -e "$RTPASS\n$RTPASS" | sudo passwd -q

if dialog --title "$DEFAULT_TITLE 17/$STEPS" --yes-label "AMD" --no-label "Intel" --yesno "\nIs your processor AMD or Inter?" 7 64; then
    #pacman -S amd-ucode
else
    #pacman -S intel-ucode
fi

printf "Downloading bootloader and other packages\n" # 18/$STEPS
#pacman -S grub networkmanager dialog wireless_tools wpa_supplicant os-prober mtools dosfstools base-devel linux-headers iwd dhcpcd

counter=0;
#pacman -S grub networkmanager dialog wireless_tools wpa_supplicant os-prober mtools dosfstools base-devel linux-headers iwd dhcpcd
pid=$!;
trap "kill $pid 2> /dev/null" EXIT;
while kill -0 $pid 2> /dev/null; do
    (( counter+=1 ))
    echo $counter | dialog --title "$DEFAULT_TITLE 18/$STEPS" --gauge "Downloading bootloader and other packages" 7 50 0;
    sleep 0.1
done;
trap - EXIT
counter=100
echo $counter | dialog --title "$DEFAULT_TITLE 18/$STEPS" --gauge "Downloading bootloader and other packages" 7 50 0
clear

printf "Installing bootloader\n"
DSKNM=$(dialog --title "$DEFAULT_TITLE 19/$STEPS" --inputbox "What is the full name of your disk (/dev/sdX): " 25 90 3>&1- 1>&2- 2>&3-); clear
#grub-install --target=i386-pc $DSKNM
#grub-mkconfig -o /boot/grub/grub.cfg

printf "INSTRUCTIONS READ BEFORE DOING\n"
printf "Please exit the installation media by typing: ${DARK_YELLOW}exit${NOCOLOR}\n"
printf "unmount the partition by typing: ${DARK_YELLOW}umount -a${NOCOLOR}\n"
printf "reboot your system by typing: ${DARK_YELLOW}reboot${NOCOLOR}\n"
printf "after the reboot activate the internet with this command:\n"
printf "${DARK_YELLOW}systemctl start NetworkManager${NOCOLOR}\n"
printf "if you are on Wi-Fi you can connect using this command: ${DARK_YELLOW}iwctl${NOCOLOR}\n"
printf "on the [iwd]# shell you will do the following to connect into a network:\n"
printf "list wireless devices names with: ${DARK_YELLOW}device list${NOCOLOR}\n"
printf "scan for networks with: ${DARK_YELLOW}station ${device} scan${NOCOLOR}\n"
printf "list all available networks with: ${DARK_YELLOW}station ${device} get-networks${NOCOLOR}\n"
printf "to connect to a network type: ${DARK_YELLOW}station ${device} connect ${SSID}${NOCOLOR}\n"
printf "clone the After First Boot script with this command:\n"
printf "${DARK_YELLOW}curl -fLo archInstallAfter.sh \"https://raw.githubusercontent.com/MiraiMindz/.dotfiles/main/archInstallAfter.sh\"${NOCOLOR}\n"
printf "Run the new script with: ${DARK_YELLOW}sh archInstallAfter.sh${NOCOLOR}\n"
printf "${DARK_GREEN}Good Luck${NOCOLOR}! \n"
exit

