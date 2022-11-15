package models

import "time"

type Config struct {
	Timeout                    time.Duration
	ApiKey                     string
	GrafanaHost                string
	GrafanaRulesDirectory      string
	AlertStatusRefreshStrategy AlertStatusRefreshStrategy
	AlertStatusTTL             time.Duration
}

type AlertStatusRefreshStrategy int32

const (
	RealTime AlertStatusRefreshStrategy = 0 //Not supported as of now
	TTL      AlertStatusRefreshStrategy = 1
)
