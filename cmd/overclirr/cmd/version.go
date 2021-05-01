package cmd

import (
	"context"
	"fmt"

	"github.com/google/go-github/v35/github"
	semver "github.com/hashicorp/go-version"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/willfantom/goverseerr/cmd/overclirr/ui"
)

const (
	defaultVersion string = "no-version"
	ghUser         string = "willfantom"
	ghRepo         string = "goverseerr"
)

var version string = defaultVersion

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of OverCLIrr",
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrettyInfo(fmt.Sprintf("OverCLIrr version: %s\n", version))
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Check if OverCLIrr is up to date",
	Run: func(cmd *cobra.Command, args []string) {
		available, err := checkForUpdate()
		if available {
			ui.PrettyInfo("An update is available", "Check the releases on GitHub")
		} else if !available && err == nil {
			ui.PrettyInfo("You are running the latest version according to GitHub")
		} else {
			ui.PrettyOops("Could not perform version check")
			ui.PrettyInfo("Perhaps this is unversioned?", "Or perhaps you can't connect to GitHub")
			logrus.WithField("extended", err.Error()).Fatalln("could not perform version check")
		}
	},
}

func checkForUpdate() (bool, error) {
	//semver local version
	semverVersion, err := semver.NewSemver(version)
	if err != nil {
		logrus.WithField("extended", err.Error()).
			Errorln("could not determine the version of overclirr")
		return false, err
	}

	//get latest github release tag
	client := github.NewClient(nil)
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), ghUser, ghRepo)
	if err != nil {
		logrus.WithField("extended", err.Error()).
			Errorln("could get repository release info from github")
		return false, err
	}
	ghVer, err := semver.NewSemver(*release.TagName)
	if err != nil {
		logrus.WithField("extended", err.Error()).
			Errorln("latest github release tag not semver compliant")
		return false, err
	}

	//compare
	if ghVer.GreaterThan(semverVersion) {
		logrus.Infoln("found a more recent release on github")
		return true, nil
	}
	logrus.Infoln("overclirr found to be latest version")
	return false, nil
}

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(updateCmd)
}
