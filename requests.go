package goverseerr

import (
	"fmt"
	"time"
)

type MediaRequestResponse struct {
	PageInfo Page            `json:"pageInfo"`
	Results  []*MediaRequest `json:"results"`
}

type NewRequest struct {
	MediaType         string `json:"mediaType"`
	MediaID           int    `json:"mediaId"`
	TVDBID            int    `json:"tvdbId"`
	Seasons           []int  `json:"seasons"`
	UHD               bool   `json:"is4k"`
	ServerID          int    `json:"serverId"`
	ProfileID         int    `json:"profileId"`
	RootFolder        string `json:"rootFolder"`
	LanguageProfileID int    `json:"languageProfileId"`
}

type MediaRequest struct {
	ID          int           `json:"id"`
	Status      RequestStatus `json:"status"`
	Media       MediaInfo     `json:"media"`
	Created     time.Time     `json:"createdAt"`
	Modified    time.Time     `json:"updatedAt"`
	Creator     User          `json:"requestedBy"`
	LastModifer User          `json:"updatedBy"`
	IsUHD       bool          `json:"is4k"`
	RootFolder  string        `json:"rootFolder"`
	ServerID    int           `json:"serverId"`
	ProfileID   int           `json:"profileID"`
}

type RequestStatus int
type RequestSort string
type RequestFilter string

type Page struct {
	Page    int `json:"page"`
	Pages   int `json:"pages"`
	Results int `json:"results"`
}

const (
	RequestStatusPending   RequestStatus = 1
	RequestStatusApproved  RequestStatus = 2
	RequestStatusDeclined  RequestStatus = 3
	RequestStatusAvailable RequestStatus = 4
)

const (
	RequestSortAdded    RequestSort = "added"
	RequestSortModified RequestSort = "modified"
)

const (
	RequestFileterAll         RequestFilter = "all"
	RequestFileterApproved    RequestFilter = "approved"
	RequestFileterPending     RequestFilter = "pending"
	RequestFileterAvailable   RequestFilter = "available"
	RequestFileterUnavailable RequestFilter = "unavailable"
)

func (o *Overseerr) GetRequests(pageNumber, pageSize int, filter RequestFilter, sort RequestSort) ([]*MediaRequest, error) {
	var requests MediaRequestResponse
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetQueryParams(map[string]string{
		"take":   fmt.Sprintf("%d", pageSize),
		"skip":   fmt.Sprintf("%d", pageSize*pageNumber),
		"filter": string(filter),
		"sort":   string(sort),
	}).SetResult(&requests).Get("/request")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return requests.Results, nil
}

func (o *Overseerr) GetRequestsByUser(pageNumber, pageSize, userID int, filter RequestFilter, sort RequestSort) (*[]MediaRequest, error) {
	var requests []MediaRequest
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetQueryParams(map[string]string{
		"take":        fmt.Sprintf("%d", pageSize),
		"skip":        fmt.Sprintf("%d", pageSize*pageNumber),
		"filter":      string(filter),
		"sort":        string(sort),
		"requestedBy": fmt.Sprintf("%d", userID),
	}).SetResult(&requests).Get("/request")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &requests, nil
}

func (o *Overseerr) CreateRequest(request NewRequest) (*MediaRequest, error) {
	var requestConfirmed MediaRequest
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetBody(request).
		SetResult(&requestConfirmed).Post("/request")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &requestConfirmed, nil
}

func (o *Overseerr) setRequestStatus(requestID int, status string) (*MediaRequest, error) {
	var requestConfirmed MediaRequest
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParams(map[string]string{
		"requestID": fmt.Sprintf("%d", requestID),
		"status":    status,
	}).
		SetResult(&requestConfirmed).Post("/request/{requestID}/{status}")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &requestConfirmed, nil
}

func (o *Overseerr) ApproveRequest(requestID int) (*MediaRequest, error) {
	return o.setRequestStatus(requestID, "approve")
}

func (o *Overseerr) DeclineRequest(requestID int) (*MediaRequest, error) {
	return o.setRequestStatus(requestID, "decline")
}

func (o *Overseerr) GetRequest(requestID int) (*MediaRequest, error) {
	var request MediaRequest
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParams(map[string]string{
		"requestID": fmt.Sprintf("%d", requestID),
	}).
		SetResult(&request).Get("/request/{requestID}")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &request, nil
}

func (o *Overseerr) DeleteRequest(requestID int) error {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("requestID", fmt.Sprintf("%d", requestID)).
		Delete("/request/{requestID}")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return nil
}

func (req MediaRequest) GetTVDetails(o *Overseerr) (*TVDetails, error) {
	if !req.Media.IsTV() {
		return nil, fmt.Errorf("request's media type is not tv")
	}
	return o.GetTV(req.Media.TMDB)
}

func (req MediaRequest) GetMovieDetails(o *Overseerr) (*MovieDetails, error) {
	if !req.Media.IsMovie() {
		return nil, fmt.Errorf("request's media type is not movie")
	}
	return o.GetMovie(req.Media.TMDB)
}

func (s RequestStatus) ToString() string {
	switch s {
	case RequestStatusApproved:
		return "Approved"
	case RequestStatusDeclined:
		return "Declined"
	case RequestStatusAvailable:
		return "Available"
	case RequestStatusPending:
		return "Pending"
	default:
		return "Unknown"
	}
}

func (s RequestStatus) ToEmoji() string {
	switch s {
	case RequestStatusApproved:
		return "✅"
	case RequestStatusDeclined:
		return "❌"
	case RequestStatusAvailable:
		return "🍿"
	case RequestStatusPending:
		return "⏱"
	default:
		return "❓"
	}
}

func StringToRequestStatus(status string) RequestStatus {
	switch status {
	case "Approved":
		return 2
	case "Declined":
		return 3
	case "Available":
		return 4
	case "Pending":
		return 1
	default:
		return 0
	}
}
