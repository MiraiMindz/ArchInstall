default_title="Arch Linux Mirai Install"
steps="19"

tgt_plat=$(dialog --title "$default_title" --no-tags --output-fd 1 --menu "Select your desired platform to install GRUB: " 7 60 0 "arm-coreboot" "arm-coreboot" "arm-efi" "arm-efi" "arm-uboot" "arm-uboot" "arm64-efi" "arm64-efi" "i386-coreboot" "i386-coreboot" "i386-efi" "i386-efi" "i386-ieee1275" "i386-ieee1275" "i386-multiboot" "i386-multiboot" "i386-pc" "i386-pc" "i386-qemu" "i386-qemu" "i386-xen" "i386-xen" "i386-xen_pvh" "i386-xen_pvh" "ia64-efi" "ia64-efi" "mips-arc" "mips-arc" "mips-qemu_mips" "mips-qemu_mips" "mipsel-arc" "mipsel-arc" "mipsel-loongson" "mipsel-loongson" "mipsel-qemu_mips" "mipsel-qemu_mips" "powerpc-ieee1275" "powerpc-ieee1275" "riscv32-efi" "riscv32-efi" "riscv64-efi" "riscv64-efi" "sparc64-ieee1275" "sparc64-ieee1275" "x86_64-efi" "x86_64-efi" "x86_64-xen" "x86_64-xen"); clear
_grb_dsk=$(dialog --title "$default_title 17/$steps" --output-fd 1 --inputbox "Enter the name of the disk to install:" 25 75); clear
if [[ "$_grb_dsk" =~ .*"/dev/".* ]]; then
    grb_dsk=$(echo "$_grb_dsk" | cut -d/ -f3-)
else
    grb_dsk=$_grb_dsk
fi

tgt_plat=$(dialog --title "$default_title" --no-tags --output-fd 1 --menu "Select your desired platform to install GRUB: " 7 60 0 "arm-coreboot" "arm-coreboot" "arm-efi" "arm-efi" "arm-uboot" "arm-uboot" "arm64-efi" "arm64-efi" "i386-coreboot" "i386-coreboot" "i386-efi" "i386-efi" "i386-ieee1275" "i386-ieee1275" "i386-multiboot" "i386-multiboot" "i386-pc" "i386-pc" "i386-qemu" "i386-qemu" "i386-xen" "i386-xen" "i386-xen_pvh" "i386-xen_pvh" "ia64-efi" "ia64-efi" "mips-arc" "mips-arc" "mips-qemu_mips" "mips-qemu_mips" "mipsel-arc" "mipsel-arc" "mipsel-loongson" "mipsel-loongson" "mipsel-qemu_mips" "mipsel-qemu_mips" "powerpc-ieee1275" "powerpc-ieee1275" "riscv32-efi" "riscv32-efi" "riscv64-efi" "riscv64-efi" "sparc64-ieee1275" "sparc64-ieee1275" "x86_64-efi" "x86_64-efi" "x86_64-xen" "x86_64-xen"); clear


#"arm-coreboot" "arm-coreboot" "arm-efi" "arm-efi" "arm-uboot" "arm-uboot" "arm64-efi" "arm64-efi" "i386-coreboot" "i386-coreboot" "i386-efi" "i386-efi" "i386-ieee1275" "i386-ieee1275" "i386-multiboot" "i386-multiboot" "i386-pc" "i386-pc" "i386-qemu" "i386-qemu" "i386-xen" "i386-xen" "i386-xen_pvh" "i386-xen_pvh" "ia64-efi" "ia64-efi" "mips-arc" "mips-arc" "mips-qemu_mips" "mips-qemu_mips" "mipsel-arc" "mipsel-arc" "mipsel-loongson" "mipsel-loongson" "mipsel-qemu_mips" "mipsel-qemu_mips" "powerpc-ieee1275" "powerpc-ieee1275" "riscv32-efi" "riscv32-efi" "riscv64-efi" "riscv64-efi" "sparc64-ieee1275" "sparc64-ieee1275" "x86_64-efi" "x86_64-efi" "x86_64-xen" "x86_64-xen"
