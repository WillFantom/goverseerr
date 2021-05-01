package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/willfantom/goverseerr"
	"github.com/willfantom/goverseerr/cmd/overclirr/ui"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Manage Overseerr caches",
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		instance = getOverseerrInstance(overseerrProfileName)
	},
	Run: func(cmd *cobra.Command, args []string) {
		caches := getAllCaches()
		var tableValues = [][]string{
			{"ID", "Name", "Hits", "Misses", "Keys", "Key Size", "Value Size"},
		}
		for _, cache := range caches {
			tableValues = append(tableValues, []string{cache.ID, cache.Name, fmt.Sprintf("%d", cache.Stats.Hits),
				fmt.Sprintf("%d", cache.Stats.Misses), fmt.Sprintf("%d", cache.Stats.Keys), fmt.Sprintf("%d", cache.Stats.KSize),
				fmt.Sprintf("%d", cache.Stats.VSize)})
		}
		ui.PrettyTable(tableValues)
		ui.PrettyInfo("To manage these caches, see " + cmd.Name() + " --help")
	},
}

var flushCacheCmd = &cobra.Command{
	Use:   "flush",
	Short: "Flush an Overseerr cache",
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		instance = getOverseerrInstance(overseerrProfileName)
	},
	Run: func(cmd *cobra.Command, args []string) {
		flushCache(selectCache(getAllCaches()))
	},
}

func getAllCaches() []*goverseerr.Cache {
	cacheStats, err := instance.GetCacheStats()
	if err != nil {
		ui.PrettyFatal("Could not get caches from Overseerr")
		ui.PrettyInfo("Do you have the right permssions?")
		logrus.WithField("extended", err.Error()).Fatalln("could not get overseerr cache stat list")
	}
	return cacheStats
}

func selectCache(caches []*goverseerr.Cache) *goverseerr.Cache {
	cacheOptions := make([]string, len(caches))
	for idx, opt := range caches {
		cacheOptions[idx] = fmt.Sprintf("%s (key size: %d | value size: %d)", opt.Name, opt.Stats.KSize, opt.Stats.VSize)
	}
	idx, _, err := ui.RunSelector("Select a cache", cacheOptions)
	if err != nil {
		logrus.Fatalln("cache selector failed")
	}
	return caches[idx]
}

func flushCache(cache *goverseerr.Cache) {
	err := instance.FlushCache(cache.ID)
	if err != nil {
		ui.PrettyFatal("Could not flush the selected cache")
		logrus.WithField("extended", err.Error()).Fatalln("could not flush the cache")
	}
	ui.PrettySuccess("Flushed Cache: " + cache.Name)
}

func init() {
	adminCmd.AddCommand(cacheCmd)
	cacheCmd.AddCommand(flushCacheCmd)
}
