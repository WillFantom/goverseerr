package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/willfantom/goverseerr"
	"github.com/willfantom/goverseerr/cmd/overclirr/ui"
)

var jobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "View, run or cancel Overseerr jobs",
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		instance = getOverseerrInstance(overseerrProfileName)
	},
	Run: func(cmd *cobra.Command, args []string) {
		allJobs := getAllJobs()
		var tableValues = [][]string{
			{"ID", "Name", "Type", "Running"},
		}
		for _, job := range allJobs {
			tableValues = append(tableValues, []string{job.ID, job.Name, string(job.Type), fmt.Sprintf("%v", job.Running)})
		}
		ui.PrettyTable(tableValues)
		ui.PrettyInfo("To manage these jobs, see " + cmd.Name() + " --help")
	},
}

var runJobCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an Overseerr job",
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		instance = getOverseerrInstance(overseerrProfileName)
	},
	Run: func(cmd *cobra.Command, args []string) {
		runJob(selectJob(getAllJobs()))
	},
}

var cancelJobCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel an Overseerr job",
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		instance = getOverseerrInstance(overseerrProfileName)
	},
	Run: func(cmd *cobra.Command, args []string) {
		cancelJob(selectJob(getAllJobs()))
	},
}

func getAllJobs() []*goverseerr.Job {
	jobs, err := instance.GetJobs()
	if err != nil {
		ui.PrettyFatal("Could not get jobs from Overseerr")
		ui.PrettyInfo("Do you have the right permssions?")
		logrus.WithField("extended", err.Error()).Fatalln("could not get overseerr job list")
	}
	return jobs
}

func selectJob(jobs []*goverseerr.Job) *goverseerr.Job {
	jobOptions := make([]string, len(jobs))
	for idx, opt := range jobs {
		jobOptions[idx] = fmt.Sprintf("%s (%v)", opt.Name, opt.Running)
	}
	idx, _, err := ui.RunSelector("Select a job", jobOptions)
	if err != nil {
		logrus.Fatalln("job selector failed")
	}
	return jobs[idx]
}

func runJob(job *goverseerr.Job) {
	if job.Running {
		ui.PrettyOops("That job is already running!")
		logrus.WithField("job", job.ID).Fatalln("attempted to run an already running job")
	}
	_, err := instance.RunJob(job.ID)
	if err != nil {
		ui.PrettyFatal("Could not start the job")
		logrus.WithField("extended", err.Error()).Fatalln("could not start the job")
	}
	ui.PrettySuccess("Running Job: " + job.Name)
}

func cancelJob(job *goverseerr.Job) {
	if !job.Running {
		ui.PrettyOops("That job is not running!")
		logrus.WithField("job", job.ID).Fatalln("attempted to cancel an already stopped job")
	}
	_, err := instance.CancelJob(job.ID)
	if err != nil {
		ui.PrettyFatal("Could not cancel the job")
		logrus.WithField("extended", err.Error()).Fatalln("could not cancel the job")
	}
	ui.PrettySuccess("Canceld Job: " + job.Name)
}

func init() {
	adminCmd.AddCommand(jobsCmd)
	jobsCmd.AddCommand(runJobCmd)
	jobsCmd.AddCommand(cancelJobCmd)
}
