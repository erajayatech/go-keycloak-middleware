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

		var isEnabled = getEnvOrDefault("KEYCLOAK_JWT_ENABLED", "false").(string)
		if strings.ToLower(isEnabled) == "false" || isEnabled == "0" {
			context.Next()
			return
		}

		s := strings.SplitN(context.Request.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			msg := "Authorization token is not found"
			middleware.abort(http.StatusUnauthorized, context, msg)
			return
		}

		var headerToken = s[1]
		unverifiedToken, err := jwt.Parse([]byte(headerToken))
		if err != nil {
			msg := err.Error()
			middleware.abort(http.StatusUnauthorized, context, msg)
			return
		}

		kid := unverifiedToken.Header().KeyID
		key, err := getPublicKey(kid)
		if err != nil {
			msg := err.Error()
			middleware.abort(http.StatusUnauthorized, context, msg)
			return
		}

		verifier, err := jwt.NewVerifierRS(jwt.RS256, key)
		if err != nil {
			msg := err.Error()
			middleware.abort(http.StatusUnauthorized, context, msg)
			return
		}

		token, err := jwt.ParseAndVerifyString(headerToken, verifier)
		if err != nil {
			msg := err.Error()
			middleware.abort(http.StatusUnauthorized, context, msg)
			return
		}

		var claims claims
		errClaims := json.Unmarshal(token.RawClaims(), &claims)
		if errClaims != nil {
			msg := errClaims.Error()
			middleware.abort(http.StatusUnauthorized, context, msg)
			return
		}

		var iss = getEnv("KEYCLOAK_JWT_ISS")
		if claims.Issuer != iss {
			msg := "Token issuer is not valid"
			middleware.abort(http.StatusUnauthorized, context, msg)
			return
		}

		if claims.ExpiresAt.Unix() < time.Now().Unix() {
			msg := "Token expired"
			middleware.abort(http.StatusUnauthorized, context, msg)
			return
		}

		if !isScopesValid(claims, scopes) {
			msg := "Access to this endpoint is not allowed"
			middleware.abort(http.StatusForbidden, context, msg)
			return
		}

		context.Next()
	}
}

func (middleware *middleware) abort(status int, context *gin.Context, message interface{}) {
	httpStatus := http.StatusOK
	if middleware.wrapperCode != 0 {
		httpStatus = status
	}
	context.AbortWithStatusJSON(httpStatus, middleware.wrapper(status, context, message, nil))
}
