package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/willfantom/goverseerr"
	"github.com/willfantom/goverseerr/cmd/overclirr/overseerr"
	"github.com/willfantom/goverseerr/cmd/overclirr/ui"
)

var (
	searchPages int
)

var searchMoviesCmd = &cobra.Command{
	Use:     "search-movies [search term]",
	Aliases: []string{"movies"},
	Short:   "Search for movies",
	Args:    cobra.MaximumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		instance = getOverseerrInstance(overseerrProfileName)
	},
	Run: func(cmd *cobra.Command, args []string) {
		searchTerm := getSearchTerm(args)
		movies := getSearchMovies(10, searchTerm)
		movie := selectMovieResult(movies)
		printMovieResult(movie)
		overseerr.MovieResultAction(&movie)
	},
}

func getSearchMovies(maxPages int, term string) []goverseerr.MovieResult {
	ui.StartLoadingSpinner()
	var allResults []goverseerr.MovieResult
	for pg := 1; pg < searchPages+1; pg++ {
		results, err := instance.SearchTyped(term, pg)
		if err != nil {
			ui.StopLoadingSpinner()
			ui.PrettyFatal("Failed to search for movies")
			logrus.WithField("extended", err.Error()).Fatalln("failed to perform a search")
		}
		allResults = append(allResults, results.Movies...)
		if pg >= results.TotalPages {
			ui.StopLoadingSpinner()
			return allResults
		}
	}
	ui.StopLoadingSpinner()
	return allResults
}

func selectMovieResult(movies []goverseerr.MovieResult) goverseerr.MovieResult {
	options := make([]string, len(movies))
	for idx, movie := range movies {
		options[idx] = fmt.Sprintf("%s (%s) [%s]", movie.Title, movie.ReleaseDate, movie.MediaInfo.Status.ToString())
	}
	idx, _, err := ui.RunSelector("Select a search result", options)
	if err != nil {
		logrus.WithField("extended", err.Error()).Fatalln("movie result selector failed")
	}
	return movies[idx]
}

func printMovieResult(movie goverseerr.MovieResult) {
	ui.PrettyHeader(movie.Title)
	ui.PrettyInfo(
		"Release Date: "+movie.ReleaseDate,
		"Overview: "+movie.Overview,
		"Status: "+movie.MediaInfo.Status.ToString(),
	)
}

func init() {
	searchMoviesCmd.Flags().IntVarP(&searchPages, "pages", "p", 10, "How many pages iof search results should be used")
	RootCmd.AddCommand(searchMoviesCmd)
}
