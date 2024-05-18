package user_service

import (
	"errors"
	"net/http"

	"github.com/DangerZombie/case-study-dealls/helper/static"
	"github.com/DangerZombie/case-study-dealls/model/entity"
	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"github.com/DangerZombie/case-study-dealls/model/request"
	"github.com/DangerZombie/case-study-dealls/model/response"
)

func (s *userServiceImpl) CreateUserInteraction(req request.CreateUserInteractionRequest) (res response.CreateUserInteractionResponse, code int, err error) {
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
		Id: req.UserId1,
	}

	user, err := s.userRepo.FindUserById(tx, userInput)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	// count swipe left
	if user.Status == static.UserFree {
		if user.SwipeCount > 0 {
			swipeCount := user.SwipeCount - 1

			updateSwipeCountInput := parameter.UpdateSwipeCountInput{
				Id:         user.Id,
				SwipeCount: swipeCount,
			}

			_, err = s.userRepo.UpdateSwipeCount(tx, updateSwipeCountInput)
			if err != nil {
				return res, http.StatusInternalServerError, err
			}
		} else {
			err = errors.New("you have reached the daily swipe")
			return res, http.StatusBadRequest, err
		}
	}

	interactionInput := parameter.CreateUserInteractionInput{
		UserInteraction: entity.UserInteraction{
			UserId1:         req.UserId1,
			UserId2:         req.UserId2,
			InteractionType: req.InteractionType,
		},
	}

	result, err := s.userRepo.CreateUserInteraction(tx, interactionInput)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	res.Message = result.Message

	if req.InteractionType == static.InteractionTypeLike {
		// Checking if another user has swipe like too
		reverseInteractionInput := parameter.FindMatchInteractionInput{
			UserId1: req.UserId1,
			UserId2: req.UserId2,
		}

		reverseInteraction, err := s.userRepo.FindMatchInteraction(tx, reverseInteractionInput)
		if err != nil {
			res.Message = ""
			return res, http.StatusInternalServerError, err
		}

		// add to match table
		if reverseInteraction.Id != "" {
			matchInput := parameter.CreateMatchUserInput{
				Match: entity.Match{
					UserId1: req.UserId1,
					UserId2: req.UserId2,
				},
			}

			_, err = s.userRepo.CreateMatchUser(tx, matchInput)
			if err != nil {
				res.Message = ""
				return res, http.StatusInternalServerError, err
			}

			res.Message = static.MessageMatch
		}
	}

	return res, http.StatusOK, nil
}
