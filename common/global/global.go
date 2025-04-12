package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Rds *redis.Client
)
