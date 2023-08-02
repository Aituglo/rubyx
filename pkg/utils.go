package pkg

import (
	"fmt"
	"regexp"
)

func ExtractDomain(input string) (string, error) {
	regex := regexp.MustCompile(`(?i)(?:https?://)?(?:\*\.)?([a-z0-9.-]+[a-z0-9])`)
	matches := regex.FindStringSubmatch(input)

	if len(matches) < 2 {
		return "", fmt.Errorf("failed to extract domain from input: %s", input)
	}

	return matches[1], nil
}
