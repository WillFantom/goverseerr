package goverseerr

import (
	"fmt"
	"time"
)

type PlexDevice struct {
	Name                   string         `json:"name"`
	Product                string         `json:"product"`
	ProductVersion         string         `json:"productVersion"`
	Platform               string         `json:"platform"`
	PlatformVersion        string         `json:"platformVersion"`
	Device                 string         `json:"device"`
	ClientIdentifier       string         `json:"clientIdentifier"`
	Created                time.Time      `json:"createdAt"`
	LastSeen               time.Time      `json:"lastSeenAt"`
	Provides               []string       `json:"provides"`
	Owned                  bool           `json:"owned"`
	OwnerID                string         `json:"ownerId"`
	Home                   bool           `json:"home"`
	SourceTitle            string         `json:"sourceTitle"`
	AccessToken            string         `json:"accessToken"`
	PublicAddress          string         `json:"publicAddress"`
	HTTPSRequired          bool           `json:"httpsRequired"`
	Synced                 bool           `json:"synced"`
	Relay                  bool           `json:"relay"`
	DNSRebindingProtection bool           `json:"dnsRebindingProtection"`
	NATLoopbackSupported   bool           `json:"natLoopbackSupported"`
	PublicAddressMatches   bool           `json:"publicAddressMatches"`
	Presence               bool           `json:"presence"`
	Connection             PlexConnection `json:"connection"`
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
	WebAppURL string        `json:"webAppUrl"`
}

type PlexSyncStatus struct {
	Running        bool          `json:"running"`
	Progress       int           `json:"progress"`
	Total          int           `json:"total"`
	CurrentLibrary PlexLibrary   `json:"currentLibrary"`
	Libraries      []PlexLibrary `json:"libraries"`
}

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
