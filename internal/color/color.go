package color

import (
	"fmt"
	"strings"
)

// Apply wraps the given ASCII art in an HTML span with the specified color.
func Apply(asciiArt, colorCode string) string {
	if colorCode == "" {
		return asciiArt
	}
	
	// Ensure that special characters like < and > are escaped if we were processing HTML,
	// but since this will be rendered inside <pre> with spans, it's safer to just wrap it.
	// For simplicity, we wrap the entire ascii string in a span with the given color.
	// We use inline CSS to apply the color.
	
	// To prevent HTML injection, we could sanitize the color code.
	safeColor := strings.ReplaceAll(colorCode, "\"", "")
	safeColor = strings.ReplaceAll(safeColor, ";", "")
	safeColor = strings.ReplaceAll(safeColor, "<", "")
	safeColor = strings.ReplaceAll(safeColor, ">", "")

	return fmt.Sprintf(`<span style="color: %s;">%s</span>`, safeColor, asciiArt)
}
