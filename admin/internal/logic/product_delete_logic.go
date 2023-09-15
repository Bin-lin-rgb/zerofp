package logic

import (
	"context"
	"errors"
	"zerofp/models"

	"zerofp/admin/internal/svc"
	"zerofp/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDeleteLogic {
	return &ProductDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductDeleteLogic) ProductDelete(req *types.ProductDeleteRequest) (resp *types.BaseReply, err error) {
	resp = &types.BaseReply{
		Code: 0,
		Msg:  "success",
	}

	var cnt int64
	err = l.svcCtx.DB.Model(new(models.Device)).Where("product_identity = ?", req.Identity).
		Count(&cnt).Error
	if err != nil {
		resp.Code = 500
		resp.Msg = "删除失败，请稍后重试"
		logx.Error("[DB ERROR] : ", err)
		return
	}
	if cnt > 0 {
		resp.Code = 500
		resp.Msg = "已关联设备，不可删除"
		err = errors.New("已关联设备，不可删除")
		return
	}

	err = l.svcCtx.DB.Debug().Where("identity = ?", req.Identity).Delete(new(models.Product)).Error
	if err != nil {
		resp.Code = 500
		resp.Msg = "删除失败，请稍后重试"
		logx.Error("[DB ERROR] : ", err)
	}
	return
}
