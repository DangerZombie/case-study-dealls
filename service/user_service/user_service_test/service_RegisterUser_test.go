package user_service_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/DangerZombie/case-study-dealls/helper/auth"
	"github.com/DangerZombie/case-study-dealls/helper/static"
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

func TestRegisterUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)
	mockBaseRepository := repository.NewMockBaseRepository(mockCtrl)
	mockUserRepository := user_repository.NewMockUserRepository(mockCtrl)

	userService := user_service.NewUserService(
		mockAuthHelper,
		mockBaseRepository,
		mockUserRepository,
	)

	registerUserRequest := request.RegisterUserRequestBody{
		Username: faker.Name(),
		Password: faker.Password(),
		Nickname: faker.Name(),
		Gender:   faker.Gender(),
		Age:      30,
		Location: faker.Sentence(),
	}

	createUserInput := parameter.CreateUserInput{
		User: entity.User{
			Username:   registerUserRequest.Username,
			Password:   registerUserRequest.Password,
			Gender:     registerUserRequest.Gender,
			Age:        registerUserRequest.Age,
			Location:   registerUserRequest.Location,
			Nickname:   registerUserRequest.Nickname,
			Status:     static.UserFree,
			SwipeCount: 10,
			Verified:   false,
		},
	}

	t.Run("Should return OK", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			CreateUser(gomock.Any(), createUserInput).
			Times(1).
			Return(parameter.CreateUserOutput{}, nil)

		mockBaseRepository.EXPECT().
			BeginCommit(gomock.Any()).
			Times(1).
			Return()

		result, code, err := userService.RegisterUser(registerUserRequest)

		require.Equal(t, "Success", result.Message)
		require.Equal(t, http.StatusOK, code)
		require.Nil(t, err)
	})

	t.Run("Should return error Internal Server Error if failed to create user", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			CreateUser(gomock.Any(), createUserInput).
			Times(1).
			Return(parameter.CreateUserOutput{}, errors.New("failed"))

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1).
			Return()

		result, code, err := userService.RegisterUser(registerUserRequest)

		require.Empty(t, result.Message)
		require.Equal(t, http.StatusInternalServerError, code)
		require.NotNil(t, err)
	})
}
