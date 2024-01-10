package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"new-fix/settings"
	"new-fix/types"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Host     string
	Port     string
	Database int
	Password string
	Client   *redis.Client
}

func NewRedis() *Redis {
	return &Redis{}
}

func (receiver *Redis) Connect() *redis.Client {
	config := settings.NewSetting()

	receiver.Host = config.DB.Section("redis").Key("host").MustString("127.0.0.1")
	receiver.Port = config.DB.Section("redis").Key("port").MustString("6379")
	receiver.Database = config.DB.Section("redis").Key("database").MustInt(0)
	receiver.Password = config.DB.Section("redis").Key("password").String()

	receiver.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", receiver.Host, receiver.Port),
		Password: receiver.Password,
		DB:       receiver.Database,
	})

	return receiver.Client
}

func (receiver *Redis) Disconnect() {
	err := receiver.Client.Close()
	if err != nil {
		log.Printf("redis 关闭失败：%s", err.Error())
	}
}

// SetValue 设置值
func (receiver Redis) SetValue(key string, value interface{}, exp time.Duration) (any, error) {
	val, err := receiver.Connect().Set(context.Background(), key, value, exp).Result()
	receiver.Disconnect()
	return val, err
}

// SetJsonValue 设置值（JSON）
func (receiver Redis) SetJsonValue(key string, value any, exp time.Duration) (any, error) {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	val, err := receiver.Connect().Set(context.Background(), key, string(jsonValue), exp).Result()
	receiver.Disconnect()
	return val, err
}

// GetValue 获取值
func (receiver *Redis) GetValue(key string) (any, error) {
	value, err := receiver.Connect().Get(context.Background(), key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return "", err
	}
	receiver.Disconnect()
	return value, nil
}

// GetSort 获取列表
func (receiver *Redis) GetSort(key string, offset, count int64, order string) ([]string, error) {
	values, err := receiver.Connect().Sort(context.Background(), key, &redis.Sort{Offset: offset, Count: count, Order: order}).Result()
	receiver.Disconnect()
	return values, err
}

// GetZRange 获取字典
func (receiver *Redis) GetZRange(key, min, max string, offset, count int64) ([]redis.Z, error) {
	values, err := receiver.Connect().ZRangeByScoreWithScores(context.Background(), key, &redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
		Count:  count,
	}).Result()
	receiver.Disconnect()
	return values, err
}

// GetZInterStore 获取字典
func (receiver *Redis) GetZInterStore(keys []string, weights types.ListFloat64) (int64, error) {
	values, err := receiver.Connect().ZInterStore(context.Background(), "out", &redis.ZStore{
		Keys:    keys,
		Weights: weights,
	}).Result()
	receiver.Disconnect()
	return values, err
}

// Do 执行命令
func (receiver *Redis) Do(command, key string, value any) (any, error) {
	res, err := receiver.Connect().Do(context.Background(), command, key, value).Result()
	receiver.Disconnect()
	return res, err
}
