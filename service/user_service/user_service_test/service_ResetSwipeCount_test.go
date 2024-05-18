package user_service_test

import (
	"errors"
	"testing"

	"github.com/DangerZombie/case-study-dealls/helper/auth"
	"github.com/DangerZombie/case-study-dealls/helper/static"
	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"github.com/DangerZombie/case-study-dealls/model/request"
	"github.com/DangerZombie/case-study-dealls/repository"
	"github.com/DangerZombie/case-study-dealls/repository/user_repository"
	"github.com/DangerZombie/case-study-dealls/service/user_service"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestResetSwipeCount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)
	mockBaseRepository := repository.NewMockBaseRepository(mockCtrl)
	mockUserRepository := user_repository.NewMockUserRepository(mockCtrl)

	userService := user_service.NewUserService(
		mockAuthHelper,
		mockBaseRepository,
		mockUserRepository,
	)

	swipeRequest := request.ResetSwipeCountRequest{
		Status:     static.UserFree,
		SwipeCount: 10,
	}

	resetInput := parameter.ResetSwipeCountInput{
		Status:     swipeRequest.Status,
		SwipeCount: swipeRequest.SwipeCount,
	}

	resetOutput := parameter.ResetSwipeCountOutput{
		Message: "Success",
	}

	t.Run("Should return OK", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			ResetSwipeCount(gomock.Any(), resetInput).
			Times(1).
			Return(resetOutput, nil)

		mockBaseRepository.EXPECT().
			BeginCommit(gomock.Any()).
			Times(1).
			Return()

		result, err := userService.ResetSwipeCount(swipeRequest)

		require.Equal(t, "Success", result.Message)
		require.Nil(t, err)
	})

	t.Run("Should return error Internal Server Error if failed to reset swipe count", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			ResetSwipeCount(gomock.Any(), resetInput).
			Times(1).
			Return(parameter.ResetSwipeCountOutput{}, errors.New("failed"))

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1).
			Return()

		result, err := userService.ResetSwipeCount(swipeRequest)

		require.Empty(t, result.Message)
		require.NotNil(t, err)
	})
}
