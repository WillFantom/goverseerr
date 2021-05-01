package ui

import (
	"strings"

	"github.com/pterm/pterm"
)

func PrettyTitle() {
	s, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("OverCLIrr")).Srender()
	pterm.Println(s)
}

func PrettyPrefix(prefix, text string) {
	oops := pterm.NewStyle(pterm.FgBlack, pterm.BgLightBlue, pterm.Underscore)
	oops.Printf(" " + prefix + " ")
	pterm.Printf("  %s\n", text)
}

func PrettyTable(values [][]string) {
	pterm.DefaultTable.WithHasHeader().WithData(values).Render()
}

func PrettyHeader(title string) {
	pterm.DefaultSection.Println(title)
}

func PrettySuccess(message string) {
	pterm.Success.Println(message)
}

func PrettyInfo(messages ...string) {
	pterm.Info.Println(strings.Join(messages, "\n"))
}

func PrettyOops(message string) {
	oops := pterm.NewStyle(pterm.FgBlack, pterm.BgLightRed, pterm.Bold)
	oops.Printf(" OOPS ")
	pterm.Printf("  %s\n", message)
}

func PrettyFatal(message string) {
	fatal := pterm.NewStyle(pterm.FgBlack, pterm.BgRed, pterm.Bold)
	fatal.Printf(" FATAL ")
	pterm.Printf(" %s\n", message)
}
