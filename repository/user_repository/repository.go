package user_repository

import "github.com/DangerZombie/case-study-dealls/repository"

type userRepo struct {
	base repository.BaseRepository
}

func NewUserRepository(br repository.BaseRepository) UserRepository {
	return &userRepo{br}
}
