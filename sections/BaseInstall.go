package sections

import (
	"ArchInstall/helpers"
	"fmt"
	"strings"
)

/*
"This step will guide you through the base configuration and installation of the system.",
"You will be prompted to select between some pre-defined setups.",
"But you can also select the between semi-manual and full-manual configuration.",
"The semi-manual configuration is similar to the full, but with some options omitted",
"to not conflict with the Desktop Environment packages.",
"The full-manual configuration includes:",
"  - Selecting between Desktop Manager and Window Manager setup.",
"    Note: Selecting a Desktop Manager will automatically trigger the \"Semi-Manual\" installation.",
"  - Common programming languages setup.",
"  - Selecting default SHELL.",
"  - Selecting default terminal emulator.",
"  - Selecting default Display Manager.",
"  - Selecting default Compositor.",
"  - Selecting desired Code Editor.",
"  - Selecting user folders (Documents, Downloads, etc...).",
"  - Compatibility Layers (Wine/Darling).",
"  - Gaming Setup.",
"  - GTK Themes:",
"    - Custom Fonts.",
"    - Color Theme.",
"    - Cursor Theme.",
"    - Icon Theme.",
"  - Other small utilities.\n",

*/

func setNetworkConnection() {
	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Network")
	helpers.PacmanInstallPackages("networkmanager", "dhclient")
	helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "systemctl", "enable", "--now", "NetworkManager")
}

func setPacman(cfgFile string) {
	//parallelDowns := helpers.JsonGetter(cfgFile, "pacmanParallelDownloads")
	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Pacman Mirrors")
	//helpers.CopyFile("/etc/pacman.d/mirrorlist", "/etc/pacman.d/mirrorlist.bak")
	//helpers.ReplaceFileLine("/etc/pacman.conf", "#ParallelDownloads", fmt.Sprintf("ParallelDownloads = %s", parallelDowns))
	//helpers.ReplaceFileLine("/etc/pacman.conf", "#Color", "Color")
	//helpers.ReplaceFileLine("/etc/pacman.conf", "#CheckSpace", "CheckSpace")
	//helpers.ReplaceFileLine("/etc/pacman.conf", "#VerbosePkgLists", "VerbosePkgLists")
	//helpers.ReplaceFileLine("/etc/pacman.conf", "#[community]", "[community]")
	//helpers.ReplaceFileLine("/etc/pacman.conf", "#[multilib]", "[multilib]")
	helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "pacman", "-Sy", "--noconfirm", "--needed")
	helpers.PacmanInstallPackages("pacman-contrib", "curl", "reflector", "grub", "git")
}

func setMakeVars() {
	memInfo := helpers.GetLine("/proc/meminfo", "MemTotal")
	totalMem := helpers.ExtractNumbers(memInfo.(string))[0] * helpers.KiB
	cpu_processors := helpers.GetOccurrences("/proc/cpuinfo", "processor")

	if totalMem > (8*helpers.GiB) && helpers.YesNo("Would you like to set custom MAKEFILE and COMPRESSXZ variables for performance in operations?") {
		helpers.ClearConsole()
		helpers.PrintHeader("Base Install", "MAKEFILE Variables")
		helpers.ReplaceFileLine("/etc/makepkg.conf", "#MAKEFLAGS=\"-j2\"", fmt.Sprintf("MAKEFLAGS=\"-j%d\"", cpu_processors))
		helpers.ReplaceFileLine("/etc/makepkg.conf", "COMPRESSXZ=(xz -c -z -)", fmt.Sprintf("COMPRESSXZ=(xz -c -T %d -z -)", cpu_processors))
	}
}

func installMicroCode() {
	cmdOut := helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "lscpu")
	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Processor Micro-Code")
	if helpers.GrepCommandOut(cmdOut, "GenuineIntel") != nil {
		fmt.Println("Installing INTEL Micro-Code")
		helpers.PacmanInstallPackages("intel-ucode")
	} else if helpers.GrepCommandOut(cmdOut, "AuthenticAMD") != nil {
		fmt.Println("Installing AMD Micro-Code")
		helpers.PacmanInstallPackages("amd-ucode")
	} else {
		fmt.Println(helpers.PrintError("UNABLE TO FIND PROCESSOR MICRO-CODE."))
	}
}

func installGPUDrivers() {
	cmdOut := helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "lspci")
	nvidiaCheck := helpers.GrepCommandOut(cmdOut, "NVIDIA") != nil || helpers.GrepCommandOut(cmdOut, "GeForce") != nil
	amdCheck := helpers.GrepCommandOut(cmdOut, "AMD") != nil || helpers.GrepCommandOut(cmdOut, "Radeon") != nil
	integratedCheck := helpers.GrepCommandOut(cmdOut, "Integrated Graphics Controller") != nil || helpers.GrepCommandOut(cmdOut, "Intel Corporation UHD") != nil || helpers.GrepCommandOut(cmdOut, "VGA") != nil
	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "GPU Drivers")
	if nvidiaCheck {
		helpers.PacmanInstallPackages("nvidia")
		helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "nvidia-xconfig")
	} else if amdCheck {
		helpers.PacmanInstallPackages("xf86-video-amdgpu")
	} else if integratedCheck {
		helpers.PacmanInstallPackages("libva-intel-driver", "libvdpau-va-gl", "lib32-vulkan-intel", "vulkan-intel", "libva-intel-driver", "libva-utils", "lib32-mesa")
	} else {
		fmt.Println(helpers.PrintError("UNABLE TO FIND GPU DRIVERS."))
	}
}

func installHelper(cfgFile string) {
	pkgHelper := helpers.JsonGetter(cfgFile, "aurHelper")
	switch pkgHelper {
	case "nix": // nix-shell -p nix-info --run "nix-info -m"
		resp := helpers.CurlResponse("https://nixos.org/nix/install")
		helpers.WriteToFile(fmt.Sprintf("%s/install_nix.sh", helpers.GetCurrDirPath()), resp, 0644)
		if helpers.YesNo("Would you like to check the downloaded script?") {
			helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "less", fmt.Sprintf("%s/install_nix.sh", helpers.GetCurrDirPath()))
		}
		if helpers.YesNo("Would you like to run the installation script?") {
			helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "sh", fmt.Sprintf("%s/install_nix.sh", helpers.GetCurrDirPath()), "--daemon")
			helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "nix-shell", "-p", "nix-info", "--run", "\"nix-info -m\"")
		}
	case "none":
		return
	default:
		helpers.InstallAurPkgWithoutHelper(pkgHelper)
	}
}

//region Base System Install

func _selectDE(cfgFile string) {
	var pkgName string
	deOpts := []helpers.ItemInfo{
		{Item: "Budgie", Info: "Designed with the modern user in mind, it focuses on simplicity and elegance."},
		{Item: "Cinnamon", Info: "Strives to provide a traditional user experience. is a fork of GNOME 3."},
		{Item: "Cutefish", Info: "New and modern desktop environment."},
		{Item: "Deepin", Info: "Feature an intuitive and elegant design. Moving around, sharing and searching etc. has become simply a joyful experience."},
		{Item: "Enlightenment", Info: "Provides an efficient environment. It supports themes, while still being capable of performing on older hardware or embedded devices."},
		{Item: "GNOME", Info: "Is an attractive and intuitive desktop environment."},
		{Item: "Plasma", Info: "Is a familiar working environment. Offers all the tools required for a modern computing experience. Be productive."},
		{Item: "LXDE", Info: "Is a fast and energy-saving desktop environment. It comes with a modern interface, and other utilities."},
		{Item: "LXQt", Info: "A lightweight, modular, blazing-fast and user-friendly desktop environment. Same as LXDE but with QT instead of GTK."},
		{Item: "MATE", Info: "Provides an intuitive and attractive desktop to Linux users using traditional metaphors"},
		{Item: "Sugar", Info: "Composed of Activities designed to help children from 5 to 12 years of age learn together through rich-media expression. Educational."},
		{Item: "UKUI", Info: "Lightweight, developed based on GTK and Qt."},
		{Item: "Xfce", Info: "Embodies the traditional UNIX philosophy of modularity and re-usability."},
		{Item: "RETURN", Info: "Go back to the previous menu"},
	}

	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Install Base System - Desktop Environment")
	_, selectedDE := helpers.PromptSelectInfo("Select your desired Desktop Environment", deOpts)

	if strings.ToLower(selectedDE) == "return" {
		_environmentSetup(cfgFile)
	} else {
		switch strings.ToLower(selectedDE) {
		case "budgie":
			pkgName = "budgie"
		case "cinnamon":
			pkgName = "cinnamon"
		case "cutefish":
			pkgName = "cutefish"
		case "deepin":
			pkgName = "deepin"
		case "enlightenment":
			pkgName = "enlightenment"
		case "gnome":
			pkgName = "gnome"
		case "plasma":
			pkgName = "plasma"
		case "lxde":
			pkgName = "lxde-gtk3"
		case "lxqt":
			pkgName = "lxqt"
		case "mate":
			pkgName = "mate"
		case "sugar":
			pkgName = "sugar|sugar-fructose"
		case "ukui":
			pkgName = "ukui"
		case "xfce":
			pkgName = "xfce4"
		}
		helpers.JsonUpdater(cfgFile, "basePackages", pkgName, false)
		_selectShell(cfgFile)
	}
}

func _selectStackWM(cfgFile string) {
	var pkgName string
	stackWMOpts := []helpers.ItemInfo{
		{Item: "Blackbox", Info: "Fast, lightweight window manager for X, without all those annoying library dependencies."},
		{Item: "Fluxbox", Info: "It is very light on resources and easy to handle but yet full of features to make an easy and extremely fast desktop experience."},
		{Item: "Gala", Info: "A beautiful window manager from elementaryos, part of Pantheon. Also as a compositing manager, based on libmutter."},
		{Item: "IceWM", Info: "The goal of IceWM is speed, simplicity, and not getting in the user's way."},
		{Item: "JWM", Info: "JWM is written in C and uses only Xlib at a minimum."},
		{Item: "KWin", Info: "The standard KDE Plasma window manager since KDE 4.0, which is also a compositing manager."},
		{Item: "lwm", Info: "Tries to keep out of your face. There are no icons, no button bars, no icon docks, no root menus. Nothing."},
		{Item: "Marco", Info: "The MATE window manager, fork of Metacity."},
		{Item: "Metacity", Info: "This window manager strives to be quiet, small, stable, get on with its job, and stay out of your attention."},
		{Item: "Mutter", Info: "Window and compositing manager for GNOME, based on Clutter, uses OpenGL."},
		{Item: "MWM", Info: "Based on the Motif toolkit."},
		{Item: "Openbox", Info: "Highly configurable window manager with extensive standards support."},
		{Item: "PekWM", Info: "Once upon a time it was based on the aewm++ window manager, but it has evolved enough that it no longer resembles aewm++ at all."},
		{Item: "twm", Info: "Simple window manager for X, the default/fallback used by Xorg since 1989."},
		{Item: "ukwm", Info: "A lightweight GTK+ window manager, the default window manager for UKUI desktop environment."},
		{Item: "Xfwm", Info: "The Xfce window manager."},
		{Item: "RETURN", Info: "Go back to the previous menu"},
	}

	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Install Base System - Stacking Window Manager")
	_, stackSelectedWM := helpers.PromptSelectInfo("Select your Stacking Window Manager", stackWMOpts)
	if strings.ToLower(stackSelectedWM) == "return" {
		_selectWM(cfgFile)
	} else {
		switch strings.ToLower(stackSelectedWM) {
		case "blackbox":
			pkgName = "blackbox"
		case "fluxbox":
			pkgName = "fluxbox"
		case "gala":
			pkgName = "gala"
		case "icewm":
			pkgName = "icewm"
		case "jwm":
			pkgName = "jwm"
		case "kwin":
			pkgName = "kwin"
		case "lwm":
			pkgName = "lwm"
		case "marco":
			pkgName = "marco"
		case "metacity":
			pkgName = "metacity"
		case "mutter":
			pkgName = "mutter"
		case "mwm":
			pkgName = "openmotif"
		case "openbox":
			pkgName = "openbox"
		case "pekwm":
			pkgName = "pekwm"
		case "twm":
			pkgName = "xorg-twm"
		case "ukwm":
			pkgName = "ukwm"
		case "xfwm":
			pkgName = "xfwm4"
		}
		helpers.JsonUpdater(cfgFile, "basePackages", pkgName, false)
		_selectShell(cfgFile)
		_selectTerminal(cfgFile)
		_selectDisplayManager(cfgFile)
		_selectCompositor(cfgFile)
	}
}

func _selectTilingWM(cfgFile string) {
	var pkgName string
	tileWMOpts := []helpers.ItemInfo{
		{Item: "Bspwm", Info: "Represents windows as the leaves of a full binary tree."},
		{Item: "Herbstluftwm", Info: "Manual tiling window manager for X11 using Xlib and Glib."},
		{Item: "i3", Info: "Tiling window manager, completely written from scratch."},
		{Item: "Notion", Info: "Tiling, tabbed window manager for the X window system that utilizes 'tiles' and 'tabbed' windows."},
		{Item: "Ratpoison", Info: "Simple Window Manager with no fat library dependencies, no fancy graphics, no window decorations, and no rodent dependence."},
		{Item: "Stumpwm", Info: "Tiling, keyboard driven X11 Window Manager written entirely in Common Lisp"},
		{Item: "RETURN", Info: "Go back to the previous menu"},
	}

	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Install Base System - Tiling Window Manager")

	_, tileSelectedWM := helpers.PromptSelectInfo("Select your Tiling Window Manager", tileWMOpts)
	if strings.ToLower(tileSelectedWM) == "return" {
		_selectWM(cfgFile)
	} else {
		switch strings.ToLower(tileSelectedWM) {
		case "bspwm":
			pkgName = "bspwm"
		case "herbstluftwm":
			pkgName = "herbstluftwm"
		case "i3":
			pkgName = "i3-wm|terminator"
		case "notion":
			pkgName = "notion"
		case "ratpoison":
			pkgName = "ratpoison"
		case "stumpwm":
			pkgName = "stumpwm"
		}

		helpers.JsonUpdater(cfgFile, "basePackages", pkgName, false)
		_selectShell(cfgFile)
		_selectTerminal(cfgFile)
		_selectDisplayManager(cfgFile)
		_selectCompositor(cfgFile)
	}
}

func _selectDynamicWM(cfgFile string) {
	var pkgName string
	dynamicWMOpts := []helpers.ItemInfo{
		{Item: "awesome", Info: "Highly configurable, next generation framework window manager for X."},
		{Item: "spectrwm", Info: "Small dynamic tiling window manager for X11, largely inspired by xmonad and dwm."},
		{Item: "Qtile", Info: "Full-featured, hackable tiling window manager written in Python."},
		{Item: "xmonad", Info: "Dynamically tiling X11 window manager that is written and configured in Haskell."},
		{Item: "RETURN", Info: "Go back to the previous menu"},
	}

	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Install Base System - Dynamic Window Manager")

	_, dynamicSelectedWM := helpers.PromptSelectInfo("Select your Dynamic Window Manager", dynamicWMOpts)
	if strings.ToLower(dynamicSelectedWM) == "return" {
		_selectWM(cfgFile)
	} else {
		switch strings.ToLower(dynamicSelectedWM) {
		case "awesome":
			pkgName = "awesome"
		case "spectrwm":
			pkgName = "spectrwm"
		case "qtile":
			pkgName = "qtile"
		case "xmonad":
			pkgName = "xmonad"
		}
		helpers.JsonUpdater(cfgFile, "basePackages", pkgName, false)
		_selectShell(cfgFile)
		_selectTerminal(cfgFile)
		_selectDisplayManager(cfgFile)
		_selectCompositor(cfgFile)
	}
}

func _selectWM(cfgFile string) {
	wmTypeOpts := []helpers.ItemInfo{
		{Item: "Stacking", Info: "Provides the traditional desktop windows \"floating\" behavior."},
		{Item: "Tiling", Info: "\"Tile\" the windows so that none are overlapping. They make use of key-bindings and have less (or no) reliance on the mouse."},
		{Item: "Dynamic", Info: "They can dynamically switch between tiling or floating window layout."},
		{Item: "RETURN", Info: "Go back to the previous menu"},
	}

	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Install Base System - Window Manager")

	_, wmSelectedType := helpers.PromptSelectInfo("Select your type of Window Manager", wmTypeOpts)

	switch strings.ToLower(wmSelectedType) {
	case "stacking":
		_selectStackWM(cfgFile)
	case "tiling":
		_selectTilingWM(cfgFile)
	case "dynamic":
		_selectDynamicWM(cfgFile)
	default:
		_environmentSetup(cfgFile)
	}
}

func _selectShell(cfgFile string) {
	var pkgName string
	shOpts := []helpers.ItemInfo{
		{Item: "Bash", Info: "Command-line history and completion, arrays, arithmetic operations, process substitution, here strings, RegEx matching and brace expansion."},
		{Item: "Dash", Info: "Descendant of the NetBSD version of the Almquist SHell (ash)"},
		{Item: "Korn", Info: "Suitable for prototyping. Has the best features of BSH and the CSH, has it's own programming language. Can improve your workflow."},
		{Item: "Oil", Info: "Oil Shell is a Bash-compatible UNIX command-line shell."},
		{Item: "Zsh", Info: "Designed for interactive use. Has a powerful scripting language. Useful features of Bash, ksh, and tcsh and original features were added."},
		{Item: "Cshell", Info: "It includes a command-line editor, programmable word completion, spelling correction, a history mechanism, job control and a C-like syntax."},
		{Item: "Elvish", Info: "Modern and expressive shell, that can carry internal structured values through pipelines. Features an expressive programming language."},
		{Item: "fish", Info: "Smart command line shell. performs syntax highlighting, completion for commands and their arguments, file existence, and history."},
		{Item: "nushell", Info: "Draws inspiration from functional programming languages, and modern CLI tools."},
		{Item: "xonsh", Info: "Python-powered shell with additional shell primitives that you are used to from Bash and IPython."},
	}

	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Install Base System - Shell")
	_, selectedShell := helpers.PromptSelectInfo("Select your desired shell", shOpts)

	switch strings.ToLower(selectedShell) {
	case "bash":
		pkgName = "bash"
	case "dash":
		pkgName = "dash"
	case "korn":
		pkgName = "ksh"
	case "oil":
		pkgName = "oil"
	case "zsh":
		pkgName = "zsh"
	case "cshell":
		pkgName = "tcsh"
	case "elvish":
		pkgName = "elvish"
	case "fish":
		pkgName = "fish"
	case "nushell":
		pkgName = "nushell"
	case "xonsh":
		pkgName = "xonsh"
	}

	upPkg := fmt.Sprintf("%s|%s", helpers.JsonGetter(cfgFile, "basePackages"), pkgName)
	helpers.JsonUpdater(cfgFile, "basePackages", upPkg, false)
}

func _selectTerminal(cfgFile string) {
	var pkgName string
	termOpts := []helpers.ItemInfo{
		{Item: "Alacritty", Info: "A cross-platform, GPU-accelerated terminal emulator."},
		{Item: "cool-retro-term", Info: "A good looking terminal emulator which mimics the old cathode display."},
		{Item: "Deepin", Info: "Terminal emulation application for Deepin desktop."},
		{Item: "Konsole", Info: "Terminal emulator included in the KDE desktop."},
		{Item: "kitty", Info: "A modern, hackable, featureful, OpenGL based terminal emulator."},
		{Item: "Liri", Info: "Material Design terminal for Liri."},
		{Item: "moserial", Info: "GTK-based serial terminal for the GNOME desktop."},
		{Item: "PuTTY", Info: "Highly configurable ssh/telnet/serial console program."},
		{Item: "QTerminal", Info: "Lightweight Qt-based terminal emulator."},
		{Item: "Station", Info: "Terminal emulation features different view modes such as split vertically and horizontally, a tabbed interface, and copy and paste commands."},
		{Item: "Terminology", Info: "Terminal emulator by the Enlightenment project team with innovative features: file thumbnails and media play like a media player."},
		{Item: "urxvt", Info: "Highly extendable unicode enabled terminal emulator featuring tabbing, url launching, a Quake style drop-down mode and pseudo-transparency."},
		{Item: "xterm", Info: "Simple terminal emulator for X. It provides DEC VT102 and Tektronix 4014 terminals for programs that cannot use the window system directly."},
		{Item: "zutty", Info: "A high-end terminal for low-end systems."},
		{Item: "GnomeConsole", Info: "Formerly known as King’s Cross, a simple user-friendly terminal emulator for the GNOME desktop."},
		{Item: "GNOMETerminal", Info: "A terminal emulator included in the GNOME desktop with support for Unicode."},
		{Item: "LXT", Info: "Desktop independent terminal emulator for LXDE."},
		{Item: "MATE", Info: "A fork of GNOME terminal for the MATE desktop."},
		{Item: "Pantheon", Info: "A super lightweight, beautiful, and simple terminal emulator. It is designed to be setup with sane defaults and little to no configuration."},
		{Item: "Terminator", Info: "Terminal emulator supporting multiple resizable terminal panels."},
		{Item: "Tilda", Info: "Configurable drop down terminal emulator."},
		{Item: "Tilix", Info: "Tiling terminal emulator for GNOME."},
		{Item: "Xfce", Info: "Terminal emulator included in the Xfce desktop with support for a colorized prompt and a tabbed interface."},
	}

	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Install Base System - Terminal")
	_, selectedTerminal := helpers.PromptSelectInfo("Select your desired terminal", termOpts)

	switch strings.ToLower(selectedTerminal) {
	case "alacritty":
		pkgName = "alacritty"
	case "cool-retro-term":
		pkgName = "cool-retro-term"
	case "deepin":
		pkgName = "deepin-terminal"
	case "konsole":
		pkgName = "konsole"
	case "kitty":
		pkgName = "kitty"
	case "liri":
		pkgName = "liri-terminal"
	case "moserial":
		pkgName = "moserial"
	case "putty":
		pkgName = "putty"
	case "qterminal":
		pkgName = "qterminal"
	case "station":
		pkgName = "maui-station"
	case "terminology":
		pkgName = "terminology"
	case "urxvt":
		pkgName = "rxvt-unicode"
	case "xterm":
		pkgName = "xterm"
	case "zutty":
		pkgName = "zutty"
	case "gnomeconsole":
		pkgName = "gnome-console"
	case "gnometerminal":
		pkgName = "gnome-terminal"
	case "lxt":
		pkgName = "lxterminal"
	case "mate":
		pkgName = "mate-terminal"
	case "pantheon":
		pkgName = "pantheon-terminal"
	case "terminator":
		pkgName = "terminator"
	case "tilda":
		pkgName = "tilda"
	case "tilix":
		pkgName = "tilix"
	case "xfce":
		pkgName = "xfce4-terminal"
	}

	upPkg := fmt.Sprintf("%s|%s", helpers.JsonGetter(cfgFile, "basePackages"), pkgName)
	helpers.JsonUpdater(cfgFile, "basePackages", upPkg, false)
}

func _selectDisplayManager(cfgFile string) {
	var pkgName string
	dmOpts := []helpers.ItemInfo{
		{Item: "GDM", Info: "GNOME display manager."},
		{Item: "LightDM", Info: "Cross-desktop display manager, can use various front-ends written in any toolkit."},
		{Item: "LXDM", Info: "LXDE display manager. Can be used independent of the LXDE desktop environment."},
		{Item: "SDDM", Info: "QML-based display manager and successor to KDM; recommended for Plasma and LXQt."},
		{Item: "SLiM", Info: "Lightweight and elegant graphical login solution. Discontinued since 2013, not fully compatible with systemd."},
		{Item: "XDM", Info: "X display manager with support for XDMCP, host chooser."},
	}

	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Install Base System - Display Manager")
	_, selectedDM := helpers.PromptSelectInfo("Select your desired display manager", dmOpts)

	switch strings.ToLower(selectedDM) {
	case "gdm":
		pkgName = "gdm"
	case "lightdm":
		pkgName = "lightdm"
	case "lxdm":
		pkgName = "lxdm-gtk3"
	case "sddm":
		pkgName = "sddm"
	case "slim":
		pkgName = "slim"
	case "xdm":
		pkgName = "xorg-xdm"
	}

	upPkg := fmt.Sprintf("%s|%s", helpers.JsonGetter(cfgFile, "basePackages"), pkgName)
	helpers.JsonUpdater(cfgFile, "basePackages", upPkg, false)
}

func _selectCompositor(cfgFile string) {
	var pkgName string
	compositorOpts := []helpers.ItemInfo{
		{Item: "Picom", Info: "Compositor (a fork of Compton)."},
		{Item: "Xcompmgr", Info: "Composite window-effects manager."},
	}

	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Install Base System - Compositor")
	fmt.Println(helpers.PrintHiYellow("Some window managers (e.g. Compiz, Enlightenment, KWin, Marco, Metacity, Muffin, Mutter, Xfwm) do compositing on their own."))
	fmt.Println(helpers.PrintHiYellow("For other window managers, a standalone composite manager can be used."))
	_, selectedCompositor := helpers.PromptSelectInfo("Select your desired compositor", compositorOpts)

	switch strings.ToLower(selectedCompositor) {
	case "picom":
		pkgName = "picom"
	case "xcompmgr":
		pkgName = "xcompmgr"
	}

	upPkg := fmt.Sprintf("%s|%s", helpers.JsonGetter(cfgFile, "basePackages"), pkgName)
	helpers.JsonUpdater(cfgFile, "basePackages", upPkg, false)
}

func __installPython() []string {
	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Install Base System")
	pythonPkgs := []string{"python"}
	fmt.Println(helpers.PrintYellow("ATTENTION: The CPython implementation is automatically installed."))
	fmt.Println(helpers.PrintYellow("These implementations are usually based on older versions of Python and are not fully compatible with CPython."))
	if helpers.YesNo("Would you like to install other python implementations?") {
		pythonImplementations := []string{"PyPy", "Jython"}

		helpers.ClearConsole()
		helpers.PrintHeader("Base Install", "Install Base System")
		selectedPythonImplementations := helpers.PromptMultiSelect("Select the implementations to install", pythonImplementations)
		for _, v := range selectedPythonImplementations {
			switch strings.ToLower(v) {
			case "pypy":
				pythonPkgs = append(pythonPkgs, "pypy")
				pythonPkgs = append(pythonPkgs, "pypy3")
			case "jython":
				pythonPkgs = append(pythonPkgs, "jython")
			}
		}
	}

	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Install Base System")
	if helpers.YesNo("Would you like to install other python interactive shells?") {
		pythonShells := []string{"Bpython", "IPython"}
		selectedPythonShells := helpers.PromptMultiSelect("Select the interactive shell to install", pythonShells)
		for _, v := range selectedPythonShells {
			switch strings.ToLower(v) {
			case "bpython":
				pythonPkgs = append(pythonPkgs, "bpython")
			case "ipython":
				pythonPkgs = append(pythonPkgs, "ipython")
			}
		}
	}
	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Install Base System")
	if helpers.YesNo("Would you like to install Jupyter related packages?") {
		pythonPkgs = append(pythonPkgs, "jupyterlab")
		pythonPkgs = append(pythonPkgs, "jupyter-notebook")
	}
	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Install Base System")
	if helpers.YesNo("Would you like to install PiP related packages?") {
		pipPkgs := []string{"pip", "pipx"}
		helpers.ClearConsole()
		helpers.PrintHeader("Base Install", "Install Base System")
		selectedPipBins := helpers.PromptMultiSelect("Select the PiP implementations to install", pipPkgs)
		for _, v := range selectedPipBins {
			switch strings.ToLower(v) {
			case "pip":
				pythonPkgs = append(pythonPkgs, "python-pip")
			case "pipx":
				pythonPkgs = append(pythonPkgs, "python-pipx")
			}
		}
	}
	return pythonPkgs
}

func __installJava() []string {
	var javaPkgs []string
	javaVersions := []string{"19", "17", "11", "8"}
	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Install Base System")
	selectedJavaVersions := helpers.PromptMultiSelect("Select the versions of java to install", javaVersions)
	installJFX := helpers.YesNo("Would you like to install JFX too?")
	for _, v := range selectedJavaVersions {
		switch v {
		case "19":
			javaPkgs = append(javaPkgs, "jre-openjdk")
			javaPkgs = append(javaPkgs, "jdk-openjdk")
			javaPkgs = append(javaPkgs, "openjdk-doc")
			javaPkgs = append(javaPkgs, "openjdk-src")
			if installJFX {
				javaPkgs = append(javaPkgs, "java-openjfx")
				javaPkgs = append(javaPkgs, "java-openjfx-doc")
				javaPkgs = append(javaPkgs, "java-openjfx-src")
			}
		case "17":
			javaPkgs = append(javaPkgs, "jre17-openjdk")
			javaPkgs = append(javaPkgs, "jdk17-openjdk")
			javaPkgs = append(javaPkgs, "openjdk17-doc")
			javaPkgs = append(javaPkgs, "openjdk17-src")
			if installJFX {
				javaPkgs = append(javaPkgs, "java17-openjfx")
				javaPkgs = append(javaPkgs, "java17-openjfx-doc")
				javaPkgs = append(javaPkgs, "java17-openjfx-src")
			}
		case "11":
			javaPkgs = append(javaPkgs, "jre11-openjdk")
			javaPkgs = append(javaPkgs, "jdk11-openjdk")
			javaPkgs = append(javaPkgs, "openjdk11-doc")
			javaPkgs = append(javaPkgs, "openjdk11-src")
			if installJFX {
				javaPkgs = append(javaPkgs, "java11-openjfx")
				javaPkgs = append(javaPkgs, "java11-openjfx-doc")
				javaPkgs = append(javaPkgs, "java11-openjfx-src")
			}
		case "8":
			javaPkgs = append(javaPkgs, "jre8-openjdk")
			javaPkgs = append(javaPkgs, "jdk8-openjdk")
			javaPkgs = append(javaPkgs, "openjdk8-doc")
			javaPkgs = append(javaPkgs, "openjdk8-src")
			if installJFX {
				helpers.ClearConsole()
				helpers.PrintHeader("Base Install", "Install Base System")
				fmt.Println(helpers.PrintYellow("JFX 8 is not available in the Arch Main Repos."))
				if helpers.YesNo("Would you like to compile it?") {
					fmt.Println("yes")
				}
			}
		}
	}
	return javaPkgs
}

func __installRust() ([]string, bool) {
	var rustPkgs []string
	var upsteamVer bool
	rustInstallOpts := []string{"Native", "Rustup"}
	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Install Base System")
	_, selectedRustInstall := helpers.PromptSelect("You want the Native or the Rustup installation?", rustInstallOpts)
	if strings.ToLower(selectedRustInstall) == "native" {
		rustPkgs = append(rustPkgs, "rust")
	} else {
		helpers.ClearConsole()
		helpers.PrintHeader("Base Install", "Install Base System")
		fmt.Println(helpers.PrintYellow("ATTENTION: 'rustup self update' will not work when installed this way, the package needs to be updated by pacman."))
		fmt.Println(helpers.PrintYellow("But, this package has the advantage that the various Rust executables live in '/usr/bin', instead of '$HOME/.cargo/bin', removing the need to add another directory to your PATH."))
		if helpers.YesNo("Would you like to use the rustup package available in the Arch main repos?") {
			rustPkgs = append(rustPkgs, "rustup")
		} else {
			helpers.ClearConsole()
			helpers.PrintHeader("Base Install", "Install Base System")
			fmt.Println(helpers.PrintYellow("ATTENTION: This will installs the UPSTREAM version from the rustup's official web page."))
			fmt.Println(helpers.PrintYellow("Pacman will not be able to update the package, so you will need to run 'rustup self update' manually."))
			resp := helpers.CurlResponse("https://sh.rustup.rs")
			helpers.WriteToFile(fmt.Sprintf("%s/rust.sh", helpers.GetCurrDirPath()), resp, 0644)
			if helpers.YesNo("Would you like to check the downloaded script?") {
				helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "less", fmt.Sprintf("%s/rust.sh", helpers.GetCurrDirPath()))
			}
			if helpers.YesNo("Would you like to run the script?") {
				helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "sh", fmt.Sprintf("%s/rust.sh", helpers.GetCurrDirPath()))
				fmt.Println(helpers.PrintYellow("You will need to manually install a toolchain afterwards."))
			} else {
				if !helpers.YesNo("Would you like to abort the rust installation?") {
					__installRust()
				}
			}
			upsteamVer = true
		}
	}
	return rustPkgs, !upsteamVer
}

func __installLisp() []string {
	var lispPkgs []string
	commonLispImplementations := []string{
		"CLISP",
		"CMUCL",
		"ECL",
		"SBCL",
	}
	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Install Base System")
	selectedLispImplemntations := helpers.PromptMultiSelect("Select the LISP implementations to install", commonLispImplementations)
	for _, v := range selectedLispImplemntations {
		switch v {
		case "clisp":
			lispPkgs = append(lispPkgs, "clisp")
		case "cmucl":
			lispPkgs = append(lispPkgs, "cmucl")
		case "ecl":
			lispPkgs = append(lispPkgs, "ecl")
		case "sbcl":
			lispPkgs = append(lispPkgs, "sbcl")
		}
	}
	if helpers.YesNo("Would you like to install 'quicklisp' library manager?") {
		lispPkgs = append(lispPkgs, "quicklisp")
	}
	return lispPkgs
}

func _selectProgrammingLanguages(cfgFile string) {
	var langsPkgs []string
	var pkgsToInstall []string
	programmingLangsOpts := []string{
		"Assembly",
		"Ada",
		"BASIC",
		"C",
		"Cpp",
		"Cs",
		"Objective-C",
		"Gambas",
		"Vala",
		"Zig",
		"Crystal",
		"D",
		"Dart",
		"Fortran",
		"Go",
		"java",
		"Groovy",
		"JavaScript",
		"Julia",
		"Kotlin",
		"Lua",
		"Nim",
		"Octave",
		"Pascal",
		"Perl",
		"PHP",
		"Python",
		"R",
		"Ruby",
		"Rust",
		"Scala",
		"Swift",
		"TCL",
		"Erlang",
		"Elixir",
		"Clojure",
		"CommonLisp",
		"Scheme",
		"Racket",
		"ML",
		"OCaml",
	}

	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Install Base System - Programming Languages | LLVM")
	useLLVM := helpers.YesNo("Would you like to install/use LLVM for compiled languages?")

	if useLLVM {
		pkgsToInstall = append(pkgsToInstall, "lld")
		pkgsToInstall = append(pkgsToInstall, "lldb")
		pkgsToInstall = append(pkgsToInstall, "llvm")
		pkgsToInstall = append(pkgsToInstall, "polly")
	}

	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Install Base System - Programming Languages")
	selectedLangs := helpers.PromptMultiSelect("Select your programming languages", programmingLangsOpts)

	for _, v := range selectedLangs {
		switch strings.ToLower(v) {
		case "assembly":
			asmFlavours := []string{
				"fasm",
				"nasm",
				"yasm",
			}

			helpers.ClearConsole()
			helpers.PrintHeader("Base Install", "Install Base System - Programming Languages | ASM")
			selectedAsmFlavours := helpers.PromptMultiSelect("Select your ASM compilers", asmFlavours)
			langsPkgs = append(langsPkgs, selectedAsmFlavours...)
		case "ada":
			langsPkgs = append(langsPkgs, "gcc-ada")
		case "basic":
			langsPkgs = append(langsPkgs, "freebasic")
		case "gambas":
			langsPkgs = append(langsPkgs, "gambas3")
		case "c":
			langsPkgs = append(langsPkgs, "gcc")
			if useLLVM {
				if !helpers.CheckExistsInStringSlice("clang", langsPkgs) {
					langsPkgs = append(langsPkgs, "clang")
				}
				langsPkgs = append(langsPkgs, "libclc")
			}
		case "cpp":
			langsPkgs = append(langsPkgs, "gcc")
			if useLLVM {
				if !helpers.CheckExistsInStringSlice("clang", langsPkgs) {
					langsPkgs = append(langsPkgs, "clang")
				}
				langsPkgs = append(langsPkgs, "libc++")
			}
		case "cs":
			cSharpPkgs := []string{
				"mono",
				"dotnet-runtime",
				"dotnet-sdk",
				"aspnet-runtime",
			}
			langsPkgs = append(langsPkgs, cSharpPkgs...)
		case "objective-c":
			langsPkgs = append(langsPkgs, "gcc-objc")
			if useLLVM && !helpers.CheckExistsInStringSlice("clang", langsPkgs) {
				langsPkgs = append(langsPkgs, "clang")
			}
		case "vala":
			langsPkgs = append(langsPkgs, "vala")
		case "zig":
			langsPkgs = append(langsPkgs, "zig")
		case "crystal":
			langsPkgs = append(langsPkgs, "crystal")
			langsPkgs = append(langsPkgs, "shards")
		case "d":
			dPkgs := []string{
				"dtools",
				"dlang-dmd",
				"ldc",
			}
			helpers.ClearConsole()
			helpers.PrintHeader("Base Install", "Install Base System - Programming Languages | D")
			if helpers.YesNo("Would you like to use the GCC front-end for D?") {
				dPkgs = append(dPkgs, "gcc-d")
			}
			langsPkgs = append(langsPkgs, dPkgs...)
		case "dart":
			langsPkgs = append(langsPkgs, "dart")
		case "fortran":
			langsPkgs = append(langsPkgs, "gcc-fortran")
		case "go":
			langsPkgs = append(langsPkgs, "go")
			helpers.ClearConsole()
			helpers.PrintHeader("Base Install", "Install Base System - Programming Languages | Go")
			if helpers.YesNo("Would you like to use the GCC front-end for Go?") {
				langsPkgs = append(langsPkgs, "gcc-go")
			}
		case "java":
			javaPkgs := __installJava()
			langsPkgs = append(langsPkgs, javaPkgs...)
		case "groovy":
			langsPkgs = append(langsPkgs, "groovy")
		case "javascript":
			langsPkgs = append(langsPkgs, "rhino")
			langsPkgs = append(langsPkgs, "nodejs")
			helpers.ClearConsole()
			helpers.PrintHeader("Base Install", "Install Base System - Programming Languages | NodeJS")
			if helpers.YesNo("Would you like to install NodeJS LTS releases?") {
				nodeLTSVersions := []string{"18", "16", "14"}
				selectedLTSVers := helpers.PromptMultiSelect("Select the NodeJS LTS versions to install", nodeLTSVersions)
				for _, v := range selectedLTSVers {
					switch v {
					case "18":
						langsPkgs = append(langsPkgs, "nodejs-lts-hydrogen")
					case "16":
						langsPkgs = append(langsPkgs, "nodejs-lts-gallium")
					case "14":
						langsPkgs = append(langsPkgs, "nodejs-lts-fermium")
					}
				}
			}
		case "julia":
			helpers.ClearConsole()
			helpers.PrintHeader("Base Install", "Install Base System - Programming Languages | Julia")
			fmt.Println(helpers.PrintYellow("Julia is not available in the Arch Main Repos."))
			if helpers.YesNo("Would you like to compile it?") {
				helpers.InstallAurPkgWithoutHelper("julia-bin")
			}
		case "kotlin":
			langsPkgs = append(langsPkgs, "kotlin")
		case "lua":
			luaPkgs := []string{
				"lua",
				"luarocks",
				"luajit",
			}
			helpers.ClearConsole()
			helpers.PrintHeader("Base Install", "Install Base System - Programming Languages | Lua")
			if helpers.YesNo("Would you like to install other versions of Lua?") {
				luaVersions := []string{"5.3", "5.2", "5.1"}
				selectedLuaVers := helpers.PromptMultiSelect("Select the versions to install", luaVersions)
				for _, v := range selectedLuaVers {
					switch v {
					case "5.3":
						luaPkgs = append(luaPkgs, "lua53")
					case "5.2":
						luaPkgs = append(luaPkgs, "lua52")
					case "5.1":
						luaPkgs = append(luaPkgs, "lua51")
					}
				}
			}
			langsPkgs = append(langsPkgs, luaPkgs...)
		case "nim":
			langsPkgs = append(langsPkgs, "nim")
		case "octave":
			langsPkgs = append(langsPkgs, "octave")
		case "pascal":
			langsPkgs = append(langsPkgs, "fpc")
		case "perl":
			langsPkgs = append(langsPkgs, "perl")
		case "php":
			phpPackages := []string{
				"php",
				"php-legacy",
				"php-cgi",
			}
			langsPkgs = append(langsPkgs, phpPackages...)
		case "python":
			pythonPkgs := __installPython()
			langsPkgs = append(langsPkgs, pythonPkgs...)
		case "r":
			langsPkgs = append(langsPkgs, "r")
			langsPkgs = append(langsPkgs, "gcc-fortran")
		case "ruby":
			rubyPkgs := []string{
				"ruby",
				"ruby-irb",
				"ruby-rdocs",
				"ruby-docs",
				"rubygems",
			}
			helpers.ClearConsole()
			helpers.PrintHeader("Base Install", "Install Base System - Programming Languages | Ruby")
			fmt.Println(helpers.PrintYellow("ATTENTION: 'Pry' will not be installed, so if you want to use it, install using gem after the system installation."))
			if helpers.YesNo("Would you like to install the Java implementation for Ruby?") {
				rubyPkgs = append(rubyPkgs, "jruby")
			}
			langsPkgs = append(langsPkgs, rubyPkgs...)
		case "rust":
			rustPkgs, appnd := __installRust()
			if appnd {
				langsPkgs = append(langsPkgs, rustPkgs...)
			}
		case "scala":
			scalaPackages := []string{
				"scala",
				"scala-docs",
				"scala-sources",
				"sbt",
				"maven",
				"gradle",
			}
			langsPkgs = append(langsPkgs, scalaPackages...)
		case "swift":
			helpers.ClearConsole()
			helpers.PrintHeader("Base Install", "Install Base  System - Programming Languages | Swift")
			fmt.Println(helpers.PrintYellow("Swift is not available in the Arch Main Repos."))
			if helpers.YesNo("Would you like to compile it?") {
				helpers.InstallAurPkgWithoutHelper("swift-bin")
			}
		case "tcl":
			langsPkgs = append(langsPkgs, "tcl")
		case "erlang":
			langsPkgs = append(langsPkgs, "erlang")
		case "elixir":
			langsPkgs = append(langsPkgs, "elixir")
		case "clojure":
			langsPkgs = append(langsPkgs, "clojure")
			langsPkgs = append(langsPkgs, "leiningen")
		case "commonlisp":
			lispPkgs := __installLisp()
			langsPkgs = append(langsPkgs, lispPkgs...)
		case "scheme":
			var schemePkgs []string
			schemeImplementations := []string{
				"Bigloo",
				"CHICKEN",
				"Gambit",
				"Gauche",
			}
			helpers.ClearConsole()
			helpers.PrintHeader("Base Install", "Install Base  System - Programming Languages | Scheme")
			selectedSchemeImplementations := helpers.PromptMultiSelect("Select the SCHEME implementation to install", schemeImplementations)
			for _, v := range selectedSchemeImplementations {
				switch strings.ToLower(v) {
				case "bigloo":
					schemePkgs = append(schemePkgs, "bigloo")
				case "chicken":
					schemePkgs = append(schemePkgs, "chicken")
				case "gambit":
					schemePkgs = append(schemePkgs, "gambit-c")
				case "gauche":
					schemePkgs = append(schemePkgs, "gauche")
				}
			}
			langsPkgs = append(langsPkgs, schemePkgs...)
		case "racket":
			langsPkgs = append(langsPkgs, "racket")
		case "ml":
			var mlPkgs []string
			mlCompilers := []string{
				"smlnj",
				"mlton",
				"polyml",
			}
			helpers.ClearConsole()
			helpers.PrintHeader("Base Install", "Install Base  System - Programming Languages | ML")
			selectedMLCompilers := helpers.PromptMultiSelect("Select the ML compilers to install", mlCompilers)
			for _, v := range selectedMLCompilers {
				switch strings.ToLower(v) {
				case "smlnj":
					mlPkgs = append(mlPkgs, "bigloo")
				case "mlton":
					mlPkgs = append(mlPkgs, "chicken")
				case "polyml":
					mlPkgs = append(mlPkgs, "gambit-c")
				}
			}
			langsPkgs = append(langsPkgs, mlPkgs...)
		case "ocaml":
			langsPkgs = append(langsPkgs, "ocaml")
		}
	}

	pkgsToInstall = append(pkgsToInstall, langsPkgs...)
	concatenatedPkgs := strings.Join(pkgsToInstall, "|")

	upPkg := fmt.Sprintf("%s|%s", helpers.JsonGetter(cfgFile, "basePackages"), concatenatedPkgs)
	helpers.JsonUpdater(cfgFile, "basePackages", upPkg, false)
}

func _environmentSetup(cfgFile string) {
	envOpts := []helpers.ItemInfo{
		{Item: "Desktop Environment", Info: "Bundles together a variety of components and integrate applications and utilities, has a complete environment."},
		{Item: "Window Manager", Info: "Controls the placement and appearance of windows within a windowing system in a GUI, has a minimal environment."},
	}

	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Install Base System - Environment")

	_, retEnv := helpers.PromptSelectInfo("You want a Desktop Environment or Window Manager setup?", envOpts)
	switch strings.ToLower(retEnv) {
	case "desktop environment":
		_selectDE(cfgFile)
	case "window manager":
		_selectWM(cfgFile)
	default:
		fmt.Println(helpers.PrintError("WRONG ENVIRONMENT SELECTION."))
		helpers.CountDown(5, "Running this section again")
		installBaseSystem(cfgFile)
	}
}

func installBaseSystem(cfgFile string) {
	var basePackages []string
	baseDefaultPackages := []string{
		"xorg",
		"usbutils",
		"binutils",
		"dosfstools",
		"linux-headers",
		"gparted",
		"gptfdisk",
		"btop",
		"os-prober",
		"pavucontrol",
		"pulseaudio",
		"pulseaudio-alsa",
		"alsa-utils",
		"alsa-firmware",
		"pulseaudio-jack",
		"pulseaudio-equalizer",
		"grub-customizer",
		"procps-ng",
		"util-linux",
		"coreutils",
	}

	textPhrases := []string{
		"This step will guide you through the basic configuration and installation of the system.",
		"That consists of:",
		"  - Selecting between Desktop Environment and Window Manager setup.",
		"  - Common programming languages setup.",
		"  - Selecting default SHELL.",
		"  - Selecting default terminal emulator.",
		"  - Selecting default Display Manager.",
		"  - Selecting default Compositor.",
		"  - Selecting user folders (Documents, Downloads, etc...).",
		"  - Other small utilities.\n",
		"The rest of these configs will be the in the USER installation. So don't worry for now.",
	}

	if helpers.JsonGetter(cfgFile, "aurHelper") == "none" {
		fmt.Println(helpers.PrintRed("YOU DIDN'T SELECTED ANY AUR/NIX HELPER, SO YOU CAN'T INSTALL PACKAGES THAT ARE NOT AVAILABLE IN THE ARCH MAIN REPO."))
	}

	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Install Base System")

	helpers.BulkPrint(textPhrases)

	if !helpers.YesNo("Can we proceed?") {
		helpers.ClearConsole()
		installBaseSystem(cfgFile)
	} else {
		//_, _, retAcr := helpers.PromptSelectInfoAcronyms("You want a DE or WM setup?", envOpts, envAcronyms, envInfos)
		_environmentSetup(cfgFile)

		helpers.ClearConsole()
		helpers.PrintHeader("Base Install", "Install Base System - User Directories")
		if helpers.YesNo("Do you want to install the user directories (Documents, Downloads, etc...)?") {
			basePackages = append(basePackages, "xdg-user-dirs")
		}

		helpers.ClearConsole()
		helpers.PrintHeader("Base Install", "Install Base System - Programming Languages")
		if helpers.YesNo("Would you like to install programming languages packages?") {
			_selectProgrammingLanguages(cfgFile)
		}

		userSelectedPackages := strings.Split(helpers.JsonGetter(cfgFile, "basePackages"), "|")

		basePackages = append(basePackages, baseDefaultPackages...)
		basePackages = append(basePackages, userSelectedPackages...)

		helpers.ClearConsole()
		helpers.PrintHeader("Base Install", "Install Base System")
		fmt.Printf("The script will install the following packages: %v\n", helpers.PrintHiBlack(fmt.Sprintf("%v", baseDefaultPackages)))
		if helpers.YesNo("Would you like to remove any of those packages?") {
			helpers.ClearConsole()
			helpers.PrintHeader("Base Install", "Install Base System - Removing Packages")
			fmt.Println(helpers.PrintYellow("ATTENTION: The following packages are essential for the system work properly"))
			fmt.Println(helpers.PrintYellow("But you can remove they if you want to."))
			fmt.Println(helpers.PrintHiBlack(baseDefaultPackages))
			pkgsToRemove := helpers.PromptMultiSelect("Select the packages to remove", basePackages)
			finalPkgs := helpers.DifferenceBetweenSlices(basePackages, pkgsToRemove)
			helpers.PacmanInstallPackages(finalPkgs...)
		} else {
			helpers.PacmanInstallPackages(basePackages...)
		}
	}
}

//endregion

func setLocaleTimezone(cfgFile string) {
	// /etc/locale.gen
	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Locale / TimeZone")
	//localeFile := "/etc/locale.gen"
	localeFile := fmt.Sprintf("%s/locales.txt", helpers.GetCurrDirPath())
	keyboardLayout := helpers.JsonGetter(cfgFile, "keyboardLayout")
	timeZone := helpers.JsonGetter(cfgFile, "timeZone")

	options := helpers.ReadLocaleFile(localeFile, "\n")
	_, selectLocale := helpers.PromptSelect("Select your locale", options)
	if helpers.YesNo(fmt.Sprintf("Selected locale: %s. Is this correct?", selectLocale)) {
		strippedLocale := strings.Split(selectLocale, " ")
		sLoc := fmt.Sprintf("LANG=\"%s\"", strippedLocale[0])
		sTz := fmt.Sprintf("LC_TIME=\"%s\"", strippedLocale[0])
		tzLoc := fmt.Sprintf("/usr/share/zoneinfo/%s", timeZone)

		helpers.ReplaceFileLine(localeFile, fmt.Sprintf("#%s", selectLocale), selectLocale)
		helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "locale-gen")
		helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "localectl", "--no-ask-password", "set-locale", sLoc, sTz)
		helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "timedatectl", "--no-ask-password", "set-timezone", timeZone)
		helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "timedatectl", "--no-ask-password", "set-ntp", "true")
		helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "ln", "-sfv", tzLoc, "/etc/localtime")
		helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "localectl", "--no-ask-password", "set-keymap", keyboardLayout)
	} else {
		setLocaleTimezone(cfgFile)
	}
}

func setUser(cfgFile string) {
	userName := helpers.JsonGetter(cfgFile, "userName")
	whoamiOutput := helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, true, "whoami")
	if strings.ToLower(whoamiOutput) == "root" {
		helpers.RunShellCommand(helpers.COMMANDS_TEST_MODE, false, "useradd", "-m", "-G", "wheel", "-s", "/bin/bash", userName)
		// COPY SCRIPTS TO USER HOME DIR
	}
}

//!/////////////////////////////!//
//! SET STEVEN BLACK HOSTS HERE
//! SET HOSTS FILE
//!/////////////////////////////!//

func setHostsFiles(cfgFile string) {
	var socialHosts, gamblingHosts, pornHosts, fakeNewsHosts bool
	hostsLock := fmt.Sprintf("%s/hosts.txt", helpers.GetCurrDirPath())
	//hostsLock := "/etc/hosts"
	hostsOptions := []string{
		"Gambling",
		"Porn",
		"Social",
		"Fake News",
	}
	hostName := helpers.JsonGetter(cfgFile, "hostName")
	helpers.ClearConsole()
	helpers.PrintHeader("Base Install", "Hosts File")

	fmt.Println(helpers.PrintYellow("This step will configure additional hosts in your /etc/hosts file (these hosts are from StevenBlack's repo)"))
	fmt.Println(helpers.PrintYellow("There are some notes to take:"))
	fmt.Println(helpers.PrintYellow("- The UNIFIED HOSTS is a collection of Malware and Adware IPs and Domains."))
	fmt.Println(helpers.PrintYellow("There are 4 'extensions' to this base host:"))
	fmt.Println(helpers.PrintYellow("- Social (this extension will block most of social media sites.)"))
	fmt.Println(helpers.PrintYellow("- Fake News (this extension will block most of fake news sites.)"))
	fmt.Println(helpers.PrintYellow("- Porn (this extension will block most of pornographic sites.)"))
	fmt.Println(helpers.PrintYellow("- Gambling (this extensions will block most of gambling sites.)"))

	baseHosts := []string{
		"# Static table lookup for hostnames.",
		"# See hosts(5) for details.",
		"# =====================================",
		"# IPv4	Config",
		"127.0.0.1\tlocalhost",
		"::1\t\tlocalhost",
		fmt.Sprintf("127.0.1.1\t%s.localdomain\t%s", hostName, hostName),
		"# =====================================",
		"# IPv6 Config",
		"::1\t\tip6-localhost",
		"::1\t\tip6-loopback",
		"fe80::1%lo0 \tlocalhost",
		"ff00::0\t\tip6-localnet",
		"ff00::0\t\tip6-mcastprefix",
		"ff02::1\t\tip6-allnodes",
		"ff02::2\t\tip6-allrouters",
		"ff02::3\t\tip6-allhosts",
		"0.0.0.0\t\t0.0.0.0",
	}

	if helpers.YesNo("Would you like to configure this custom hosts?") {
		selectedHosts := helpers.PromptMultiSelect("Select your desired Custom Hosts extensions: ", hostsOptions)
		for _, v := range selectedHosts {
			switch strings.ToLower(v) {
			case "gambling":
				gamblingHosts = true
			case "porn":
				pornHosts = true
			case "social":
				socialHosts = true
			case "fake news":
				fakeNewsHosts = true
			}
		}
		helpers.EmptyFile(hostsLock)
		if gamblingHosts && pornHosts && socialHosts && fakeNewsHosts {
			resp := helpers.CurlResponse("https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/fakenews-gambling-porn-social/hosts")
			helpers.EmptyFile(hostsLock)
			helpers.WriteToFile(hostsLock, resp, 0644)
		} else if gamblingHosts && pornHosts && socialHosts {
			resp := helpers.CurlResponse("https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/gambling-porn-social/hosts")
			helpers.EmptyFile(hostsLock)
			helpers.WriteToFile(hostsLock, resp, 0644)
		} else if gamblingHosts && pornHosts && fakeNewsHosts {
			resp := helpers.CurlResponse("https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/fakenews-gambling-porn/hosts")
			helpers.EmptyFile(hostsLock)
			helpers.WriteToFile(hostsLock, resp, 0644)
		} else if gamblingHosts && socialHosts && fakeNewsHosts {
			resp := helpers.CurlResponse("https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/fakenews-gambling-social/hosts")
			helpers.EmptyFile(hostsLock)
			helpers.WriteToFile(hostsLock, resp, 0644)
		} else if pornHosts && socialHosts && fakeNewsHosts {
			resp := helpers.CurlResponse("https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/fakenews-porn-social/hosts")
			helpers.EmptyFile(hostsLock)
			helpers.WriteToFile(hostsLock, resp, 0644)
		} else if gamblingHosts && pornHosts {
			resp := helpers.CurlResponse("https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/gambling-porn/hosts")
			helpers.EmptyFile(hostsLock)
			helpers.WriteToFile(hostsLock, resp, 0644)
		} else if gamblingHosts && socialHosts {
			resp := helpers.CurlResponse("https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/gambling-social/hosts")
			helpers.EmptyFile(hostsLock)
			helpers.WriteToFile(hostsLock, resp, 0644)
		} else if gamblingHosts && fakeNewsHosts {
			resp := helpers.CurlResponse("https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/fakenews-gambling/hosts")
			helpers.EmptyFile(hostsLock)
			helpers.WriteToFile(hostsLock, resp, 0644)
		} else if pornHosts && socialHosts {
			resp := helpers.CurlResponse("https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/porn-social/hosts")
			helpers.EmptyFile(hostsLock)
			helpers.WriteToFile(hostsLock, resp, 0644)
		} else if pornHosts && fakeNewsHosts {
			resp := helpers.CurlResponse("https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/fakenews-porn/hosts")
			helpers.EmptyFile(hostsLock)
			helpers.WriteToFile(hostsLock, resp, 0644)
		} else if socialHosts && fakeNewsHosts {
			resp := helpers.CurlResponse("https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/fakenews-social/hosts")
			helpers.EmptyFile(hostsLock)
			helpers.WriteToFile(hostsLock, resp, 0644)
		} else if gamblingHosts {
			resp := helpers.CurlResponse("https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/gambling/hosts")
			helpers.EmptyFile(hostsLock)
			helpers.WriteToFile(hostsLock, resp, 0644)
		} else if pornHosts {
			resp := helpers.CurlResponse("https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/porn/hosts")
			helpers.EmptyFile(hostsLock)
			helpers.WriteToFile(hostsLock, resp, 0644)
		} else if socialHosts {
			resp := helpers.CurlResponse("https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/social/hosts")
			helpers.EmptyFile(hostsLock)
			helpers.WriteToFile(hostsLock, resp, 0644)
		} else if fakeNewsHosts {
			resp := helpers.CurlResponse("https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/fakenews/hosts")
			helpers.EmptyFile(hostsLock)
			helpers.WriteToFile(hostsLock, resp, 0644)
		} else {
			resp := helpers.CurlResponse("https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts")
			helpers.EmptyFile(hostsLock)
			helpers.WriteToFile(hostsLock, resp, 0644)
		}

		helpers.ReplaceFileLine(hostsLock, "127.0.0.1 localhost.localdomain", fmt.Sprintf("127.0.0.1 localhost.localdomain\n127.0.1.1 %s.localdomain %s", hostName, hostName))
	} else {
		helpers.EmptyFile(hostsLock)
		for _, v := range baseHosts {
			helpers.AppendToFile(hostsLock, v)
		}
	}
}

func BaseInstall() {
	var CONFIG_DIR string = fmt.Sprintf("%s/config", helpers.GetCurrDirPath())
	var CONFIG_FILE string = fmt.Sprintf("%s/config.json", CONFIG_DIR)
	//setLocaleTimezone(CONFIG_FILE)
	//installHelper(CONFIG_FILE)
	//installBaseSystem(CONFIG_FILE)
	setHostsFiles(CONFIG_FILE)
}
