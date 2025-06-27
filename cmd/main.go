package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Cyb3r3x3r/header-inspector/internal/inspector"
	"github.com/Cyb3r3x3r/header-inspector/internal/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: header-inspector <URL>")
		return
	}

	rawURL := os.Args[1]
	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		rawURL = "https://" + rawURL
	}

	if !utils.IsValidURL(rawURL) {
		fmt.Printf("Invalid URL: %s", rawURL)
		return
	}

	headers, err := inspector.FetchHeaders(rawURL)
	if err != nil {
		fmt.Printf("Error fetching headers: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Security Headers: for %s\n", rawURL)
	for name, status := range headers {
		fmt.Printf(" %-30s %s\n", name+":", status)
	}
}
