package goverseerr

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/go-resty/resty/v2"
)

const (
	apiPrefix      string = "/api/v1"
	userCookieName string = "connect.sid"
)

type Overseerr struct {
	URL        string
	restClient *resty.Client
	locale     string
}

func new(url string, customHeaders map[string]string, locale string) *Overseerr {
	url = strings.TrimSuffix(url, "/")
	oversr := Overseerr{
		restClient: resty.New(),
	}
	oversr.URL = url
	oversr.restClient.SetHostURL(url + apiPrefix)
	oversr.restClient.SetHeaders(customHeaders)
	oversr.locale = locale
	return &oversr
}

// NewKeyAuth creates a new Overseerr client with the X-Api-Key header set.
// An error is returned alongside the client if an auth check fails.
func NewKeyAuth(url string, customHeaders map[string]string, locale string, apikey string) (*Overseerr, error) {
	o := new(url, customHeaders, locale)
	o.restClient.SetHeader("X-Api-Key", apikey)
	_, err := o.GetLoggedInUser()
	return o, err
}

// NewLocalAuth creates a new Overseerr client with the session cookie for auth.
// The cookie is generated via an API call that requires and email and password.
// An error is returned alongside the client if an auth check fails.
func NewLocalAuth(url string, customHeaders map[string]string, locale string, email, password string) (*Overseerr, error) {
	o := new(url, customHeaders, locale)
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(map[string]string{
			"email":    email,
			"password": password,
		}).Post("/auth/plex")
	if err != nil {
		return o, err
	}
	if resp.StatusCode() != 200 {
		return o, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	for _, cookie := range resp.Cookies() {
		if cookie.Name == userCookieName {
			o.restClient.SetCookie(cookie)
			return o, nil
		}
	}
	return o, fmt.Errorf("no auth cookie contained in response")
}

// NewPlexAuth creates a new Overseerr client with the session cookie for auth.
// The cookie is generated via an API call that requires a plex toke.
// An error is returned alongside the client if an auth check fails.
func NewPlexAuth(url string, customHeaders map[string]string, locale string, plexToken string) (*Overseerr, error) {
	o := new(url, customHeaders, locale)
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(map[string]string{
			"authToken": plexToken,
		}).Post("/auth/plex")
	if err != nil {
		return o, err
	}
	if resp.StatusCode() != 200 {
		return o, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	for _, cookie := range resp.Cookies() {
		if cookie.Name == userCookieName {
			o.restClient.SetCookie(cookie)
			return o, nil
		}
	}
	return o, fmt.Errorf("no auth cookie contained in response")
}

type Status struct {
	Version   string `json:"version"`
	CommitTag string `json:"commitTag"`
}

type AppData struct {
	Configured bool   `json:"appData"`
	Path       string `json:"appDataPath"`
}

func (o *Overseerr) Status() (*Status, error) {
	var status Status
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&status).Get("/status")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &status, nil
}

func (o *Overseerr) GetAppData() (*AppData, error) {
	var appdata AppData
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&appdata).Get("/status/appdata")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return &appdata, nil
}

// RegisterPreRequestMiddleware allows for a custom function to be called
// before each request. Useful for logging.
func (o *Overseerr) RegisterPreRequestMiddleware(middleware func(c *resty.Client, req *resty.Request) error) {
	o.restClient.OnBeforeRequest(middleware)
}

// RegisterPostResponseMiddleware allows for a custom function to be called
// after each request. Useful for logging.
func (o *Overseerr) RegisterPostResponseMiddleware(middleware func(c *resty.Client, resp *resty.Response) error) {
	o.restClient.OnAfterResponse(middleware)
}

// RegisterRequestErrorMiddleware allows for a custom function to be called
// if a request encounters an error. Useful for logging.
func (o *Overseerr) RegisterRequestErrorMiddleware(middleware func(req *resty.Request, err error)) {
	o.restClient.OnError(middleware)
}

// SetProxy forces all requests to the Overseerr to go via the given proxy
// and removes the proxy if a blank string is given
func (o *Overseerr) SetProxy(proxyURL string) error {
	if proxyURL == "" {
		o.restClient.RemoveProxy()
		return nil
	}
	if _, err := url.Parse(proxyURL); err == nil {
		o.restClient.SetProxy(proxyURL)
		return nil
	}
	return fmt.Errorf("proxy url is not valid")
}

// SetUserAgent allows requests to the Overseerr instance to bypass any
// security that checks the user agent
func (o *Overseerr) SetUserAgent(ua string) {
	o.restClient.SetHeader("User-Agent", ua)
}

// SetBasicAuth allows requests to the Overseerr instance to get past
// basic authentication using a given username and password
func (o *Overseerr) SetBasicAuth(user, pass string) {
	o.restClient.SetBasicAuth(user, pass)
}

// HealthCheck ensures an Overseerr instance is accessible and that the API
// key provided is valid
func (o *Overseerr) HealthCheck() bool {
	if _, err := o.GetAbout(); err != nil {
		return false
	}
	return true
}
