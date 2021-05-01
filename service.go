package goverseerr

import "fmt"

func (o *Overseerr) GetRadarrServers() ([]*RadarrSettings, error) {
	var settings []*RadarrSettings
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&settings).Get("/services/radarr")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return settings, nil
}

func (o *Overseerr) GetRadarrProfiles(radarrID int) ([]*RadarrSettings, error) {
	var settings []*RadarrSettings
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&settings).Get("/services/radarr")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return settings, nil
}
