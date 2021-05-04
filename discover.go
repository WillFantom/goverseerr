package goverseerr

import "fmt"

func (o *Overseerr) DiscoverMovies(pageNumber int) (*SearchResults, error) {
	if pageNumber < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResults
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetQueryParams(map[string]string{
		"page":     fmt.Sprintf("%d", pageNumber),
		"language": o.locale,
	}).SetResult(&results).Get("/discover/movies")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &results, nil
}

func (o *Overseerr) DiscoverTV(pageNumber int) (*SearchResults, error) {
	if pageNumber < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResults
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetQueryParams(map[string]string{
		"page":     fmt.Sprintf("%d", pageNumber),
		"language": o.locale,
	}).SetResult(&results).Get("/discover/tv")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &results, nil
}

func (o *Overseerr) DiscoverMoviesByGenre(pageNumber, genreID int) (*SearchResults, error) {
	if pageNumber < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResults
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetQueryParams(map[string]string{
		"page":     fmt.Sprintf("%d", pageNumber),
		"language": o.locale,
	}).SetPathParam("genreID", fmt.Sprintf("%d", genreID)).SetResult(&results).
		Get("/discover/movies/genre/{genreID}")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &results, nil
}

func (o *Overseerr) DiscoverMoviesByStudio(pageNumber, studioID int) (*SearchResults, error) {
	if pageNumber < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResults
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetQueryParams(map[string]string{
		"page":     fmt.Sprintf("%d", pageNumber),
		"language": o.locale,
	}).SetPathParam("studioID", fmt.Sprintf("%d", studioID)).SetResult(&results).
		Get("/discover/movies/studio/{studioID}")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &results, nil
}

func (o *Overseerr) DiscoverUpcomingMovies(pageNumber int) (*SearchResults, error) {
	if pageNumber < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResults
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetQueryParams(map[string]string{
		"page":     fmt.Sprintf("%d", pageNumber),
		"language": o.locale,
	}).SetResult(&results).Get("/discover/movies/upcoming")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &results, nil
}

func (o *Overseerr) DiscoverTVByGenre(pageNumber, genreID int) (*SearchResults, error) {
	if pageNumber < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResults
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetQueryParams(map[string]string{
		"page":     fmt.Sprintf("%d", pageNumber),
		"language": o.locale,
	}).SetPathParam("genreID", fmt.Sprintf("%d", genreID)).SetResult(&results).
		Get("/discover/tv/genre/{genreID}")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &results, nil
}

func (o *Overseerr) DiscoverTVByNetwork(pageNumber, networkID int) (*SearchResults, error) {
	if pageNumber < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResults
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetQueryParams(map[string]string{
		"page":     fmt.Sprintf("%d", pageNumber),
		"language": o.locale,
	}).SetPathParam("networkID", fmt.Sprintf("%d", networkID)).SetResult(&results).
		Get("/discover/tv/network/{networkID}")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &results, nil
}

func (o *Overseerr) DiscoverUpcomingTV(pageNumber int) (*SearchResults, error) {
	if pageNumber < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResults
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetQueryParams(map[string]string{
		"page":     fmt.Sprintf("%d", pageNumber),
		"language": o.locale,
	}).SetResult(&results).Get("/discover/tv/upcoming")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &results, nil
}

func (o *Overseerr) DiscoverTrending(pageNumber int) (*SearchResults, error) {
	if pageNumber < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResults
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetQueryParams(map[string]string{
		"page":     fmt.Sprintf("%d", pageNumber),
		"language": o.locale,
	}).SetResult(&results).Get("/discover/trending")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &results, nil
}
