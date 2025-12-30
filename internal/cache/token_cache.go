package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// TokenCache 管理 JWT Token 的 Redis 存储
// 用于支持 Token 主动失效（登出、踢下线等）
type TokenCache struct {
	redisClient redis.UniversalClient // 使用 UniversalClient 支持单机/集群
}

// NewTokenStore 创建一个新的 TokenStore 实例
// client: 已配置好的 Redis 客户端（建议全局复用）
func NewTokenCache(client redis.UniversalClient) *TokenCache {
	return &TokenCache{
		redisClient: client,
	}
}

// Save 将 TokenID 与 UserID 的映射存入 Redis
// ttl: Token 有效期（应与 JWT 的过期时间一致）
func (ts *TokenCache) Save(ctx context.Context, tokenID string, userID uint, ttl time.Duration) error {
	if tokenID == "" || userID == 0 {
		return fmt.Errorf("invalid tokenID or userID")
	}
	key := ts.buildKey(tokenID)
	value := fmt.Sprintf("%d", userID)
	err := ts.redisClient.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return fmt.Errorf("failed to save token to redis: %w", err)
	}
	return nil
}

func (ts *TokenCache) Exists(ctx context.Context, tokenID string) (bool, error) {
	if tokenID == "" {
		return false, fmt.Errorf("tokenID is empty")
	}
	key := ts.buildKey(tokenID)
	// EXISTS 返回 1 表示存在，0 表示不存在
	exists, err := ts.redisClient.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check token in redis: %w", err)
	}
	if exists == 0 {
		return false, fmt.Errorf("token does not exist")
	}
	return true, nil
}

func (ts *TokenCache) Delete(ctx context.Context, tokenID string) error {
	if tokenID == "" {
		return fmt.Errorf("tokenID is empty")
	}
	key := ts.buildKey(tokenID)
	err := ts.redisClient.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete token from redis: %w", err)
	}
	return nil
}

// 通过 TokenID 获取对应的 UserID
func (ts *TokenCache) GetUserID(ctx context.Context, tokenID string) (uint, error) {
	if tokenID == "" {
		return 0, fmt.Errorf("tokenID is empty")
	}
	key := ts.buildKey(tokenID)
	value, err := ts.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, fmt.Errorf("token does not exist")
		}
		return 0, fmt.Errorf("failed to get token from redis: %w", err)
	}
	var userID uint
	_, err = fmt.Sscanf(value, "%d", &userID)
	if err != nil {
		return 0, fmt.Errorf("invalid userID format in redis: %w", err)
	}
	return userID, nil
}

// RefreshTTL 刷新 Token 在 Redis 中的过期时间
func (ts *TokenCache) RefreshTTL(ctx context.Context, tokenID string, ttl time.Duration) error {
	if tokenID == "" {
		return fmt.Errorf("tokenID is empty")
	}
	key := ts.buildKey(tokenID)
	err := ts.redisClient.Expire(ctx, key, ttl).Err()
	if err != nil {
		return fmt.Errorf("failed to refresh token TTL in redis: %w", err)
	}
	return nil
}

// DeleteByUserID 根据 UserID 删除所有相关的 Token
func (ts *TokenCache) DeleteByUserID(ctx context.Context, userID uint) error {
	if userID == 0 {
		return fmt.Errorf("invalid userID")
	}
	pattern := "auth:token:*"
	iter := ts.redisClient.Scan(ctx, 0, pattern, 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		value, err := ts.redisClient.Get(ctx, key).Result()
		if err != nil {
			continue
		}
		var uid uint
		_, err = fmt.Sscanf(value, "%d", &uid)
		if err != nil {
			continue
		}
		if uid == userID {
			err = ts.redisClient.Del(ctx, key).Err()
			if err != nil {
				return fmt.Errorf("failed to delete token for userID %d: %w", userID, err)
			}
		}
	}
	if err := iter.Err(); err != nil {
		return fmt.Errorf("error during scanning tokens: %w", err)
	}
	return nil
}

// buildKey 构建 Redis Key（统一前缀，避免冲突）
func (ts *TokenCache) buildKey(tokenID string) string {
	return "auth:token:" + tokenID
}
