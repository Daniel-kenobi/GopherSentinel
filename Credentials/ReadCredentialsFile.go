package Credentials

import (
	"GopherSentinel/Credentials/Models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadCredentialsFile(completePath string) (Models.CredentialFileModel, error) {

	fmt.Println("ReadCredentialsFile starting")

	if len(completePath) <= 0 {
		completePath = "Credentials.json"
	}

	jsonFile, err := os.Open(completePath)

	if err != nil {
		return Models.CredentialFileModel{}, err
	}

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println("ReadCredentialsFile error")
			log.Println(err.Error())
		}
	}(jsonFile)

	byteArray, err := io.ReadAll(jsonFile)

	if err != nil {
		return Models.CredentialFileModel{}, err
	}

	var credentialsFile Models.CredentialFileModel
	err = json.Unmarshal(byteArray, &credentialsFile)

	if err != nil {
		return Models.CredentialFileModel{}, err
	}

	return credentialsFile, nil
}
