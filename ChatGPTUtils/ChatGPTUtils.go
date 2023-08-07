package ChatGPTUtils

import "fmt"

func getChatGPTBaseUrl() string {
	return "https://api.openai.com/v1/"
}

func GetChatGPTConversationURL() string {
	return fmt.Sprint(getChatGPTBaseUrl(), "chat/completions")
}
