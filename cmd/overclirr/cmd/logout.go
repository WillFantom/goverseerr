package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/willfantom/goverseerr/cmd/overclirr/overseerr"
	"github.com/willfantom/goverseerr/cmd/overclirr/ui"
)

var deleteProfileCmd = &cobra.Command{
	Use:     "logout",
	Aliases: []string{"del-profile"},
	Short:   "Delete a login profile from the configuration",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrettyHeader("Remove an Overseerr config")
		profiles, err := overseerr.GetOverseerrProfiles()
		if err != nil {
			ui.PrettyFatal("Can not load in existing profiles!")
			logrus.WithField("extended", err.Error()).Fatalln("failed to get current overseerr profiles from the configuration")
		}
		if len(profiles) == 0 {
			ui.PrettyFatal("No profiles to delete")
			return
		}
		var keys []string
		for key := range profiles {
			keys = append(keys, key)
		}
		_, profile, err := ui.RunSelector("Which profile do you want to delete", keys)
		if err != nil {
			logrus.Fatalln("profile selector failed")
		}
		if err := overseerr.RemoveOverseerProfile(profile); err != nil {
			ui.PrettyFatal("Could not remvoe the profile from the configuration file")
			logrus.WithField("extended", err.Error()).Fatalln("failed to get remove overseerr profile from the configuration")
		}
		ui.PrettySuccess("Configuration Deleted!")
	},
}

func init() {
	RootCmd.AddCommand(deleteProfileCmd)
}
