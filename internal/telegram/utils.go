package telegram

import "regexp"

func getTextWithoutCommand(text string) string {
	re := regexp.MustCompile(`^/\w+\s*(.*)$`)
	matches := re.FindStringSubmatch(text)
	if len(matches) < 2 {
		return ""
	}
	return matches[1]
}
