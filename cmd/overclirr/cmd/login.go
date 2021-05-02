package cmd

import (
	"errors"
	"regexp"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/willfantom/goverseerr/cmd/overclirr/overseerr"
	"github.com/willfantom/goverseerr/cmd/overclirr/ui"
)

func nonEmptyValidator(in string) error {
	if len(in) == 0 {
		return errors.New("cannot be empty")
	}
	return nil
}

func emailValidator(in string) error {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(in) {
		return errors.New("must be a valid email address")
	}
	return nil
}

var loginCmd = &cobra.Command{
	Use:     "login",
	Aliases: []string{"new-profile", "add-profile"},
	Short:   "Add a login/user profile for an Overseerr instance",
	Long:    "Add a login/user profile for an Overseer instance with a TUI session or by using optional command flags",
	Args:    cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		currentProfiles, err := overseerr.GetOverseerrProfiles()
		if err != nil {
			ui.PrettyFatal("Could not interact with the configuration file correctly")
			logrus.WithField("extended", err.Error()).Fatalln("Could not interact with the configuration file correctly")
		}
		for name := range currentProfiles {
			if name == overseerrProfileName {
				ui.PrettyInfo("This will overtire the profile with the name " + name)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrettyHeader("Login to an Overseerr instance")
		var er error
		url, _ := cmd.Flags().GetString("url")
		if err := overseerr.URLValidator(url); err != nil {
			url, er = ui.GetInput("Enter the URL of the Overseerr instance", overseerr.URLValidator)
			if er != nil {
				logrus.WithField("extended", er.Error()).Fatalln("url input failed")
			}
		}
		locale, _ := cmd.Flags().GetString("locale")
		if err := overseerr.LocaleValidator(locale); err != nil {
			locale, er = ui.GetInput("Enter the locale for the Overseerr instance (e.g. 'en')", overseerr.LocaleValidator)
			if er != nil {
				logrus.WithField("extended", er.Error()).Fatalln("locale input failed")
			}
		}
		authTypes := []string{"API Key (for admins)", "Email/Password", "Plex Token"}
		selection, _, er := ui.RunSelector("How would you like to authenticate with Overseerr", authTypes)
		if er != nil {
			logrus.WithField("extended", er.Error()).Fatalln("auth type selection failed")
		}
		var config overseerr.OverseerrProfile
		config.URL = url
		config.Locale = locale
		switch selection {
		case 0:
			key, _ := cmd.Flags().GetString("apikey")
			if err := nonEmptyValidator(key); err != nil {
				key, er = ui.GetInput("Enter the API Key", nonEmptyValidator)
				if er != nil {
					logrus.WithField("extended", er.Error()).Fatalln("api key input failed")
				}
			}
			config.Auth.Key = key
			config.Auth.Type = overseerr.OverseerrAuthTypeKey
		case 1:
			email, _ := cmd.Flags().GetString("email")
			if err := emailValidator(email); err != nil {
				email, er = ui.GetInput("Enter the account Email", emailValidator)
				if er != nil {
					logrus.WithField("extended", er.Error()).Fatalln("email input failed")
				}
			}
			password, _ := cmd.Flags().GetString("password")
			if err := nonEmptyValidator(password); err != nil {
				password, er = ui.GetMaskedInput("Enter the account Password", nonEmptyValidator)
				if er != nil {
					logrus.WithField("extended", er.Error()).Fatalln("password input failed")
				}
			}
			config.Auth.Email = email
			config.Auth.Password = password
			config.Auth.Type = overseerr.OverseerrAuthTypeLocal
		case 2:
			ui.PrettyInfo("You can find your Plex user's token easily", "See: https://support.plex.tv/articles/204059436-finding-an-authentication-token-x-plex-token/")
			token, _ := cmd.Flags().GetString("plex-token")
			if err := nonEmptyValidator(token); err != nil {
				token, er = ui.GetMaskedInput("Enter the account's Plex token", nonEmptyValidator)
				if er != nil {
					logrus.WithField("extended", er.Error()).Fatalln("plex token input failed")
				}
			}
			config.Auth.PlexToken = token
			config.Auth.Type = overseerr.OverseerrAuthTypePlex
		default:
			ui.PrettyFatal("Hmm, something went really wrong with the auth type selector!")
			logrus.Panicln("invalid auth type selected... somehow")
		}
		name, _ := cmd.Flags().GetString("name")
		if name != defaultProfile {
			ui.PrettyInfo("A login profile needs a name", "Call it `"+defaultProfile+"` to set it as you default login profile")
			if err := nonEmptyValidator(name); err != nil {
				name, er = ui.GetInput("Enter a name for this login profile", nonEmptyValidator)
				if er != nil {
					logrus.WithField("extended", er.Error()).Fatalln("name input failed")
				}
			}
		}
		if err := overseerr.AddOverseerrProfile(name, config, true); err != nil {
			ui.PrettyFatal("This login profile could not be added to the configuration")
			logrus.WithField("extended", er.Error()).Fatalln("could not add overseerr profile to configuration")
		}
		if _, err := overseerr.GetOverseerrFromProfile(name); err != nil {
			ui.PrettyFatal("Could not login using the given information!")
			logrus.WithField("extended", er.Error()).Fatalln("login failed with newly added profile")
		} else {
			ui.PrettySuccess("Added and checked login profile " + name)
		}
	},
}

func init() {
	loginCmd.Flags().String("url", "", "The full URL of the Overseer instance (optional)")
	loginCmd.Flags().String("locale", "en", "The locale for the Overseerr instance (optional)")
	loginCmd.Flags().String("apikey", "", "The API key for admin access (optional)")
	loginCmd.Flags().String("email", "", "The email to use for local login (optional)")
	loginCmd.Flags().String("password", "", "The password to use for local login (optional)")
	loginCmd.Flags().String("plex-token", "", "The plex token to use for plex auth login (optional)")

	loginCmd.Flags().String("name", "default", "The name for this login in the configuration file (optional)")
	RootCmd.AddCommand(loginCmd)
}
