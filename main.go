package main

import (
	"embed"
	"fmt"

	"github.com/Alwin18/algo/cmd"
)

// Embed seluruh file template di dalam folder `cmd/templates`
//
//go:embed cmd/*
var templates embed.FS

func main() {
	// Debug: Print embedded files
	entries, _ := templates.ReadDir("templates")
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
	cmd.Execute()
}
