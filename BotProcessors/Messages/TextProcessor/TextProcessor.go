package TextProcessor

import (
	"GopherSentinel/Utils"
	"strings"
)

type TextProcessor struct {
}

func (tp *TextProcessor) IsMessageAppropriated(message string) bool {
	badWords := Utils.PortugueseBadWordList()
	splitedInputStrings := strings.Split(message, " ")

	for _, badWord := range badWords {
		for _, splitedInputString := range splitedInputStrings {
			if strings.ToLower(badWord) == splitedInputString {
				return false
			}
		}
	}

	return true
}
