package overseerr

import (
	"errors"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/willfantom/goverseerr"
	"github.com/willfantom/goverseerr/cmd/overclirr/ui"
	"golang.org/x/text/language"
)

type OverseerrAuthType string

const (
	OverseerrAuthTypeKey   OverseerrAuthType = "key"
	OverseerrAuthTypeLocal OverseerrAuthType = "local"
	OverseerrAuthTypePlex  OverseerrAuthType = "plex"
)

type OverseerrAuthConfig struct {
	Type      OverseerrAuthType `json:"type"`
	Key       string            `json:"key,omitempty"`
	Email     string            `json:"email,omitempty"`
	Password  string            `json:"password,omitempty"`
	PlexToken string            `json:"plexToken,omitempty"`
}

type OverseerrProfile struct {
	URL           string              `json:"url"`
	CustomHeaders map[string]string   `json:"customHeaders"`
	Locale        string              `json:"locale" default:"en"`
	Auth          OverseerrAuthConfig `json:"auth"`
}

func (c OverseerrProfile) quickValidate() error {
	if err := URLValidator(c.URL); err != nil {
		return err
	}
	if err := LocaleValidator(c.Locale); err != nil {
		return err
	}
	return nil
}

func URLValidator(u string) error {
	if len(u) == 0 {
		return errors.New("url must not be empty")
	}
	if parsedURL, err := url.Parse(u); err != nil {
		return errors.New("not parseable as a url")
	} else {
		if !parsedURL.IsAbs() {
			return errors.New("url must be absolute (have a scheme such as https://)")
		}
	}
	return nil
}

func LocaleValidator(l string) error {
	if _, err := language.Parse(l); err != nil {
		return errors.New("locale not known")
	}
	return nil
}

var preRequestMiddleware = func(c *resty.Client, req *resty.Request) error {
	logrus.WithFields(logrus.Fields{
		"url":        c.HostURL,
		"authscheme": c.AuthScheme,
		"requestUrl": req.URL,
	}).Traceln("made api request")
	return nil
}

var postResponseMiddleware = func(c *resty.Client, resp *resty.Response) error {
	logrus.WithFields(logrus.Fields{
		"url":            c.HostURL,
		"authscheme":     c.AuthScheme,
		"responseStatus": resp.Status(),
	}).Traceln("received api response")
	return nil
}

var requestErrorMiddleware = func(req *resty.Request, err error) {
	if v, ok := err.(*resty.ResponseError); ok {
		logrus.WithFields(logrus.Fields{
			"lastResponse":  v.Response,
			"originalError": v.Err,
		}).Errorln("api request error")
	}
	logrus.Errorln("inextendable api request error")
}

func GetOverseerrProfiles() (map[string]OverseerrProfile, error) {
	overseerrProfiles := make(map[string]OverseerrProfile)
	if err := viper.UnmarshalKey("overseerrs", &overseerrProfiles); err != nil {
		ui.PrettyFatal("Failed to parse Overseerr configs from the file")
		logrus.WithField("extended", err.Error()).Errorln("failed to parse/unmarshal overseerr configs from file/viper")
		return nil, err
	}
	return overseerrProfiles, nil
}

func GetOverseerrFromProfile(name string) (*goverseerr.Overseerr, error) {
	overseerrProfiles, err := GetOverseerrProfiles()
	if err != nil {
		return nil, err
	}
	if config, ok := overseerrProfiles[name]; !ok {
		return nil, errors.New("no config found named " + name)
	} else {
		if err := config.quickValidate(); err != nil {
			return nil, err
		}
		switch config.Auth.Type {
		case OverseerrAuthTypeKey:
			return goverseerr.NewKeyAuth(config.URL, config.CustomHeaders, config.Locale, config.Auth.Key)
		case OverseerrAuthTypeLocal:
			return goverseerr.NewLocalAuth(config.URL, config.CustomHeaders, config.Locale, config.Auth.Email, config.Auth.Password)
		case OverseerrAuthTypePlex:
			return goverseerr.NewPlexAuth(config.URL, config.CustomHeaders, config.Locale, config.Auth.PlexToken)
		default:
			return nil, errors.New("no valid auth found for config " + name)
		}
	}
}

func AddWrappersToOverseerr(o *goverseerr.Overseerr) {
	o.RegisterPostResponseMiddleware(postResponseMiddleware)
	o.RegisterPreRequestMiddleware(preRequestMiddleware)
	o.RegisterRequestErrorMiddleware(requestErrorMiddleware)
}

func AddOverseerrProfile(name string, config OverseerrProfile, overwrite bool) error {
	overseerrProfiles, err := GetOverseerrProfiles()
	if err != nil {
		return err
	}
	if _, ok := overseerrProfiles[name]; ok && !overwrite {
		return errors.New("overseer profile with that name already exists")
	}
	overseerrProfiles[name] = config
	viper.Set("overseerrs", overseerrProfiles)
	if err := viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}

func RemoveOverseerProfile(name string) error {
	overseerrProfiles, err := GetOverseerrProfiles()
	if err != nil {
		return err
	}
	if _, ok := overseerrProfiles[name]; !ok {
		return errors.New("no overseerr profile with name `" + name + "` exists to delete")
	}
	delete(overseerrProfiles, name)
	viper.Set("overseerrs", overseerrProfiles)
	if err := viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}
