package config

import (
	"time"
)

type View interface {
	GetString(string) string
	GetInt(string) int
	GetInt64(string) int64
	GetFloat64(string) float64
	GetStringSlice(string) []string
	GetBool(string) bool
	GetDuration(string) time.Duration
}
