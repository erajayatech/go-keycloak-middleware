package keycloakmiddleware

import (
	"encoding/json"
	"github.com/cristalhq/jwt/v3"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type middleware struct {
	wrapperCode int // 0: default, 1:standard, 2:traceable
}

func Construct(wrapperCode int) middleware {
	return middleware{wrapperCode: wrapperCode}
}

func (middleware *middleware) Validate(scopes []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		s := strings.SplitN(context.Request.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			msg := "Authorization token is not found"
			context.AbortWithStatusJSON(http.StatusUnauthorized, middleware.wrapper(context, msg, nil))
			return
		}

		var headerToken = s[1]
		unverifiedToken, err := jwt.Parse([]byte(headerToken))
		if err != nil {
			msg := err.Error()
			context.AbortWithStatusJSON(http.StatusInternalServerError, middleware.wrapper(context, msg, nil))
			return
		}

		kid := unverifiedToken.Header().KeyID
		key, err := getPublicKey(kid)
		if err != nil {
			msg := err.Error()
			context.AbortWithStatusJSON(http.StatusUnauthorized, middleware.wrapper(context, msg, nil))
			return
		}

		verifier, err := jwt.NewVerifierRS(jwt.RS256, key)
		if err != nil {
			msg := err.Error()
			context.AbortWithStatusJSON(http.StatusInternalServerError, middleware.wrapper(context, msg, nil))
			return
		}

		token, err := jwt.ParseAndVerifyString(headerToken, verifier)
		if err != nil {
			msg := err.Error()
			context.AbortWithStatusJSON(http.StatusInternalServerError, middleware.wrapper(context, msg, nil))
			return
		}

		var claims claims
		errClaims := json.Unmarshal(token.RawClaims(), &claims)
		if errClaims != nil {
			msg := errClaims.Error()
			context.AbortWithStatusJSON(http.StatusUnauthorized, middleware.wrapper(context, msg, nil))
			return
		}

		var iss = getEnv("KEYCLOAK_JWT_ISS")
		if claims.Issuer != iss {
			msg := "Token issuer is not valid"
			context.AbortWithStatusJSON(http.StatusUnauthorized, middleware.wrapper(context, msg, nil))
			return
		}

		if claims.ExpiresAt.Unix() < time.Now().Unix() {
			msg := "Token expired"
			context.AbortWithStatusJSON(http.StatusUnauthorized, middleware.wrapper(context, msg, nil))
			return
		}

		if !isScopesValid(claims, scopes) {
			msg := "Access to this endpoint is not allowed"
			context.AbortWithStatusJSON(http.StatusForbidden, middleware.wrapper(context, msg, nil))
			return
		}
	}
}
