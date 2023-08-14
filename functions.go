package eztable

import (
	"regexp"
	"strings"

	"github.com/mattn/go-runewidth"
)

// Length will find the true length of a provided string.
func Length(v string) int {
	return runewidth.StringWidth(strings.ReplaceAll(regexp.MustCompile("[\u001B\u009B\x1b][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))").ReplaceAllString(v, ""), "<escape>", ""))
}

// Strip will strip an item of text of all escape codes entirley.
func Strip(v string) string {
	return regexp.MustCompile("[\u001B\u009B\x1b][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))").ReplaceAllString(v, "")
}

// Repeat will repeat a specified string with the provided amount of loops.
func Repeat(s string, t int) string {
	var out string

	// We use this function instead of the actual string repeater because if a negative number is provided, it will not crash.
	for i := 0; i < t; i++ {
		out += s
	}

	return out
}

// TotalBodys will figure out how many body values we actually have.
func (t *Table) TotalBodys() int {
	var count int = 0

	for _, c := range t.Cells {
		if len(c.Bodys) > count {
			count = len(c.Bodys)
		}
	}

	return count
}

// Split will split the text character by character, except escape codes which are stored as a whole item in itself.
func Split(input string) []string {
	ansiPattern := `\x1b\[[0-9;]*m`

	// Compile the regular expression pattern
	re := regexp.MustCompile(ansiPattern)

	// Find all occurrences of ANSI escape codes in the input string
	ansiCodes := re.FindAllStringIndex(input, -1)

	// Initialize the result array
	var result []string

	// Start position for slicing the input string
	start := 0

	// Process each ANSI escape code and the text between them
	for _, indices := range ansiCodes {
		startIndex, endIndex := indices[0], indices[1]

		// Append the text between the previous escape code and the current one
		rawChars := input[start:startIndex]
		for _, char := range rawChars {
			result = append(result, string(char))
		}

		// Append the ANSI escape code itself
		result = append(result, input[startIndex:endIndex])

		// Update the start position for the next iteration
		start = endIndex
	}

	// Append the remaining text after the last escape code, if any
	rawChars := input[start:]
	for _, char := range rawChars {
		result = append(result, string(char))
	}

	return result
}

// FindSpacing will find the current spacing of a cell.
func FindSpacing(c *Cell) int {
	var Spacing int = 0

	for _, b := range c.Bodys {
		if Length(b) > Spacing {
			Spacing = Length(b)
		}
	}

	if Length(c.Title) > Spacing {
		Spacing = Length(c.Title)
	}

	return Spacing
}
