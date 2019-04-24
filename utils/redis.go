package utils

import (
	"github.com/go-redis/redis"
	"math/rand"
	"time"
)

type RedisConn struct {
	Client *redis.Client
	convertor *Convertor
}
const idKey = "id"
const startId = 14776336
const fixedExp = 60
const randomExp = 120

func NewClient(addr, password string, db int, largestKeyword string) (*RedisConn) {
	redisConn := new(RedisConn)
	redisConn.Client = redis.NewClient(&redis.Options{
		Addr:		addr,
		Password: 	password,
		DB:			db,
	})
	redisConn.convertor = NewConvertor()
	if "" == largestKeyword {
		redisConn.Client.Set(idKey, startId, 0)
	} else {
		redisConn.Client.Set(idKey, redisConn.convertor.ConvertToNum(largestKeyword), 0)
	}
	return redisConn
}
/**
从redis中查询短键对应的长网址
*/
func (redisConn *RedisConn) GetUrlFromCache(keyword string) (string, bool) {
	url, err := redisConn.Client.Get(keyword).Result()
	if err == redis.Nil {
		return "", false
	}
	return url, true
}
/**
设置redis的键值对
*/
func (redisConn *RedisConn) SetCache(keyword, url string) bool {
	if err := redisConn.Client.Set(keyword, url, time.Duration(fixedExp + rand.Intn(randomExp))*time.Second).Err(); err != nil {
		panic(err)
		return false
	}
	return true
}
/**
自增id
 */
func (redisConn *RedisConn) NextKeyword() string {
	id, _ := redisConn.Client.Incr(idKey).Result()
	return redisConn.convertor.ConvertToString(uint(id))
}