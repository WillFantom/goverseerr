package goverseerr

import "fmt"

type RadarrService struct {
	Server      RadarrSettings   `json:"server"`
	Profiles    []ServiceProfile `json:"profiles"`
	RootFolders []RootFolder     `json:"rootFolders"`
}

type SonarrService struct {
	Server           SonarrSettings    `json:"server"`
	Profiles         []ServiceProfile  `json:"profiles"`
	RootFolders      []RootFolder      `json:"rootFolders"`
	LanguageProfiles []LanguageProfile `json:"languageProfiles"`
}

type ServiceProfile struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type LanguageProfile struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	UpgradeAllowed bool   `json:"upgradeAllowed"`
}

type RootFolder struct {
	ID        int    `json:"id"`
	Path      string `json:"path"`
	FreeSpace int    `json:"freeSpace"`
}

func (o *Overseerr) GetRadarrServers() ([]*RadarrSettings, error) {
	var settings []*RadarrSettings
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&settings).Get("/service/radarr")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return settings, nil
}

func (o *Overseerr) GetRadarrProfiles(radarrID int) (*RadarrService, error) {
	var services RadarrService
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("radarrID", fmt.Sprintf("%d", radarrID)).
		SetResult(&services).Get("/service/radarr/{radarrID}")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &services, nil
}

func (o *Overseerr) GetSonarrServers() ([]*SonarrSettings, error) {
	var settings []*SonarrSettings
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&settings).Get("/service/sonarr")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return settings, nil
}

func (o *Overseerr) GetSonarrProfiles(sonarrID int) (*SonarrService, error) {
	var services SonarrService
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("sonarrID", fmt.Sprintf("%d", sonarrID)).
		SetResult(&services).Get("/service/sonarr/{sonarrID}")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &services, nil
}
