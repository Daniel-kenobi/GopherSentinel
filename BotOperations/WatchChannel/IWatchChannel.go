package WatchChannel

import (
	"github.com/bwmarrin/discordgo"
)

type IBotOperations interface {
	SendMessageToChannel(messageText string, userIDReply string)
	DeleteMessage(channelId string, messageId string) error
	HandleChannelMessages(s *discordgo.Session, m *discordgo.MessageCreate)
	RetrieveBase64FromImage(string, error)
	StartBot()
}
