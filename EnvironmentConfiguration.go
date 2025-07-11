package conf

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type EnvironmentConfiguration struct {
}

func NewEnvironmentConfiguration() *EnvironmentConfiguration {
	return &EnvironmentConfiguration{}
}

// GetStringWithFile
// try get data from file located by `<id>_file` environment variable.
// When `<id>_file` is NOT found or empty, returns defaultValue.
// if `<id>_file` environment variable is empty, try get value directly from `<id>` environment variable.
// When variable not found or empty - returns defaultValue
func (e EnvironmentConfiguration) GetStringWithFile(id string, defaultValue string) string {
	valFile := os.Getenv(id + "_file")
	if valFile != "" {
		content, err := os.ReadFile(valFile)
		if err != nil || len(content) == 0 {
			return defaultValue
		}

		return string(content)
	}

	val := os.Getenv(id)
	if val == "" {
		return defaultValue
	}
	return val
}

// GetString
// Try get string directly from `<id>` environment variable.
// When variable not found or is empty - returns defaultValue
func (e EnvironmentConfiguration) GetString(id string, defaultValue string) string {
	val := os.Getenv(id)
	if val == "" {
		return defaultValue
	}
	return val
}

// GetStrings
// get array (slice) of strings, from keys `<id>__<IDX>` starting IDX from 0
// first missing (empty or unexisting) value is considered as end of array.
// Max length of array is 1024 elements
func (e EnvironmentConfiguration) GetStrings(id string) []string {
	ret := make([]string, 0, 8)

	idx := 0
	for idx < 1024 {
		val := os.Getenv(fmt.Sprintf("%s__%d", id, idx))
		if val == "" {
			return ret
		}
		ret = append(ret, val)
		idx += 1
	}

	return ret
}

// GetInt
// Try get int directly from `<id>` environment variable.
// When variable not found or is empty - returns defaultValue
// When value can't be parsed into int, also returns defaultValue
func (e EnvironmentConfiguration) GetInt(id string, defaultValue int) int {
	val := os.Getenv(id)
	if val == "" {
		return defaultValue
	}

	v, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}
	return v
}

func (d EnvironmentConfiguration) GetDuration(id string, defaultValue int, unit time.Duration) time.Duration {
	val := os.Getenv(id)
	if val == "" {
		return time.Duration(defaultValue) * unit
	}

	v, err := strconv.Atoi(val)
	if err != nil {
		return time.Duration(defaultValue) * unit
	}
	return time.Duration(v) * unit
}

func (e EnvironmentConfiguration) GetDurationS(id string, defaultValueS int) time.Duration {
	val := os.Getenv(id)
	if val == "" {
		return time.Duration(defaultValueS) * time.Second
	}

	v, err := strconv.Atoi(val)
	if err != nil {
		return time.Duration(defaultValueS) * time.Second
	}
	return time.Second * time.Duration(v)
}

func (e EnvironmentConfiguration) GetStringRequired(id string) string {
	val := os.Getenv(id)
	if val == "" {
		panic(fmt.Sprintf("error: environment variable '%s' is required, but NOT set", id))
	}
	return val
}

var _ Configuration = EnvironmentConfiguration{}
