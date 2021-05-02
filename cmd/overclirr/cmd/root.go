package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/willfantom/goverseerr"
	"github.com/willfantom/goverseerr/cmd/overclirr/newui"
	"github.com/willfantom/goverseerr/cmd/overclirr/overseerr"
	"github.com/willfantom/goverseerr/cmd/overclirr/ui"
)

const (
	defaultProfile string = "default"
)

var (
	logLevel             string
	overseerrProfileName string
	noTitle              bool
	instance             *goverseerr.Overseerr
)

var RootCmd = &cobra.Command{
	Use:     "overclirr",
	Aliases: []string{"ocrr", "overseerr", "overseerr-cli"},
	Short:   "Manage Overseerr(s) from the command line",
	Long:    `A simple command line tool for managing media server(s) with Overseerr!`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		setupLogger()
		if !noTitle {
			newui.TitleBox("OverCLIrr", "An Overseerr Management Tool")
		}
		logrus.WithFields(logrus.Fields{
			"command": cmd.Name(),
			"args":    args,
		}).Debugln("running command")
	},
	Run: func(cmd *cobra.Command, args []string) {
		profiles, err := overseerr.GetOverseerrProfiles()
		if err != nil {
			logrus.WithField("extended", err.Error()).Errorln("failed to get login profiles")
			ui.PrettyFatal("Could not get user profiles from configuration")
		}
		for name := range profiles {
			profile, err := overseerr.GetOverseerrFromProfile(name)
			if err != nil {
				logrus.WithField("extended", err.Error()).Errorln("failed to get profile from profile config")
				ui.PrettyOops("Could not build profile: " + name)
				continue
			}
			if !profile.HealthCheck() {
				logrus.WithField("profile", name).Errorln("could not establish a connection using profile")
				ui.PrettyOops("Could not establish a connection with profile: " + name)
				continue
			}
			newui.Success("Established connection with profile: " + name)
		}
	},
}

func setupLogger() {
	if level, err := logrus.ParseLevel(logLevel); err != nil {
		ui.PrettyOops("invalid log level given: " + logLevel)
		ui.PrettyInfo("Using 'panic'")
		logrus.SetLevel(logrus.PanicLevel)
	} else {
		logrus.SetLevel(level)
	}
	if logrus.GetLevel() != logrus.PanicLevel {
		viper.Set("showLoadingSpinner", false)
	}
}

// getOverseerrInstance returns an Overseerr struct for a profile or exits with errors
func getOverseerrInstance(profile string) *goverseerr.Overseerr {
	ui.StartLoadingSpinner()
	o, err := overseerr.GetOverseerrFromProfile(profile)
	if err != nil {
		ui.StopLoadingSpinner()
		ui.PrettyFatal("Could not create OversCLIrr instance with profile: " + profile)
		logrus.WithField("extended", err.Error()).Fatalln("Could not create OversCLIrr instance with profile: " + profile)
	}
	overseerr.AddWrappersToOverseerr(o)
	if !o.HealthCheck() {
		ui.StopLoadingSpinner()
		ui.PrettyFatal("Could not connect and authorize OversCLIrr with profile: " + profile)
		logrus.Fatalln("Could not connect and authorize OversCLIrr with profile: " + profile)
	}
	ui.StopLoadingSpinner()
	return o
}

func init() {
	RootCmd.PersistentFlags().BoolVar(&noTitle, "no-title", false, "stop printing the big fu*king title")
	RootCmd.PersistentFlags().StringVar(&logLevel, "log", "panic", "set the log level (fatal, error, info, debug, trace)")
	RootCmd.PersistentFlags().StringVar(&overseerrProfileName, "profile", defaultProfile, "use a specific overseerr login profile name")
	viper.BindPFlag("log", RootCmd.PersistentFlags().Lookup("log"))
}
