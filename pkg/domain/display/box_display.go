package display

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/svoosh/q/pkg/domain/config"
	"github.com/svoosh/q/pkg/logic"
)

// BoxChars defines the characters used for drawing boxes
type BoxChars struct {
	TopLeft     string
	TopRight    string
	BottomLeft  string
	BottomRight string
	Horizontal  string
	Vertical    string
	TDown       string
	TUp         string
	TRight      string
	TLeft       string
	Cross       string
}

// DefaultBoxChars returns the default box drawing characters
func DefaultBoxChars() BoxChars {
	return BoxChars{
		TopLeft:     "┌",
		TopRight:    "┐",
		BottomLeft:  "└",
		BottomRight: "┘",
		Horizontal:  "─",
		Vertical:    "│",
		TDown:       "┬",
		TUp:         "┴",
		TRight:      "├",
		TLeft:       "┤",
		Cross:       "┼",
	}
}

// Column represents a column in the display
type Column struct {
	Name      string
	Width     int
	Alignment string // "left" or "right"
}

// Config represents the display configuration
type Config struct {
	Columns  []Column
	BoxChars BoxChars
	LongFormat bool
	HumanReadable bool
	All bool
}

// DefaultConfig returns the default display configuration
var DefaultConfig = Config{
	Columns: []Column{
		{Name: "#", Width: 4, Alignment: "right"},
		{Name: "Name", Width: 30, Alignment: "left"},
		{Name: "Type", Width: 10, Alignment: "left"},
		{Name: "Size", Width: 12, Alignment: "right"},
	},
	BoxChars: DefaultBoxChars(),
	LongFormat: false,
	HumanReadable: false,
	All: false,
}

// LongFormatConfig returns the configuration for long format display
var LongFormatConfig = Config{
	Columns: []Column{
		{Name: "Permissions", Width: 10, Alignment: "left"},
		{Name: "Size", Width: 12, Alignment: "right"},
		{Name: "Modified", Width: 20, Alignment: "left"},
		{Name: "Name", Width: 30, Alignment: "left"},
	},
	BoxChars: DefaultBoxChars(),
	LongFormat: true,
	HumanReadable: false,
	All: false,
}

// BoxDisplay handles the display formatting
type BoxDisplay struct {
	config Config
	// Color functions
	dirColor    *color.Color
	fileColor   *color.Color
	imageColor  *color.Color
	videoColor  *color.Color
	boxColor    *color.Color
	emptyColor  *color.Color
	titleColor  *color.Color
	customColors map[string]*color.Color
}

// hexToColor converts a hex color string to a color.Color
func hexToColor(hex string) *color.Color {
	// Remove # if present
	if strings.HasPrefix(hex, "#") {
		hex = hex[1:]
	}

	// Parse hex values
	var r, g, b uint8
	fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)

	// Map RGB values to closest color attributes
	// This is a simple mapping - you might want to implement a more sophisticated color matching algorithm
	var attrs []color.Attribute

	// Add foreground color
	if r > 128 && g > 128 && b > 128 {
		attrs = append(attrs, color.FgWhite)
	} else if r > 128 && g < 128 && b < 128 {
		attrs = append(attrs, color.FgRed)
	} else if r < 128 && g > 128 && b < 128 {
		attrs = append(attrs, color.FgGreen)
	} else if r < 128 && g < 128 && b > 128 {
		attrs = append(attrs, color.FgBlue)
	} else if r > 128 && g > 128 && b < 128 {
		attrs = append(attrs, color.FgYellow)
	} else if r > 128 && g < 128 && b > 128 {
		attrs = append(attrs, color.FgMagenta)
	} else if r < 128 && g > 128 && b > 128 {
		attrs = append(attrs, color.FgCyan)
	} else {
		attrs = append(attrs, color.FgWhite)
	}

	// Add bold if the color is bright
	if r > 192 || g > 192 || b > 192 {
		attrs = append(attrs, color.Bold)
	}

	return color.New(attrs...)
}

// NewBoxDisplay creates a new BoxDisplay instance
func NewBoxDisplay(displayConfig Config, userConfig *config.Config) *BoxDisplay {
	// Convert hex colors to color.Color
	boxColor := hexToColor(userConfig.Colors.Box)
	titleColor := hexToColor(userConfig.Colors.Title)
	fileColor := hexToColor(userConfig.Colors.Default)
	emptyColor := hexToColor(userConfig.Colors.Empty)
	dirColor := hexToColor(userConfig.Colors.Types.Directory)
	imageColor := hexToColor(userConfig.Colors.Types.Image)
	videoColor := hexToColor(userConfig.Colors.Types.Video)

	// Convert custom colors
	customColors := make(map[string]*color.Color)
	for fileType, hexColor := range userConfig.Colors.Types.Custom {
		customColors[string(fileType)] = hexToColor(hexColor)
	}

	return &BoxDisplay{
		config: displayConfig,
		dirColor: dirColor,
		fileColor: fileColor,
		imageColor: imageColor,
		videoColor: videoColor,
		boxColor: boxColor,
		emptyColor: emptyColor,
		titleColor: titleColor,
		customColors: customColors,
	}
}

// formatSize formats a size in bytes to a human-readable string
func formatSize(size int64, humanReadable bool) string {
	if !humanReadable {
		return fmt.Sprintf("%d", size)
	}

	units := []string{"B", "KB", "MB", "GB"}
	value := float64(size)
	unitIndex := 0

	for value >= 1024 && unitIndex < len(units)-1 {
		value /= 1024
		unitIndex++
	}

	return fmt.Sprintf("%.1f %s", value, units[unitIndex])
}

// padString pads a string to fit within the specified width
func padString(str string, width int, alignment string) string {
	if len(str) >= width {
		return str[:width]
	}

	padding := width - len(str)
	switch alignment {
	case "right":
		return strings.Repeat(" ", padding) + str
	default: // left
		return str + strings.Repeat(" ", padding)
	}
}

// createBorder creates the top or bottom border of the box
func (b *BoxDisplay) createBorder(isTop bool) string {
	var parts []string
	if isTop {
		parts = append(parts, b.config.BoxChars.TopLeft)
	} else {
		parts = append(parts, b.config.BoxChars.BottomLeft)
	}

	for i, col := range b.config.Columns {
		parts = append(parts, strings.Repeat(b.config.BoxChars.Horizontal, col.Width))
		if i < len(b.config.Columns)-1 {
			if isTop {
				parts = append(parts, b.config.BoxChars.TDown)
			} else {
				parts = append(parts, b.config.BoxChars.TUp)
			}
		}
	}

	if isTop {
		parts = append(parts, b.config.BoxChars.TopRight)
	} else {
		parts = append(parts, b.config.BoxChars.BottomRight)
	}
	return b.boxColor.Sprint(strings.Join(parts, ""))
}

// getColorForType returns the appropriate color for a file type
func (b *BoxDisplay) getColorForType(entryType logic.FileType) *color.Color {
	// Check custom colors first
	if customColor, ok := b.customColors[string(entryType)]; ok {
		return customColor
	}

	// Use predefined colors
	switch entryType {
	case logic.TypeDirectory:
		return b.dirColor
	case logic.TypeImage:
		return b.imageColor
	case logic.TypeVideo:
		return b.videoColor
	default:
		return b.fileColor
	}
}

// createRow creates a row of data
func (b *BoxDisplay) createRow(index int, entry logic.PathEntry) string {
	var parts []string
	parts = append(parts, b.boxColor.Sprint(b.config.BoxChars.Vertical))

	if b.config.LongFormat {
		// Permissions (simplified for now)
		perms := "-rw-r--r--"
		if entry.EntryType == logic.TypeDirectory {
			perms = "drwxr-xr-x"
		}
		parts = append(parts, padString(perms, b.config.Columns[0].Width, b.config.Columns[0].Alignment))
		parts = append(parts, b.boxColor.Sprint(b.config.BoxChars.Vertical))

		// Size
		sizeStr := formatSize(entry.Size, b.config.HumanReadable)
		parts = append(parts, padString(sizeStr, b.config.Columns[1].Width, b.config.Columns[1].Alignment))
		parts = append(parts, b.boxColor.Sprint(b.config.BoxChars.Vertical))

		// Modified time (simplified for now)
		modified := time.Now().Format("Jan 02 15:04")
		parts = append(parts, padString(modified, b.config.Columns[2].Width, b.config.Columns[2].Alignment))
		parts = append(parts, b.boxColor.Sprint(b.config.BoxChars.Vertical))

		// Name
		name := entry.Name
		colorFunc := b.getColorForType(entry.EntryType)
		parts = append(parts, colorFunc.Sprint(padString(name, b.config.Columns[3].Width, b.config.Columns[3].Alignment)))
	} else {
		// Index
		indexStr := fmt.Sprintf("%d", index+1)
		parts = append(parts, padString(indexStr, b.config.Columns[0].Width, b.config.Columns[0].Alignment))
		parts = append(parts, b.boxColor.Sprint(b.config.BoxChars.Vertical))

		// Name
		name := entry.Name
		colorFunc := b.getColorForType(entry.EntryType)
		parts = append(parts, colorFunc.Sprint(padString(name, b.config.Columns[1].Width, b.config.Columns[1].Alignment)))
		parts = append(parts, b.boxColor.Sprint(b.config.BoxChars.Vertical))

		// Type
		parts = append(parts, colorFunc.Sprint(padString(string(entry.EntryType), b.config.Columns[2].Width, b.config.Columns[2].Alignment)))
		parts = append(parts, b.boxColor.Sprint(b.config.BoxChars.Vertical))

		// Size
		sizeStr := formatSize(entry.Size, b.config.HumanReadable)
		parts = append(parts, padString(sizeStr, b.config.Columns[3].Width, b.config.Columns[3].Alignment))
	}

	parts = append(parts, b.boxColor.Sprint(b.config.BoxChars.Vertical))
	return strings.Join(parts, "")
}

// DisplayEntries displays the entries in a box format
func (b *BoxDisplay) DisplayEntries(entries []logic.PathEntry) error {
	var lines []string

	// If no entries, display "Empty" message
	if len(entries) == 0 {
		// Create a simple box with "Empty" message
		width := 20
		emptyBox := []string{
			b.boxColor.Sprint(b.config.BoxChars.TopLeft + strings.Repeat(b.config.BoxChars.Horizontal, width-2) + b.config.BoxChars.TopRight),
			b.boxColor.Sprint(b.config.BoxChars.Vertical) + b.emptyColor.Sprint(padString("Empty", width-2, "center")) + b.boxColor.Sprint(b.config.BoxChars.Vertical),
			b.boxColor.Sprint(b.config.BoxChars.BottomLeft + strings.Repeat(b.config.BoxChars.Horizontal, width-2) + b.config.BoxChars.BottomRight),
		}
		fmt.Println(strings.Join(emptyBox, "\n"))
		return nil
	}

	// Top border
	lines = append(lines, b.createBorder(true))

	// Header
	var headerParts []string
	headerParts = append(headerParts, b.boxColor.Sprint(b.config.BoxChars.Vertical))
	for _, col := range b.config.Columns {
		headerParts = append(headerParts, b.boxColor.Sprint(padString(col.Name, col.Width, col.Alignment)))
		headerParts = append(headerParts, b.boxColor.Sprint(b.config.BoxChars.Vertical))
	}
	lines = append(lines, strings.Join(headerParts, ""))

	// Separator
	var separatorParts []string
	separatorParts = append(separatorParts, b.boxColor.Sprint(b.config.BoxChars.TRight))
	for i, col := range b.config.Columns {
		separatorParts = append(separatorParts, b.boxColor.Sprint(strings.Repeat(b.config.BoxChars.Horizontal, col.Width)))
		if i < len(b.config.Columns)-1 {
			separatorParts = append(separatorParts, b.boxColor.Sprint(b.config.BoxChars.Cross))
		}
	}
	separatorParts = append(separatorParts, b.boxColor.Sprint(b.config.BoxChars.TLeft))
	lines = append(lines, strings.Join(separatorParts, ""))

	// Data rows
	for i, entry := range entries {
		lines = append(lines, b.createRow(i, entry))
	}

	// Bottom border
	lines = append(lines, b.createBorder(false))

	// Print all lines
	fmt.Println(strings.Join(lines, "\n"))
	return nil
} 