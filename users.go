package goverseerr

import (
	"fmt"
	"time"
)

type UserType int

const (
	UserTypePlex  UserType = 1
	UserTypeLocal UserType = 2
)

type User struct {
	ID           int            `json:"id"`
	Email        string         `json:"email"`
	PlexToken    string         `json:"plexToken"`
	UserType     UserType       `json:"userType"`
	RequestCount int            `json:"requestCount"`
	Requests     []MediaRequest `json:"requests"`
	Avatar       string         `json:"avatar"`
	Created      time.Time      `json:"createdAt"`
	Modified     time.Time      `json:"updatedAt"`
	Permissions  int            `json:"permissions"`
	Settings     UserSettings   `json:"settings"`
}

type UsersResponse struct {
	PageInfo Page    `json:"pageInfo"`
	Results  []*User `json:"results"`
}

type UserSettings struct {
	NotificationAgents int    `json:"notificationAgents"`
	ID                 int    `json:"id"`
	DiscordID          string `json:"discordId"`
	Region             string `json:"region"`
	Language           string `json:"originalLanguage"`
	PGPKey             string `json:"pgpKey"`
	TelegramChatID     int    `json:"telegramChatId"`
	TelegramSilent     bool   `json:"telegramSendSilently"`
}

type GenerealUserSettings struct {
	Username        string `json:"username"`
	MovieQuotaLimit int    `json:"movieQuotaLimit"`
	MovieQuotaDays  int    `json:"movieQuotaDays"`
	TVQuotaLimit    int    `json:"tvQuotaLimit"`
	TVQuotaDays     int    `json:"tvQuotaDays"`
}

type UserQuota struct {
	MovieQuota MediaQuota `json:"movie"`
	TVQuota    MediaQuota `json:"tv"`
}

type MediaQuota struct {
	Days       int  `json:"days"`
	Limit      int  `json:"limit"`
	Used       int  `json:"used"`
	Remaining  int  `json:"remaining"`
	Restricted bool `json:"restricted"`
}

// User

func (o *Overseerr) GetAllUsers(pageSize, pageNumber int) ([]*User, *Page, error) {
	var usersResponse UsersResponse
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetQueryParams(map[string]string{
			"take": fmt.Sprintf("%d", pageSize),
			"skip": fmt.Sprintf("%d", pageSize*pageNumber),
		}).
		SetResult(&usersResponse).Get("/user")
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return usersResponse.Results, &usersResponse.PageInfo, nil
}

func (o *Overseerr) GetUser(userID int) (*User, error) {
	var user User
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("userID", fmt.Sprintf("%d", userID)).
		SetResult(&user).Get("/user/{userID}")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &user, nil
}

func (o *Overseerr) CreateNewUser(newUser User) (*User, error) {
	var user User
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetBody(newUser).
		SetResult(&user).Post("/user")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 201 {
		return nil, fmt.Errorf("received non-201 status code (%d)", resp.StatusCode())
	}
	return &user, nil
}

func (o *Overseerr) UpdateUser(userID int, updatedUser User) (*User, error) {
	var user User
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("userID", fmt.Sprintf("%d", userID)).
		SetBody(updatedUser).SetResult(&user).Put("/user/{userID}")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &user, nil
}

func (o *Overseerr) DeleteUser(userID int) (*User, error) {
	var user User
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("userID", fmt.Sprintf("%d", userID)).
		SetResult(&user).Delete("/user/{userID}")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &user, nil
}

func (o *Overseerr) GetLoggedInUser() (*User, error) {
	var user User
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&user).Get("/auth/me")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &user, nil
}

func (o *Overseerr) GetUserQuota(userID int) (*UserQuota, error) {
	var quota UserQuota
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("userID", fmt.Sprintf("%d", userID)).
		SetResult(&quota).Get("/user/{userID}/quota")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &quota, nil
}

func (o *Overseerr) GetUserRequests(userID int, pageNumber, pageSize int) ([]*MediaRequest, *Page, error) {
	var requests MediaRequestResponse
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("userID", fmt.Sprintf("%d", userID)).
		SetQueryParams(map[string]string{
			"take": fmt.Sprintf("%d", pageSize),
			"skip": fmt.Sprintf("%d", pageSize*pageNumber),
		}).
		SetResult(&requests).Get("/user/{userID}/requests")
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return requests.Results, &requests.PageInfo, nil
}

func (o *Overseerr) GetUserGeneralSettings(userID int) (*GenerealUserSettings, error) {
	var settings GenerealUserSettings
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("userID", fmt.Sprintf("%d", userID)).
		SetResult(&settings).Get("/user/{userID}/settings/main")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &settings, nil
}

func (o *Overseerr) SetUserGeneralSettings(userID int, new GenerealUserSettings) error {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("userID", fmt.Sprintf("%d", userID)).
		SetBody(new).Post("/user/{userID}/settings/main")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return nil
}

func (o *Overseerr) ImportPlexUsers() ([]*User, error) {
	var newUsers []*User
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&newUsers).Post("/user/import-from-plex")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 201 {
		return nil, fmt.Errorf("received non-201 status code (%d)", resp.StatusCode())
	}
	return newUsers, nil
}

func (t UserType) ToString() string {
	switch t {
	case UserTypePlex:
		return "Plex User"
	default:
		return "Unknown"
	}
}
