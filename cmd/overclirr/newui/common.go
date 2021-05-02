package newui

import (
	"os"
	"strings"
	"time"

	box "github.com/Delta456/box-cli-maker/v2"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/pterm/pterm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/willfantom/reticulating-go"
)

const (
	maxLineLength int = 50
	selectorSize  int = 5
)

var (
	loadingSpinner     *spinner.Spinner
	loadingMessageFunc func() string
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

	miniBox = box.New(box.Config{
		Type:         "Round",
		Color:        "White",
		ContentAlign: "Center",
		TitlePos:     "Inside",
		Px:           2,
		Py:           0,
	})
)

func init() {
	loadingSpinner = spinner.New(spinner.CharSets[3], 100*time.Millisecond, spinner.WithWriter(os.Stderr))
	loadingSpinner.HideCursor = true
	loadingSpinner.Color("white")
	loadingMessageFunc = reticulating.GetLoadingMessage
}

func stringToParagraph(in string) string {
	return pterm.DefaultParagraph.WithMaxWidth(maxLineLength).Sprint(in)
}

func TitleBox(title, sub string) {
	StopLoadingSpinner()
	titleBox.Println(title, sub)
}

func ContentBox(title, contents string) {
	StopLoadingSpinner()
	contentsBox.Println(title, contents)
}

func SingleLineHeader(header string) {
	color.New(color.FgMagenta, color.Underline, color.Bold).Printf(" %s \n", header)
}

func SingleLineSubHeader(sub string) {
	color.New(color.FgWhite).Printf("%s \n", sub)
}

func MiniBox(title, content string) {
	StopLoadingSpinner()
	miniBox.Println(title, content)
}

func Success(message string) {
	StopLoadingSpinner()
	color.New(color.FgBlack, color.BgGreen).Printf(" SUCCESS ")
	color.New(color.FgGreen).Printf(" %s\n", message)
}

func Error(message string) {
	StopLoadingSpinner()
	color.New(color.FgBlack, color.BgRed).Printf("  ERROR  ")
	color.New(color.FgRed).Printf(" %s\n", message)
}

func Fatal(message string, err error) {
	StopLoadingSpinner()
	color.New(color.FgBlack, color.BgRed, color.Bold).Printf(" FATAL  ")
	color.New(color.FgRed).Printf(" %s\n", message)
	logrus.WithField("extended", err.Error()).Fatalln(strings.ToLower(message))
}

func HiddenFatal(message string, err error) {
	logrus.WithField("extended", err.Error()).Fatalln(strings.ToLower(message))
}

func StartSpinner() {
	if viper.GetBool("showLoadingSpinner") {
		loadingSpinner.Start()
		go updateLoadingMessage()
	}
}

func updateLoadingMessage() {
	for loadingSpinner.Active() {
		loadingSpinner.Suffix = "  " + loadingMessageFunc()
		time.Sleep(500 * time.Millisecond)
	}
}

func StopLoadingSpinner() {
	loadingSpinner.Stop()
}

func Selector(title string, options []string) (int, string, error) {
	prompt := promptui.Select{
		Label: title,
		Items: options,
		Size:  selectorSize,
	}
	logrus.WithField("title", title).Traceln("running selector menu")
	index, option, err := prompt.Run()
	if err != nil {
		return 0, "", err
	}
	logrus.WithField("selection", option).Traceln("ending selector menu")
	return index, option, err
}

func DestructiveConfirmation() {
	prompt := promptui.Prompt{
		Label:     "What you are about to do could be destructive, continue?",
		IsConfirm: true,
	}
	result, err := prompt.Run()
	if err != nil {
		HiddenFatal("destructive confirmation failed", err)
	}
	if strings.ToLower(result) != "y" {
		Error("Aborted!")
		HiddenFatal("destructive confirmation rejected", err)
	}
}
