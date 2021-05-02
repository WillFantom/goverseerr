package cmd

import (
	"errors"
	"fmt"
	"sync"

	"github.com/spf13/cobra"
	"github.com/willfantom/goverseerr"
	"github.com/willfantom/goverseerr/cmd/overclirr/newui"
	"github.com/willfantom/goverseerr/cmd/overclirr/overseerr"
	"github.com/willfantom/goverseerr/cmd/overclirr/ui"
)

var (
	reqPageSize   int
	reqPageNumber int
	reqAll        bool
)

var requestsAdminCmd = &cobra.Command{
	Use:   "requests",
	Short: "View or manage Overseerr media requests",
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		instance = getOverseerrInstance(overseerrProfileName)
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var requestsAdminManageCmd = &cobra.Command{
	Use:   "manage",
	Short: "View or manage individual Overseerr media requests",
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		instance = getOverseerrInstance(overseerrProfileName)
	},
	Run: func(cmd *cobra.Command, args []string) {
		newui.StartSpinner()
		var requests []*goverseerr.MediaRequest
		var pages int
		if !reqAll {
			requests, pages = getRequests(reqPageNumber-1, reqPageSize)
		} else {
			requests, pages = getAllRequests()
		}
		tvDetails, movieDetails := getRequestDetails(requests)
		newui.StopLoadingSpinner()
		request := newui.RequestSelector(requests, tvDetails, movieDetails, reqPageNumber, pages)
		overseerr.PrintRequest(instance, requests[request])
		//overseerr.RequestAction(request, instance)
	},
}

func getRequestDetails(requests []*goverseerr.MediaRequest) ([]*goverseerr.TVDetails, []*goverseerr.MovieDetails) {
	movieDetails := make([]*goverseerr.MovieDetails, len(requests))
	tvDetails := make([]*goverseerr.TVDetails, len(requests))
	var wg sync.WaitGroup
	for idx, request := range requests {
		wg.Add(1)
		go func(i int, r *goverseerr.MediaRequest) {
			switch r.Media.MediaType {
			case goverseerr.MediaTypeMovie:
				details, err := r.GetMovieDetails(instance)
				if err != nil {
					newui.Fatal("could not get movie details for request", err)
				}
				movieDetails[i] = details
			case goverseerr.MediaTypeTV:
				details, err := r.GetTVDetails(instance)
				if err != nil {
					newui.Fatal("could not get tv show details for request", err)
				}
				tvDetails[i] = details
			default:
				newui.Fatal("could not get request details for request with unknown media type", errors.New(string(r.Media.MediaType)+" is not a valid media type"))
			}
			wg.Done()
		}(idx, request)
	}
	wg.Wait()
	return tvDetails, movieDetails
}

func getRequests(page, pgSize int) ([]*goverseerr.MediaRequest, int) {
	results, pageInfo, err := instance.GetRequests(page, pgSize, goverseerr.RequestFileterAll, goverseerr.RequestSortAdded)
	if err != nil {
		newui.Fatal("Could not get requests from Overseerr", err)
	}
	return results, pageInfo.Pages
}

func getAllRequests() ([]*goverseerr.MediaRequest, int) {
	var allRequests []*goverseerr.MediaRequest
	pg := 0
	for {
		requests, pages := getRequests(pg, reqPageSize)
		allRequests = append(allRequests, requests...)
		pg++
		if pg >= pages {
			return allRequests, 1
		}
	}
}

var retryAllRequestsCmd = &cobra.Command{
	Use:   "retry-all",
	Short: "Resend all non-available requests to sonarr/radarr",
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		instance = getOverseerrInstance(overseerrProfileName)
	},
	Run: func(cmd *cobra.Command, args []string) {
		newui.DestructiveConfirmation()
		requests, _ := getAllRequests()
		var wg sync.WaitGroup
		ui.StartLoadingSpinner()
		for idx, req := range requests {
			wg.Add(1)
			go func(c int, req *goverseerr.MediaRequest) {
				if req.Media.Status == goverseerr.MediaStatsAvailable {
					wg.Done()
					return
				}
				_, err := instance.RetryRequest(req.ID)
				if err != nil {
					newui.Fatal("Could not retry request: "+fmt.Sprintf("%d", req.ID), err)
				}
				wg.Done()
			}(idx, req)
		}
		wg.Wait()
		ui.StopLoadingSpinner()
		newui.Success("Non-available requests have been retried")
	},
}

func init() {
	requestsAdminManageCmd.Flags().IntVarP(&reqPageNumber, "page", "p", 1, "The page of requests that is used")
	requestsAdminManageCmd.Flags().IntVarP(&reqPageSize, "page-size", "s", 25, "The number of requests per page")
	requestsAdminManageCmd.Flags().BoolVarP(&reqAll, "all", "a", false, "Ignore pages, get ALL requests")
	requestsAdminCmd.AddCommand(requestsAdminManageCmd)
	requestsAdminCmd.AddCommand(retryAllRequestsCmd)
	adminCmd.AddCommand(requestsAdminCmd)
}
