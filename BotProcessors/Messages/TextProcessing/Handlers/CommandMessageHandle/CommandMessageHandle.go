package CommandMessageHandle

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type CommandMessageHandle struct {
}

func (mh *CommandMessageHandle) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !strings.HasPrefix(m.Content, "/") {
		return
	}
}