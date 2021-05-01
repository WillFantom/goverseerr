package goverseerr

import (
	"fmt"
	"time"
)

type Job struct {
	ID          int
	Name        string
	Type        JobType
	NextRunTime time.Time
	Running     bool
}

type JobType string

const (
	JobTypeProcess JobType = "process"
	JobTypeCommand JobType = "command"
)

func (o *Overseerr) GetJobs() ([]*Job, error) {
	var jobs []*Job
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&jobs).Get("/settings/jobs")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return jobs, nil
}

func (o *Overseerr) RunJob(jobID int) (*Job, error) {
	var job *Job
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("jobID", fmt.Sprintf("%d", jobID)).
		SetResult(&job).Post("/settings/jobs/{jobID}/run")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return job, nil
}

func (o *Overseerr) CancelJob(jobID int) (*Job, error) {
	var job *Job
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetPathParam("jobID", fmt.Sprintf("%d", jobID)).
		SetResult(&job).Post("/settings/jobs/{jobID}/cancel")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return job, nil
}
