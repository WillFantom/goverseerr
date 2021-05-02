package newui

import (
	"strings"

	box "github.com/Delta456/box-cli-maker/v2"
	"github.com/gookit/color"
	"github.com/pterm/pterm"
	"github.com/sirupsen/logrus"
)

const (
	maxLineLength int = 50
)

var (
	titleBox = box.New(box.Config{
		Type:         "Double",
		Color:        "HiMagenta",
		ContentAlign: "Left",
		TitlePos:     "Inside",
		Px:           3,
		Py:           0,
	})

	contentsBox = box.New(box.Config{
		Type:         "Round",
		Color:        "White",
		ContentAlign: "Left",
		TitlePos:     "Top",
		Px:           2,
		Py:           1,
	})
)

func stringToParagraph(in string) string {
	return pterm.DefaultParagraph.WithMaxWidth(maxLineLength).Sprint(in)
}

func TitleBox(title, sub string) {
	titleBox.Println(title, sub)
}

func ContentBox(title, contents string) {
	contentsBox.Println(title, contents)
}

func Success(message string) {
	color.New(color.FgBlack, color.BgGreen).Printf(" SUCCESS ")
	color.New(color.FgGreen).Printf(" %s\n", message)
}

func Error(message string) {
	color.New(color.FgBlack, color.BgRed).Printf("  ERROR  ")
	color.New(color.FgRed).Printf(" %s\n", message)
}

func Fatal(message string, err error) {
	color.New(color.FgBlack, color.BgRed, color.Bold).Printf(" FATAL  ")
	color.New(color.FgRed).Printf(" %s\n", message)
	logrus.WithField("extended", err.Error()).Fatalln(strings.ToLower(message))
}
