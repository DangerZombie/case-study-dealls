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

func TestLogin(t *testing.T) {
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
	username := faker.Name()
	password := faker.Password()
	token := faker.Jwt()
	loginRequest := request.LoginRequestBody{
		Username: username,
		Password: password,
	}

	loginEmptyRequest := request.LoginRequestBody{
		Username: "",
		Password: "",
	}

	findUserByUsernameAndPasswordInput := parameter.FindUserByUsernameAndPasswordInput{
		Username: username,
		Password: password,
	}

	findUserByUsernameAndPasswordOutput := parameter.FindUserByUsernameAndPasswordOutput{
		BaseModel: base.BaseModel{
			Id: id,
		},
		Username: username,
		Password: password,
	}

	t.Run("Should return OK", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserByUsernameAndPassword(gomock.Any(), findUserByUsernameAndPasswordInput).
			Times(1).
			Return(findUserByUsernameAndPasswordOutput, nil)

		mockAuthHelper.EXPECT().
			GenerateJWT(id).
			Times(1).
			Return(token, nil)

		mockBaseRepository.EXPECT().
			BeginCommit(gomock.Any()).
			Times(1).
			Return()

		result, code, err := userService.Login(loginRequest)

		require.Equal(t, token, result.Token)
		require.Equal(t, http.StatusOK, code)
		require.Empty(t, err)
	})

	t.Run("Should return error Bad Request if username or password is empty", func(t *testing.T) {
		result, code, err := userService.Login(loginEmptyRequest)

		require.Empty(t, result.Token)
		require.Equal(t, http.StatusBadRequest, code)
		require.NotEmpty(t, err)
	})

	t.Run("Should return error Internal Server Error if failed to find user", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserByUsernameAndPassword(gomock.Any(), findUserByUsernameAndPasswordInput).
			Times(1).
			Return(parameter.FindUserByUsernameAndPasswordOutput{}, errors.New("failed"))

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1).
			Return()

		result, code, err := userService.Login(loginRequest)

		require.Empty(t, result.Token)
		require.Equal(t, http.StatusInternalServerError, code)
		require.NotEmpty(t, err)
	})

	t.Run("Should return error Internal Server Error if failed to generate token", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserByUsernameAndPassword(gomock.Any(), findUserByUsernameAndPasswordInput).
			Times(1).
			Return(findUserByUsernameAndPasswordOutput, nil)

		mockAuthHelper.EXPECT().
			GenerateJWT(id).
			Times(1).
			Return("", errors.New("failed"))

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1).
			Return()

		result, code, err := userService.Login(loginRequest)

		require.Empty(t, result.Token)
		require.Equal(t, http.StatusInternalServerError, code)
		require.NotEmpty(t, err)
	})
}
