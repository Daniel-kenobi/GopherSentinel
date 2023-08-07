package DiscordUtils

import (
	"fmt"
	"net/http"
)

func getBaseUrl() string {
	return "https://discord.com/api/"
}

func GetPushCommandsUrl(discordAppId string) string {
	return fmt.Sprint(getBaseUrl(), "applications/", discordAppId, "/commands")
}

func GetBotAuthorizationHeader(discordToken string) string {
	return fmt.Sprint("Bot ", discordToken)
}

func GetClientAuthorizationHeader(discordToken string) string {
	return fmt.Sprint("Bearer ", discordToken)
}

func GetKuteGoUrl(botToken string) string {
	return fmt.Sprint("https://kutego-api-", botToken, "-ew.a.run.app")
}

func CreateHTTPHeaders(header http.Header, discordToken string) {
	header.Add("Authorization", GetBotAuthorizationHeader(discordToken))
	header.Add("Content-Type", "application/json")
}
