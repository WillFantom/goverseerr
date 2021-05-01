package goverseerr

import (
	"fmt"
	"net/url"
	"sort"
	"sync"

	"github.com/mitchellh/mapstructure"
)

const PosterPathBase string = "https://image.tmdb.org/t/p/w600_and_h900_bestv2"

type MovieResult struct {
	ID               int       `json:"id" mapstructure:"id"`
	MediaType        string    `json:"mediaType" mapstructure:"mediaType"`
	Popularity       int       `json:"popularity" mapstructure:"popularity"`
	PosterPath       string    `json:"posterPath" mapstructure:"posterPath"`
	BackdropPath     string    `json:"backdropPath" mapstructure:"backdropPath"`
	VoteCount        int       `json:"voteCount" mapstructure:"voteCount"`
	VoteAverage      int       `json:"voteAverage" mapstructure:"voteAverage"`
	GenreIDs         []int     `json:"genreIds" mapstructure:"genreIds"`
	Overview         string    `json:"overview" mapstructure:"overview"`
	OriginalLanguage string    `json:"originalLanguage" mapstructure:"originalLanguage"`
	Title            string    `json:"title" mapstructure:"title"`
	OriginalTitle    string    `json:"originalTitle" mapstructure:"originalTitle"`
	ReleaseDate      string    `json:"releaseDate" mapstructure:"releaseDate"`
	Adult            bool      `json:"adult" mapstructure:"adult"`
	Video            bool      `json:"video" mapstructure:"video"`
	MediaInfo        MediaInfo `json:"mediaInfo" mapstructure:"mediaInfo"`

	OriginalIndex int
}

type TVResult struct {
	ID               int       `json:"id" mapstructure:"id"`
	MediaType        string    `json:"mediaType" mapstructure:"mediaType"`
	Popularity       int       `json:"popularity" mapstructure:"popularity"`
	PosterPath       string    `json:"posterPath" mapstructure:"posterPath"`
	BackdropPath     string    `json:"backdropPath" mapstructure:"backdropPath"`
	VoteCount        int       `json:"voteCount" mapstructure:"voteCount"`
	VoteAverage      int       `json:"voteAverage" mapstructure:"voteAverage"`
	GenreIDs         []int     `json:"genreIds" mapstructure:"genreIds"`
	Overview         string    `json:"overview" mapstructure:"overview"`
	OriginalLanguage string    `json:"originalLanguage" mapstructure:"originalLanguage"`
	Name             string    `json:"name" mapstructure:"name"`
	OriginalName     string    `json:"originalName" mapstructure:"originalName"`
	OriginCounty     []string  `json:"originCountry" mapstructure:"originCountry"`
	FirstAiredDate   string    `json:"firstAirDate" mapstructure:"firstAirDate"`
	MediaInfo        MediaInfo `json:"mediaInfo" mapstructure:"mediaInfo"`
	Adult            bool      `json:"adult" mapstructure:"adult"`

	OriginalIndex int
}

type KnownFor struct {
	Movie MovieResult
	TV    TVResult
}

type PersonResult struct {
	ID          int        `json:"id" mapstructure:"id"`
	Name        string     `json:"name" mapstructure:"name"`
	Popularity  int        `json:"popularity" mapstructure:"popularity"`
	ProfilePath string     `json:"profilePath" mapstructure:"profilePath"`
	Adult       bool       `json:"adult" mapstructure:"adult"`
	MediaType   string     `json:"mediaType" mapstructure:"mediaType"`
	KnownFor    []KnownFor `json:"knownFor" mapstructure:"knownFor"`

	OriginalIndex int
}

type SearchResult struct {
	Page         int               `json:"page"`
	TotalPages   int               `json:"totalPages"`
	TotalResults int               `json:"totalResults"`
	Genre        Genre             `json:"genre,omitempty"`
	Studio       ProductionCompany `json:"studio,omitempty"`
	Network      Network           `json:"network,omitempty"`
	Results      []interface{}     `json:"results"`
}

type TypedSearchResult struct {
	SearchResult
	Movies []MovieResult
	TV     []TVResult
	People []PersonResult
}

func (r SearchResult) castResults() *TypedSearchResult {
	var results TypedSearchResult
	var wg sync.WaitGroup
	for idx, result := range r.Results {
		wg.Add(3)
		go func(idx int, result interface{}) {
			defer wg.Done()
			var tv TVResult
			tvErr := mapstructure.Decode(result, &tv)
			if tvErr == nil && tv.MediaType == MediaTypeTV {
				tv.OriginalIndex = idx
				results.TV = append(results.TV, tv)
			}
		}(idx, result)
		go func(idx int, result interface{}) {
			defer wg.Done()
			var movie MovieResult
			movieErr := mapstructure.Decode(result, &movie)
			if movieErr == nil && movie.MediaType == MediaTypeMovie {
				movie.OriginalIndex = idx
				results.Movies = append(results.Movies, movie)
			}
		}(idx, result)
		go func(idx int, result interface{}) {
			defer wg.Done()
			var person PersonResult
			personErr := mapstructure.Decode(result, &person)
			if personErr == nil && person.MediaType == MediaTypePerson {
				person.OriginalIndex = idx
				results.People = append(results.People, person)
			}
		}(idx, result)
	}
	wg.Wait()
	wg.Add(3)
	go func() {
		defer wg.Done()
		sort.Slice(results.Movies, func(p, q int) bool {
			return results.Movies[p].OriginalIndex < results.Movies[q].OriginalIndex
		})
	}()
	go func() {
		defer wg.Done()
		sort.Slice(results.TV, func(p, q int) bool {
			return results.TV[p].OriginalIndex < results.TV[q].OriginalIndex
		})
	}()
	go func() {
		defer wg.Done()
		sort.Slice(results.People, func(p, q int) bool {
			return results.People[p].OriginalIndex < results.People[q].OriginalIndex
		})
	}()
	wg.Wait()
	results.Page = r.Page
	results.TotalPages = r.TotalPages
	results.TotalResults = r.TotalResults
	results.Genre = r.Genre
	results.Studio = r.Studio
	results.Network = r.Network
	return &results
}

func (o *Overseerr) Search(query string, pageNumber int) (*SearchResult, error) {
	if pageNumber < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResult
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetQueryParams(map[string]string{
		"query":    url.QueryEscape(query),
		"page":     fmt.Sprintf("%d", pageNumber),
		"language": o.locale,
	}).SetResult(&results).Get("/search")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &results, nil
}

func (o *Overseerr) SearchTyped(query string, pageNumber int) (*TypedSearchResult, error) {
	results, err := o.Search(query, pageNumber)
	if err != nil {
		return nil, err
	}
	return results.castResults(), nil
}

func (o *Overseerr) DiscoverMovies(pageNumber int) (*TypedSearchResult, error) {
	if pageNumber < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResult
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
	return results.castResults(), nil
}

func (o *Overseerr) DiscoverTV(pageNumber int) (*TypedSearchResult, error) {
	if pageNumber < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResult
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
	return results.castResults(), nil
}

func (o *Overseerr) DiscoverMoviesByGenre(pageNumber, genreID int) (*TypedSearchResult, error) {
	if pageNumber < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResult
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
	return results.castResults(), nil
}

func (o *Overseerr) DiscoverMoviesByStudio(pageNumber, studioID int) (*TypedSearchResult, error) {
	if pageNumber < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResult
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
	return results.castResults(), nil
}

func (o *Overseerr) DiscoverUpcomingMovies(pageNumber int) (*TypedSearchResult, error) {
	if pageNumber < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResult
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
	return results.castResults(), nil
}

func (o *Overseerr) DiscoverTVByGenre(pageNumber, genreID int) (*TypedSearchResult, error) {
	if pageNumber < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResult
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
	return results.castResults(), nil
}

func (o *Overseerr) DiscoverTVByNetwork(pageNumber, networkID int) (*TypedSearchResult, error) {
	if pageNumber < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResult
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetQueryParams(map[string]string{
		"page":     fmt.Sprintf("%d", pageNumber),
		"language": o.locale,
	}).SetPathParam("networkID", fmt.Sprintf("%d", networkID)).SetResult(&results).
		Get("/discover/tv/studio/{networkID}")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return results.castResults(), nil
}

func (o *Overseerr) DiscoverUpcomingTV(pageNumber int) (*TypedSearchResult, error) {
	if pageNumber < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResult
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
	return results.castResults(), nil
}

func (o *Overseerr) DiscoverTrending(pageNumber int) (*SearchResult, error) {
	if pageNumber < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResult
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

func (o *Overseerr) DiscoverTrendingTyped(pageNumber int) (*TypedSearchResult, error) {
	results, err := o.DiscoverTrending(pageNumber)
	if err != nil {
		return nil, err
	}
	return results.castResults(), nil
}

func (o *Overseerr) MovieGenres() ([]*Genre, error) {
	var genres []*Genre
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetQueryParams(map[string]string{
		"language": o.locale,
	}).SetResult(&genres).Get("/discover/genreslider/movie")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return genres, nil
}

func (o *Overseerr) TVGenres() ([]*Genre, error) {
	var genres []*Genre
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetQueryParams(map[string]string{
		"language": o.locale,
	}).SetResult(&genres).Get("/discover/genreslider/tv")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return genres, nil
}
