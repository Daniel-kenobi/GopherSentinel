package WatchMessageHandle

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type WatchMessageHandle struct {
}

func (mh *WatchMessageHandle) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "/") {
		return
	}
}
