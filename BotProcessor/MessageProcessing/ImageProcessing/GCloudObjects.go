package ImageProcessing

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
