package user_service_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/DangerZombie/case-study-dealls/helper/auth"
	"github.com/DangerZombie/case-study-dealls/model/base"
	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"github.com/DangerZombie/case-study-dealls/model/request"
	"github.com/DangerZombie/case-study-dealls/repository"
	"github.com/DangerZombie/case-study-dealls/repository/user_repository"
	"github.com/DangerZombie/case-study-dealls/service/user_service"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetUserToSwiper(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)
	mockBaseRepository := repository.NewMockBaseRepository(mockCtrl)
	mockUserRepository := user_repository.NewMockUserRepository(mockCtrl)

	userService := user_service.NewUserService(
		mockAuthHelper,
		mockBaseRepository,
		mockUserRepository,
	)

	getUserRequest := request.GetUserToSwipeRequest{
		Id: faker.UUIDHyphenated(),
	}

	findUserToSwipeInput := parameter.FindUserToSwipeInput{
		Id: getUserRequest.Id,
	}

	findUserToSwipeOutput := parameter.FindUserToSwipeOutput{
		BaseModel: base.BaseModel{
			Id: faker.UUIDHyphenated(),
		},
		Username: faker.Username(),
		Nickname: faker.Name(),
		Gender:   faker.Gender(),
		Age:      30,
		Location: faker.Sentence(),
	}

	t.Run("Should return OK", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserToSwipe(gomock.Any(), findUserToSwipeInput).
			Times(1).
			Return(findUserToSwipeOutput, nil)

		mockBaseRepository.EXPECT().
			BeginCommit(gomock.Any()).
			Times(1).
			Return()

		result, code, err := userService.GetUserToSwipe(getUserRequest)

		require.NotEmpty(t, result)
		require.Equal(t, http.StatusOK, code)
		require.Nil(t, err)
	})

	t.Run("Should return error Internal Server Error if failed to fetch user to swipe", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserToSwipe(gomock.Any(), findUserToSwipeInput).
			Times(1).
			Return(parameter.FindUserToSwipeOutput{}, errors.New("failed"))

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1).
			Return()

		result, code, err := userService.GetUserToSwipe(getUserRequest)

		require.Empty(t, result)
		require.Equal(t, http.StatusInternalServerError, code)
		require.NotNil(t, err)
	})
}
