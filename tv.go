package goverseerr

import "fmt"

type TVDetails struct {
	ID                  int                 `json:"id"`
	Name                string              `json:"name"`
	TagLine             string              `json:"tagline"`
	IMDBID              string              `json:"imdbId"`
	InProduction        bool                `json:"inProduction"`
	Genres              []Genre             `json:"genres"`
	Overview            string              `json:"overview"`
	Creator             []Creator           `json:"createdBy"`
	FirstAired          string              `json:"firstAirDate"`
	EpisodeRuntime      []int               `json:"episodeRunTime"`
	Homepage            string              `json:"homepage"`
	Languages           []string            `json:"languages"`
	LastAired           string              `json:"lastAirDate"`
	LastAiredEpisode    Episode             `json:"lastEpisodeToAir"`
	NextEpisodeToAir    Episode             `json:"nextEpisodeToAir"`
	Networks            []Network           `json:"networks"`
	EpisodeCount        int                 `json:"numberOfEpisodes"`
	SeasonCount         int                 `json:"numberOfSeasons"`
	OriginCountry       []string            `json:"originCountry"`
	OriginalLanguage    string              `json:"originalLanguage"`
	OriginalName        string              `json:"originalName"`
	Popularity          float64             `json:"popularity"`
	ProductionCompanies []ProductionCompany `json:"productionCompanies"`
	SpokenLanguages     []SpokenLanguage    `json:"spokenLanguages"`
	Seasons             []Season            `json:"seasons"`
	Status              string              `json:"status"`
	Type                string              `json:"type"`
	VoteAverage         float64             `json:"voteAverage"`
	VoteCount           int                 `json:"voteCount"`
	Credits             Credits             `json:"credits"`
	ExternalIDs         ExternalIDs         `json:"externalIds"`
	KeyWords            []Keyword           `json:"keywords"`
	MediaInfo           MediaInfo           `json:"mediaInfo"`
	RelatedVideos       []RelatedVideo      `json:"relatedVideos"`
}

type Season struct {
	ID           int       `json:"id"`
	AirDate      string    `json:"airDate"`
	EpisodeCount int       `json:"episodeCount"`
	Name         string    `json:"name"`
	Overview     string    `json:"overview"`
	PosterPath   string    `json:"posterPath"`
	Number       int       `json:"seasonNumber"`
	Episodes     []Episode `json:"episodes"`
}

type Episode struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	AirDate        string  `json:"airDate"`
	Number         int     `json:"episodeNumber"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"productionCode"`
	SeasonNumber   int     `json:"seasonNumber"`
	ShowID         int     `json:"showId"`
	StillPath      string  `json:"stillPath"`
	VoteAverage    float64 `json:"voteAverage"`
	VoteCount      float64 `json:"voteCount"`
}

type Network struct {
	ID            int    `json:"id"`
	LogoPath      string `json:"logoPath"`
	OriginCountry string `json:"originCountry"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Headquaters   string `json:"headquesters"`
	Homepage      string `json:"homepage"`
}

func (o *Overseerr) GetTVDetails(tvID int) (*TVDetails, error) {
	var details TVDetails
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("tvID", fmt.Sprintf("%d", tvID)).
		SetQueryParam("language", o.locale).SetResult(&details).Get("/tv/{tvID}")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &details, nil
}

func (o *Overseerr) GetTVSeason(tvID, seasonID int) (*Season, error) {
	var details Season
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParams(map[string]string{
		"tvID":     fmt.Sprintf("%d", tvID),
		"seasonID": fmt.Sprintf("%d", seasonID),
	}).
		SetQueryParam("language", o.locale).SetResult(&details).Get("/tv/{tvID}/season/{seasonID}")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &details, nil
}

func (o *Overseerr) GetTVRecommendations(tvID, page int) (*SearchResults, error) {
	if page < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResults
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("tvID", fmt.Sprintf("%d", tvID)).
		SetQueryParams(map[string]string{
			"language": o.locale,
			"page":     fmt.Sprintf("%d", page),
		}).SetResult(&results).Get("/tv/{tvID}/recommendations")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &results, nil
}

func (o *Overseerr) GetTVSimilar(tvID, page int) (*SearchResults, error) {
	if page < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResults
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("tvID", fmt.Sprintf("%d", tvID)).
		SetQueryParams(map[string]string{
			"language": o.locale,
			"page":     fmt.Sprintf("%d", page),
		}).SetResult(&results).Get("/tv/{tvID}/similar")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &results, nil
}

func (o *Overseerr) GetTVRatings(tvID int) (*Rating, error) {
	var rating Rating
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("tvID", fmt.Sprintf("%d", tvID)).
		SetResult(&rating).Get("/tv/{tvID}/ratings")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &rating, nil
}
