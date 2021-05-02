package ui

import (
	"github.com/sirupsen/logrus"
	"github.com/willfantom/goverseerr"
)

func selectNSRadarr(o *goverseerr.Overseerr) *goverseerr.RadarrSettings {
	servers, err := o.GetRadarrServers()
	if err != nil {
		PrettyFatal("Could not get list of Radarr servers")
		logrus.WithField("extended", err.Error()).Fatalln("could not get list of radarr services")
	}
	options := make([]string, len(servers))
	for idx, svr := range servers {
		options[idx] = svr.Name + " (" + svr.ExternalURL + ")"
	}
	idx, _, err := RunSelector("Select a Radarr service", options)
	if err != nil {
		logrus.WithField("extended", err.Error()).Fatalln("radarr selector failed")
	}
	return servers[idx]
}

func selectRadarrProfile(o *goverseerr.Overseerr, id int) *goverseerr.ServiceProfile {
	radarrs, err := o.GetRadarrProfiles(id)
	if err != nil {
		PrettyFatal("Could not get list of Radarr profiles")
		logrus.WithField("extended", err.Error()).Fatalln("could not get list of radarr profiles")
	}
	profiles := radarrs.Profiles
	options := make([]string, len(profiles))
	for idx, pr := range profiles {
		options[idx] = pr.Name
	}
	idx, _, err := RunSelector("Select a Radarr profile", options)
	if err != nil {
		logrus.WithField("extended", err.Error()).Fatalln("radarr profile selector failed")
	}
	return &profiles[idx]
}

func selectRadarrRootFolder(o *goverseerr.Overseerr, id int) *goverseerr.RootFolder {
	radarrs, err := o.GetRadarrProfiles(id)
	if err != nil {
		PrettyFatal("Could not get list of Radarr folders")
		logrus.WithField("extended", err.Error()).Fatalln("could not get list of radarr folders")
	}
	folders := radarrs.RootFolders
	options := make([]string, len(folders))
	for idx, pr := range folders {
		options[idx] = pr.Path
	}
	idx, _, err := RunSelector("Select a Radarr folder", options)
	if err != nil {
		logrus.WithField("extended", err.Error()).Fatalln("radarr folder selector failed")
	}
	return &folders[idx]
}
