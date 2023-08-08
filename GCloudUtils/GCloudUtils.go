package GCloudUtils

import (
	"GopherSentinel/Utils"
	"fmt"
	"net/http"
)

type GCloudVisionRequestModel struct {
	Requests []GClouRequest `json:"requests"`
}

type GClouRequest struct {
	Image    Image     `json:"image"`
	Features []Feature `json:"features"`
}

type Image struct {
	Content string `json:"content"`
}

type Feature struct {
	Type string `json:"type"`
}

type GoogleVisionResponse struct {
	Responses []GoogleSafeSearchAnnotations `json:"responses"`
}

type GoogleSafeSearchAnnotations struct {
	SafeSearchAnnotation GoogleAnnotations `json:"safeSearchAnnotation"`
}

type GoogleAnnotations struct {
	Adult    string `json:"adult"`
	Spoof    string `json:"spoof"`
	Medical  string `json:"medical"`
	Violence string `json:"violence"`
	Racy     string `json:"racy"`
}

func GetVisionUrl() string {
	return "https://vision.googleapis.com/v1/images:annotate"
}

func AddGcloudHTTPHeaders(header http.Header) {
	configFile, err := Utils.ReadCredentialsFile("")

	if err != nil {
		fmt.Println(err.Error())
	}

	header.Add("Authorization", fmt.Sprint("Bearer ", configFile.Google_bearer_token))
	header.Add("Content-Type", "application/json")
	header.Add("x-goog-user-project", configFile.Google_User_Project)
}

func notAceptedProbabilitys() []string {
	return []string{
		/*"UNKNOWN",
		"VERY_UNLIKELY",
		"UNLIKELY",
		"UNLIKELY",*/
		"LIKELY",
		"VERY_LIKELY",
	}
}

func IsImageInnapropriated(gVisionObject GoogleVisionResponse) bool {
	probabilityItens := notAceptedProbabilitys()
	responseList := gVisionObject.Responses

	for _, response := range responseList {
		for _, probabilityItem := range probabilityItens {
			if response.SafeSearchAnnotation.Adult == probabilityItem ||
				response.SafeSearchAnnotation.Medical == probabilityItem ||
				response.SafeSearchAnnotation.Racy == probabilityItem ||
				response.SafeSearchAnnotation.Spoof == probabilityItem ||
				response.SafeSearchAnnotation.Violence == probabilityItem {
				return true
			}
		}
	}

	return false
}

func CreateGCloudRequestObject(base64Image string) GCloudVisionRequestModel {
	return GCloudVisionRequestModel{
		Requests: []GClouRequest{
			{
				Image: Image{
					Content: base64Image,
				},
				Features: []Feature{
					{
						Type: "SAFE_SEARCH_DETECTION",
					},
				},
			},
		},
	}
}
