package Models

type CredentialFileModel struct {
	Bot_Token           string `json:"BOT_TOKEN"`
	App_Id              string `json:"APP_ID"`
	Public_key          string `json:"PUBLIC_KEY"`
	Google_User_Project string `json:"GOOGLE_USER_PROJ"`
	Google_bearer_token string `json:"GOOGLE_BEARER_TOKEN"`
}
