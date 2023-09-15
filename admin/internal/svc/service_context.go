package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"zerofp/admin/internal/config"
	"zerofp/models"
	user2 "zerofp/user/rpc/types/user"
	user "zerofp/user/rpc/user_client"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	RpcUser  user.User
	AuthUser *user2.UserAuthReply
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.InitDB()
	return &ServiceContext{
		Config:  c,
		DB:      models.DB,
		RpcUser: user.NewUser(zrpc.MustNewClient(c.Z)),
	}
}
