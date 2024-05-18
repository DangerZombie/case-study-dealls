package auth

import (
	"time"

	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func (h *authHelperImpl) GenerateJWT(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	tx := h.baseRepo.GetBegin()
	user, err := h.userRepo.FindUserById(tx, parameter.FindUserByIdInput{
		Id: id,
	})

	if err != nil {
		h.baseRepo.BeginRollback(tx)
		return "", err
	} else {
		h.baseRepo.BeginCommit(tx)
	}

	claims := token.Claims.(jwt.MapClaims)
	now := time.Now()

	claims["iss"] = user.Id
	claims["iat"] = now.Unix()
	claims["sub"] = user.Username
	claims["exp"] = now.Add(24 * time.Hour).Unix()
	claims["usr"] = user.Nickname

	secret := []byte(viper.GetString("jwt.secret-key"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
