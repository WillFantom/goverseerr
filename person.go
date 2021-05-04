package goverseerr

import "fmt"

type PersonDetails struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	DeathDay     string   `json:"deathday"`
	AKA          []string `json:"alsoKnownAs"`
	Gender       string   `json:"gender"`
	Bio          string   `json:"biography"`
	Popularity   string   `json:"popularity"`
	PlaceOfBirth string   `json:"placeOfBirth"`
	ProfilePath  string   `json:"profilePath"`
	Adult        bool     `json:"adult"`
	IMDBID       string   `json:"imdbId"`
	Homepage     string   `json:"homepage"`
}

func (o *Overseerr) GetPersonDetails(personID int) (*PersonDetails, error) {
	var details PersonDetails
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("personID", fmt.Sprintf("%d", personID)).
		SetQueryParam("language", o.locale).SetResult(&details).Get("/person/{personID}")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &details, nil
}
