package goverseerr

import (
	"fmt"
)

type MainSettings struct {
	APIKey                 string `json:"apiKey"`
	AppTitle               string `json:"applicationTitle"`
	AppURL                 string `json:"applicationUrl"`
	TrustProxy             bool   `json:"trustProxy"`
	CSRFProtection         bool   `json:"csrfProtection"`
	HideAvailable          bool   `json:"hideAvailable"`
	PartialRequestsEnabled bool   `json:"partialRequestsEnabled"`
	LocalLogin             bool   `json:"localLogin"`
	DefaultPermissions     int    `json:"defaultPermissions"`
}

type PublicSettings struct {
	Initialized bool `json:"initialized"`
}

type About struct {
	Version         string `json:"version"`
	TotalRequests   int    `json:"totalRequests"`
	TotalMediaItems int    `json:"totalMediaItems"`
	TimeZone        string `json:"tz"`
}

func (o *Overseerr) GetMainSettings() (*MainSettings, error) {
	var settings MainSettings
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&settings).Get("/settings/main")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &settings, nil
}

func (o *Overseerr) UpdateMainSettings(newSettings MainSettings) (*MainSettings, error) {
	var settings MainSettings
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(newSettings).SetResult(&settings).Post("/settings/main")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &settings, nil
}

func (o *Overseerr) RegenerateMainSettings() (*MainSettings, error) {
	var settings MainSettings
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&settings).Get("/settings/main/regenerate")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &settings, nil
}

func (o *Overseerr) GetPublicSettings() (*PublicSettings, error) {
	var settings PublicSettings
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&settings).Get("/settings/public")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &settings, nil
}

func (o *Overseerr) GetAbout() (*About, error) {
	var about About
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&about).Get("/settings/about")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &about, nil
}
