package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/willfantom/goverseerr"
	"github.com/willfantom/goverseerr/cmd/overclirr/ui"
)

var genreCmd = &cobra.Command{
	Use:       "genres [tv/movie]",
	Aliases:   []string{"genre"},
	Short:     "Get a list of all tv/movie genres available",
	ValidArgs: []string{"tv", "movie"},
	Args:      cobra.ExactValidArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		instance = getOverseerrInstance(overseerrProfileName)
	},
	Run: func(cmd *cobra.Command, args []string) {
		var genreList []*goverseerr.Genre
		var err error
		switch args[0] {
		case "tv":
			genreList, err = instance.TVGenres()
		case "movie":
			genreList, err = instance.MovieGenres()
		default:
			logrus.WithField("givenArg", args[0]).
				Panicln("Invalid genre argument got past argument validator")
		}
		if err != nil {
			logrus.WithField("extended", err.Error()).Errorln("could not get genre list from instance")
		}
		logrus.WithField("genreCount", len(genreList)).Debug("collected genre list")
		var tableValues = [][]string{
			{"ID", "Name"},
		}
		for _, genre := range genreList {
			tableValues = append(tableValues, []string{fmt.Sprintf("%d", genre.ID), genre.Name})
		}
		ui.PrettyTable(tableValues)
	},
}

func init() {
	RootCmd.AddCommand(genreCmd)
}
