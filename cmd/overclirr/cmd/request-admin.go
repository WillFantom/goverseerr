package cmd

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/willfantom/goverseerr"
	"github.com/willfantom/goverseerr/cmd/overclirr/overseerr"
	"github.com/willfantom/goverseerr/cmd/overclirr/ui"
)

var requestsAdminCmd = &cobra.Command{
	Use:   "requests",
	Short: "View or manage Overseerr media requests",
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		instance = getOverseerrInstance(overseerrProfileName)
	},
	Run: func(cmd *cobra.Command, args []string) {
		requests := getAllRequests()
		request := overseerr.SelectRequest(instance, requests)
		overseerr.PrintRequest(instance, request)
		overseerr.RequestAction(request, instance)
	},
}

var retryAllRequestsCmd = &cobra.Command{
	Use:   "retry-all",
	Short: "Resend all requests to sonarr/radarr",
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		instance = getOverseerrInstance(overseerrProfileName)
	},
	Run: func(cmd *cobra.Command, args []string) {
		ui.DestructiveConfirmation()
		requests := getAllRequests()
		var wg sync.WaitGroup
		ui.StartLoadingSpinner()
		for c := 0; c < len(requests); c++ {
			wg.Add(1)
			go func(c int, reqs []*goverseerr.MediaRequest) {
				overseerr.RetryRequest(instance, reqs[c], false)
				wg.Done()
			}(c, requests)
			if c > 10 {
				wg.Wait()
			}
		}
		ui.StopLoadingSpinner()
		wg.Wait()
		ui.PrettyInfo("Done, but errors will not have been handled")

	},
}

func getAllRequests() []*goverseerr.MediaRequest {
	var allRequests []*goverseerr.MediaRequest
	pg := 0
	for {
		requests, page, err := instance.GetRequests(pg, 20, goverseerr.RequestFileterAll, goverseerr.RequestSortAdded)
		if err != nil {
			ui.PrettyFatal("Could not get requests from Overseerr")
			logrus.WithField("extended", err.Error()).Fatalln("could not get overseerr request list")
		}
		allRequests = append(allRequests, requests...)
		pg++
		if pg >= page.Pages {
			return allRequests
		}
	}
}

func init() {
	requestsAdminCmd.AddCommand(retryAllRequestsCmd)
	adminCmd.AddCommand(requestsAdminCmd)
}
