package helpers

import (
	"fmt"

	"github.com/fatih/color"
)

var PrintYellow = color.New(color.FgYellow).SprintFunc()
var PrintBlack = color.New(color.FgBlack).SprintFunc()
var PrintCyan = color.New(color.FgCyan).SprintFunc()
var PrintBlue = color.New(color.FgBlue).SprintFunc()
var PrintRed = color.New(color.FgRed).SprintFunc()

func PrintAsciiArt() {

	coloredArchLinux1 := PrintBlue("    _             _       _     _                  ")
	coloredArchLinux2 := PrintBlue("   / \\   _ __ ___| |__   | |   (_)_ __  _   ___  __")
	coloredArchLinux3 := PrintBlue("  / _ \\ | '__/ __| '_ \\  | |   | | '_ \\| | | \\ \\/ /")
	coloredArchLinux4 := PrintBlue(" / ___ \\| | | (__| | | | | |___| | | | | |_| |>  < ")
	coloredArchLinux5 := PrintBlue("/_/   \\_\\_|  \\___|_| |_| |_____|_|_| |_|\\__,_/_/\\_\\")
	ghostWrittenInGo1 := "__        __    _ _   _               _          ____       "
	ghostWrittenInGo2 := "\\ \\      / / __(_) |_| |_ ___ _ __   (_)_ __    / ___| ___  "
	ghostWrittenInGo3 := " \\ \\ /\\ / / '__| | __| __/ _ \\ '_ \\  | | '_ \\  | |  _ / _ \\ "
	ghostWrittenInGo4 := "  \\ V  V /| |  | | |_| ||  __/ | | | | | | | | | |_| | (_) |"
	ghostWrittenInGo5 := "   \\_/\\_/ |_|  |_|\\__|\\__\\___|_| |_| |_|_| |_|  \\____|\\___/ "
	coloredWrittenInGo1 := fmt.Sprintf("__        __    _ _   _               _       %s", PrintCyan("   ____       "))
	coloredWrittenInGo2 := fmt.Sprintf("\\ \\      / / __(_) |_| |_ ___ _ __   (_)_ __  %s", PrintCyan("  / ___| ___  "))
	coloredWrittenInGo3 := fmt.Sprintf(" \\ \\ /\\ / / '__| | __| __/ _ \\ '_ \\  | | '_ \\ %s", PrintCyan(" | |  _ / _ \\ "))
	coloredWrittenInGo4 := fmt.Sprintf("  \\ V  V /| |  | | |_| ||  __/ | | | | | | | |%s", PrintCyan(" | |_| | (_) |"))
	coloredWrittenInGo5 := fmt.Sprintf("   \\_/\\_/ |_|  |_|\\__|\\__\\___|_| |_| |_|_| |_|%s", PrintCyan("  \\____|\\___/ "))

	ghostArchlinuxInstaller1 := "    _             _       _     _                    ___           _        _ _           "
	ghostArchlinuxInstaller2 := "   / \\   _ __ ___| |__   | |   (_)_ __  _   ___  __ |_ _|_ __  ___| |_ __ _| | | ___ _ __ "
	ghostArchlinuxInstaller3 := "  / _ \\ | '__/ __| '_ \\  | |   | | '_ \\| | | \\ \\/ /  | || '_ \\/ __| __/ _` | | |/ _ \\ '__|"
	ghostArchlinuxInstaller4 := " / ___ \\| | | (__| | | | | |___| | | | | |_| |>  <   | || | | \\__ \\ || (_| | | |  __/ |   "
	ghostArchlinuxInstaller5 := "/_/   \\_\\_|  \\___|_| |_| |_____|_|_| |_|\\__,_/_/\\_\\ |___|_| |_|___/\\__\\__,_|_|_|\\___|_|   "
	coloredArchLinuxInstaller1 := fmt.Sprintf("%s  ___           _        _ _           ", coloredArchLinux1)
	coloredArchLinuxInstaller2 := fmt.Sprintf("%s |_ _|_ __  ___| |_ __ _| | | ___ _ __ ", coloredArchLinux2)
	coloredArchLinuxInstaller3 := fmt.Sprintf("%s  | || '_ \\/ __| __/ _` | | |/ _ \\ '__|", coloredArchLinux3)
	coloredArchLinuxInstaller4 := fmt.Sprintf("%s  | || | | \\__ \\ || (_| | | |  __/ |   ", coloredArchLinux4)
	coloredArchLinuxInstaller5 := fmt.Sprintf("%s |___|_| |_|___/\\__\\__,_|_|_|\\___|_|   ", coloredArchLinux5)
	fmt.Println(CenterSprint(coloredArchLinuxInstaller1, ghostArchlinuxInstaller1))
	fmt.Println(CenterSprint(coloredArchLinuxInstaller2, ghostArchlinuxInstaller2))
	fmt.Println(CenterSprint(coloredArchLinuxInstaller3, ghostArchlinuxInstaller3))
	fmt.Println(CenterSprint(coloredArchLinuxInstaller4, ghostArchlinuxInstaller4))
	fmt.Println(CenterSprint(coloredArchLinuxInstaller5, ghostArchlinuxInstaller5))

	fmt.Println(CenterSprint(coloredWrittenInGo1, ghostWrittenInGo1))
	fmt.Println(CenterSprint(coloredWrittenInGo2, ghostWrittenInGo2))
	fmt.Println(CenterSprint(coloredWrittenInGo3, ghostWrittenInGo3))
	fmt.Println(CenterSprint(coloredWrittenInGo4, ghostWrittenInGo4))
	fmt.Println(CenterSprint(coloredWrittenInGo5, ghostWrittenInGo5))
}
