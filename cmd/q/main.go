package main

import (
	"log"
	"os"

	"github.com/svoosh/q/pkg/domain/config"
	"github.com/svoosh/q/pkg/domain/display"
	"github.com/svoosh/q/pkg/domain/flags"
	"github.com/svoosh/q/pkg/logic"
)

func main() {
	// Load user configuration
	userConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Parse command line flags
	displayFlags, targetPath, err := flags.ParseFlags(os.Args[1:])
	if err != nil {
		log.Fatalf("Error parsing flags: %v", err)
	}

	// Create display configuration
	displayConfig := display.Config{
		Columns: display.DefaultConfig.Columns,
		BoxChars: display.DefaultBoxChars(),
		LongFormat: displayFlags.LongFormat,
		HumanReadable: displayFlags.HumanReadable,
		All: displayFlags.All,
	}

	// Create path handler
	logger := log.New(os.Stdout, "", log.LstdFlags)
	pathHandler := logic.NewPathHandler(logger)

	// Get entries for the target path
	entries, err := pathHandler.GetEntries(targetPath, displayFlags.All)
	if err != nil {
		log.Fatalf("Error getting entries: %v", err)
	}

	// Create box display with user configuration
	boxDisplay := display.NewBoxDisplay(displayConfig, userConfig)

	// Display entries
	if err := boxDisplay.DisplayEntries(entries); err != nil {
		log.Fatalf("Error displaying entries: %v", err)
	}
} 