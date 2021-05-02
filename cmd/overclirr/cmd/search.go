package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/willfantom/goverseerr/cmd/overclirr/ui"
)

func getSearchTerm(args []string) string {
	searchTerm := ""
	if len(args) == 0 {
		term, err := ui.GetInput("What would you like to search for?", nonEmptyValidator)
		if err != nil {
			logrus.WithField("extended", err.Error()).Fatalln("search term input failed")
		}
		searchTerm = term
	} else {
		if len(args[0]) > 0 {
			searchTerm = args[0]
		} else {
			logrus.Fatalln("search term command arg must not be an empty string")
		}
	}
	return searchTerm
}
