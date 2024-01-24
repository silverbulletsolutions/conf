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

func (d *DefaultConfiguration) GetInt(id string, defaultValue int) int {
	v, ok := d.requiredConfiguration[id]
	if ok {
		x, _ := strconv.Atoi(v)
		return x
	}
	return defaultValue
}

func (d *DefaultConfiguration) GetDurationS(id string, defaultValueS time.Duration) time.Duration {
	v, ok := d.requiredConfiguration[id]
	if ok {
		t, _ := strconv.Atoi(v)
		return time.Duration(t) * time.Second
	}
	return defaultValueS
}

func (d *DefaultConfiguration) GetStringRequired(id string) string {
	v, ok := d.requiredConfiguration[id]
	if !ok {
		panic(fmt.Sprintf("configuration '%s' is required, but is missing in default configuration constructor", id))
	}

	return v
}
