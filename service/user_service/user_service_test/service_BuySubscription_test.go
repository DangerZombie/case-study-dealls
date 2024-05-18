package user_service_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/DangerZombie/case-study-dealls/helper/auth"
	"github.com/DangerZombie/case-study-dealls/helper/static"
	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"github.com/DangerZombie/case-study-dealls/model/request"
	"github.com/DangerZombie/case-study-dealls/repository"
	"github.com/DangerZombie/case-study-dealls/repository/user_repository"
	"github.com/DangerZombie/case-study-dealls/service/user_service"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestBuySubscription(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)
	mockBaseRepository := repository.NewMockBaseRepository(mockCtrl)
	mockUserRepository := user_repository.NewMockUserRepository(mockCtrl)

	userService := user_service.NewUserService(
		mockAuthHelper,
		mockBaseRepository,
		mockUserRepository,
	)

	id := faker.UUIDHyphenated()
	buySubscriptionRequest := request.BuySubscriptionRequestBody{
		Id: id,
	}

	userInput := parameter.FindUserByIdInput{
		Id: buySubscriptionRequest.Id,
	}

	userOutput := parameter.FindUserByIdOutput{
		Id:         id,
		Username:   faker.Username(),
		Nickname:   faker.Name(),
		Status:     static.UserFree,
		SwipeCount: 10,
	}

	userPremiumOutput := parameter.FindUserByIdOutput{
		Id:         id,
		Username:   faker.Username(),
		Nickname:   faker.Name(),
		Status:     static.UserPremium,
		SwipeCount: -1,
	}

	t.Run("Should return OK", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserById(gomock.Any(), userInput).
			Times(1).
			Return(userOutput, nil)

		mockUserRepository.EXPECT().
			UpdateSubscription(gomock.Any(), gomock.Any()).
			Times(1).
			Return(parameter.UpdateSubscriptionOutput{
				Message: "Success",
			}, nil)

		mockBaseRepository.EXPECT().
			BeginCommit(gomock.Any()).
			Times(1).
			Return()

		result, code, err := userService.BuySubscription(buySubscriptionRequest)

		require.NotEmpty(t, result)
		require.Equal(t, http.StatusOK, code)
		require.Nil(t, err)
	})

	t.Run("Should return error Internal Server Error if failed to fetch user", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserById(gomock.Any(), userInput).
			Times(1).
			Return(parameter.FindUserByIdOutput{}, errors.New("failed"))

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1).
			Return()

		result, code, err := userService.BuySubscription(buySubscriptionRequest)

		require.Empty(t, result)
		require.Equal(t, http.StatusInternalServerError, code)
		require.NotNil(t, err)
	})

	t.Run("Should return error Bad Request if user already premium user", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserById(gomock.Any(), userInput).
			Times(1).
			Return(userPremiumOutput, nil)

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1).
			Return()

		result, code, err := userService.BuySubscription(buySubscriptionRequest)

		require.Empty(t, result)
		require.Equal(t, http.StatusBadRequest, code)
		require.NotNil(t, err)
	})

	t.Run("Should return Internal Server Error if failed to subscribe", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserById(gomock.Any(), userInput).
			Times(1).
			Return(userOutput, nil)

		mockUserRepository.EXPECT().
			UpdateSubscription(gomock.Any(), gomock.Any()).
			Times(1).
			Return(parameter.UpdateSubscriptionOutput{}, errors.New("failed"))

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1).
			Return()

		result, code, err := userService.BuySubscription(buySubscriptionRequest)

		require.Empty(t, result)
		require.Equal(t, http.StatusInternalServerError, code)
		require.NotNil(t, err)
	})
}
