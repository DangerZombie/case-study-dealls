package auth_test

import (
	"errors"
	"testing"

	"github.com/DangerZombie/case-study-dealls/helper/auth"
	"github.com/DangerZombie/case-study-dealls/helper/static"
	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"github.com/DangerZombie/case-study-dealls/repository"
	"github.com/DangerZombie/case-study-dealls/repository/user_repository"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestAuthGenerateJWT(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockBaseRepository := repository.NewMockBaseRepository(mockCtrl)
	mockUserRepository := user_repository.NewMockUserRepository(mockCtrl)

	authHelper := auth.NewAuthHelper(
		mockBaseRepository,
		mockUserRepository,
	)

	id := faker.UUIDHyphenated()
	findUserByIdInput := parameter.FindUserByIdInput{
		Id: id,
	}

	findUserByIdOutput := parameter.FindUserByIdOutput{
		Id:         id,
		Username:   faker.Username(),
		Status:     static.UserFree,
		Nickname:   faker.Name(),
		SwipeCount: 10,
	}

	t.Run("Should return token", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserById(gomock.Any(), findUserByIdInput).
			Times(1).
			Return(findUserByIdOutput, nil)

		mockBaseRepository.EXPECT().
			BeginCommit(gomock.Any()).
			Times(1)

		result, err := authHelper.GenerateJWT(id)

		require.NotEmpty(t, result)
		require.Empty(t, err)
	})

	t.Run("Should return error if failed to fetch user", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserById(gomock.Any(), findUserByIdInput).
			Times(1).
			Return(parameter.FindUserByIdOutput{}, errors.New("failed"))

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1)

		result, err := authHelper.GenerateJWT(id)

		require.Empty(t, result)
		require.NotEmpty(t, err)
	})
}
