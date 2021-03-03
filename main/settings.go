package main

import "os"

var Settings = struct {
	Port string
}{
	Port: "8080",
}

func readSettings() {
	port := os.Getenv(`PORT`)
	if port != `` {
		Settings.Port = port
	}
}
