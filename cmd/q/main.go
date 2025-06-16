package main

import (
	"log"
	"os"

	"github.com/MagnusVestvik/q/pkg/domain/config"
	"github.com/MagnusVestvik/q/pkg/domain/display"
	"github.com/MagnusVestvik/q/pkg/domain/flags"
	"github.com/MagnusVestvik/q/pkg/logic"
)

var Version = "dev"

func main() {
	userConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	displayFlags, targetPath, err := flags.ParseFlags(os.Args[1:])
	if err != nil {
		log.Fatalf("Error parsing flags: %v", err)
	}

	displayConfig := display.Config{
		Columns:       display.DefaultConfig.Columns,
		BoxChars:      display.DefaultBoxChars(),
		LongFormat:    displayFlags.LongFormat,
		HumanReadable: displayFlags.HumanReadable,
		All:           displayFlags.All,
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)
	pathHandler := logic.NewPathHandler(logger)

	pathEntries, err := pathHandler.GetEntries(targetPath, displayFlags.All)
	if err != nil {
		log.Fatalf("Error getting entries: %v", err)
	}

	boxDisplay := display.NewBoxDisplay(displayConfig, userConfig)

	if err := boxDisplay.DisplayEntries(pathEntries); err != nil {
		log.Fatalf("Error displaying entries: %v", err)
	}
}
