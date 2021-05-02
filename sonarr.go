package goverseerr

import "fmt"

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
