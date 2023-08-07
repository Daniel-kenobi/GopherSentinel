package ChatGPTRequests

import (
	"GopherSentinel/ChatGPTObjects"
	"GopherSentinel/ChatGPTUtils"
	"GopherSentinel/Utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func ChatGPTConversation(conversationObject ChatGPTObjects.ChatGPTConversationObject) error {
	conversationUrl := ChatGPTUtils.GetChatGPTConversationURL()

	marsahaledReq, err := json.Marshal(conversationObject)

	if err != nil {
		return err
	}

	reqByteArray := bytes.NewReader(marsahaledReq)
	req, err := http.NewRequest("POST", conversationUrl, reqByteArray)

	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	credentiallFiles, err := Utils.ReadCredentialsFile("")

	if err != nil {
		fmt.Println("Naõ foi possivel ler o arquivo de credenciais %s", err.Error())
		panic(err.Error())
	}

	req.Header.Add("Authorization", fmt.Sprint("Bearer ", credentiallFiles.ChatGPT_Secret_Key))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	fmt.Println(res.Body)

	return nil
}
