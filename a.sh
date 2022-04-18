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