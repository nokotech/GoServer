package redis

import "github.com/garyburd/redigo/redis"

type Redis struct {
	conn redis.Conn
}

var sharedInstance *Redis = newRedis()

func newRedis() *Redis {
	c, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		panic(err)
	}
	return &Redis{c}
}

func GetInstance() *Redis {
	return sharedInstance
}

func (this *Redis) Set(key string, value string) {
	this.conn.Do("SET", key, value)
}

func (this *Redis) Get(key string) string {
	word, err := redis.String(this.conn.Do("GET", key))
	if err != nil {
		panic(err)
	}
	return word
}
