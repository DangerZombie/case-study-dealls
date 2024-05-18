package user_service

import (
	"errors"
	"net/http"

	"github.com/DangerZombie/case-study-dealls/helper/static"
	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"github.com/DangerZombie/case-study-dealls/model/request"
	"github.com/DangerZombie/case-study-dealls/model/response"
)

func (s *userServiceImpl) BuySubscription(req request.BuySubscriptionRequestBody) (res response.BuySubscriptionResponse, code int, err error) {
	tx := s.baseRepo.GetBegin()
	defer func() {
		if err != nil {
			s.baseRepo.BeginRollback(tx)
		} else {
			s.baseRepo.BeginCommit(tx)
		}
	}()

	// check user
	userInput := parameter.FindUserByIdInput{
		Id: req.Id,
	}

	user, err := s.userRepo.FindUserById(tx, userInput)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	if user.Status == static.UserPremium {
		return res, http.StatusBadRequest, errors.New("user already premium user")
	}

	updateSubscriptionInput := parameter.UpdateSubscriptionInput{
		Id: user.Id,
	}

	result, err := s.userRepo.UpdateSubscription(tx, updateSubscriptionInput)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	res = response.BuySubscriptionResponse{
		Message: result.Message,
	}

	return res, http.StatusOK, nil
}
