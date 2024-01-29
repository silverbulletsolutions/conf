package conf

import "time"

type Configuration interface {
	GetStringWithFile(id string, defaultValue string) string
	GetString(id string, defaultValue string) string
	GetStrings(id string) []string
	GetStringRequired(id string) string
	GetInt(id string, defaultValue int) int
	GetDurationS(id string, defaultValueS time.Duration) time.Duration
}
