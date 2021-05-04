package goverseerr

import (
	"fmt"
	"net/url"
	"sort"
)

const PosterPathBase string = "https://image.tmdb.org/t/p/w600_and_h900_bestv2"

type MovieResult struct {
	Title       string `json:"title" mapstructure:"title"`
	ReleaseDate string `json:"releaseDate" mapstructure:"releaseDate"`
	Video       bool   `json:"video" mapstructure:"video"`
}

type TVResult struct {
	OriginCounty   []string `json:"originCountry" mapstructure:"originCountry"`
	FirstAiredDate string   `json:"firstAirDate" mapstructure:"firstAirDate"`
}

type PersonResult struct {
	ProfilePath string        `json:"profilePath" mapstructure:"profilePath"`
	KnownFor    []interface{} `json:"knownFor" mapstructure:"knownFor"`
}

type SearchResults struct {
	Page         int                   `json:"page"`
	TotalPages   int                   `json:"totalPages"`
	TotalResults int                   `json:"totalResults"`
	Genre        Genre                 `json:"genre,omitempty"`
	Studio       ProductionCompany     `json:"studio,omitempty"`
	Network      Network               `json:"network,omitempty"`
	Results      []GenericSearchResult `json:"results"`
}

// GenericSearchResult represents a single search result item
type GenericSearchResult struct {
	ID               int       `json:"id" mapstructure:"id"`
	MediaType        MediaType `json:"mediaType" mapstructure:"mediaType"`
	MediaInfo        MediaInfo `json:"mediaInfo" mapstructure:"mediaInfo"`
	Name             string    `json:"name" mapstructure:"name"`
	Popularity       float64   `json:"popularity" mapstructure:"popularity"`
	Adult            bool      `json:"adult" mapstructure:"adult"`
	PosterPath       string    `json:"posterPath" mapstructure:"posterPath"`
	BackdropPath     string    `json:"backdropPath" mapstructure:"backdropPath"`
	VoteCount        int       `json:"voteCount" mapstructure:"voteCount"`
	VoteAverage      float64   `json:"voteAverage" mapstructure:"voteAverage"`
	GenreIDs         []int     `json:"genreIds" mapstructure:"genreIds"`
	Overview         string    `json:"overview" mapstructure:"overview"`
	OriginalLanguage string    `json:"originalLanguage" mapstructure:"originalLanguage"`
	OriginalTitle    string    `json:"originalTitle" mapstructure:"originalTitle"`
	MovieResult
	TVResult
	PersonResult
}

// Result Helpers

func (results *SearchResults) SortByPopularity() {
	sort.Slice(results.Results, func(i, j int) bool {
		return results.Results[i].Popularity > results.Results[j].Popularity
	})
}

func (results *SearchResults) SortByID() {
	sort.Slice(results.Results, func(i, j int) bool {
		return results.Results[i].ID < results.Results[j].ID
	})
}

func (results *SearchResults) SortByAdult() {
	sort.Slice(results.Results, func(i, j int) bool {
		return results.Results[i].Adult
	})
}

// Endpoints

func (o *Overseerr) Search(query string, pageNumber int) (*SearchResults, error) {
	if pageNumber < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResults
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

// Genre Slider

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
