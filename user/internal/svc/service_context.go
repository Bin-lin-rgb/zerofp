package svc

import (
	"gorm.io/gorm"
	"zerofp/models"
	"zerofp/user/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.InitDB()
	return &ServiceContext{
		Config: c,
		DB:     models.DB,
	}
}
