package main

import (
	"GopherSentinel/DiscordOperations"
	"GopherSentinel/DiscordUtils"
	"GopherSentinel/Utils"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	credentials, err := Utils.ReadCredentialsFile("")

	if err != nil {
		panic(err.Error())
	}

	bot, err := discordgo.New(DiscordUtils.GetBotAuthorizationHeader(credentials.Bot_Token))

	if err != nil {
		panic(err.Error())
	}

	bot.AddHandler(DiscordOperations.HandleSendedMessages)
	bot.Identify.Intents = discordgo.IntentsGuildMessages

	err = bot.Open()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("GopherSentinel em execução. Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	bot.Close()
}
