package Credentials

import (
	"GopherSentinel/Credentials/Models"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

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

func SaveCredentialsFile(completePath string, updatedFile Models.CredentialFileModel) error {
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
