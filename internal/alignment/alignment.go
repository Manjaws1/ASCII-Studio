package alignment

import (
	"strings"
)

// Apply handles alignment by determining the width of the ASCII art
// and padding it based on a standard terminal/container width.

func Apply(asciiArt, align string) string {
	if align == "" || align == "left" {
		return asciiArt
	}

	// For Web UI, the simplest and most robust way to align the ASCII block
	// is to wrap it in a div that applies standard text alignment.
	// The frontend will render this inside a parent container.
	var style string
	switch align {
	case "center":
		style = "text-align: center;"
	case "right":
		style = "text-align: right;"
	case "justify":
		style = "text-align: justify;"
	default:
		return asciiArt
	}

	// We apply a block display so text-align works properly on the text inside.
	return "<div style=\"" + style + "\">\n" + strings.TrimRight(asciiArt, "\n") + "\n</div>"
}
