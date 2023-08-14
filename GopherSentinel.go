package main

import (
	"GopherSentinel/Credentials"
	"GopherSentinel/Discord/Channel"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func getChannelOperations() Channel.IChannelOperations {
	return &Channel.Operations{}
}

func main() {
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

		channelOps := getChannelOperations()

		botExecution.AddHandler(channelOps.HandleChannelMessages)
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
