package overseerr

import (
	"github.com/pkg/browser"
	"github.com/sirupsen/logrus"
	"github.com/willfantom/goverseerr"
	"github.com/willfantom/goverseerr/cmd/overclirr/ui"
)

func MovieResultAction(movie *goverseerr.MovieResult) {
	options := []string{"view poster", "request", "plex", "radarr"}
	_, opt, err := ui.RunSelector("Pick an action", options)
	if err != nil {
		logrus.WithField("extended", err.Error()).Fatalln("movie result action selector failed")
	}
	switch opt {
	case options[0]:
		if movie.PosterPath != "" {
			ui.ShowMediaPoster(movie.PosterPath)
		} else {
			ui.PrettyFatal("No poster can be found for this title")
			logrus.WithField("movieID", movie.ID).Fatalln("movie result poster can not be opened")
		}

	case options[1]:
		// make request

	case options[2]:
		if movie.MediaInfo.PlexURL != "" && movie.MediaInfo.Status == goverseerr.MediaStatsAvailable {
			browser.OpenURL(movie.MediaInfo.PlexURL)
		} else {
			ui.PrettyFatal("Movie can not be opened in Plex")
			logrus.WithField("movieID", movie.ID).Fatalln("movie result cant be opened in plex")
		}

	case options[3]:
		if movie.MediaInfo.ServiceURL != "" {
			browser.OpenURL(movie.MediaInfo.ServiceURL)
		} else {
			ui.PrettyFatal("Movie can not be opened in Radarr")
			logrus.WithField("movieID", movie.ID).Fatalln("movie result cant be opened in radarr")
		}
	}
}
