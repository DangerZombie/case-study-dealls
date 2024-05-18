package user_service

import (
	"net/http"

	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"github.com/DangerZombie/case-study-dealls/model/request"
	"github.com/DangerZombie/case-study-dealls/model/response"
)

func (s userServiceImpl) GetUserToSwipe(req request.GetUserToSwipeRequest) (res response.GetUserToSwipeResponse, code int, err error) {
	tx := s.baseRepo.GetBegin()
	defer func() {
		if err != nil {
			s.baseRepo.BeginRollback(tx)
		} else {
			s.baseRepo.BeginCommit(tx)
		}
	}()

	findUserToSwipeInput := parameter.FindUserToSwipeInput{
		Id: req.Id,
	}

	// user ready to swipe
	user, err := s.userRepo.FindUserToSwipe(tx, findUserToSwipeInput)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	res = response.GetUserToSwipeResponse{
		Id:       user.Id,
		Nickname: user.Nickname,
		Gender:   user.Gender,
		Age:      user.Age,
		Location: user.Location,
	}

	return res, http.StatusOK, nil
}
