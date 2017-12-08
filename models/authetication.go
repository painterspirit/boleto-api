package models

// Authentication autenticação para entrada na API do banco
type Authentication struct {
	Username           string `json:"Username,omitempty"`
	Password           string `json:"Password,omitempty"`
	AuthorizationToken string `json:"AuthorizationToken,omitempty"`
	AccessKey          string `json:"AccessKey,omitempty"`
}
