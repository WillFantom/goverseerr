package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/willfantom/goverseerr"
	"github.com/willfantom/goverseerr/cmd/overclirr/ui"
)

var (
	logFetchLevel string
	logsToGet     int
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Get logs from an Overseerr instance",
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		instance = getOverseerrInstance(overseerrProfileName)
	},
	Run: func(cmd *cobra.Command, args []string) {
		level, err := goverseerr.StringToLogLevel(logFetchLevel)
		if err != nil {
			ui.PrettyFatal("An invalid log level was given")
			ui.PrettyInfo(fmt.Sprintf("Try: %s, %s, %s, %s",
				string(goverseerr.LogLevelDebug), string(goverseerr.LogLevelInfo),
				string(goverseerr.LogLevelWarn), string(goverseerr.LogLevelError)))
			logrus.WithField("givenLevel", logFetchLevel).Fatalln("invalid log level given")
		}
		logs, err := instance.GetLogs(logsToGet, 0, level)
		if err != nil {
			ui.PrettyFatal("Could not get logs from Overseerr")
			logrus.WithField("extended", err.Error()).Fatalln("logs could not be retrived")
		}
		for _, log := range logs {
			fmt.Printf("%+v\n", log)
		}
	},
}

func init() {
	logsCmd.Flags().StringVar(&logFetchLevel, "level", "info", "Log level to fetch at")
	logsCmd.Flags().IntVarP(&logsToGet, "number", "n", 5, "Fetch the most recent N log entries")
	adminCmd.AddCommand(logsCmd)
}
