package global

import (
	"LibraryManagementSys/config"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

var (
	LMS_DB     *gorm.DB
	LMS_REDIS  *redis.Client
	LMS_CONFIG config.Server
	LMS_VP     *viper.Viper
	LMS_LOG    *zap.SugaredLogger // sugaredLogger可以通过printf格式记录语句
	Lock       sync.RWMutex
)
