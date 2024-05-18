package user_service

import (
	"github.com/DangerZombie/case-study-dealls/helper/auth"
	"github.com/DangerZombie/case-study-dealls/repository"
	"github.com/DangerZombie/case-study-dealls/repository/user_repository"
)

type userServiceImpl struct {
	authHelper auth.AuthHelper
	baseRepo   repository.BaseRepository
	userRepo   user_repository.UserRepository
}

func NewUserService(
	ah auth.AuthHelper,
	br repository.BaseRepository,
	ur user_repository.UserRepository,
) UserService {
	return &userServiceImpl{ah, br, ur}
}
