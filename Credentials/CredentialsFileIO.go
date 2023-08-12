package Credentials

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

func SaveCredentialsFile(completePath string, updatedFile CredentialFileModel) error {
	if len(completePath) <= 0 {
		completePath = "Credentials.json"
	}

	updatedJSON, err := json.MarshalIndent(updatedFile, "", "  ")

	if err != nil {
		fmt.Println("Erro ao codificar JSON:", err)
		return err
	}

	// Escrever o JSON atualizado de volta no arquivo
	err = os.WriteFile(completePath, updatedJSON, 0644)

	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return err
	}

	return nil
}

func ReadCredentialsFile(completePath string) (CredentialFileModel, error) {

	if len(completePath) <= 0 {
		completePath = "Credentials.json"
	}

	jsonFile, err := os.Open(completePath)

	if err != nil {
		return CredentialFileModel{}, err
	}

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)

	byteArray, err := io.ReadAll(jsonFile)

	if err != nil {
		return CredentialFileModel{}, err
	}

	var credentialsFile CredentialFileModel
	err = json.Unmarshal(byteArray, &credentialsFile)

	if err != nil {
		return CredentialFileModel{}, err
	}

	return credentialsFile, nil
}

func SaveBearerOnCredentialsFile(completePath string, bearerToken string) error {
	if len(bearerToken) == 0 {
		return errors.New("token não pode ser nulo")
	}

	outdatedCredentials, err := ReadCredentialsFile(completePath)

	if err != nil {
		return errors.New("não foi possível atualizar o arquivo de credenciais")
	}

	outdatedCredentials.Google_bearer_token = bearerToken

	err = SaveCredentialsFile("", outdatedCredentials)

	if err != nil {
		return errors.New("não foi possível atualizar o arquivo de credenciais")
	}

	return nil
}
