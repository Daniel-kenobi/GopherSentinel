package TextProcessing

import (
	"GopherSentinel/Utils"
	"strings"
)

type MessageProcessor struct {
}

func (mp *MessageProcessor) IsMessageAppropriated(message string) bool {
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
