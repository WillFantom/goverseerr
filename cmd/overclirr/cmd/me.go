package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/willfantom/goverseerr/cmd/overclirr/ui"
)

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Get information about the currently logged in user",
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		instance = getOverseerrInstance(overseerrProfileName)
	},
	Run: func(cmd *cobra.Command, args []string) {
		user, err := instance.GetLoggedInUser()
		if err != nil {
			ui.PrettyFatal("Could not fetch user information")
			logrus.New().WithField("extended", err.Error()).Fatalln("Could not fetch user information")
		}
		ui.PrettyHeader("User Information")
		ui.PrettyInfo(
			fmt.Sprintf("Email: %s", user.Email),
			fmt.Sprintf("ID: %d", user.ID),
			fmt.Sprintf("Avatar Path: %s", user.Avatar),
			fmt.Sprintf("Created: %s", user.Created.Local().String()),
			fmt.Sprintf("Permissions: %d", user.Permissions),
			fmt.Sprintf("Requests Made: %d", user.RequestCount),
		)
	},
}

func init() {
	RootCmd.AddCommand(meCmd)
}
