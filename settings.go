package goverseerr

import (
	"fmt"
	"time"
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

type PlexDevice struct {
	Name             string         `json:"name"`
	Product          string         `json:"product"`
	ProductVersion   string         `json:"productVersion"`
	Platform         string         `json:"platform"`
	PlatformVersion  string         `json:"platformVersion"`
	Device           string         `json:"device"`
	ClientIdentifier string         `json:"clientIdentifier"`
	Created          time.Time      `json:"createdAt"`
	LastSeen         time.Time      `json:"lastSeenAt"`
	Owned            bool           `json:"owned"`
	OwnerID          string         `json:"ownerId"`
	Home             bool           `json:"home"`
	SourceTitle      string         `json:"sourceTitle"`
	AccessToken      string         `json:"accessToken"`
	PublicAddress    string         `json:"publicAddress"`
	HTTPSRequired    bool           `json:"httpsRequired"`
	Synced           bool           `json:"synced"`
	Relay            bool           `json:"relay"`
	Connection       PlexConnection `json:"connection"`
}

type PlexConnection struct {
	Protocol string `json:"protocol"`
	Address  string `json:"address"`
	Port     int    `json:"port"`
	URI      string `json:"uri"`
	Local    bool   `json:"local"`
	Status   int    `json:"status"`
	Message  string `json:"message"`
}

type PlexLibrary struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

type PlexSettings struct {
	Name      string        `json:"name"`
	MachineID string        `json:"machineId"`
	IP        string        `json:"ip"`
	Port      int           `json:"port"`
	Libraries []PlexLibrary `json:"libraries"`
}

type PlexSyncStatus struct {
	Running        bool          `json:"running"`
	Progress       int           `json:"progress"`
	Total          int           `json:"total"`
	CurrentLibrary PlexLibrary   `json:"currentLibrary"`
	Libraries      []PlexLibrary `json:"libraries"`
}

type ServiceProfile struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

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

type SonarrSettings struct {
	ID                     int    `json:"id"`
	Name                   string `json:"name"`
	Hostname               string `json:"hostname"`
	Port                   int    `json:"port"`
	APIKey                 string `json:"apiKey"`
	SSL                    bool   `json:"useSsl"`
	BaseURL                string `json:"baseUrl"`
	ProfileID              int    `json:"activeProfileId"`
	ProfileName            string `json:"activeProfileName"`
	Directory              string `json:"activeDirectory"`
	LanguageProfileId      int    `json:"activeLanguageProfileId"`
	AnimeProfileID         int    `json:"activeAnimeProfileId"`
	AnimeLanguageProfileId int    `json:"activeAnimeLanguageProfileId"`
	AnimeProfileName       string `json:"activeAnimeProfileName"`
	AnimeDirectory         string `json:"activeAnimeDirectory"`
	UHD                    bool   `json:"is4k"`
	MinAvailability        string `json:"minimumAvailability"`
	Default                bool   `json:"isDefault"`
	ExternalURL            string `json:"externalUrl"`
	SyncEnabled            bool   `json:"syncEnabled"`
	PreventSearch          bool   `json:"preventSearch"`
}

// Main Settings

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

// Plex

func (o *Overseerr) GetPlexSettings() (*PlexSettings, error) {
	var settings PlexSettings
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&settings).Get("/settings/plex")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &settings, nil
}

func (o *Overseerr) UpdatePlexSettings(newSettings PlexSettings) error {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(newSettings).Post("/settings/plex")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return nil
}

func (o *Overseerr) GetPlexLibraries() ([]*PlexLibrary, error) {
	var libraries []*PlexLibrary
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&libraries).Get("/settings/plex/library")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return libraries, nil
}

func (o *Overseerr) GetPlexSyncStatus() (*PlexSyncStatus, error) {
	var status PlexSyncStatus
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&status).Get("/settings/plex/sync")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &status, nil
}

func (o *Overseerr) GetPlexServers() ([]*PlexDevice, error) {
	var devices []*PlexDevice
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&devices).Get("/settings/plex/devices/servers")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return devices, nil
}

func (o *Overseerr) TriggerPlexSync() error {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(map[string]bool{
			"start":  true,
			"cancel": false,
		}).Post("/settings/plex/sync")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return nil
}

func (o *Overseerr) CancelPlexSync() error {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(map[string]bool{
			"start":  false,
			"cancel": true,
		}).Post("/settings/plex/sync")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return nil
}

// Radarr

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
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &settingResponse, nil
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

func (o *Overseerr) GetRadarrProfile(radarrID int) ([]*ServiceProfile, error) {
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

// Sonarr

func (o *Overseerr) GetSonarrSettings() ([]*SonarrSettings, error) {
	var settings []*SonarrSettings
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&settings).Get("/settings/sonarr")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return settings, nil
}

func (o *Overseerr) AddSonarr(settings SonarrSettings) (*SonarrSettings, error) {
	var settingResponse SonarrSettings
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetResult(&settingResponse).
		SetBody(settings).Post("/settings/sonarr")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &settingResponse, nil
}

func (o *Overseerr) TestSonarr(settings SonarrSettings) error {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(settings).Post("/settings/sonarr/test")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return nil
}

func (o *Overseerr) UpdateSonarrSettings(newSettings SonarrSettings, sonarrID int) error {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("sonarrID", fmt.Sprintf("%d", sonarrID)).
		SetBody(newSettings).Post("/settings/sonarr/{sonarrID}")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return nil
}

func (o *Overseerr) DeleteSonarr(sonarrID int) error {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("sonarrID", fmt.Sprintf("%d", sonarrID)).
		Delete("/settings/sonarr/{sonarrID}")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return nil
}
