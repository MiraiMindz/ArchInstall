package utils

import "github.com/charmbracelet/lipgloss"

var DefaultBackgroundColor = lipgloss.CompleteAdaptiveColor{
    Light: lipgloss.CompleteColor{TrueColor: "#F2F2F2", ANSI256: "255", ANSI: "15"},
    Dark:  lipgloss.CompleteColor{TrueColor: "#0C0C0C", ANSI256: "232", ANSI: "0"},
}

var DefaultTextColor = lipgloss.CompleteAdaptiveColor{
    Light: lipgloss.CompleteColor{TrueColor: "#767676", ANSI256: "232", ANSI: "0"},
    Dark:  lipgloss.CompleteColor{TrueColor: "#CCCCCC", ANSI256: "255", ANSI: "15"},
}

var DefaultRedColor = lipgloss.CompleteAdaptiveColor{
    Light: lipgloss.CompleteColor{TrueColor: "#C50F1F", ANSI256: "88", ANSI: "1"},
    Dark:  lipgloss.CompleteColor{TrueColor: "#E74856", ANSI256: "196", ANSI: "9"},
}

var DefaultGreenColor = lipgloss.CompleteAdaptiveColor{
    Light: lipgloss.CompleteColor{TrueColor: "#13A10E", ANSI256: "28", ANSI: "2"},
    Dark:  lipgloss.CompleteColor{TrueColor: "#16C60C", ANSI256: "41", ANSI: "10"},
}

var DefaultYellowColor = lipgloss.CompleteAdaptiveColor{
    Light: lipgloss.CompleteColor{TrueColor: "#C19C00", ANSI256: "101", ANSI: "3"},
    Dark:  lipgloss.CompleteColor{TrueColor: "#F9F1A5", ANSI256: "226", ANSI: "11"},
}

var DefaultBlueColor = lipgloss.CompleteAdaptiveColor{
    Light: lipgloss.CompleteColor{TrueColor: "#0037DA", ANSI256: "17", ANSI: "4"},
    Dark:  lipgloss.CompleteColor{TrueColor: "#F9F1A5", ANSI256: "57", ANSI: "12"},
}

var DefaultMagentaColor = lipgloss.CompleteAdaptiveColor{
    Light: lipgloss.CompleteColor{TrueColor: "#881798", ANSI256: "90", ANSI: "5"},
    Dark:  lipgloss.CompleteColor{TrueColor: "#B4009E", ANSI256: "93", ANSI: "13"},
}

var DefaultCyanColor = lipgloss.CompleteAdaptiveColor{
    Light: lipgloss.CompleteColor{TrueColor: "#3A96DD", ANSI256: "29", ANSI: "6"},
    Dark:  lipgloss.CompleteColor{TrueColor: "#61D6D6", ANSI256: "81", ANSI: "14"},
}