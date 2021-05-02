package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/willfantom/goverseerr"
	"github.com/willfantom/goverseerr/cmd/overclirr/ui"
)

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "View and mange users on an Overseerr instance",
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		instance = getOverseerrInstance(overseerrProfileName)
	},
	Run: func(cmd *cobra.Command, args []string) {
		viewUsers(getAllUsers())
	},
}

func getAllUsers() []*goverseerr.User {
	ui.StartLoadingSpinner()
	var allUsers []*goverseerr.User
	pg := 0
	for {
		users, page, err := instance.GetAllUsers(20, pg)
		if err != nil {
			ui.StopLoadingSpinner()
			ui.PrettyFatal("Could not fetch all users")
			logrus.WithField("extended", err.Error()).Fatalln("could not fetch all users")
		}
		allUsers = append(allUsers, users...)
		pg++
		if pg >= page.Pages {
			ui.StopLoadingSpinner()
			return allUsers
		}
	}
}

func viewUsers(users []*goverseerr.User) {
	ui.StartLoadingSpinner()
	var tableValues = [][]string{
		{"ID", "Email", "Permissions", "Requests"},
	}
	for _, user := range users {
		tableValues = append(tableValues, []string{fmt.Sprintf("%v", user.ID), user.Email, fmt.Sprintf("%v", user.Permissions), fmt.Sprintf("%v", user.RequestCount)})
	}
	ui.StopLoadingSpinner()
	ui.PrettyTable(tableValues)
}

var viewUsersCmd = &cobra.Command{
	Use:   "view",
	Short: "View users from an Overseerr instance",
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		instance = getOverseerrInstance(overseerrProfileName)
	},
	Run: func(cmd *cobra.Command, args []string) {
		viewUsers(getAllUsers())
	}}

func init() {
	usersCmd.AddCommand(viewUsersCmd)
	adminCmd.AddCommand(usersCmd)
}
