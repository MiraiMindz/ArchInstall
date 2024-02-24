package utils

import (
	"os"
	"golang.org/x/term"
)

func GetTerminalSize() (int, int, error) {
	fd := int(os.Stdout.Fd())
    width, height, err := term.GetSize(fd)

	return width, height, err
}

func DivideTerminalWidth(num int) (int, int) {
    if num%2 == 0 {
        return num / 2, num / 2
    } else {
        return num / 2, (num / 2) + 1
    }
}