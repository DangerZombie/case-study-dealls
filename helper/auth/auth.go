package auth

import (
	"github.com/DangerZombie/case-study-dealls/repository"
	"github.com/DangerZombie/case-study-dealls/repository/user_repository"
)

type authHelperImpl struct {
	baseRepo repository.BaseRepository
	userRepo user_repository.UserRepository
}

func NewAuthHelper(br repository.BaseRepository, ur user_repository.UserRepository) AuthHelper {
	return &authHelperImpl{br, ur}
}
