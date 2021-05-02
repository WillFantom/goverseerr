package ui

import (
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/viper"
	"github.com/willfantom/reticulating-go"
)

var (
	loadingSpinner     *spinner.Spinner
	loadingMessageFunc func() string
)

func init() {
	loadingSpinner = spinner.New(spinner.CharSets[3], 100*time.Millisecond, spinner.WithWriter(os.Stderr))
	loadingSpinner.HideCursor = true
	loadingSpinner.Color("white")
	loadingMessageFunc = reticulating.GetLoadingMessage
}

func StartLoadingSpinner() {
	if viper.GetBool("showTerminalUIComponenets") {
		loadingSpinner.Start()
		go updateLoadingMessage()
	}
}

func StopLoadingSpinner() {
	if viper.GetBool("showTerminalUIComponenets") {
		loadingSpinner.Stop()
	}
}

func updateLoadingMessage() {
	for loadingSpinner.Active() {
		loadingSpinner.Suffix = "  " + loadingMessageFunc()
		time.Sleep(500 * time.Millisecond)
	}
}
