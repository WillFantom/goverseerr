package goverseerr

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type LogResponse struct {
	PageInfo Page          `json:"pageInfo"`
	Entries  []*LogMessage `json:"results"`
}

type LogMessage struct {
	Label     string    `json:"label"`
	Level     LogLevel  `json:"level"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type LogLevel string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
)

func StringToLogLevel(level string) (LogLevel, error) {
	switch strings.ToLower(level) {
	case string(LogLevelDebug):
		return LogLevelDebug, nil
	case string(LogLevelInfo):
		return LogLevelInfo, nil
	case string(LogLevelWarn):
		return LogLevelWarn, nil
	case string(LogLevelError):
		return LogLevelError, nil
	default:
		return "", errors.New("level given is not a valid log level")
	}
}

func (o *Overseerr) GetLogs(take, skip int, filter LogLevel) ([]*LogMessage, error) {
	var logs LogResponse
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
	return logs.Entries, nil
}
