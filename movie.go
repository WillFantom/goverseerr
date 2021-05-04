package goverseerr

import "fmt"

type MovieDetails struct {
	ID                  int                 `json:"id"`
	Title               string              `json:"title"`
	Adult               bool                `json:"adult"`
	IMDBID              string              `json:"imdbId"`
	ReleaseDate         string              `json:"releaseDate"`
	Genres              []Genre             `json:"genres"`
	Overview            string              `json:"overview"`
	BackdropPath        string              `json:"backdropPath"`
	PosterPath          string              `json:"posterPath"`
	Budget              int                 `json:"budget"`
	Homepage            string              `json:"homepage"`
	RelatedVideos       []RelatedVideo      `json:"relatedVideos"`
	OriginalLanguage    string              `json:"originalLanguage"`
	OriginalTitle       string              `json:"originalTitle"`
	Popularity          float64             `json:"popularity"`
	ProductionCompanies []ProductionCompany `json:"productionCompanies"`
	Revenue             int                 `json:"revenue"`
	Runtime             int                 `json:"runtime"`
	SpokenLanguages     []SpokenLanguage    `json:"spokenLanguages"`
	Status              string              `json:"status"`
	Tagline             string              `json:"tagline"`
	Video               bool                `json:"video"`
	VoteAverage         float64             `json:"voteAverage"`
	VoteCount           int                 `json:"voteCount"`
	Credits             Credits             `json:"credits"`
	Collection          Collection          `json:"collection"`
	ExternalIDs         ExternalIDs         `json:"externalIds"`
	MediaInfo           MediaInfo           `json:"mediaInfo"`
}

type ProductionCompany struct {
	ID            int    `json:"id"`
	LogoPath      string `json:"logoPath"`
	OriginCountry string `json:"originCountry"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Headquaters   string `json:"headquesters"`
	Homepage      string `json:"homepage"`
}

func (o *Overseerr) GetMovieDetails(movieID int) (*MovieDetails, error) {
	var details MovieDetails
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("movieID", fmt.Sprintf("%d", movieID)).
		SetQueryParam("language", o.locale).SetResult(&details).Get("/movie/{movieID}")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &details, nil
}

func (o *Overseerr) GetMovieRecommendations(movieID, page int) (*SearchResults, error) {
	if page < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResults
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("movieID", fmt.Sprintf("%d", movieID)).
		SetQueryParams(map[string]string{
			"language": o.locale,
			"page":     fmt.Sprintf("%d", page),
		}).SetResult(&results).Get("/movie/{movieID}/recommendations")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &results, nil
}

func (o *Overseerr) GetMovieSimilar(movieID, page int) (*SearchResults, error) {
	if page < 1 {
		return nil, fmt.Errorf("page number must be 1 or higher")
	}
	var results SearchResults
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("movieID", fmt.Sprintf("%d", movieID)).
		SetQueryParams(map[string]string{
			"language": o.locale,
			"page":     fmt.Sprintf("%d", page),
		}).SetResult(&results).Get("/movie/{movieID}/similar")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &results, nil
}

func (o *Overseerr) GetMovieRatings(movieID int) (*Rating, error) {
	var rating Rating
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("movieID", fmt.Sprintf("%d", movieID)).
		SetResult(&rating).Get("/movie/{movieID}/ratings")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &rating, nil
}
