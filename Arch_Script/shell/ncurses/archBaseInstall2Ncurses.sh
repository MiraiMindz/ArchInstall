#######################################################################################################################
#                                                                                                                     #
# NOTE: ON NCURSES INSTALL, create way simillar to the timezone for the user to select the keymap                     #
# ┌──────────────┬────────────────┬──────────────┬──────────────────────────────────────────────────────────────────┐ #
# │ [x] CLI Auto │ [ ] CLI Guided │ [x] TUI Auto │ [ ] TUI Guided │                       STEPS          ┌──────────┤ #
# │     [x]      │       [ ]      │      [x]     │       [ ]      │  01 - Load keys                      │ SCRIPT 1 │ #
# │     [x]      │       [ ]      │      [x]     │       [ ]      │  02 - Update system clock            │ SCRIPT 1 │ #
# │     [x]      │       [ ]      │      [x]     │       [ ]      │  03 - Install base packages          │ SCRIPT 1 │ #
# │     [x]      │       [ ]      │      [x]     │       [ ]      │  04 - Generate FSTab                 │ SCRIPT 1 │ #
# │     [x]      │       [ ]      │      [x]     │       [ ]      │  05 - Check FSTab                    │ SCRIPT 1 │ #
# │     [x]      │       [ ]      │      [x]     │       [ ]      │  06 - Chroot                         │ SCRIPT 1 │ #
# │     [x]      │       [ ]      │      [x]     │       [x]      │  07 - Update ZoneInfo                └──────────┤ #
# │     [x]      │       [ ]      │      [x]     │       [x]      │  08 - Sync hardware clock                       │ #
# │     [x]      │       [ ]      │      [x]     │       [x]      │  09 - Generate locales                          │ #
# │     [x]      │       [ ]      │      [x]     │       [x]      │  10 - Save Locales                              │ #
# │     [x]      │       [ ]      │      [x]     │       [ ]      │  11 - Save keyboard layout                      │ #
# │     [x]      │       [ ]      │      [x]     │       [x]      │  12 - Set hostname                              │ #
# │     [x]      │       [ ]      │      [x]     │       [x]      │  13 - Hosts file and custom hosts               │ #
# │     [x]      │       [ ]      │      [x]     │       [x]      │  14 - Root password                             │ #
# │     [x]      │       [ ]      │      [x]     │       [x]      │  15 - Processor micro-code                      │ #
# │     [x]      │       [ ]      │      [x]     │       [x]      │  16 - Download bootloader and more              │ #
# │     [x]      │       [ ]      │      [x]     │       [ ]      │  17 - Install bootloader                        │ #
# └─────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘ #
#######################################################################################################################

installmtd=""
default_title="Arch Linux Mirai Install"
steps="17"

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
        echo "LANG=pt_BR.UTF-8" >> /etc/locale.conf
    fi
    printf "Saving the keyboard layout in /etc/vconsole.conf\n"
    if [ -e /etc/vconsole.conf ]; then
        echo "KEYMAP=br-abnt2" >> /etc/vconsole.conf
    else
        touch /etc/vconsole.conf
        echo "KEYMAP=br-abnt2" >> /etc/vconsole.conf
    fi

    printf "Generating the hostname file\n"
    read -e -p "Enter this machine hostname: " hstnm
    if [[ -e /etc/hostname ]]; then
        echo $hstnm >> /etc/hostname
    else
        touch /etc/hostname
        echo $hstnm >> /etc/hostname
    fi

    printf "Generating the hosts file\n"
    echo "# =====================================" >> /etc/hosts
    echo "# IPv4	Config" >> /etc/hosts
    echo "127.0.0.1	localhost" >> /etc/hosts
    echo "::1		localhost" >> /etc/hosts
    echo "127.0.1.1	${hstnm}.localdomain	${hstnm}" >> /etc/hosts
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
    read -e -p "What is the name of your disk: " dsknm
    grub-install --target=i386-pc $dsknm
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

ncursesAutoInstall2() {
    sleep 5
    dialog --title "$default_title 7/$steps" --msgbox "\nUpdating the ZoneInfo to America/Sao_Paulo" 25 90 && clear
    #ln -sf /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime

    dialog --title "$default_title 8/$steps" --msgbox "\nSyncronizing the hardware clock to the system clock" 25 90 && clear
    #hwclock --systohc

    dialog --title "$default_title 9/$steps" --msgbox "\nEditing /etc/locale.gen to pt_BR.UTF-8" 25 90 && clear
    #sed -i "s/#pt_BR.UTF-8 UTF-8/pt_BR.UTF-8 UTF-8/" /etc/locale.gen
    #locale-gen
    dialog --title "$default_title 10/$steps" --msgbox "\nSaving the locale in /etc/locale.conf" 25 90 && clear
    #if [ -e /etc/locale.conf ]; then
    #    echo "LANG=pt_BR.UTF-8" >> /etc/locale.conf
    #else
    #    touch /etc/locale.conf
    #    echo "LANG=pt_BR.UTF-8" >> /etc/locale.conf
    #fi
    dialog --title "$default_title 11/$steps" --msgbox "\nSaving the keyboard layout in /etc/vconsole.conf" 25 90 && clear
    #if [ -e /etc/vconsole.conf ]; then
    #    echo "KEYMAP=br-abnt2" >> /etc/vconsole.conf
    #else
    #    touch /etc/vconsole.conf
    #    echo "KEYMAP=br-abnt2" >> /etc/vconsole.conf
    #fi

    hstnm=$(dialog --title "$default_title 12/$steps" --output-fd 1 --inputbox "Enter this machine hostname: " 25 90); clear
    #if [[ -e /etc/hostname ]]; then
    #    echo $hstnm >> /etc/hostname
    #else
    #    touch /etc/hostname
    #    echo $hstnm >> /etc/hostname
    #fi

    dialog --title "$default_title 13/$steps" --msgbox "\nGenerating the hosts file" 25 90 && clear
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

    if dialog --title "$default_title 13/$steps" --yes-label "Yes, add" --no-label "No, don't add" --yesno "\nDo you want to add Custom Hosts to this file too?" 7 64; then
        clear
        counter=0;
        #curl -fL "https://raw.githubusercontent.com/MiraiMindz/.dotfiles/main/misc/MiraiHosts.txt" >> /etc/hosts
        pid=$!;
        trap "kill $pid 2> /dev/null" EXIT;
        while kill -0 $pid 2> /dev/null; do
            (( counter+=1 ))
            echo $counter | dialog --title "$default_title 13/$steps" --gauge "Adding custom hosts" 7 50 0;
            sleep 0.1
        done;
        trap - EXIT
        counter=100
        echo $counter | dialog --title "$default_title 13/$steps" --gauge "Adding custom hosts" 7 50 0
        clear
    else
        dialog --title "$default_title 13/$steps" --msgbox "\nProceeding with the installation" 5 12 && clear
    fi

    rtpass=$(dialog --title "$default_title 14/$steps" --output-fd 1 --passwordbox "Enter the root password: " 12 50); clear
    #echo -e "$rtpass\n$rtpass" | sudo passwd -q

    if dialog --title "$default_title 15/$steps" --yes-label "Intel" --no-label "AMD" --yesno "\nIs your processor Intel or AMD?" 7 64; then
        #pacman -S intel-ucode
    else
        #pacman -S amd-ucode
    fi

    #printf "Downloading bootloader and other packages\n" # 18/$steps
    #pacman -S grub networkmanager dialog wireless_tools wpa_supplicant os-prober mtools dosfstools base-devel linux-headers iwd dhcpcd

    counter=0;
    #pacman -S grub networkmanager dialog wireless_tools wpa_supplicant os-prober mtools dosfstools base-devel linux-headers iwd dhcpcd
    pid=$!;
    trap "kill $pid 2> /dev/null" EXIT;
    while kill -0 $pid 2> /dev/null; do
        (( counter+=1 ))
        echo $counter | dialog --title "$default_title 16/$steps" --gauge "Downloading bootloader and other packages" 7 50 0;
        sleep 0.1
    done;
    trap - EXIT
    counter=100
    echo $counter | dialog --title "$default_title 16/$steps" --gauge "Downloading bootloader and other packages" 7 50 0
    clear

    printf "Installing bootloader\n"
    dsknm=$(dialog --title "$default_title 17/$steps" --output-fd 1 --inputbox "What is the full name of your disk (/dev/sdX): " 25 90); clear
    #grub-install --target=i386-pc $dsknm
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
}

ncursesManualInstall2() {
    localtmzn=$(dialog --title "$default_title 8/$steps" --output-fd 1 --inputbox "Type 'MENU' to use a menu selection\nEnter the location of your timezone:\nExample: America/Sao_Paulo" 25 50); clear
    lwcs_localtmzn=$(echo "$localtmzn" | tr '[:upper:]' '[:lower:]')
    if [[ "$lwcs_localtmzn" == "menu" ]]; then # really big one line menu
        localtmzn=$(dialog --title "Mirai Arch" --no-tags --output-fd 1 --menu "Select your timezone" 20 50 0 "Africa/Abidjan" "Africa/Abidjan" "Africa/Accra" "Africa/Accra" "Africa/Addis_Ababa" "Africa/Addis_Ababa" "Africa/Algiers" "Africa/Algiers" "Africa/Asmara" "Africa/Asmara" "Africa/Asmera" "Africa/Asmera" "Africa/Bamako" "Africa/Bamako" "Africa/Bangui" "Africa/Bangui" "Africa/Banjul" "Africa/Banjul" "Africa/Bissau" "Africa/Bissau" "Africa/Blantyre" "Africa/Blantyre" "Africa/Brazzaville" "Africa/Brazzaville" "Africa/Bujumbura" "Africa/Bujumbura" "Africa/Cairo" "Africa/Cairo" "Africa/Casablanca" "Africa/Casablanca" "Africa/Ceuta" "Africa/Ceuta" "Africa/Conakry" "Africa/Conakry" "Africa/Dakar" "Africa/Dakar" "Africa/Dar_es_Salaam" "Africa/Dar_es_Salaam" "Africa/Djibouti" "Africa/Djibouti" "Africa/Douala" "Africa/Douala" "Africa/El_Aaiun" "Africa/El_Aaiun" "Africa/Freetown" "Africa/Freetown" "Africa/Gaborone" "Africa/Gaborone" "Africa/Harare" "Africa/Harare" "Africa/Johannesburg" "Africa/Johannesburg" "Africa/Juba" "Africa/Juba" "Africa/Kampala" "Africa/Kampala" "Africa/Khartoum" "Africa/Khartoum" "Africa/Kigali" "Africa/Kigali" "Africa/Kinshasa" "Africa/Kinshasa" "Africa/Lagos" "Africa/Lagos" "Africa/Libreville" "Africa/Libreville" "Africa/Lome" "Africa/Lome" "Africa/Luanda" "Africa/Luanda" "Africa/Lubumbashi" "Africa/Lubumbashi" "Africa/Lusaka" "Africa/Lusaka" "Africa/Malabo" "Africa/Malabo" "Africa/Maputo" "Africa/Maputo" "Africa/Maseru" "Africa/Maseru" "Africa/Mbabane" "Africa/Mbabane" "Africa/Mogadishu" "Africa/Mogadishu" "Africa/Monrovia" "Africa/Monrovia" "Africa/Nairobi" "Africa/Nairobi" "Africa/Ndjamena" "Africa/Ndjamena" "Africa/Niamey" "Africa/Niamey" "Africa/Nouakchott" "Africa/Nouakchott" "Africa/Ouagadougou" "Africa/Ouagadougou" "Africa/Porto-Novo" "Africa/Porto-Novo" "Africa/Sao_Tome" "Africa/Sao_Tome" "Africa/Timbuktu" "Africa/Timbuktu" "Africa/Tripoli" "Africa/Tripoli" "Africa/Tunis" "Africa/Tunis" "Africa/Windhoek" "Africa/Windhoek" "America/Adak" "America/Adak" "America/Anchorage" "America/Anchorage" "America/Anguilla" "America/Anguilla" "America/Antigua" "America/Antigua" "America/Araguaina" "America/Araguaina" "America/Argentina/Buenos_Aires" "America/Argentina/Buenos_Aires" "America/Argentina/Catamarca" "America/Argentina/Catamarca" "America/Argentina/ComodRivadavia" "America/Argentina/ComodRivadavia" "America/Argentina/Cordoba" "America/Argentina/Cordoba" "America/Argentina/Jujuy" "America/Argentina/Jujuy" "America/Argentina/La_Rioja" "America/Argentina/La_Rioja" "America/Argentina/Mendoza" "America/Argentina/Mendoza" "America/Argentina/Rio_Gallegos" "America/Argentina/Rio_Gallegos" "America/Argentina/Salta" "America/Argentina/Salta" "America/Argentina/San_Juan" "America/Argentina/San_Juan" "America/Argentina/San_Luis" "America/Argentina/San_Luis" "America/Argentina/Tucuman" "America/Argentina/Tucuman" "America/Argentina/Ushuaia" "America/Argentina/Ushuaia" "America/Aruba" "America/Aruba" "America/Asuncion" "America/Asuncion" "America/Atikokan" "America/Atikokan" "America/Atka" "America/Atka" "America/Bahia" "America/Bahia" "America/Bahia_Banderas" "America/Bahia_Banderas" "America/Barbados" "America/Barbados" "America/Belem" "America/Belem" "America/Belize" "America/Belize" "America/Blanc-Sablon" "America/Blanc-Sablon" "America/Boa_Vista" "America/Boa_Vista" "America/Bogota" "America/Bogota" "America/Boise" "America/Boise" "America/Buenos_Aires" "America/Buenos_Aires" "America/Cambridge_Bay" "America/Cambridge_Bay" "America/Campo_Grande" "America/Campo_Grande" "America/Cancun" "America/Cancun" "America/Caracas" "America/Caracas" "America/Catamarca" "America/Catamarca" "America/Cayenne" "America/Cayenne" "America/Cayman" "America/Cayman" "America/Chicago" "America/Chicago" "America/Chihuahua" "America/Chihuahua" "America/Coral_Harbour" "America/Coral_Harbour" "America/Cordoba" "America/Cordoba" "America/Costa_Rica" "America/Costa_Rica" "America/Creston" "America/Creston" "America/Cuiaba" "America/Cuiaba" "America/Curacao" "America/Curacao" "America/Danmarkshavn" "America/Danmarkshavn" "America/Dawson" "America/Dawson" "America/Dawson_Creek" "America/Dawson_Creek" "America/Denver" "America/Denver" "America/Detroit" "America/Detroit" "America/Dominica" "America/Dominica" "America/Edmonton" "America/Edmonton" "America/Eirunepe" "America/Eirunepe" "America/El_Salvador" "America/El_Salvador" "America/Ensenada" "America/Ensenada" "America/Fort_Nelson" "America/Fort_Nelson" "America/Fort_Wayne" "America/Fort_Wayne" "America/Fortaleza" "America/Fortaleza" "America/Glace_Bay" "America/Glace_Bay" "America/Godthab" "America/Godthab" "America/Goose_Bay" "America/Goose_Bay" "America/Grand_Turk" "America/Grand_Turk" "America/Grenada" "America/Grenada" "America/Guadeloupe" "America/Guadeloupe" "America/Guatemala" "America/Guatemala" "America/Guayaquil" "America/Guayaquil" "America/Guyana" "America/Guyana" "America/Halifax" "America/Halifax" "America/Havana" "America/Havana" "America/Hermosillo" "America/Hermosillo" "America/Indiana/Indianapolis" "America/Indiana/Indianapolis" "America/Indiana/Knox" "America/Indiana/Knox" "America/Indiana/Marengo" "America/Indiana/Marengo" "America/Indiana/Petersburg" "America/Indiana/Petersburg" "America/Indiana/Tell_City" "America/Indiana/Tell_City" "America/Indiana/Vevay" "America/Indiana/Vevay" "America/Indiana/Vincennes" "America/Indiana/Vincennes" "America/Indiana/Winamac" "America/Indiana/Winamac" "America/Indianapolis" "America/Indianapolis" "America/Inuvik" "America/Inuvik" "America/Iqaluit" "America/Iqaluit" "America/Jamaica" "America/Jamaica" "America/Jujuy" "America/Jujuy" "America/Juneau" "America/Juneau" "America/Kentucky/Louisville" "America/Kentucky/Louisville" "America/Kentucky/Monticello" "America/Kentucky/Monticello" "America/Knox_IN" "America/Knox_IN" "America/Kralendijk" "America/Kralendijk" "America/La_Paz" "America/La_Paz" "America/Lima" "America/Lima" "America/Los_Angeles" "America/Los_Angeles" "America/Louisville" "America/Louisville" "America/Lower_Princes" "America/Lower_Princes" "America/Maceio" "America/Maceio" "America/Managua" "America/Managua" "America/Manaus" "America/Manaus" "America/Marigot" "America/Marigot" "America/Martinique" "America/Martinique" "America/Matamoros" "America/Matamoros" "America/Mazatlan" "America/Mazatlan" "America/Mendoza" "America/Mendoza" "America/Menominee" "America/Menominee" "America/Merida" "America/Merida" "America/Metlakatla" "America/Metlakatla" "America/Mexico_City" "America/Mexico_City" "America/Miquelon" "America/Miquelon" "America/Moncton" "America/Moncton" "America/Monterrey" "America/Monterrey" "America/Montevideo" "America/Montevideo" "America/Montreal" "America/Montreal" "America/Montserrat" "America/Montserrat" "America/Nassau" "America/Nassau" "America/New_York" "America/New_York" "America/Nipigon" "America/Nipigon" "America/Nome" "America/Nome" "America/Noronha" "America/Noronha" "America/North_Dakota/Beulah" "America/North_Dakota/Beulah" "America/North_Dakota/Center" "America/North_Dakota/Center" "America/North_Dakota/New_Salem" "America/North_Dakota/New_Salem" "America/Nuuk" "America/Nuuk" "America/Ojinaga" "America/Ojinaga" "America/Panama" "America/Panama" "America/Pangnirtung" "America/Pangnirtung" "America/Paramaribo" "America/Paramaribo" "America/Phoenix" "America/Phoenix" "America/Port-au-Prince" "America/Port-au-Prince" "America/Port_of_Spain" "America/Port_of_Spain" "America/Porto_Acre" "America/Porto_Acre" "America/Porto_Velho" "America/Porto_Velho" "America/Puerto_Rico" "America/Puerto_Rico" "America/Punta_Arenas" "America/Punta_Arenas" "America/Rainy_River" "America/Rainy_River" "America/Rankin_Inlet" "America/Rankin_Inlet" "America/Recife" "America/Recife" "America/Regina" "America/Regina" "America/Resolute" "America/Resolute" "America/Rio_Branco" "America/Rio_Branco" "America/Rosario" "America/Rosario" "America/Santa_Isabel" "America/Santa_Isabel" "America/Santarem" "America/Santarem" "America/Santiago" "America/Santiago" "America/Santo_Domingo" "America/Santo_Domingo" "America/Sao_Paulo" "America/Sao_Paulo" "America/Scoresbysund" "America/Scoresbysund" "America/Shiprock" "America/Shiprock" "America/Sitka" "America/Sitka" "America/St_Barthelemy" "America/St_Barthelemy" "America/St_Johns" "America/St_Johns" "America/St_Kitts" "America/St_Kitts" "America/St_Lucia" "America/St_Lucia" "America/St_Thomas" "America/St_Thomas" "America/St_Vincent" "America/St_Vincent" "America/Swift_Current" "America/Swift_Current" "America/Tegucigalpa" "America/Tegucigalpa" "America/Thule" "America/Thule" "America/Thunder_Bay" "America/Thunder_Bay" "America/Tijuana" "America/Tijuana" "America/Toronto" "America/Toronto" "America/Tortola" "America/Tortola" "America/Vancouver" "America/Vancouver" "America/Virgin" "America/Virgin" "America/Whitehorse" "America/Whitehorse" "America/Winnipeg" "America/Winnipeg" "America/Yakutat" "America/Yakutat" "America/Yellowknife" "America/Yellowknife" "Antarctica/Casey" "Antarctica/Casey" "Antarctica/Davis" "Antarctica/Davis" "Antarctica/DumontDUrville" "Antarctica/DumontDUrville" "Antarctica/Macquarie" "Antarctica/Macquarie" "Antarctica/Mawson" "Antarctica/Mawson" "Antarctica/McMurdo" "Antarctica/McMurdo" "Antarctica/Palmer" "Antarctica/Palmer" "Antarctica/Rothera" "Antarctica/Rothera" "Antarctica/South_Pole" "Antarctica/South_Pole" "Antarctica/Syowa" "Antarctica/Syowa" "Antarctica/Troll" "Antarctica/Troll" "Antarctica/Vostok" "Antarctica/Vostok" "Arctic/Longyearbyen" "Arctic/Longyearbyen" "Asia/Aden" "Asia/Aden" "Asia/Almaty" "Asia/Almaty" "Asia/Amman" "Asia/Amman" "Asia/Anadyr" "Asia/Anadyr" "Asia/Aqtau" "Asia/Aqtau" "Asia/Aqtobe" "Asia/Aqtobe" "Asia/Ashgabat" "Asia/Ashgabat" "Asia/Ashkhabad" "Asia/Ashkhabad" "Asia/Atyrau" "Asia/Atyrau" "Asia/Baghdad" "Asia/Baghdad" "Asia/Bahrain" "Asia/Bahrain" "Asia/Baku" "Asia/Baku" "Asia/Bangkok" "Asia/Bangkok" "Asia/Barnaul" "Asia/Barnaul" "Asia/Beirut" "Asia/Beirut" "Asia/Bishkek" "Asia/Bishkek" "Asia/Brunei" "Asia/Brunei" "Asia/Calcutta" "Asia/Calcutta" "Asia/Chita" "Asia/Chita" "Asia/Choibalsan" "Asia/Choibalsan" "Asia/Chongqing" "Asia/Chongqing" "Asia/Chungking" "Asia/Chungking" "Asia/Colombo" "Asia/Colombo" "Asia/Dacca" "Asia/Dacca" "Asia/Damascus" "Asia/Damascus" "Asia/Dhaka" "Asia/Dhaka" "Asia/Dili" "Asia/Dili" "Asia/Dubai" "Asia/Dubai" "Asia/Dushanbe" "Asia/Dushanbe" "Asia/Famagusta" "Asia/Famagusta" "Asia/Gaza" "Asia/Gaza" "Asia/Harbin" "Asia/Harbin" "Asia/Hebron" "Asia/Hebron" "Asia/Ho_Chi_Minh" "Asia/Ho_Chi_Minh" "Asia/Hong_Kong" "Asia/Hong_Kong" "Asia/Hovd" "Asia/Hovd" "Asia/Irkutsk" "Asia/Irkutsk" "Asia/Istanbul" "Asia/Istanbul" "Asia/Jakarta" "Asia/Jakarta" "Asia/Jayapura" "Asia/Jayapura" "Asia/Jerusalem" "Asia/Jerusalem" "Asia/Kabul" "Asia/Kabul" "Asia/Kamchatka" "Asia/Kamchatka" "Asia/Karachi" "Asia/Karachi" "Asia/Kashgar" "Asia/Kashgar" "Asia/Kathmandu" "Asia/Kathmandu" "Asia/Katmandu" "Asia/Katmandu" "Asia/Khandyga" "Asia/Khandyga" "Asia/Kolkata" "Asia/Kolkata" "Asia/Krasnoyarsk" "Asia/Krasnoyarsk" "Asia/Kuala_Lumpur" "Asia/Kuala_Lumpur" "Asia/Kuching" "Asia/Kuching" "Asia/Kuwait" "Asia/Kuwait" "Asia/Macao" "Asia/Macao" "Asia/Macau" "Asia/Macau" "Asia/Magadan" "Asia/Magadan" "Asia/Makassar" "Asia/Makassar" "Asia/Manila" "Asia/Manila" "Asia/Muscat" "Asia/Muscat" "Asia/Nicosia" "Asia/Nicosia" "Asia/Novokuznetsk" "Asia/Novokuznetsk" "Asia/Novosibirsk" "Asia/Novosibirsk" "Asia/Omsk" "Asia/Omsk" "Asia/Oral" "Asia/Oral" "Asia/Phnom_Penh" "Asia/Phnom_Penh" "Asia/Pontianak" "Asia/Pontianak" "Asia/Pyongyang" "Asia/Pyongyang" "Asia/Qatar" "Asia/Qatar" "Asia/Qostanay" "Asia/Qostanay" "Asia/Qyzylorda" "Asia/Qyzylorda" "Asia/Rangoon" "Asia/Rangoon" "Asia/Riyadh" "Asia/Riyadh" "Asia/Saigon" "Asia/Saigon" "Asia/Sakhalin" "Asia/Sakhalin" "Asia/Samarkand" "Asia/Samarkand" "Asia/Seoul" "Asia/Seoul" "Asia/Shanghai" "Asia/Shanghai" "Asia/Singapore" "Asia/Singapore" "Asia/Srednekolymsk" "Asia/Srednekolymsk" "Asia/Taipei" "Asia/Taipei" "Asia/Tashkent" "Asia/Tashkent" "Asia/Tbilisi" "Asia/Tbilisi" "Asia/Tehran" "Asia/Tehran" "Asia/Tel_Aviv" "Asia/Tel_Aviv" "Asia/Thimbu" "Asia/Thimbu" "Asia/Thimphu" "Asia/Thimphu" "Asia/Tokyo" "Asia/Tokyo" "Asia/Tomsk" "Asia/Tomsk" "Asia/Ujung_Pandang" "Asia/Ujung_Pandang" "Asia/Ulaanbaatar" "Asia/Ulaanbaatar" "Asia/Ulan_Bator" "Asia/Ulan_Bator" "Asia/Urumqi" "Asia/Urumqi" "Asia/Ust-Nera" "Asia/Ust-Nera" "Asia/Vientiane" "Asia/Vientiane" "Asia/Vladivostok" "Asia/Vladivostok" "Asia/Yakutsk" "Asia/Yakutsk" "Asia/Yangon" "Asia/Yangon" "Asia/Yekaterinburg" "Asia/Yekaterinburg" "Asia/Yerevan" "Asia/Yerevan" "Atlantic/Azores" "Atlantic/Azores" "Atlantic/Bermuda" "Atlantic/Bermuda" "Atlantic/Canary" "Atlantic/Canary" "Atlantic/Cape_Verde" "Atlantic/Cape_Verde" "Atlantic/Faeroe" "Atlantic/Faeroe" "Atlantic/Faroe" "Atlantic/Faroe" "Atlantic/Jan_Mayen" "Atlantic/Jan_Mayen" "Atlantic/Madeira" "Atlantic/Madeira" "Atlantic/Reykjavik" "Atlantic/Reykjavik" "Atlantic/South_Georgia" "Atlantic/South_Georgia" "Atlantic/St_Helena" "Atlantic/St_Helena" "Atlantic/Stanley" "Atlantic/Stanley" "Australia/ACT" "Australia/ACT" "Australia/Adelaide" "Australia/Adelaide" "Australia/Brisbane" "Australia/Brisbane" "Australia/Broken_Hill" "Australia/Broken_Hill" "Australia/Canberra" "Australia/Canberra" "Australia/Currie" "Australia/Currie" "Australia/Darwin" "Australia/Darwin" "Australia/Eucla" "Australia/Eucla" "Australia/Hobart" "Australia/Hobart" "Australia/LHI" "Australia/LHI" "Australia/Lindeman" "Australia/Lindeman" "Australia/Lord_Howe" "Australia/Lord_Howe" "Australia/Melbourne" "Australia/Melbourne" "Australia/NSW" "Australia/NSW" "Australia/North" "Australia/North" "Australia/Perth" "Australia/Perth" "Australia/Queensland" "Australia/Queensland" "Australia/South" "Australia/South" "Australia/Sydney" "Australia/Sydney" "Australia/Tasmania" "Australia/Tasmania" "Australia/Victoria" "Australia/Victoria" "Australia/West" "Australia/West" "Australia/Yancowinna" "Australia/Yancowinna" "Brazil/Acre" "Brazil/Acre" "Brazil/DeNoronha" "Brazil/DeNoronha" "Brazil/East" "Brazil/East" "Brazil/West" "Brazil/West" "CET" "CET" "CST6CDT" "CST6CDT" "Canada/Atlantic" "Canada/Atlantic" "Canada/Central" "Canada/Central" "Canada/Eastern" "Canada/Eastern" "Canada/Mountain" "Canada/Mountain" "Canada/Newfoundland" "Canada/Newfoundland" "Canada/Pacific" "Canada/Pacific" "Canada/Saskatchewan" "Canada/Saskatchewan" "Canada/Yukon" "Canada/Yukon" "Chile/Continental" "Chile/Continental" "Chile/EasterIsland" "Chile/EasterIsland" "Cuba" "Cuba" "EET" "EET" "EST" "EST" "EST5EDT" "EST5EDT" "Egypt" "Egypt" "Eire" "Eire" "Etc/GMT" "Etc/GMT" "Etc/GMT+0" "Etc/GMT+0" "Etc/GMT+1" "Etc/GMT+1" "Etc/GMT+10" "Etc/GMT+10" "Etc/GMT+11" "Etc/GMT+11" "Etc/GMT+12" "Etc/GMT+12" "Etc/GMT+2" "Etc/GMT+2" "Etc/GMT+3" "Etc/GMT+3" "Etc/GMT+4" "Etc/GMT+4" "Etc/GMT+5" "Etc/GMT+5" "Etc/GMT+6" "Etc/GMT+6" "Etc/GMT+7" "Etc/GMT+7" "Etc/GMT+8" "Etc/GMT+8" "Etc/GMT+9" "Etc/GMT+9" "Etc/GMT-0" "Etc/GMT-0" "Etc/GMT-1" "Etc/GMT-1" "Etc/GMT-10" "Etc/GMT-10" "Etc/GMT-11" "Etc/GMT-11" "Etc/GMT-12" "Etc/GMT-12" "Etc/GMT-13" "Etc/GMT-13" "Etc/GMT-14" "Etc/GMT-14" "Etc/GMT-2" "Etc/GMT-2" "Etc/GMT-3" "Etc/GMT-3" "Etc/GMT-4" "Etc/GMT-4" "Etc/GMT-5" "Etc/GMT-5" "Etc/GMT-6" "Etc/GMT-6" "Etc/GMT-7" "Etc/GMT-7" "Etc/GMT-8" "Etc/GMT-8" "Etc/GMT-9" "Etc/GMT-9" "Etc/GMT0" "Etc/GMT0" "Etc/Greenwich" "Etc/Greenwich" "Etc/UCT" "Etc/UCT" "Etc/UTC" "Etc/UTC" "Etc/Universal" "Etc/Universal" "Etc/Zulu" "Etc/Zulu" "Europe/Amsterdam" "Europe/Amsterdam" "Europe/Andorra" "Europe/Andorra" "Europe/Astrakhan" "Europe/Astrakhan" "Europe/Athens" "Europe/Athens" "Europe/Belfast" "Europe/Belfast" "Europe/Belgrade" "Europe/Belgrade" "Europe/Berlin" "Europe/Berlin" "Europe/Bratislava" "Europe/Bratislava" "Europe/Brussels" "Europe/Brussels" "Europe/Bucharest" "Europe/Bucharest" "Europe/Budapest" "Europe/Budapest" "Europe/Busingen" "Europe/Busingen" "Europe/Chisinau" "Europe/Chisinau" "Europe/Copenhagen" "Europe/Copenhagen" "Europe/Dublin" "Europe/Dublin" "Europe/Gibraltar" "Europe/Gibraltar" "Europe/Guernsey" "Europe/Guernsey" "Europe/Helsinki" "Europe/Helsinki" "Europe/Isle_of_Man" "Europe/Isle_of_Man" "Europe/Istanbul" "Europe/Istanbul" "Europe/Jersey" "Europe/Jersey" "Europe/Kaliningrad" "Europe/Kaliningrad" "Europe/Kiev" "Europe/Kiev" "Europe/Kirov" "Europe/Kirov" "Europe/Lisbon" "Europe/Lisbon" "Europe/Ljubljana" "Europe/Ljubljana" "Europe/London" "Europe/London" "Europe/Luxembourg" "Europe/Luxembourg" "Europe/Madrid" "Europe/Madrid" "Europe/Malta" "Europe/Malta" "Europe/Mariehamn" "Europe/Mariehamn" "Europe/Minsk" "Europe/Minsk" "Europe/Monaco" "Europe/Monaco" "Europe/Moscow" "Europe/Moscow" "Europe/Nicosia" "Europe/Nicosia" "Europe/Oslo" "Europe/Oslo" "Europe/Paris" "Europe/Paris" "Europe/Podgorica" "Europe/Podgorica" "Europe/Prague" "Europe/Prague" "Europe/Riga" "Europe/Riga" "Europe/Rome" "Europe/Rome" "Europe/Samara" "Europe/Samara" "Europe/San_Marino" "Europe/San_Marino" "Europe/Sarajevo" "Europe/Sarajevo" "Europe/Saratov" "Europe/Saratov" "Europe/Simferopol" "Europe/Simferopol" "Europe/Skopje" "Europe/Skopje" "Europe/Sofia" "Europe/Sofia" "Europe/Stockholm" "Europe/Stockholm" "Europe/Tallinn" "Europe/Tallinn" "Europe/Tirane" "Europe/Tirane" "Europe/Tiraspol" "Europe/Tiraspol" "Europe/Ulyanovsk" "Europe/Ulyanovsk" "Europe/Uzhgorod" "Europe/Uzhgorod" "Europe/Vaduz" "Europe/Vaduz" "Europe/Vatican" "Europe/Vatican" "Europe/Vienna" "Europe/Vienna" "Europe/Vilnius" "Europe/Vilnius" "Europe/Volgograd" "Europe/Volgograd" "Europe/Warsaw" "Europe/Warsaw" "Europe/Zagreb" "Europe/Zagreb" "Europe/Zaporozhye" "Europe/Zaporozhye" "Europe/Zurich" "Europe/Zurich" "Factory" "Factory" "GB" "GB" "GB-Eire" "GB-Eire" "GMT" "GMT" "GMT+0" "GMT+0" "GMT-0" "GMT-0" "GMT0" "GMT0" "Greenwich" "Greenwich" "HST" "HST" "Hongkong" "Hongkong" "Iceland" "Iceland" "Indian/Antananarivo" "Indian/Antananarivo" "Indian/Chagos" "Indian/Chagos" "Indian/Christmas" "Indian/Christmas" "Indian/Cocos" "Indian/Cocos" "Indian/Comoro" "Indian/Comoro" "Indian/Kerguelen" "Indian/Kerguelen" "Indian/Mahe" "Indian/Mahe" "Indian/Maldives" "Indian/Maldives" "Indian/Mauritius" "Indian/Mauritius" "Indian/Mayotte" "Indian/Mayotte" "Indian/Reunion" "Indian/Reunion" "Iran" "Iran" "Israel" "Israel" "Jamaica" "Jamaica" "Japan" "Japan" "Kwajalein" "Kwajalein" "Libya" "Libya" "MET" "MET" "MST" "MST" "MST7MDT" "MST7MDT" "Mexico/BajaNorte" "Mexico/BajaNorte" "Mexico/BajaSur" "Mexico/BajaSur" "Mexico/General" "Mexico/General" "NZ" "NZ" "NZ-CHAT" "NZ-CHAT" "Navajo" "Navajo" "PRC" "PRC" "PST8PDT" "PST8PDT" "Pacific/Apia" "Pacific/Apia" "Pacific/Auckland" "Pacific/Auckland" "Pacific/Bougainville" "Pacific/Bougainville" "Pacific/Chatham" "Pacific/Chatham" "Pacific/Chuuk" "Pacific/Chuuk" "Pacific/Easter" "Pacific/Easter" "Pacific/Efate" "Pacific/Efate" "Pacific/Enderbury" "Pacific/Enderbury" "Pacific/Fakaofo" "Pacific/Fakaofo" "Pacific/Fiji" "Pacific/Fiji" "Pacific/Funafuti" "Pacific/Funafuti" "Pacific/Galapagos" "Pacific/Galapagos" "Pacific/Gambier" "Pacific/Gambier" "Pacific/Guadalcanal" "Pacific/Guadalcanal" "Pacific/Guam" "Pacific/Guam" "Pacific/Honolulu" "Pacific/Honolulu" "Pacific/Johnston" "Pacific/Johnston" "Pacific/Kanton" "Pacific/Kanton" "Pacific/Kiritimati" "Pacific/Kiritimati" "Pacific/Kosrae" "Pacific/Kosrae" "Pacific/Kwajalein" "Pacific/Kwajalein" "Pacific/Majuro" "Pacific/Majuro" "Pacific/Marquesas" "Pacific/Marquesas" "Pacific/Midway" "Pacific/Midway" "Pacific/Nauru" "Pacific/Nauru" "Pacific/Niue" "Pacific/Niue" "Pacific/Norfolk" "Pacific/Norfolk" "Pacific/Noumea" "Pacific/Noumea" "Pacific/Pago_Pago" "Pacific/Pago_Pago" "Pacific/Palau" "Pacific/Palau" "Pacific/Pitcairn" "Pacific/Pitcairn" "Pacific/Pohnpei" "Pacific/Pohnpei" "Pacific/Ponape" "Pacific/Ponape" "Pacific/Port_Moresby" "Pacific/Port_Moresby" "Pacific/Rarotonga" "Pacific/Rarotonga" "Pacific/Saipan" "Pacific/Saipan" "Pacific/Samoa" "Pacific/Samoa" "Pacific/Tahiti" "Pacific/Tahiti" "Pacific/Tarawa" "Pacific/Tarawa" "Pacific/Tongatapu" "Pacific/Tongatapu" "Pacific/Truk" "Pacific/Truk" "Pacific/Wake" "Pacific/Wake" "Pacific/Wallis" "Pacific/Wallis" "Pacific/Yap" "Pacific/Yap" "Poland" "Poland" "Portugal" "Portugal" "ROC" "ROC" "ROK" "ROK" "Singapore" "Singapore" "Turkey" "Turkey" "UCT" "UCT" "US/Alaska" "US/Alaska" "US/Aleutian" "US/Aleutian" "US/Arizona" "US/Arizona" "US/Central" "US/Central" "US/East-Indiana" "US/East-Indiana" "US/Eastern" "US/Eastern" "US/Hawaii" "US/Hawaii" "US/Indiana-Starke" "US/Indiana-Starke" "US/Michigan" "US/Michigan" "US/Mountain" "US/Mountain" "US/Pacific" "US/Pacific" "US/Samoa" "US/Samoa" "UTC" "UTC" "Universal" "Universal" "W-SU" "W-SU" "WET" "WET" "Zulu" "Zulu"); clear
        echo "$localtmzn"
    else
        if [[ "$localtmzn" =~ .*"/usr/share/zoneinfo/".* ]]; then
            tmzn=$(echo "$localtmzn" | cut -d/ -f5-)
        else
            tmzn=$localtmzn
        fi
    fi
    #ln -sf /usr/share/zoneinfo/$tmzn /etc/localtime

    dialog --title "$default_title 9/$steps" --infobox "\nSyncronizing the hardware clock to the system clock" 25 90 && clear
    #hwclock --systohc

    localelang=$(dialog --title "$default_title 10/$steps" --output-fd 1 --inputbox "Type 'MENU' to use a menu selection\nEnter the languages that you want to use on your system:\nNOTE: They need to be separated by comma (,) without spaces in between\nExample: pt_BR.UTF-8 UTF-8,en_US ISO-8859-1" 25 75); clear
    IFS=',' read -r -a splt_localelang <<< "$localelang"
    usrlocaleslist=()
    for element in "${splt_localelang[@]}"; do
        sed -i "s/#$element/$element/" #/etc/locale.gen
        usrlocaleslist+="\"$element\" \"$element\" "
    done
    cmd=("dialog --title \"Mirai Arch\" --no-tags --output-fd 1 --menu \"Select the language that you will use on your system\" 25 75 0")
    _cmd=$(echo "${cmd[@]}" "${usrlocaleslist[@]}")
    usrlocale=$(bash -c "$_cmd")
    clear
    dialog --title "$default_title 10/$steps" --infobox "Saving the locale in /etc/locale.conf" 25 90 && clear
    if [[ -e /etc/locale.conf ]]; then
        echo "LANG=$usrlocale" >> /etc/locale.conf
    else
        touch /etc/locale.conf
        echo "LANG=$usrlocale" >> /etc/locale.conf
    fi

    #######################################################################################################################################
    # SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP #
    # SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP #
    # SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP #
    # SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP #
    # SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP SAVE KEYMAP #
    #######################################################################################################################################

    hstnm=$(dialog --title "$default_title 12/$steps" --output-fd 1 --inputbox "Enter your hostname:" 25 75); clear
    if [[ -e /etc/hostname ]]; then
        echo $hstnm >> /etc/hostname
    else
        touch /etc/hostname
        echo $hstnm >> /etc/hostname
    fi

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
    if dialog --title "$default_title 13/$steps" --yes-label "Yes" --no-label "No" --yesno "Creating hosts files\nDo you want to add custom hosts to it?" 7 64; then
        clear
        counter=0;
        #curl -fL "https://raw.githubusercontent.com/MiraiMindz/.dotfiles/main/misc/MiraiHosts.txt" >> /etc/hosts
        pid=$!;
        trap "kill $pid 2> /dev/null" EXIT;
        while kill -0 $pid 2> /dev/null; do
            (( counter+=1 ))
            echo $counter | dialog --title "$default_title 13/$steps" --gauge "Adding custom hosts" 7 50 0;
            sleep 0.1
        done;
        trap - EXIT
        counter=100
        echo $counter | dialog --title "$default_title 13/$steps" --gauge "Adding custom hosts" 7 50 0
        clear
    fi

    rtpass=$(dialog --title "$default_title 14/$steps" --output-fd 1 --passwordbox "Enter the root password: " 12 50); clear
    #echo -e "$rtpass\n$rtpass" | sudo passwd -q

    if dialog --title "$default_title 15/$steps" --yes-label "Intel" --no-label "AMD" --yesno "Is your processor Intel or AMD?" 7 64; then
        #pacman -S intel-ucode
    else
        #pacman -S amd-ucode
    fi

    instpkgs=$(dialog --title "$default_title 16/$steps" --output-fd 1 --inputbox "Type \"AUTO\" to install the packages from the auto installation\nAUTO=grub networkmanager dialog wireless_tools wpa_supplicant os-prober mtools dosfstools base-devel linux-headers iwd dhcpcd\nEnter the packages that you want to install:" 25 75); clear
    lwcs_instpkgs=$(echo "$instpkgs" | tr '[:upper:]' '[:lower:]')
    if [[ "$lwcs_instpkgs" == "auto" ]]; then
        #pacman -S grub networkmanager dialog wireless_tools wpa_supplicant os-prober mtools dosfstools base-devel linux-headers iwd dhcpcd
    else
        IFS=' ' read -r -a splt_pkgs <<< "$instpkgs"
        #pacman -S "${instpkgs[@]}"
    fi

    if [[ -e $(command -v grub) ]]; fi # FINISH
    tgt_plat=$(dialog --title "$default_title" --no-tags --output-fd 1 --menu "Select your desired platform to install GRUB: " 7 60 0 "arm-coreboot" "arm-coreboot" "arm-efi" "arm-efi" "arm-uboot" "arm-uboot" "arm64-efi" "arm64-efi" "i386-coreboot" "i386-coreboot" "i386-efi" "i386-efi" "i386-ieee1275" "i386-ieee1275" "i386-multiboot" "i386-multiboot" "i386-pc" "i386-pc" "i386-qemu" "i386-qemu" "i386-xen" "i386-xen" "i386-xen_pvh" "i386-xen_pvh" "ia64-efi" "ia64-efi" "mips-arc" "mips-arc" "mips-qemu_mips" "mips-qemu_mips" "mipsel-arc" "mipsel-arc" "mipsel-loongson" "mipsel-loongson" "mipsel-qemu_mips" "mipsel-qemu_mips" "powerpc-ieee1275" "powerpc-ieee1275" "riscv32-efi" "riscv32-efi" "riscv64-efi" "riscv64-efi" "sparc64-ieee1275" "sparc64-ieee1275" "x86_64-efi" "x86_64-efi" "x86_64-xen" "x86_64-xen"); clear
    _grb_dsk=$(dialog --title "$default_title 17/$steps" --output-fd 1 --inputbox "Enter the name of the disk to install GRUB:" 25 75); clear
    if [[ "$_grb_dsk" =~ .*"/dev/".* ]]; then
        grb_dsk=$(echo "$_grb_dsk" | cut -d/ -f3-)
    else
        grb_dsk=$_grb_dsk
    fi
    #grub-install --target=$tgt_plat /dev/$grb_dsk
    #grub-mkconfig -o /boot/grub/grub.cfg
}

# Checks if dialog exists, if not go with default cli installation
if [[ ! -e $(command -v dialog) || ! -f /usr/bin/dialog || ! -f /bin/dialog ]]; then
    printf "%s\n" "dialog not found, proceeding with default cli installation"
    cliInstall
fi

# Checks the method used in the first script
if [[ "$installmtd" == "auto" ]]; then
    ncursesAutoInstall2
elif [[ "$installmtd" == "manu" ]]; then
    ncursesManualInstall2
else
    if [[ -e $(command -v dialog) || -f /usr/bin/dialog || -f /bin/dialog ]]; then
        if dialog --title "$default_title" --yes-label "Automated" --no-label "Manual" --yesno "\nPrevious install method not found\nDo you want the Automated Install or the Manual Install?" 7 64; then
            clear
            ncursesAutoInstall
        else
            clear
            ncursesManualInstall2
        fi
    fi
fi


