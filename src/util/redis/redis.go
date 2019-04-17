package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"tutor_platform/src/config"
	"tutor_platform/src/util/log"
)

var tokenClient *redis.Client
var messageClient *redis.Client

func init() {
	tokenClient = redis.NewClient(&redis.Options{
		Addr:     config.TokenRedisConf.Addr,
		Password: config.TokenRedisConf.Password,
		DB:       config.TokenRedisConf.DB,
		PoolSize: 5,
	})

	messageClient = redis.NewClient(&redis.Options{
		Addr:     config.MessageRedisConf.Addr,
		Password: config.MessageRedisConf.Password,
		DB:       config.MessageRedisConf.DB,
		PoolSize: 5,
	})

	pong, err := tokenClient.Ping().Result()
	if err != nil {
		log.Logger.Error("REDIS", zap.String("e", err.Error()))
	} else {
		log.Logger.Info("REDIS", zap.String("url", config.TokenRedisConf.Addr), zap.String("ping", pong))
	}
}

func GetClient() *redis.Client {
	return tokenClient
}

func AddValue(key string, val string) {
	tokenClient.Append(key, val)
	fmt.Println(key, val)
}

func UpdateValue(key, val string) error {
	err := tokenClient.Set(key, val, -1).Err()
	if err != nil {
		panic(err)
	}
	return err
}

func GetValue(key string) string {
	return tokenClient.Get(key).Val()
}

func DelValue(key string) {
	tokenClient.Del(key)
}

func SAdd(key, value string) {
	tokenClient.SAdd(key, value)
}

func SDel(key, value string) {
	tokenClient.SRem(key, value)
}

func SMembers(key string) []string {
	val, err := tokenClient.SMembers(key).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func SJudgeExist(key string, val string) bool {
	ok, err := tokenClient.SIsMember(key, val).Result()
	if err != nil {
		panic(err)
	}
	return ok
}

func ZAdd(key, member string, score int) error {
	err := messageClient.ZAdd(key, redis.Z{Score: float64(score), Member: member}).Err()
	if err != nil {
		panic(err)
	}
	return nil
}

func ZDel(key, member string) error {
	err := messageClient.ZRem(key, member).Err()
	if err != nil {
		panic(err)
	}
	return err
}

func ZRange(key string, num int) map[string]int {
	res, err := messageClient.ZRange(key, 0, int64(num-1)).Result()
	if err != nil {
		return nil
	}
	memberScoreMap := make(map[string]int)
	for _, member := range res {
		score, err := messageClient.ZScore(key, member).Result()
		if err != nil {
			return nil
		}
		memberScoreMap[member] = int(score)
	}
	return memberScoreMap
}

func ZAddScore(key, member string) error {
	return messageClient.ZIncrBy(key, 1, member).Err()
}

func ZRevRange(key string, num int) map[string]int {
	res, err := messageClient.ZRevRange(key, 0, int64(num-1)).Result()
	if err != nil {
		return nil
	}
	memberScoreMap := make(map[string]int)
	for _, member := range res {
		score, err := messageClient.ZScore(key, member).Result()
		if err != nil {
			return nil
		}
		memberScoreMap[member] = int(score)
	}
	return memberScoreMap
}
