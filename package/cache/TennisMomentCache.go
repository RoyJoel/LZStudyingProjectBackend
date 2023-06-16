package cache

import (
	"context"
	"time"

	"github.com/RoyJoel/LZStudyingProject/package/config"

	"github.com/go-redis/redis/v8"
)

type TennisMomentCacheDAOImpl struct {
	db *redis.Client
}

type TennisMomentCacheDAO interface {
	// set一个
	AddDeviceJwtMapping(client *redis.Client, deviceID string, jwt string) error
	// 根据ID获取一个
	GetJwtByDeviceID(ctx context.Context, client *redis.Client, deviceID string) (string, error)

	GetDeviceIDByJwt(ctx context.Context, client *redis.Client, jwt string) (string, error)
}

func NewTennisMomentCacheDAOImpl() *TennisMomentCacheDAOImpl {
	return &TennisMomentCacheDAOImpl{db: config.RDB}
}

// 把设备ID和JWT添加到映射表中，并设置过期时间
func AddDeviceJwtMapping(client *redis.Client, deviceID string, jwt string) error {
	expiration := time.Hour * 24 * 7 // 设置过期时间为7天
	err := client.Set(context.Background(), deviceID, jwt, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// 根据设备ID获取对应的JWT
func GetJwtByDeviceID(ctx context.Context, client *redis.Client, deviceID string) (string, error) {
	jwt, err := client.Get(ctx, deviceID).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return jwt, nil
}

// 根据JWT获取对应的设备ID
func GetDeviceIDByJwt(ctx context.Context, client *redis.Client, jwt string) (string, error) {
	deviceID, err := client.Get(ctx, jwt).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return deviceID, nil
}
