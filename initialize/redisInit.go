package initialize

import (
	"LibraryManagementSys/global"
	"context"
	"github.com/go-redis/redis/v9"
	"os"
)

func RedisInit() {
	redisConf := global.LMS_CONFIG.Redis

	rDB := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr,
		Password: redisConf.Password, // no password set
		DB:       redisConf.Db,       // use default DB
	})
	pong, err := rDB.Ping(context.Background()).Result()
	if err != nil {
		global.LMS_LOG.Errorf("redis connect ping failed, err: %s", err)
		os.Exit(0)
	} else {
		global.LMS_LOG.Infof("redis connect ping response: %s", pong)
		// 设置全局
		global.LMS_REDIS = rDB
	}
}
