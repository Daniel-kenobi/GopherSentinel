package ImageProcessing

import (
	"GopherSentinel/Credentials"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type ImageProcessor struct {
}

func (g *ImageProcessor) IsImageAppropriated(base64DecodedImage string) (bool, error) {
	safeSearchRequestObj := GCloudVisionRequestModel{
		Requests: []GCloudRequest{
			{
				Image: Image{
					Content: base64DecodedImage,
				},
				Features: []Feature{
					{
						Type: "SAFE_SEARCH_DETECTION",
					},
				},
			},
		},
	}

	byteRequest, err := json.Marshal(safeSearchRequestObj)

	if err != nil {
		fmt.Println(err.Error())
		return true, err
	}

	reader := bytes.NewReader(byteRequest)
	req, err := http.NewRequest("POST", GoogleVisionUrl(), reader)

	addHTTPHeaders(req.Header)

	if err != nil {
		fmt.Println(err.Error())
		return true, err
	}

	var httpClient = &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		return true, err
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == 401 {
			updatedToken := g.GetPlatformToken()
			err := Credentials.SaveBearerOnCredentialsFile("", updatedToken)

			if err != nil {
				return false, err
			}

			return g.IsImageAppropriated(base64DecodedImage)
		} else {
			return true, errors.New(http.StatusText(resp.StatusCode))
		}
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var responseBody GoogleVisionResponse
	err = json.Unmarshal(body, &responseBody)

	if err != nil {
		fmt.Println(err.Error())
		return true, err
	}

	return checkImageProbability(responseBody), nil
}

func (g *ImageProcessor) GetPlatformToken() string {
	cmd := exec.Command("gcloud", "auth", "print-access-token")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("Erro ao executar o comando: %s\n", err)
		log.Printf("Saída de erro: %s\n", stderr.String())
		panic(err.Error())
	}

	output := strings.TrimSuffix(stdout.String(), "\r\n")

	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}

	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}

	return output
}

func addHTTPHeaders(header http.Header) {
	configFile, err := Credentials.ReadCredentialsFile("")

	if err != nil {
		fmt.Println(err.Error())
	}

	header.Add("Authorization", fmt.Sprint("Bearer ", configFile.Google_bearer_token))
	header.Add("Content-Type", "application/json")
	header.Add("x-goog-user-project", configFile.Google_User_Project)
}

func notAceptedProbabilitys() []string {
	return []string{
		"LIKELY",
		"VERY_LIKELY",
	}
}

func checkImageProbability(gVisionResponse GoogleVisionResponse) bool {
	probabilityItens := notAceptedProbabilitys()
	responseList := gVisionResponse.Responses

	for _, response := range responseList {
		for _, probabilityItem := range probabilityItens {
			if response.SafeSearchAnnotation.Adult == probabilityItem ||
				response.SafeSearchAnnotation.Medical == probabilityItem ||
				response.SafeSearchAnnotation.Racy == probabilityItem ||
				response.SafeSearchAnnotation.Spoof == probabilityItem ||
				response.SafeSearchAnnotation.Violence == probabilityItem {
				return false
			}
		}
	}

	return true
}
