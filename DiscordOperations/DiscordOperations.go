package DiscordOperations

import (
	"GopherSentinel/DiscordObjects"
	"GopherSentinel/DiscordRequests"
	"GopherSentinel/Utils"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func CreateCommandOperation(commands []DiscordObjects.CreateCommand) {
	response, err := DiscordRequests.SetAppCommands(commands)

	if err != nil {
		fmt.Println("Erro ao enviar request: ", err.Error())
		panic(err.Error())
	}

	fmt.Println("Sucesso ao enviar requisição: ", response)
}

func HandleSendedMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if len(m.Content) <= 0 {
		return
	}

	invalidMessage := Utils.IsInapropriatedWord(m.Content)

	// Valida se a mensagem está de acordo com o esperado
	if invalidMessage {
		err := s.ChannelMessageDelete(m.ChannelID, m.Message.ID)

		if err != nil {
			fmt.Println(err.Error())
			panic(err.Error())
		}

		_, err = s.ChannelMessageSend(m.ChannelID, "Mensagem deletada pois contém palavras inapropriadas")

		if err != nil {
			fmt.Println(err.Error())
			panic(err.Error())
		}
	}
}
