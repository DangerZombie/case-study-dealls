package user_service

import (
	"errors"
	"net/http"

	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"github.com/DangerZombie/case-study-dealls/model/request"
	"github.com/DangerZombie/case-study-dealls/model/response"
)

func (s *userServiceImpl) Login(req request.LoginRequestBody) (res response.LoginResponse, code int, err error) {
	if req.Username == "" || req.Password == "" {
		return res, http.StatusBadRequest, errors.New("username and password required")
	}

	tx := s.baseRepo.GetBegin()
	defer func() {
		if err != nil {
			s.baseRepo.BeginRollback(tx)
		} else {
			s.baseRepo.BeginCommit(tx)
		}
	}()

	findUserByUsernameAndPasswordInput := parameter.FindUserByUsernameAndPasswordInput{
		Username: req.Username,
		Password: req.Password,
	}

	user, err := s.userRepo.FindUserByUsernameAndPassword(tx, findUserByUsernameAndPasswordInput)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	token, err := s.authHelper.GenerateJWT(user.Id)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	res = response.LoginResponse{
		Token: token,
	}

	return res, http.StatusOK, nil
}
