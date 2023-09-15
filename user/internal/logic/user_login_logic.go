package logic

import (
	"context"
	"errors"
	"zerofp/models"
	"zerofp/pkg"

	"zerofp/user/internal/svc"
	"zerofp/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginRequest) (resp *types.UserLoginReply, err error) {
	resp = new(types.UserLoginReply)
	ub := new(models.User)
	err = l.svcCtx.DB.Where("name = ? AND password = ?",
		req.Username, pkg.Md5(req.Password)).First(ub).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		err = errors.New("用户名或密码不正确")
		return
	}
	token, err := pkg.GenerateToken(ub.ID, ub.Identity, ub.Name, 3600*24*30)
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		err = errors.New("用户名或密码不正确")
		return
	}
	resp.Token = token
	return
}
