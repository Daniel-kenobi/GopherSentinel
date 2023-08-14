package Channel

import (
	"GopherSentinel/BotProcessors/Messages/ImageProcessor"
	"GopherSentinel/BotProcessors/Messages/TextProcessor"
	"GopherSentinel/Utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"log"
	"net/http"
)

var session *discordgo.Session
var message *discordgo.MessageCreate

type Operations struct{}

func getRequiredDependencies(s *discordgo.Session, m *discordgo.MessageCreate) (ImageProcessor.IImageProcessor, TextProcessor.ITextProcessor) {
	session, message = s, m

	return &ImageProcessor.ImageProcessor{}, &TextProcessor.TextProcessor{}
}

func (co *Operations) SendMessageToChannel(messageText string, userIDReply string) error {
	var finalMessage string

	if len(userIDReply) > 0 {
		finalMessage = fmt.Sprint("<@", userIDReply, "> ")
	}

	finalMessage += messageText
	_, err := session.ChannelMessageSend(message.ChannelID, finalMessage)

	if err != nil {
		log.Println("Failed to send message to channel")
		log.Println(err.Error())
		return err
	}

	return nil
}

func (co *Operations) DeleteMessageFromChannel(channelId string, messageId string) error {
	err := session.ChannelMessageDelete(channelId, messageId)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (co *Operations) HandleChannelMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || (len(m.Content) <= 0 && len(m.Attachments) <= 0) {
		return
	}

	imageProcessor, messageProcessor := getRequiredDependencies(s, m)

	var appropriatedMessage = true
	var appropriatedAttachment = true

	if len(m.Content) > 0 {
		appropriatedMessage = messageProcessor.IsMessageAppropriated(m.Content)
	}

	if len(m.Attachments) > 0 {
		for _, attachment := range m.Attachments {
			base64Image, err := co.GetBase64FromImageUrl(attachment.URL)

			if err != nil || len(base64Image) == 0 {
				continue
			}

			appropriatedAttachment, err = imageProcessor.IsImageAppropriated(base64Image)

			if err != nil {
				continue
			}
		}
	}

	if !appropriatedMessage || !appropriatedAttachment {
		err := co.DeleteMessageFromChannel(m.Message.ChannelID, m.Message.ID)

		if err != nil {
			log.Println(err.Error())
			return
		}

		err = co.SendMessageToChannel("Mensagem deletada pois contém palavras e/ou anexos inapropriados", message.Author.ID)

		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}

func (co *Operations) GetBase64FromImageUrl(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return Utils.ByteArrayToBase64(bytes), nil
}
