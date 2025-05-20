package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/thineshsubramani/sheet-cli/internal/cache"
	"github.com/thineshsubramani/sheet-cli/internal/config"
	"github.com/thineshsubramani/sheet-cli/internal/output"
	"github.com/thineshsubramani/sheet-cli/internal/usage"
)

func main() {
	var (
		backendURL   string
		detailed     bool
		refreshCache bool
		refreshAll   bool
	)

	flag.StringVar(&backendURL, "backend", "https://raw.githubusercontent.com/thineshsubramani/cheatsheet/main/", "Backend base URL or local cache path")
	flag.BoolVar(&detailed, "detailed", false, "Show detailed output with descriptions")
	flag.BoolVar(&refreshCache, "refresh", false, "Force refresh cache for specified language")
	flag.BoolVar(&refreshAll, "refresh-all", false, "Refresh cache for all languages")
	flag.Usage = printUsage
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	langInput := strings.ToLower(args[0])
	sectionInputs := []string{"all"}
	if len(args) > 1 {
		sectionInputs = config.ParseSectionInputs(args[1:])
	}

	mainConfig, err := config.LoadMainConfig(backendURL, refreshCache)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading main config: %v\n", err)
		os.Exit(1)
	}

	if refreshAll {
		if err := cache.RefreshAll(mainConfig, backendURL); err != nil {
			fmt.Fprintf(os.Stderr, "Error refreshing all caches: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("All caches refreshed successfully!")
		return
	}

	lang := config.FindLanguage(mainConfig, langInput)
	if lang == nil {
		fmt.Fprintf(os.Stderr, "Language not supported: %s\n", langInput)
		os.Exit(1)
	}

	sections, err := cache.LoadLanguageSections(backendURL, lang.Database, refreshCache)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading language data for %s: %v\n", lang.Name, err)
		os.Exit(1)
	}

	go usage.Track(lang.Name)

	if len(sectionInputs) == 1 && strings.ToLower(sectionInputs[0]) == "help" {
		output.PrintHelpSections(sections, lang.Name)
		return
	}

	output.PrintSections(sections, sectionInputs, lang.Name, detailed)
}

func printUsage() {
	fmt.Println("Usage: sheet [language] [sections...]")
	fmt.Println("Example: sheet go map slice or sheet python help")
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
}
