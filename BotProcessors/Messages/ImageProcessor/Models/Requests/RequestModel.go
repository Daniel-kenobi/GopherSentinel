package Requests

type GCloudVisionRequestModel struct {
	Requests []GCloudRequest `json:"requests"`
}

type GCloudRequest struct {
	Image    Image     `json:"image"`
	Features []Feature `json:"features"`
}

type Image struct {
	Content string `json:"content"`
}

type Feature struct {
	Type string `json:"type"`
}
