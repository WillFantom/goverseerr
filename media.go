package goverseerr

import (
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

type Rating struct {
	Title          string `json:"title"`
	Year           int    `json:"year"`
	URL            string `json:"url"`
	CriticScore    int    `json:"criticsScore"`
	CriticRating   string `json:"criticsRating"`
	AudienceScore  int    `json:"audienceScore"`
	AudienceRating string `json:"audienceRating"`
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

type Keyword struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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
	MediaStatusAvailable  MediaStatus = 5
)

func (s MediaStatus) ToString() string {
	switch s {
	case MediaStatusAvailable:
		return "Available"
	case MediaStatusPartial:
		return "Part-Available"
	case MediaStatusProcessing:
		return "Processing"
	case MediaStatusPending:
		return "Pending"
	case MediaStatusUnknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

func (s MediaStatus) ToEmoji() string {
	switch s {
	case MediaStatusAvailable:
		return "✅"
	case MediaStatusPartial:
		return "✔️"
	case MediaStatusProcessing:
		return "🧑‍💻"
	case MediaStatusPending:
		return "⏱"
	case MediaStatusUnknown:
		return "❓"
	default:
		return "❓"
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

func (o *Overseerr) GetMovie(movieID int) (*MovieDetails, []GenericSearchResult, []GenericSearchResult, *Rating, error) {
	details, err := o.GetMovieDetails(movieID)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	recommendations, err := o.GetMovieRecommendations(movieID, 1)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	similar, err := o.GetMovieSimilar(movieID, 1)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	ratings, err := o.GetMovieRatings(movieID)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	return details, recommendations.Results, similar.Results, ratings, nil
}

func (o *Overseerr) GetTV(tvID int) (*TVDetails, []GenericSearchResult, []GenericSearchResult, *Rating, error) {
	details, err := o.GetTVDetails(tvID)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	recommendations, err := o.GetTVRecommendations(tvID, 1)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	similar, err := o.GetTVSimilar(tvID, 1)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	ratings, err := o.GetTVRatings(tvID)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	return details, recommendations.Results, similar.Results, ratings, nil
}
