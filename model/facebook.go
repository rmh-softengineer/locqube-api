package model

type Post struct {
	Message string
	Link    *string
	Image   *string
}

type ValidationTokenResponse struct {
	Data ValidationTokenData `json:"data"`
}

type ValidationTokenData struct {
	AppID     string   `json:"app_id"`
	UserID    string   `json:"user_id"`
	IsValid   bool     `json:"is_valid"`
	ExpiresAt int      `json:"expires_at"`
	Scopes    []string `json:"scopes"`
}

type LoginRequest struct {
	AccessToken string `json:"access_token"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
