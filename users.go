package goverseerr

import (
	"fmt"
	"time"
)

type UserType int

const (
	UserTypePlex UserType = 1
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
	DiscordID string `json:"discordId"`
	Region    string `json:"region"`
	Language  string `json:"language"`
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

func (o *Overseerr) GetAllUsers(pageSize, pageNumber int) ([]*User, error) {
	var usersResponse UsersResponse
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetQueryParams(map[string]string{
			"take": fmt.Sprintf("%d", pageSize),
			"skip": fmt.Sprintf("%d", pageSize*pageNumber),
		}).
		SetResult(&usersResponse).Get("/user")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return usersResponse.Results, nil
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

func (o *Overseerr) GetUserRequests(userID int, pageNumber, pageSize int) ([]*MediaRequest, error) {
	var requests MediaRequestResponse
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("userID", fmt.Sprintf("%d", userID)).
		SetQueryParams(map[string]string{
			"take": fmt.Sprintf("%d", pageSize),
			"skip": fmt.Sprintf("%d", pageSize*pageNumber),
		}).
		SetResult(&requests).Get("/user/{userID}/requests")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return requests.Results, nil
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
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
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
