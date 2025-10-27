package service

import "github.com/trannhutphat23/Go-Backend-EcommerceApp-API/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetInfoUserService() string {
	return us.userRepo.GetInfoUser()
}