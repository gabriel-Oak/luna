package listening

import (
	"regexp"
	"strings"
)

var detectionRegex = regexp.MustCompile(`(?i)(^([A-z\d](.*))[^\]\)]$)`)
var calledDetectionRegex = regexp.MustCompile(`(?i)l[uo]na`)
var speechRunes = []rune{rune(27), rune(91), rune(50), rune(75), rune(13)}

func containsRune(runes []rune, r rune) bool {
	for _, runeValue := range runes {
		if runeValue == r {
			return true
		}
	}
	return false
}

func removeRunes(s string, runes []rune) string {
	return strings.Map(func(rn rune) rune {
		if int(rn) >= 65 && int(rn) <= 90 {
			return rn
		}

		if containsRune(runes, rn) {
			return -1
		}

		return rn
	}, s)
}

func getLastIndex(length int) int {
	if length <= 0 {
		return 0
	}

	return length - 1
}
