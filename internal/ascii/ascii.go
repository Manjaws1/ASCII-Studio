package ascii

import (
	"fmt"
	"os"
	"strings"
)

const charHeight = 8

// GetBanner loads a banner from the local file system.
func GetBanner(name string) ([]string, error) {
	validBanners := map[string]bool{
		"standard":   true,
		"shadow":     true,
		"thinkertoy": true,
	}

	if !validBanners[name] {
		return nil, fmt.Errorf("invalid banner style: %s", name)
	}

	// Assuming the server is run from the root of the project
	filePath := "banners/" + name + ".txt"
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read banner file: %v", err)
	}

	rawText := strings.ReplaceAll(string(data), "\r\n", "\n")
	return strings.Split(rawText, "\n"), nil
}

// Convert turns text into ASCII art using the specified banner.
func Convert(input string, banner []string) string {
	var result strings.Builder

	input = strings.ReplaceAll(input, "\\n", "\n")
	input = strings.ReplaceAll(input, "\r\n", "\n")

	lines := strings.Split(input, "\n")

	for li, text := range lines {
		if text == "" {
			if li < len(lines)-1 {
				result.WriteString("\n")
			}
			continue
		}

		for row := 0; row < charHeight; row++ {
			for _, r := range text {
				if r < 32 || r > 126 {
					continue
				}

				index := (int(r)-32)*9 + 1 + row

				if index < len(banner) {
					result.WriteString(banner[index])
				}
			}
			result.WriteString("\n")
		}
	}

	return result.String()
}
