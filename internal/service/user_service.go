package service

import (
	"errors"
	"go-flash-sale/internal/model"
	"go-flash-sale/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Register(email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &model.User{
		Email:    email,
		Password: string(hashedPassword), // 注意：实际应用中应对密码进行哈希处理
	}
	return s.userRepo.Create(user)
}

func (s *UserService) Login(email, password string) (bool, error) {
	//从用户仓库查询用户信息 后改为从数据库查询
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return false, errors.New("invalid email or password")
	}
	//
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, errors.New("invalid email or password")
	}
	return true, nil
}
