package Handlers

import "github.com/bwmarrin/discordgo"

type IMessageHandle interface {
	Handle(s *discordgo.Session, m *discordgo.MessageCreate)
}
