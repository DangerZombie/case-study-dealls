package user_service

import (
	"net/http"

	"github.com/DangerZombie/case-study-dealls/helper/static"
	"github.com/DangerZombie/case-study-dealls/model/entity"
	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"github.com/DangerZombie/case-study-dealls/model/request"
	"github.com/DangerZombie/case-study-dealls/model/response"
)

func (s *userServiceImpl) RegisterUser(req request.RegisterUserRequestBody) (res response.RegisterUserResponse, code int, err error) {
	tx := s.baseRepo.GetBegin()
	defer func() {
		if err != nil {
			s.baseRepo.BeginRollback(tx)
		} else {
			s.baseRepo.BeginCommit(tx)
		}
	}()

	createUserInput := parameter.CreateUserInput{
		User: entity.User{
			Username:   req.Username,
			Password:   req.Password,
			Gender:     req.Gender,
			Age:        req.Age,
			Location:   req.Location,
			Nickname:   req.Nickname,
			Status:     static.UserFree,
			SwipeCount: 10,
			Verified:   false,
		},
	}

	_, err = s.userRepo.CreateUser(tx, createUserInput)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	res = response.RegisterUserResponse{
		Message: "Success",
	}

	return res, http.StatusOK, nil
}
