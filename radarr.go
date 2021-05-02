package goverseerr

import "fmt"

type RadarrSettings struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Hostname        string `json:"hostname"`
	Port            int    `json:"port"`
	APIKey          string `json:"apiKey"`
	SSL             bool   `json:"useSsl"`
	BaseURL         string `json:"baseUrl"`
	ProfileID       int    `json:"activeProfileId"`
	ProfileName     string `json:"activeProfileName"`
	Directory       string `json:"activeDirectory"`
	UHD             bool   `json:"is4k"`
	MinAvailability string `json:"minimumAvailability"`
	Default         bool   `json:"isDefault"`
	ExternalURL     string `json:"externalUrl"`
	SyncEnabled     bool   `json:"syncEnabled"`
	PreventSearch   bool   `json:"preventSearch"`
}

func (o *Overseerr) GetRadarrSettings() ([]*RadarrSettings, error) {
	var settings []*RadarrSettings
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&settings).Get("/settings/radarr")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return settings, nil
}

func (o *Overseerr) AddRadarr(settings RadarrSettings) (*RadarrSettings, error) {
	var settingResponse RadarrSettings
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetResult(&settingResponse).
		SetBody(settings).Post("/settings/radarr")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 201 {
		return nil, fmt.Errorf("received non-201 status code (%d)", resp.StatusCode())
	}
	return &settingResponse, nil
}

func (o *Overseerr) UpdateRadarrSettings(newSettings RadarrSettings, radarrID int) error {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("radarrID", fmt.Sprintf("%d", radarrID)).
		SetBody(newSettings).Post("/settings/radarr/{radarrID}")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return nil
}
func (o *Overseerr) DeleteRadarr(radarrID int) error {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("radarrID", fmt.Sprintf("%d", radarrID)).
		Delete("/settings/radarr/{radarrID}")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return nil
}

func (o *Overseerr) TestRadarr(settings RadarrSettings) error {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(settings).Post("/settings/radarr/test")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return nil
}

func (o *Overseerr) GetAllRadarrProfiles(radarrID int) ([]*ServiceProfile, error) {
	var profiles []*ServiceProfile
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("radarrID", fmt.Sprintf("%d", radarrID)).
		SetResult(&profiles).Get("/settings/radarr/{radarrID}/profiles")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return profiles, nil
}
