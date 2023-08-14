package Credentials

import (
	"GopherSentinel/Credentials/Models"
	"encoding/json"
	"io"
	"os"
)

func ReadCredentialsFile(completePath string) (Models.CredentialFileModel, error) {

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
