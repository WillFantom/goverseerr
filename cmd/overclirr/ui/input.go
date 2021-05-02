package ui

import (
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
)

func RunSelector(title string, options []string) (int, string, error) {
	prompt := promptui.Select{
		Label: title,
		Items: options,
	}
	logrus.WithField("title", title).Traceln("running selector menu")
	resIdx, resString, err := prompt.Run()
	if err != nil {
		return 0, "", err
	}
	logrus.WithField("selection", resString).Traceln("ending selector menu")

	return resIdx, resString, err
}

func GetInput(title string, validator func(string) error) (string, error) {
	prompt := promptui.Prompt{
		Label:    title,
		Validate: validator,
	}
	result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}

func GetMaskedInput(title string, validator func(string) error) (string, error) {
	prompt := promptui.Prompt{
		Label:    title,
		Validate: validator,
		Mask:     '*',
	}
	result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}

func DestructiveConfirmation() {
	prompt := promptui.Prompt{
		Label:     "This is a destructive action, do you want to continue? (yes/no)",
		IsConfirm: true,
	}
	result, err := prompt.Run()
	if err != nil {
		logrus.WithField("extended", err.Error()).Fatalln("destructive confirmation failed")
	}
	if result == "n" {
		logrus.Fatalln("refused to accept destructive confirmation")
	}
}
