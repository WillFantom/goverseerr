package ui

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/willfantom/goverseerr"
)

func PrettyTVRequest(request *goverseerr.MediaRequest, details *goverseerr.TVDetails) {
	pterm.DefaultSection.Println(details.Name)
	pterm.DefaultParagraph.Println(details.Overview)
	PrettyInfo(
		fmt.Sprintf("%s: %s", "First Aired", details.FirstAired),
		fmt.Sprintf("%s: %s", "Last Aired", details.LastAired),
		fmt.Sprintf("%s: %d", "Seasons", details.SeasonCount),
		fmt.Sprintf("%s: %d", "Episodes", details.EpisodeCount),
	)
	PrettyRequest(request)
}

func PrettyMovieRequest(request *goverseerr.MediaRequest, details *goverseerr.MovieDetails) {
	pterm.DefaultSection.Println(details.Title)
	pterm.DefaultParagraph.Println(details.Overview)
	PrettyInfo(
		fmt.Sprintf("%s: %s", "Released", details.ReleaseDate),
	)
	PrettyRequest(request)
}

func PrettyRequest(request *goverseerr.MediaRequest) {
	pterm.DefaultSection.WithLevel(2).Println("Request")
	PrettyInfo(
		fmt.Sprintf("%s: %s", "Email", request.Creator.Email),
		fmt.Sprintf("%s: %s", "Created At", request.Created.Local().Format("Mon Jan 2 15:04:05 -0700 MST 2006")),
		fmt.Sprintf("%s: %s", "Status", request.Status.ToString()),
		fmt.Sprintf("%s: %s", "Media Status", request.Media.Status.ToString()),
	)
}
