package app

import (
	"go-flash-sale/internal/cache"
	"go-flash-sale/internal/initialization"
	"go-flash-sale/internal/repository"
	"go-flash-sale/internal/service"

	"github.com/redis/go-redis/v9"
)

type Dependencies struct {
	// Add fields for dependencies here, e.g. database connections, caches, etc.
	RedisClient redis.UniversalClient //reids依赖
	UserService *service.UserService  //用户服务依赖
}

func BuildDependencies() (*Dependencies, error) {
	// 初始化数据库链接
	db := initialization.InitDB()
	// 初始化Redis链接
	redisClient := initialization.InitRedis()
	// 自动迁移模式，创建或更新表结构
	initialization.InitTableAutoMigrate(db)
	// 初始化Token依赖
	tokenCache := cache.NewTokenCache(redisClient)

	//初始化业务服务
	userRepo := repository.NewUserRepository(db)
	userCache := cache.NewUserCache(redisClient)
	userService := service.NewUserService(userRepo, userCache, tokenCache)

	return &Dependencies{
		RedisClient: redisClient,
		UserService: userService,
	}, nil
}
