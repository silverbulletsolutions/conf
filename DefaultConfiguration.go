package conf

import (
	"fmt"
	"strconv"
	"time"
)

type DefaultConfiguration struct {
	requiredConfiguration map[string]string
}

func NewDefaultConfiguration(requiredConfiguration map[string]string) *DefaultConfiguration {
	if requiredConfiguration == nil {
		requiredConfiguration = make(map[string]string)
	}
	return &DefaultConfiguration{
		requiredConfiguration: requiredConfiguration,
	}
}

func (d *DefaultConfiguration) GetStringWithFile(id string, defaultSecret string) string {
	v, ok := d.requiredConfiguration[id]
	if ok {
		return v
	}
	return defaultSecret
}

func (d *DefaultConfiguration) GetString(id string, defaultValue string) string {
	v, ok := d.requiredConfiguration[id]
	if ok {
		return v
	}
	return defaultValue
}

// GetStrings
// get array (slice) of strings, from keys `<id>__<IDX>` starting IDX from 0
// first missing (empty or unexisting) value is considered as end of array.
// Max length of array is 1024 elements
func (d *DefaultConfiguration) GetStrings(id string) []string {
	ret := make([]string, 0, 8)

	idx := 0
	for idx < 1024 {
		v, ok := d.requiredConfiguration[fmt.Sprintf("%s__%d", id, idx)]
		if !ok || v == "" {
			return ret
		}
		ret = append(ret, v)
		idx += 1
	}

	return ret
}

func (d *DefaultConfiguration) GetInt(id string, defaultValue int) int {
	v, ok := d.requiredConfiguration[id]
	if ok {
		x, _ := strconv.Atoi(v)
		return x
	}
	return defaultValue
}

func (d *DefaultConfiguration) GetDuration(id string, defaultValue int, unit time.Duration) time.Duration {
	v, ok := d.requiredConfiguration[id]
	if ok {
		t, _ := strconv.Atoi(v)
		return time.Duration(t) * unit
	}
	return time.Duration(defaultValue) * unit
}

func (d *DefaultConfiguration) GetDurationS(id string, defaultValueS int) time.Duration {
	v, ok := d.requiredConfiguration[id]
	if ok {
		t, _ := strconv.Atoi(v)
		return time.Duration(t) * time.Second
	}
	return time.Duration(defaultValueS) * time.Second
}

func (d *DefaultConfiguration) GetStringRequired(id string) string {
	v, ok := d.requiredConfiguration[id]
	if !ok {
		panic(fmt.Sprintf("configuration '%s' is required, but is missing in default configuration constructor", id))
	}

	return v
}

var _ Configuration = (*DefaultConfiguration)(nil)
