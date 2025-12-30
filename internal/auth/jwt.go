package auth

import (
	"context"
	"errors"
	"go-flash-sale/internal/config"
	myErr "go-flash-sale/internal/errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type MyClaims struct {
	UserId uint `json:"user_id"`
	// Username string `json:"username"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type JWTService interface {
	GenerateToken(userID uint, email string) (string, string, error)
	ValidateToken(ctx context.Context, tokenString string) (*MyClaims, error)
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() *jwtService {
	if len(config.JWTSecretKey) < 32 {
		panic("JWT secret key must be at least 32 characters")
	}
	return &jwtService{
		secretKey: config.JWTSecretKey,
		issuer:    config.JWTIssuer,
	}
}

func (j *jwtService) GenerateToken(userID uint, email string) (string, string, error) {
	jti := uuid.New().String()
	claims := &MyClaims{
		UserId: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 设置过期时间为24小时
			ID:        jti,                                                // 每个token一个唯一ID
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 以HS256加密方式生成token
	//
	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", "", myErr.ErrCreateTokenFailed
	}
	return tokenString, jti, nil
}

func (j *jwtService) ValidateToken(ctx context.Context, tokenString string) (*MyClaims, error) {
	// 解析token 并验证是否生效，是否过期，是否在合理时间范围内
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil {
		// 检查是否是 token 本身的错误（如过期、未生效）
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, myErr.ErrTokenExpired
		}
		if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, myErr.ErrTokenNotActive
		}
		// 其他错误（签名无效、格式错误等）
		return nil, myErr.ErrInvalidToken
	}
	claims, ok := token.Claims.(*MyClaims)
	if !ok {
		return nil, myErr.ErrInvalidToken
	}
	if !token.Valid {
		return nil, myErr.ErrInvalidToken
	}
	return claims, nil
}
