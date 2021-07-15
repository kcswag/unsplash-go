package proto

type Auth struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	Scope string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
	CreatedAt int64
}
