package auth

import (
	"net/http"

	"github.com/DangerZombie/case-study-dealls/model/parameter"
)

type AuthHelper interface {
	GenerateJWT(id string) (string, error)
	VerifyJWT(headers http.Header) (output parameter.JwtClaims, err error)
}
