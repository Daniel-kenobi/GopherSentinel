package Responses

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
