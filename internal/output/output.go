package output

import (
	"fmt"
	"strings"

	"github.com/thineshsubramani/sheet-cli/internal/types"
)

const (
	colorReset = "\033[0m"
	colorTitle = "\033[1;36m"
	colorDesc  = "\033[0;33m"
	colorCode  = "\033[0;32m"
)

// PrintHelpSections displays available sections for a language.
func PrintHelpSections(sections []types.Section, lang string) {
	fmt.Printf("Available sections for '%s':\n", lang)
	for _, sec := range sections {
		if strings.ToLower(sec.Name) == "all" {
			continue
		}
		if len(sec.Alias) > 0 {
			fmt.Printf(" - %s | alias = %s\n", sec.Name, strings.Join(sec.Alias, ", "))
		} else {
			fmt.Printf(" - %s\n", sec.Name)
		}
	}
}

// PrintSections displays the requested sections.
func PrintSections(sections []types.Section, sectionInputs []string, lang string, detailed bool) {
	for _, secName := range sectionInputs {
		sec := findSection(sections, secName)
		if sec == nil {
			fmt.Printf("Section '%s' not found for language '%s'\n", secName, lang)
			continue
		}

		fmt.Printf("%s---- %s ----%s\n", colorTitle, strings.Title(sec.Name), colorReset)
		if detailed {
			fmt.Printf("%s%s%s\n\n", colorDesc, sec.Description, colorReset)
		}
		fmt.Printf("%s%s%s\n\n", colorCode, sec.Syntax, colorReset)
	}
}

// findSection finds a section by name or alias.
func findSection(sections []types.Section, name string) *types.Section {
	name = strings.ToLower(name)
	for _, sec := range sections {
		if strings.EqualFold(sec.Name, name) || contains(sec.Alias, name) {
			return &sec
		}
	}
	return nil
}

// contains checks if a string exists in a slice case-insensitively.
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, item) {
			return true
		}
	}
	return false
}
