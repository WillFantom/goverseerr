package overseerr

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/willfantom/goverseerr"
	"github.com/willfantom/goverseerr/cmd/overclirr/ui"
)

func SelectRequest(o *goverseerr.Overseerr, requestSet []*goverseerr.MediaRequest) *goverseerr.MediaRequest {
	ui.StartLoadingSpinner()
	options := make([]string, len(requestSet))
	var wg sync.WaitGroup
	for idx, request := range requestSet {
		wg.Add(1)
		go func(idx int, request *goverseerr.MediaRequest) {
			defer wg.Done()
			options[idx] = ToOptionsString(o, request)
		}(idx, request)
	}
	wg.Wait()
	ui.StopLoadingSpinner()
	idx, _, err := ui.RunSelector("Select a request", options)
	if err != nil {
		logrus.WithField("extended", err.Error()).Fatalln("request selector failed")
	}
	return requestSet[idx]
}

func ToOptionsString(o *goverseerr.Overseerr, request *goverseerr.MediaRequest) string {
	// ID - title example (tv/movie) - abc@123.com [pending]
	template := "%d - %s (%s %s) - %s [%s %s]"
	switch request.Media.MediaType {
	case goverseerr.MediaTypeMovie:
		details, err := request.GetMovieDetails(o)
		if err != nil {
			ui.PrettyFatal("Failed to get movie details")
			logrus.WithField("extended", err.Error()).Fatalln("failed to get movie details")
		}
		return fmt.Sprintf(template, request.ID, details.Title, request.Media.MediaType.ToEmoji(),
			string(request.Media.MediaType), request.Creator.Email, request.Status.ToEmoji(), request.Status.ToString())
	case goverseerr.MediaTypeTV:
		details, err := request.GetTVDetails(o)
		if err != nil {
			ui.PrettyFatal("Failed to get tv details")
			logrus.WithField("extended", err.Error()).Fatalln("failed to get tv details")
		}
		return fmt.Sprintf(template, request.ID, details.Name, request.Media.MediaType.ToEmoji(),
			string(request.Media.MediaType), request.Creator.Email, request.Status.ToEmoji(), request.Status.ToString())
	}
	return "unusable request"
}

func PrintRequest(o *goverseerr.Overseerr, request *goverseerr.MediaRequest) {
	switch request.Media.MediaType {
	case goverseerr.MediaTypeMovie:
		details, err := request.GetMovieDetails(o)
		if err != nil {
			ui.PrettyFatal("Failed to get movie details")
			logrus.WithField("extended", err.Error()).Fatalln("failed to get movie details")
		}
		ui.PrettyMovieRequest(request, details)
	case goverseerr.MediaTypeTV:
		details, err := request.GetTVDetails(o)
		if err != nil {
			ui.PrettyFatal("Failed to get tv details")
			logrus.WithField("extended", err.Error()).Fatalln("failed to get tv details")
		}
		ui.PrettyTVRequest(request, details)
	default:
		ui.PrettyRequest(request)
	}
}

func RetryRequest(o *goverseerr.Overseerr, request *goverseerr.MediaRequest, exit bool) {
	if request.Media.Status == goverseerr.MediaStatsAvailable {
		return
	}
	_, err := o.RetryRequest(request.ID)
	if err != nil {
		ui.PrettyFatal(fmt.Sprintf("Failed to retry this request: %d", request.ID))
		if exit {
			ui.PrettyInfo("Do you have the right permissions?")
			logrus.WithField("extended", err.Error()).Fatalln("request retry failed")
		}
	}
}

func RequestAction(request *goverseerr.MediaRequest, o *goverseerr.Overseerr) {
	if request.Media.IsMovie() {
		details, err := request.GetMovieDetails(o)
		if err != nil {
			ui.PrettyFatal("Failed to get the movie details from Overseerr")
			logrus.WithField("extended", err.Error()).Fatalln("failed to get movie details")
		}
		RequestActionMovie(o, request, details)
	} else if request.Media.IsTV() {
		details, err := request.GetTVDetails(o)
		if err != nil {
			ui.PrettyFatal("Failed to get the show details from Overseerr")
			logrus.WithField("extended", err.Error()).Fatalln("failed to get tv show details")
		}
		RequestActionTV(o, request, details)
	}
	logrus.WithField("request", request.ID).Fatalln("request not a tc show or movie")
}

func RequestActionMovie(o *goverseerr.Overseerr, request *goverseerr.MediaRequest, details *goverseerr.MovieDetails) {
	options := []string{"approve", "decline", "retry"}
	_, opt, err := ui.RunSelector("Pick an action", options)
	if err != nil {
		logrus.WithField("extended", err.Error()).Fatalln("movie result action selector failed")
	}
	switch opt {
	case options[0]:
		if request.Status == goverseerr.RequestStatusPending || request.Status == goverseerr.RequestStatusDeclined {
			_, err := o.ApproveRequest(request.ID)
			if err != nil {
				ui.PrettyFatal("Failed to approve this request")
				ui.PrettyInfo("Do you have the right permissions?")
				logrus.WithField("extended", err.Error()).Fatalln("request approval failed")
			}
		} else {
			ui.PrettyFatal("A request must be pending or declined to be approved")
			logrus.Fatalln("request must be pending or declined to be approved")
		}
	case options[1]:
		if request.Status == goverseerr.RequestStatusPending || request.Status == goverseerr.RequestStatusApproved {
			_, err := o.DeclineRequest(request.ID)
			if err != nil {
				ui.PrettyFatal("Failed to decline this request")
				ui.PrettyInfo("Do you have the right permissions?")
				logrus.WithField("extended", err.Error()).Fatalln("request decline failed")
			}
		} else {
			ui.PrettyFatal("A request must be pending or approved to be declined")
			logrus.Fatalln("request must be pending or approved to be declined")
		}
	case options[2]:
		_, err := o.RetryRequest(request.ID)
		if err != nil {
			ui.PrettyFatal("Failed to retry this request")
			ui.PrettyInfo("Do you have the right permissions?")
			logrus.WithField("extended", err.Error()).Fatalln("request retry failed")
		}
	}

}

func RequestActionTV(o *goverseerr.Overseerr, request *goverseerr.MediaRequest, details *goverseerr.TVDetails) {
	options := []string{"approve", "decline", "retry"}
	_, opt, err := ui.RunSelector("Pick an action", options)
	if err != nil {
		logrus.WithField("extended", err.Error()).Fatalln("tv request action selector failed")
	}
	switch opt {
	case options[0]:
		if request.Status == goverseerr.RequestStatusPending || request.Status == goverseerr.RequestStatusDeclined {
			_, err := o.ApproveRequest(request.ID)
			if err != nil {
				ui.PrettyFatal("Failed to approve this request")
				ui.PrettyInfo("Do you have the right permissions?")
				logrus.WithField("extended", err.Error()).Fatalln("request approval failed")
			}
		} else {
			ui.PrettyFatal("A request must be pending or declined to be approved")
			logrus.Fatalln("request must be pending or declined to be approved")
		}
	case options[1]:
		if request.Status == goverseerr.RequestStatusPending || request.Status == goverseerr.RequestStatusApproved {
			_, err := o.DeclineRequest(request.ID)
			if err != nil {
				ui.PrettyFatal("Failed to decline this request")
				ui.PrettyInfo("Do you have the right permissions?")
				logrus.WithField("extended", err.Error()).Fatalln("request decline failed")
			}
		} else {
			ui.PrettyFatal("A request must be pending or approved to be declined")
			logrus.Fatalln("request must be pending or approved to be declined")
		}
	case options[2]:
		_, err := o.RetryRequest(request.ID)
		if err != nil {
			ui.PrettyFatal("Failed to retry this request")
			ui.PrettyInfo("Do you have the right permissions?")
			logrus.WithField("extended", err.Error()).Fatalln("request retry failed")
		}
	}
}
