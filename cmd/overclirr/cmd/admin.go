package cmd

import (
	"github.com/spf13/cobra"
)

var adminCmd = &cobra.Command{
	Use:     "admin",
	Aliases: []string{"adm"},
	Short:   "Manage an Overseerr instance",
	Args:    cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		instance = getOverseerrInstance(overseerrProfileName)
	},
	Run: func(cmd *cobra.Command, args []string) {
		//Show Overview of Instance stats
	},
}

func init() {
	RootCmd.AddCommand(adminCmd)
}
