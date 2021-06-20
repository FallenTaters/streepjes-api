package main

import (
	"os"
	"strings"
)

var settings = struct {
	Port          string
	DisableSecure bool
}{
	Port:          "81",
	DisableSecure: false,
}

func readSettings() {
	settings.Port = readString(`PORT`, settings.Port)
	settings.DisableSecure = readBool(`DISABLE_SECURE`, settings.DisableSecure)
}

func readString(name, defaultValue string) string {
	v := os.Getenv(name)
	if v == `` {
		return defaultValue
	}
	return v
}

func readBool(name string, defaultValue bool) bool {
	v := os.Getenv(name)
	v = strings.ToLower(v)
	switch v {
	case `true`:
		return true
	case `false`:
		return false
	}
	return defaultValue
}
