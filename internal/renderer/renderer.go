package renderer

import (
	"ascii-studio/internal/alignment"
	"ascii-studio/internal/ascii"
	"ascii-studio/internal/color"
	"fmt"
)

// Process takes all parameters and orchestrates the generation of the final output.
func Process(text, bannerName, colorCode, align string) (string, error) {
	// Get Banner
	banner, err := ascii.GetBanner(bannerName)
	if err != nil {
		return "", fmt.Errorf("error loading banner: %v", err)
	}

	// Convert text to ASCII
	rawAscii := ascii.Convert(text, banner)

	// Apply Color
	coloredAscii := color.Apply(rawAscii, colorCode)

	// Apply Alignment
	finalAscii := alignment.Apply(coloredAscii, align)

	return finalAscii, nil
}
