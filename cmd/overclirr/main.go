package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/willfantom/goverseerr/cmd/overclirr/cmd"
	"github.com/willfantom/goverseerr/cmd/overclirr/ui"
)

const (
	initialLogLevel string = "panic"
)

func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		if logrus.GetLevel() < logrus.WarnLevel {
			ui.PrettyFatal("Try setting the log level higher (e.g. info) to see what is going on!")
		}
		logrus.WithField("extended", err.Error()).
			Fatalln("an error occurred executing the command")
	}
}

// init adds the config information to the global viper
func init() {
	//set initial log level
	lvl, _ := logrus.ParseLevel(initialLogLevel)
	logrus.SetLevel(lvl)
	// define configuration file info
	viper.SetConfigName("overclirr")
	viper.SetConfigType("json")
	viper.AddConfigPath("$HOME/.config/")
	viper.AddConfigPath(".")

	// set default configuration values
	viper.SetDefault("showTerminalUIComponenets", true)
	viper.SetDefault("logLevel", "panic")
	viper.SetDefault("itemsPerPage", 25)
	viper.SetDefault("overseerrs", nil)

	// create config file
	if err := viper.SafeWriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); !ok {
			ui.PrettyFatal("Configuration file could not be created!")
			logrus.WithField("extended", err.Error()).Fatalln("configuration file does not exist and can not be written to")
		}
	}

	// read in existing config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			ui.PrettyFatal("Existing configuration file could not be read in")
			ui.PrettyInfo("Try removing the file and readding the configuration values", "Or check the wiki on GitHub...")
			logrus.WithField("extended", err.Error()).Errorln("configuration file found but could not be used!")
		} else {
			ui.PrettyFatal("No configuration file could be found")
			ui.PrettyInfo("This should have been created automatically...")
			logrus.WithField("extended", err.Error()).Errorln("could not find a config file")
		}
	}

	logrus.Traceln("configuration init success")
}
