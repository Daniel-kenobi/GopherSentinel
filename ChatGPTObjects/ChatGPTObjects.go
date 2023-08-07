package ChatGPTObjects

type ChatGPTMEssageObject struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTConversationObject struct {
	Model    string                 `json:"model"`
	Messages []ChatGPTMEssageObject `json:"messages"`
}
