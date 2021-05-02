package newui

import (
	"fmt"
	"sync"

	"github.com/willfantom/goverseerr"
)

func RequestSelector(requests []*goverseerr.MediaRequest, tvDetails []*goverseerr.TVDetails, movieDetails []*goverseerr.MovieDetails, pageNumber, pages int) int {
	content := fmt.Sprintf("showing requests page %d of %d", pageNumber, pages)
	SingleLineHeader("Requests")
	SingleLineSubHeader(content)

	options := make([]string, len(requests))
	var wg sync.WaitGroup
	for idx, req := range requests {
		wg.Add(1)
		go func(i int, request *goverseerr.MediaRequest, tvDetail *goverseerr.TVDetails, movieDetail *goverseerr.MovieDetails) {
			options[i] = requestToOptionString(request, tvDetail, movieDetail)
			wg.Done()
		}(idx, req, tvDetails[idx], movieDetails[idx])
	}
	wg.Wait()

	index, _, err := Selector("Select a Media Request to view/manage", options)
	if err != nil {
		HiddenFatal("request selector failed", err)
	}
	return index
}

func requestToOptionString(request *goverseerr.MediaRequest, tvDetail *goverseerr.TVDetails, movieDetail *goverseerr.MovieDetails) string {
	template := "%d - %s (%s %s) - %s [%s %s]"
	switch request.Media.MediaType {
	case goverseerr.MediaTypeMovie:
		return fmt.Sprintf(template, request.ID, movieDetail.Title, request.Media.MediaType.ToEmoji(),
			string(request.Media.MediaType), request.Creator.Email, request.Status.ToEmoji(), request.Status.ToString())
	case goverseerr.MediaTypeTV:
		return fmt.Sprintf(template, request.ID, tvDetail.Name, request.Media.MediaType.ToEmoji(),
			string(request.Media.MediaType), request.Creator.Email, request.Status.ToEmoji(), request.Status.ToString())
	}
	return "unusable request"
}
