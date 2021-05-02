package newui

import (
	"fmt"

	"github.com/willfantom/goverseerr"
)

const (
	maxGenres int = 3
)

func PrintMovieDetails(details *goverseerr.MovieDetails) {
	content := ""
	content += fmt.Sprintf("%s\n\n", stringToParagraph(details.Overview))
	content += fmt.Sprintf("Release Date: %s\n\n", details.ReleaseDate)

	content += "Genres\n"
	for i, g := range details.Genres {
		if i >= maxGenres {
			break
		}
		content += fmt.Sprintf(" [%s] ", g.Name)
	}
	content += "\n\n"

	ContentBox(details.Title, content)
}

func PrintMovieResult(result *goverseerr.MovieResult) {
	content := ""
	content += fmt.Sprintf("%s\n\n", stringToParagraph(result.Overview))
	content += fmt.Sprintf("Release Date: %s\n\n", result.ReleaseDate)

	ContentBox(result.Title, content)
}
