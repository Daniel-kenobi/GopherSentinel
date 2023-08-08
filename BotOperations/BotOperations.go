package BotOperations

import (
	"GopherSentinel/DiscordObjects"
	"GopherSentinel/DiscordRequests"
	"GopherSentinel/DiscordUtils"
	"GopherSentinel/GCloudRequests"
	"GopherSentinel/GCloudUtils"
	"GopherSentinel/Utils"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func createCommandOperation(commands []DiscordObjects.CreateCommand) {
	response, err := DiscordRequests.SetAppCommands(commands)

	if err != nil {
		fmt.Println("Erro ao enviar request: ", err.Error())
		panic(err.Error())
	}

	fmt.Println("Sucesso ao enviar requisição: ", response)
}

func isOffensiveImage(base64Image string) bool {
	resp, err := GCloudRequests.SendSafeSearchRequest(base64Image)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return GCloudUtils.IsImageInnapropriated(resp)
}

func handlechanelMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || (len(m.Content) <= 0 && len(m.Attachments) <= 0) {
		return
	}

	inapropriatedMessage := Utils.IsInapropriatedWord(m.Content)
	innapropriatedAttachment := false

	if len(m.Attachments) > 0 {
		for _, attach := range m.Attachments {
			base64StringImage, err := DiscordRequests.RetrieveBase64FromImage(attach.URL)

			if err != nil {
				continue
			}

			innapropriatedAttachment = isOffensiveImage(base64StringImage)
		}
	}

	if inapropriatedMessage || innapropriatedAttachment {
		err := s.ChannelMessageDelete(m.ChannelID, m.Message.ID)

		if err != nil {
			fmt.Println(err.Error())
			panic(err.Error())
		}
		_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprint("<@", m.Author.ID, ">", " Mensagem deletada pois contém palavras inapropriadas"))

		if err != nil {
			fmt.Println(err.Error())
			panic(err.Error())
		}
	}
}

func StartBot() {
	credentials, err := Utils.ReadCredentialsFile("")

	if err != nil {
		panic(err.Error())
	}

	bot, err := discordgo.New(DiscordUtils.GetBotAuthorizationHeader(credentials.Bot_Token))

	if err != nil {
		panic(err.Error())
	}

	bot.AddHandler(handlechanelMessages)
	bot.Identify.Intents = discordgo.IntentsGuildMessages

	err = bot.Open()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("GopherSentinel em execução. Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()
}
