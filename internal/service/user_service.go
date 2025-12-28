package service

import (
	"context"
	"errors"
	"go-flash-sale/internal/auth"
	"go-flash-sale/internal/model"
	"go-flash-sale/internal/repository"
	"time"

	"github.com/go-sql-driver/mysql"
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
	err = s.userRepo.Create(user)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 { //捕获唯一键冲突错误，返回“邮箱已存在”
			return errors.New("email already registered")
		}
		return err
	}
	return nil
}

func (s *UserService) Login(email, password string) (string, error) {
	//从用户仓库查询用户信息 后改为从数据库查询
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	//验证成功 签发token
	jwtService := auth.NewJWTService()
	tokenString, tokenID, err := jwtService.GenerateToken(user.ID, email)
	if err != nil {
		return "", err
	}
	//将tokenID存入redis
	tokenStore := auth.NewTokenStore()
	err = tokenStore.Save(context.Background(), tokenID, user.ID, 24*time.Hour)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
