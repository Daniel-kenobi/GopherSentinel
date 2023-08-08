package GCloudRequests

import (
	"GopherSentinel/GCloudUtils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendSafeSearchRequest(base64Image string) (GCloudUtils.GoogleVisionResponse, error) {
	byteRequest, err := json.Marshal(GCloudUtils.CreateGCloudRequestObject(base64Image))

	if err != nil {
		fmt.Println(err.Error())
		return GCloudUtils.GoogleVisionResponse{}, err
	}

	reader := bytes.NewReader(byteRequest)
	req, err := http.NewRequest("POST", GCloudUtils.GetVisionUrl(), reader)

	GCloudUtils.AddGcloudHTTPHeaders(req.Header)

	if err != nil {
		fmt.Println(err.Error())
		return GCloudUtils.GoogleVisionResponse{}, err
	}

	var httpClient = &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		return GCloudUtils.GoogleVisionResponse{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var responseBody GCloudUtils.GoogleVisionResponse
	err = json.Unmarshal(body, &responseBody)

	if err != nil {
		fmt.Println(err.Error())
		return GCloudUtils.GoogleVisionResponse{}, err
	}

	return responseBody, nil
}
