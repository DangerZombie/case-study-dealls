package user_service

import (
	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"github.com/DangerZombie/case-study-dealls/model/request"
	"github.com/DangerZombie/case-study-dealls/model/response"
)

func (s *userServiceImpl) ResetSwipeCount(req request.ResetSwipeCountRequest) (res response.ResetSwipeCountResponse, err error) {
	tx := s.baseRepo.GetBegin()
	defer func() {
		if err != nil {
			s.baseRepo.BeginRollback(tx)
		} else {
			s.baseRepo.BeginCommit(tx)
		}
	}()

	resetInput := parameter.ResetSwipeCountInput{
		Status:     req.Status,
		SwipeCount: req.SwipeCount,
	}

	result, err := s.userRepo.ResetSwipeCount(tx, resetInput)
	if err != nil {
		return res, err
	}

	res = response.ResetSwipeCountResponse{
		Message: result.Message,
	}

	return res, nil
}
