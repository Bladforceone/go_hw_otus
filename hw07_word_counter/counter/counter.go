package counter

import (
	"regexp"
	"strings"
)

func CountWords(s string) map[string]int {
	if s == "" {
		return nil
	}

	re := regexp.MustCompile(`[[:punct:]]`)
	s = re.ReplaceAllString(s, "")

	strs := strings.Fields(s)

	ans := make(map[string]int)
	for _, word := range strs {
		word = strings.ToLower(word)
		if word != "" {
			ans[word]++
		}
	}

	return ans
}
