package models

type GoogleAccountCredential struct {
	AccessToken  string `json:"accessToken"`
	Created      int    `json:"created"`
	Scope        string `json:"scope"`
	IDToken      string `json:"idToken"`
	ExpiresIn    int    `json:"expiresIn"`
	RefreshToken string `json:"refreshToken"`
}
