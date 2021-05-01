package goverseerr

import "fmt"

type LogMessage struct {
	Label     string   `json:"label"`
	Level     LogLevel `json:"level"`
	Message   string   `json:"message"`
	Timestamp string   `json:"timestamp"`
}

type LogLevel string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
)

func (o *Overseerr) GetLogs(take, skip int, filter LogLevel) ([]*LogMessage, error) {
	var logs []*LogMessage
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").SetQueryParams(map[string]string{
		"take":   fmt.Sprintf("%d", take),
		"skip":   fmt.Sprintf("%d", skip),
		"filter": string(filter),
	}).
		SetResult(&logs).Get("/settings/logs")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return logs, nil
}
