package cmd

import (
	"fmt"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/willfantom/goverseerr"
	"github.com/willfantom/goverseerr/cmd/overclirr/ui"
)

var genreSearchCmd = &cobra.Command{
	Use:   "genre-search [search_term]",
	Short: "Check if a genre exists",
	Args:  cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		instance = getOverseerrInstance(overseerrProfileName)
	},
	Run: func(cmd *cobra.Command, args []string) {
		tvGenreList, movieGenreList := searchGenre(args[0])
		var tableValues = [][]string{
			{"Type", "ID", "Name"},
		}
		for _, genre := range tvGenreList {
			if fuzzy.MatchNormalizedFold(args[0], genre.Name) {
				tableValues = append(tableValues, []string{"TV", fmt.Sprintf("%d", genre.ID), genre.Name})
			}
		}
		for _, genre := range movieGenreList {
			if fuzzy.MatchNormalizedFold(args[0], genre.Name) {
				tableValues = append(tableValues, []string{"Movie", fmt.Sprintf("%d", genre.ID), genre.Name})
			}
		}
		ui.PrettyTable(tableValues)
	},
}

func searchGenre(searchTerm string) ([]*goverseerr.Genre, []*goverseerr.Genre) {
	logrus.WithField("searchTerm", searchTerm).Traceln("searching for genre id")
	tvGenreList, err := instance.TVGenres()
	if err != nil {
		ui.PrettyFatal("Could not get Genre list from Overseerr instance")
		logrus.WithField("extended", err.Error()).Fatal("could not get tv genre list from instance")
	}
	movieGenreList, err := instance.MovieGenres()
	if err != nil {
		ui.PrettyFatal("Could not get Genre list from Overseerr instance")
		logrus.WithField("extended", err.Error()).Fatal("could not get movie genre list from instance")
	}
	logrus.WithFields(logrus.Fields{
		"tv-genres":    len(tvGenreList),
		"movie-genres": len(movieGenreList),
	}).Debug("collected genre lists")
	return tvGenreList, movieGenreList
}

func init() {
	RootCmd.AddCommand(genreSearchCmd)
}
