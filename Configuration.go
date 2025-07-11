package conf

import "time"

type Configuration interface {
	GetStringWithFile(id string, defaultValue string) string
	GetString(id string, defaultValue string) string
	GetStrings(id string) []string
	GetStringRequired(id string) string
	GetInt(id string, defaultValue int) int
	GetDuration(id string, defaultValue int, unit time.Duration) time.Duration

	// Deprecated: use GetDuration instead (with a directly specified unit)
	GetDurationS(id string, defaultValueS int) time.Duration
}
