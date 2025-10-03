package templatetagger

import (
	"fmt"
	"os"
	"regexp"
)

// TagTemplateText parses the given template file in the template folder,
// finds all text between '>' and '<', and replaces it with a {{ .tag }}
// where 'tag' is the surrounding HTML tag name.
func TagTemplateText(path string) error {
	contentBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	content := string(contentBytes)

	// Regex to match: <tag ...>text</tag>
	re := regexp.MustCompile(`<([a-zA-Z0-9]+)([^>]*)>([^<>]+)</([a-zA-Z0-9]+)>`)
	result := re.ReplaceAllStringFunc(content, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		if len(submatches) < 5 {
			return match
		}
		tag := submatches[1]
		attrs := submatches[2]
		closingTag := submatches[4]
		if tag != closingTag {
			return match
		}

		// Replace the text between tags with {{ .tag }}
		return fmt.Sprintf("<%s%s>{{ .%s }}</%s>", tag, attrs, tag, tag)
	})

	os.WriteFile(path, []byte(result), 0644)

	return nil
}
