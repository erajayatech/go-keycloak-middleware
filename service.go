package keycloakmiddleware

import (
	"crypto/rsa"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
)

func getPublicKey(kid string) (*rsa.PublicKey, error) {
	var keysUrl = getEnv("KEYCLOAK_JWT_JWK_ENDPOINT")
	keysRequest, err := http.NewRequest("GET", keysUrl, nil)
	if err != nil {
		return nil, err
	}

	keysResponse, err := http.DefaultClient.Do(keysRequest)
	if err != nil {
		return nil, err
	}

	keysResponseBody, err := ioutil.ReadAll(keysResponse.Body)
	if err != nil {
		return nil, err
	}

	var jwk keycloakJWK
	err = json.Unmarshal([]byte(keysResponseBody), &jwk)
	if err != nil {
		return nil, err
	}

	var n *big.Int
	var e int
	for _, key := range jwk.Keys {
		if key.Kid == kid {
			n = decodeBase64BigInt(key.N)
			e = int(decodeBase64BigInt(key.E).Int64())
			break
		}
	}

	if n == nil || e == 0 {
		return nil, err
	}

	jwtKey := &rsa.PublicKey{
		N: n,
		E: e,
	}
	return jwtKey, nil
}

func isScopesValid(claims claims, scopes []string) bool {
	scopeMap := make(map[string]struct{})

	for _, search := range scopes {
		scopeMap[search] = struct{}{}
	}

	for _, permission := range claims.Authorization.Permissions {
		for _, scope := range permission.Scopes {
			if _, exists := scopeMap[scope]; exists {
				return true
			}
		}
	}

	return false
}
