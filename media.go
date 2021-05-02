package goverseerr

import (
	"fmt"
	"time"
)

type MediaStatus int
type MediaType string
type RelatedVideoType string
type RelatedVideoSite string

type MediaInfo struct {
	ID         int            `json:"id"`
	TMDB       int            `json:"tmdbID"`
	TVDB       int            `json:"tvdbID"`
	MediaType  MediaType      `json:"mediaType"`
	Status     MediaStatus    `json:"status"`
	Created    time.Time      `json:"createdAt"`
	Modified   time.Time      `json:"updatedAt"`
	Requests   []MediaRequest `json:"requests"`
	PlexURL    string         `json:"plexUrl"`
	ServiceURL string         `json:"serviceUrl"`
}

type Genre struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Backdrops []string `json:"backdrops,omitempty"`
}

type RelatedVideo struct {
	URL  string           `json:"url"`
	Key  string           `json:"key"`
	Name string           `json:"name"`
	Size int              `json:"size"`
	Type RelatedVideoType `json:"type"`
	Site RelatedVideoSite `json:"site"`
}

type ProductionCompany struct {
	ID            int    `json:"id"`
	LogoPath      string `json:"logoPath"`
	OriginCountry string `json:"originCountry"`
	Name          string `json:"name"`
}

type SpokenLanguage struct {
	EnglishName string `json:"englishName"`
	Code        string `json:"iso_639_1"`
	Name        string `json:"name"`
}

type Collection struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Overview     string        `json:"overview"`
	PosterPath   string        `json:"posterPath"`
	BackdropPath string        `json:"backdropPath"`
	Parts        []MovieResult `json:"parts"`
}

type ExternalIDs struct {
	Facebook  string `json:"facebookId"`
	Freebase  string `json:"freebaseId"`
	IMDB      string `json:"imdbId"`
	Instagram string `json:"instagramId"`
	TVDB      int    `json:"tvdbId"`
	Twitter   string `json:"twitterId"`
}

type Network struct {
	ID            int    `json:"id"`
	LogoPath      string `json:"logoPath"`
	OriginCountry string `json:"originCountry"`
	Name          string `json:"name"`
}

type Cast struct {
	ID          int    `json:"id"`
	CastID      int    `json:"castId"`
	Character   string `json:"character"`
	CreditID    string `json:"creditId"`
	Gender      int    `json:"gender"`
	Name        string `json:"name"`
	Order       int    `json:"order"`
	ProfilePath string `json:"profilePath"`
}

type Crew struct {
	ID          int    `json:"id"`
	CreditID    string `json:"creditId"`
	Gender      int    `json:"gender"`
	Name        string `json:"name"`
	Job         string `json:"job"`
	Department  string `json:"department"`
	ProfilePath string `json:"profilePath"`
}

type Credits struct {
	Cast []Cast `json:"cast"`
	Crew []Crew `json:"crew"`
}

type Creator struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Gender      int    `json:"gender"`
	ProfilePath string `json:"profilePath"`
}

type Episode struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	AirDate        string `json:"airDate"`
	Number         int    `json:"episodeNumber"`
	Overview       string `json:"overview"`
	ProductionCode string `json:"productionCode"`
	SeasonNumber   int    `json:"seasonNumber"`
	ShowID         int    `json:"showId"`
	StillPath      string `json:"stillPath"`
	VoteAverage    int    `json:"voteAverage"`
	VoteCount      int    `json:"voteCount"`
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

type Keyword struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

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
}

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

const (
	RelatedVideoTypeClip            RelatedVideoType = "Clip"
	RelatedVideoTypeTeaser          RelatedVideoType = "Teaser"
	RelatedVideoTypeTrailer         RelatedVideoType = "Trailer"
	RelatedVideoTypeFeaturette      RelatedVideoType = "Featurette"
	RelatedVideoTypeOpeningCredits  RelatedVideoType = "Opening Credits"
	RelatedVideoTypeBehindTheScenes RelatedVideoType = "Behind the Scenes"
	RelatedVideoTypeBloopers        RelatedVideoType = "Bloopers"
)
const (
	RelatedVideoSiteYoutube RelatedVideoSite = "YouTube"
)
const (
	MediaTypeTV     MediaType = "tv"
	MediaTypeMovie  MediaType = "movie"
	MediaTypePerson MediaType = "person"
)

const (
	MediaStatusUnknown    MediaStatus = 0
	MediaStatusPending    MediaStatus = 1
	MediaStatusProcessing MediaStatus = 3
	MediaStatusPartial    MediaStatus = 4
	MediaStatsAvailable   MediaStatus = 5
)

func (s MediaStatus) ToString() string {
	switch s {
	case MediaStatsAvailable:
		return "available"
	case MediaStatusPartial:
		return "part-available"
	case MediaStatusProcessing:
		return "processing"
	case MediaStatusPending:
		return "pending"
	case MediaStatusUnknown:
		return "unknown"
	default:
		return "unknown"
	}
}

func (i MediaInfo) IsTV() bool {
	return i.MediaType == MediaTypeTV
}

func (i MediaInfo) IsMovie() bool {
	return i.MediaType == MediaTypeMovie
}

func (i MediaInfo) IsPerson() bool {
	return i.MediaType == MediaTypePerson
}

func (mt MediaType) ToEmoji() string {
	switch mt {
	case MediaTypeMovie:
		return "🎬"
	case MediaTypeTV:
		return "📺"
	case MediaTypePerson:
		return "👤"
	default:
		return "❓"
	}
}

func (o *Overseerr) GetMovie(movieID int) (*MovieDetails, error) {
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

func (o *Overseerr) GetTV(tvID int) (*TVDetails, error) {
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
