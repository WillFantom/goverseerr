package ui

import (
	"github.com/sirupsen/logrus"
	"github.com/willfantom/goverseerr"
)

func selectNSSonarr(o *goverseerr.Overseerr) *goverseerr.SonarrSettings {
	servers, err := o.GetSonarrServers()
	if err != nil {
		PrettyFatal("Could not get list of Sonarr servers")
		logrus.WithField("extended", err.Error()).Fatalln("could not get list of sonarr services")
	}
	options := make([]string, len(servers))
	for idx, svr := range servers {
		options[idx] = svr.Name + " (" + svr.ExternalURL + ")"
	}
	idx, _, err := RunSelector("Select a Sonarr service", options)
	if err != nil {
		logrus.WithField("extended", err.Error()).Fatalln("sonarr selector failed")
	}
	return servers[idx]
}

func selectSonarrProfile(o *goverseerr.Overseerr, id int) *goverseerr.ServiceProfile {
	sonarrs, err := o.GetSonarrProfiles(id)
	if err != nil {
		PrettyFatal("Could not get list of Sonarr profiles")
		logrus.WithField("extended", err.Error()).Fatalln("could not get list of sonarr profiles")
	}
	profiles := sonarrs.Profiles
	options := make([]string, len(profiles))
	for idx, pr := range profiles {
		options[idx] = pr.Name
	}
	idx, _, err := RunSelector("Select a Sonarr profile", options)
	if err != nil {
		logrus.WithField("extended", err.Error()).Fatalln("sonarr profile selector failed")
	}
	return &profiles[idx]
}

func selectSonarrRootFolder(o *goverseerr.Overseerr, id int) *goverseerr.RootFolder {
	sonarrs, err := o.GetSonarrProfiles(id)
	if err != nil {
		PrettyFatal("Could not get list of Sonarr folders")
		logrus.WithField("extended", err.Error()).Fatalln("could not get list of sonarr folders")
	}
	folders := sonarrs.RootFolders
	options := make([]string, len(folders))
	for idx, pr := range folders {
		options[idx] = pr.Path
	}
	idx, _, err := RunSelector("Select a Sonarr folder", options)
	if err != nil {
		logrus.WithField("extended", err.Error()).Fatalln("sonarr folder selector failed")
	}
	return &folders[idx]
}

func selectSonarrLanguageProfile(o *goverseerr.Overseerr, id int) *goverseerr.LanguageProfile {
	sonarrs, err := o.GetSonarrProfiles(id)
	if err != nil {
		PrettyFatal("Could not get list of Sonarr profiles")
		logrus.WithField("extended", err.Error()).Fatalln("could not get list of sonarr profiles")
	}
	profiles := sonarrs.LanguageProfiles
	options := make([]string, len(profiles))
	for idx, pr := range profiles {
		options[idx] = pr.Name
	}
	idx, _, err := RunSelector("Select a Sonarr profile", options)
	if err != nil {
		logrus.WithField("extended", err.Error()).Fatalln("sonarr profile selector failed")
	}
	return &profiles[idx]
}
