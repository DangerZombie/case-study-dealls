package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func (h *authHelperImpl) VerifyJWT(headers http.Header) (output parameter.JwtClaims, err error) {
	if headers.Get("Authorization") == "" {
		err = errors.New("token is null, need valid token")
		return
	}

	tokenString := strings.Split(headers["Authorization"][0], " ")[1]

	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Here you need to specify the function that will be used to verify the key.
		// In this case, we are using a shared secret key.
		return []byte(viper.GetString("jwt.secret-key")), nil
	})

	// Verify the token
	if err != nil {
		return
	}

	// Access the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = fmt.Errorf("errors: %s", "token invalid")
		return
	}

	output = parameter.JwtClaims{
		Issuer:  claims["iss"].(string),
		Subject: claims["sub"].(string),
		User:    claims["usr"].(string),
	}

	return
}
