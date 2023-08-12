package WatchChannel

import (
	"GopherSentinel/BotProcessor/MessageProcessing/ImageProcessing"
	"GopherSentinel/BotProcessor/MessageProcessing/TextProcessing"
	"GopherSentinel/Credentials"
	"GopherSentinel/Utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var session *discordgo.Session
var message *discordgo.MessageCreate

type WatchChannel struct{}

func getRequiredDependencies(s *discordgo.Session, m *discordgo.MessageCreate) (*ImageProcessing.ImageProcessor, *TextProcessing.MessageProcessor) {
	session, message = s, m

	return &ImageProcessing.ImageProcessor{}, &TextProcessing.MessageProcessor{}
}

func (watchSession *WatchChannel) RetrieveBase64FromImage(url string) (string, error) {
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

func (watchSession *WatchChannel) HandleChannelMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
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
			base64Image, err := watchSession.RetrieveBase64FromImage(attachment.URL)

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
		err := watchSession.DeleteMessage(m.Message.ChannelID, m.Message.ID)

		if err != nil {
			log.Println(err.Error())
			return
		}

		err = watchSession.SendMessageToChannel("Mensagem deletada pois contém palavras e/ou anexos inapropriados", message.Author.ID)

		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}

func (watchSession *WatchChannel) SendMessageToChannel(messageText string, userIDReply string) error {
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

func (watchSession *WatchChannel) DeleteMessage(channelId string, messageId string) error {
	err := session.ChannelMessageDelete(channelId, messageId)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (watchSession *WatchChannel) StartBot() {
	instanceRunTimes := 0

	fmt.Println(fmt.Sprintln("GopherSentinel starting - instance: ", instanceRunTimes))

	for instanceRunTimes <= 5 {
		credentials, err := Credentials.ReadCredentialsFile("")

		if err != nil {
			instanceRunTimes++
			log.Println(err.Error())
			continue
		}

		botExecution, err := discordgo.New(fmt.Sprint("Bot ", credentials.Bot_Token))

		if err != nil {
			instanceRunTimes++
			log.Println(err.Error())
			continue
		}

		botExecution.AddHandler(watchSession.HandleChannelMessages)
		botExecution.Identify.Intents = discordgo.IntentsGuildMessages

		err = botExecution.Open()

		if err != nil {
			instanceRunTimes++
			log.Println(err.Error())
			continue
		}

		fmt.Println("GopherSentinel is running...")

		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-sc

		err = botExecution.Close()

		if err != nil {
			return
		}
	}

	log.Println("Errors occurred while trying run GopherSentinel. See logs for more details")
}
