package renderer

import (
	"ascii-studio/internal/alignment"
	"ascii-studio/internal/ascii"
	"ascii-studio/internal/color"
	"fmt"
)

// Process takes all parameters and orchestrates the generation of the final output.
func Process(text, bannerName, colorCode, align string) (string, error) {
	// 1. Get Banner
	banner, err := ascii.GetBanner(bannerName)
	if err != nil {
		return "", fmt.Errorf("error loading banner: %v", err)
	}

	// 2. Convert text to ASCII
	rawAscii := ascii.Convert(text, banner)

	// 3. Apply Color
	coloredAscii := color.Apply(rawAscii, colorCode)

	// 4. Apply Alignment
	finalAscii := alignment.Apply(coloredAscii, align)

	return finalAscii, nil
}
