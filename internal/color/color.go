package color

import (
	"fmt"
	"strings"
)

// Apply wraps the given ASCII art in an HTML span with the specified color.
func Apply(asciiArt, colorCode string) string {
	if colorCode == "" {
		return asciiArt
	

	// To prevent HTML injection, we could sanitize the color code.
	safeColor := strings.ReplaceAll(colorCode, "\"", "")
	safeColor = strings.ReplaceAll(safeColor, ";", "")
	safeColor = strings.ReplaceAll(safeColor, "<", "")
	safeColor = strings.ReplaceAll(safeColor, ">", "")

	return fmt.Sprintf(`<span style="color: %s;">%s</span>`, safeColor, asciiArt)
}
