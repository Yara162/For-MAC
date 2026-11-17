// Package macfmt normalizes MAC addresses to the AA:BB:CC:DD:EE:FF form.
package macfmt

import (
	"regexp"
	"strings"
)

var nonHex = regexp.MustCompile(`[^a-fA-F0-9]`)

const (
	ErrTooFew  = "ERROR: too few characters"
	ErrTooMany = "ERROR: too many characters"
)

// FormatLine strips every non-hex character from line and, if exactly 12 hex
// digits remain, returns them uppercased and colon-separated. Blank input
// yields "", and the wrong number of digits yields an error string.
func FormatLine(line string) string {
	cleaned := strings.ToUpper(nonHex.ReplaceAllString(line, ""))
	switch {
	case cleaned == "":
		return ""
	case len(cleaned) < 12:
		return ErrTooFew
	case len(cleaned) > 12:
		return ErrTooMany
	}
	parts := make([]string, 0, 6)
	for i := 0; i < 12; i += 2 {
		parts = append(parts, cleaned[i:i+2])
	}
	return strings.Join(parts, ":")
}

// FormatAll formats each line of input independently, keeping line positions
// so output rows line up with input rows. Output lines are joined with sep.
func FormatAll(input, sep string) string {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	for i, l := range lines {
		lines[i] = FormatLine(l)
	}
	return strings.Join(lines, sep)
}
