package keycloakmiddleware

import "github.com/cristalhq/jwt/v3"

// Set all model to private

type claims struct {
	jwt.StandardClaims
	Authorization authorization `json:"authorization,omitempty"`
	Username      string        `json:"preferred_username,omitempty"`
	Name          string        `json:"name,omitempty"`
	Email         string        `json:"email,omitempty"`
}

type authorization struct {
	Permissions []permission `json:"permissions,omitempty"`
}

type permission struct {
	RsID   string   `json:"rsid,omitempty"`
	RsName string   `json:"rsname,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
}

type keycloakJWKDetail struct {
	Key     string   `json:"kty"`
	Kid     string   `json:"kid"`
	Use     string   `json:"sig"`
	Alg     string   `json:"alg"`
	N       string   `json:"n"`
	E       string   `json:"e"`
	X5c     []string `json:"x5c"`
	X5t     string   `json:"x5t"`
	X5tS256 string   `json:"x5t#S256"`
}

type keycloakJWK struct {
	Keys []keycloakJWKDetail `json:"keys"`
}
