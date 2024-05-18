package user_service_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/DangerZombie/case-study-dealls/helper/auth"
	"github.com/DangerZombie/case-study-dealls/helper/static"
	"github.com/DangerZombie/case-study-dealls/model/base"
	"github.com/DangerZombie/case-study-dealls/model/entity"
	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"github.com/DangerZombie/case-study-dealls/model/request"
	"github.com/DangerZombie/case-study-dealls/repository"
	"github.com/DangerZombie/case-study-dealls/repository/user_repository"
	"github.com/DangerZombie/case-study-dealls/service/user_service"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateUserInteraction(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)
	mockBaseRepository := repository.NewMockBaseRepository(mockCtrl)
	mockUserRepository := user_repository.NewMockUserRepository(mockCtrl)

	userService := user_service.NewUserService(
		mockAuthHelper,
		mockBaseRepository,
		mockUserRepository,
	)

	id1 := faker.UUIDHyphenated()
	id2 := faker.UUIDHyphenated()
	createUserInteractionRequest := request.CreateUserInteractionRequest{
		UserId1:         id1,
		UserId2:         id2,
		InteractionType: static.InteractionTypeLike,
	}

	userInput := parameter.FindUserByIdInput{
		Id: id1,
	}

	userOutput := parameter.FindUserByIdOutput{
		Id:         id1,
		Username:   faker.Username(),
		Nickname:   faker.Name(),
		Status:     static.UserFree,
		SwipeCount: 10,
	}

	userInsufficientSwipeOutput := parameter.FindUserByIdOutput{
		Id:         id1,
		Username:   faker.Username(),
		Nickname:   faker.Name(),
		Status:     static.UserFree,
		SwipeCount: 0,
	}

	interactionInput := parameter.CreateUserInteractionInput{
		UserInteraction: entity.UserInteraction{
			UserId1:         id1,
			UserId2:         id2,
			InteractionType: static.InteractionTypeLike,
		},
	}

	interactionOutput := parameter.CreateUserInteractionOutput{
		Message: static.MessageLike,
	}

	reverseInteractionInput := parameter.FindMatchInteractionInput{
		UserId1: id1,
		UserId2: id2,
	}

	reverseInteractionOutput := parameter.FindMatchInteractionOutput{
		BaseModel: base.BaseModel{
			Id: faker.UUIDHyphenated(),
		},
		UserId1: id1,
		UserId2: id2,
	}

	matchInput := parameter.CreateMatchUserInput{
		Match: entity.Match{
			UserId1: id1,
			UserId2: id2,
		},
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
			UpdateSwipeCount(gomock.Any(), gomock.Any()).
			Times(1).
			Return(parameter.UpdateSwipeCountOutput{}, nil)

		mockUserRepository.EXPECT().
			CreateUserInteraction(gomock.Any(), interactionInput).
			Times(1).
			Return(interactionOutput, nil)

		mockUserRepository.EXPECT().
			FindMatchInteraction(gomock.Any(), reverseInteractionInput).
			Times(1).
			Return(reverseInteractionOutput, nil)

		mockUserRepository.EXPECT().
			CreateMatchUser(gomock.Any(), matchInput).
			Times(1).
			Return(parameter.CreateMatchUserOutput{}, nil)

		mockBaseRepository.EXPECT().
			BeginCommit(gomock.Any()).
			Times(1).
			Return()

		result, code, err := userService.CreateUserInteraction(createUserInteractionRequest)

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

		result, code, err := userService.CreateUserInteraction(createUserInteractionRequest)

		require.Empty(t, result)
		require.Equal(t, http.StatusInternalServerError, code)
		require.NotNil(t, err)
	})

	t.Run("Should return error Internal Server Error if failed to update swipe count", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserById(gomock.Any(), userInput).
			Times(1).
			Return(userOutput, nil)

		mockUserRepository.EXPECT().
			UpdateSwipeCount(gomock.Any(), gomock.Any()).
			Times(1).
			Return(parameter.UpdateSwipeCountOutput{}, errors.New("failed"))

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1).
			Return()

		result, code, err := userService.CreateUserInteraction(createUserInteractionRequest)

		require.Empty(t, result)
		require.Equal(t, http.StatusInternalServerError, code)
		require.NotNil(t, err)
	})

	t.Run("Should return error Bad Request if insufficient swipe count", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserById(gomock.Any(), userInput).
			Times(1).
			Return(userInsufficientSwipeOutput, nil)

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1).
			Return()

		result, code, err := userService.CreateUserInteraction(createUserInteractionRequest)

		require.Empty(t, result)
		require.Equal(t, http.StatusBadRequest, code)
		require.NotNil(t, err)
	})

	t.Run("Should return error Internal Server Error if failed to create user interaction", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserById(gomock.Any(), userInput).
			Times(1).
			Return(userOutput, nil)

		mockUserRepository.EXPECT().
			UpdateSwipeCount(gomock.Any(), gomock.Any()).
			Times(1).
			Return(parameter.UpdateSwipeCountOutput{}, nil)

		mockUserRepository.EXPECT().
			CreateUserInteraction(gomock.Any(), interactionInput).
			Times(1).
			Return(parameter.CreateUserInteractionOutput{}, errors.New("failed"))

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1).
			Return()

		result, code, err := userService.CreateUserInteraction(createUserInteractionRequest)

		require.Empty(t, result)
		require.Equal(t, http.StatusInternalServerError, code)
		require.NotNil(t, err)
	})

	t.Run("Should return error Internal Server Error if failed to fetch matching partner", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserById(gomock.Any(), userInput).
			Times(1).
			Return(userOutput, nil)

		mockUserRepository.EXPECT().
			UpdateSwipeCount(gomock.Any(), gomock.Any()).
			Times(1).
			Return(parameter.UpdateSwipeCountOutput{}, nil)

		mockUserRepository.EXPECT().
			CreateUserInteraction(gomock.Any(), interactionInput).
			Times(1).
			Return(interactionOutput, nil)

		mockUserRepository.EXPECT().
			FindMatchInteraction(gomock.Any(), reverseInteractionInput).
			Times(1).
			Return(parameter.FindMatchInteractionOutput{}, errors.New("failed"))

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1).
			Return()

		result, code, err := userService.CreateUserInteraction(createUserInteractionRequest)

		require.Empty(t, result)
		require.Equal(t, http.StatusInternalServerError, code)
		require.NotNil(t, err)
	})

	t.Run("Should return error Internal Server Error if failed to create matches partner", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserById(gomock.Any(), userInput).
			Times(1).
			Return(userOutput, nil)

		mockUserRepository.EXPECT().
			UpdateSwipeCount(gomock.Any(), gomock.Any()).
			Times(1).
			Return(parameter.UpdateSwipeCountOutput{}, nil)

		mockUserRepository.EXPECT().
			CreateUserInteraction(gomock.Any(), interactionInput).
			Times(1).
			Return(interactionOutput, nil)

		mockUserRepository.EXPECT().
			FindMatchInteraction(gomock.Any(), reverseInteractionInput).
			Times(1).
			Return(reverseInteractionOutput, nil)

		mockUserRepository.EXPECT().
			CreateMatchUser(gomock.Any(), matchInput).
			Times(1).
			Return(parameter.CreateMatchUserOutput{}, errors.New("failed"))

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1).
			Return()

		result, code, err := userService.CreateUserInteraction(createUserInteractionRequest)

		require.Empty(t, result)
		require.Equal(t, http.StatusInternalServerError, code)
		require.NotNil(t, err)
	})
}
