package DiscordRequests

import (
	"GopherSentinel/DiscordObjects"
	"GopherSentinel/DiscordUtils"
	"GopherSentinel/Utils"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func createPutCommandHttpRequest(serializedCommands []byte) (*http.Request, error) {
	credentialsFile, err := Utils.ReadCredentialsFile("")

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, DiscordUtils.GetPushCommandsUrl(credentialsFile.App_Id), bytes.NewBuffer(serializedCommands))

	if err != nil {
		return nil, err
	}

	DiscordUtils.CreateHTTPHeaders(req.Header, credentialsFile.Bot_Token)
	return req, nil
}

func SetAppCommands(commands []DiscordObjects.CreateCommand) (string, error) {

	serializedCommands, err := json.Marshal(commands)

	if err != nil {
		return "", err
	}

	request, err := createPutCommandHttpRequest(serializedCommands)

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	return string(body), nil
}

func RetrieveBase64FromImage(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return toBase64(bytes), nil
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
