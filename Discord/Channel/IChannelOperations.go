package Channel

import (
	"github.com/bwmarrin/discordgo"
)

type IChannelOperations interface {
	SendMessageToChannel(messageText string, userIDReply string) error
	DeleteMessageFromChannel(channelId string, messageId string) error
	HandleChannelMessages(s *discordgo.Session, m *discordgo.MessageCreate)
	GetBase64FromImageUrl(url string) (string, error)
}
