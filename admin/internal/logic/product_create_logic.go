package logic

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"zerofp/models"

	"zerofp/admin/internal/svc"
	"zerofp/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCreateLogic {
	return &ProductCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCreateLogic) ProductCreate(req *types.ProductCreateRequest) (resp *types.BaseReply, err error) {
	resp = &types.BaseReply{
		Code: 0,
		Msg:  "success",
	}
	// 如果名字重复，返回错误
	res := l.svcCtx.DB.Where("name = ?", req.Name).First(&models.Product{})
	if res.RowsAffected > 0 {
		resp.Code = 600
		resp.Msg = "不允许新增重复项"
		err = errors.New("不允许新增重复项")
		return
	}
	err = l.svcCtx.DB.Create(&models.Product{
		Identity: uuid.New().String(),
		Key:      uuid.New().String(),
		Name:     req.Name,
		Desc:     req.Desc,
	}).Error
	if err != nil {
		resp.Code = 500
		resp.Msg = "创建失败，请稍后重试"
		err = errors.New("创建失败，请稍后重试")
		logx.Error("[DB ERROR] : ", err)
		return
	}

	return
}
