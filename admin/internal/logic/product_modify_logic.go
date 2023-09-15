package logic

import (
	"context"
	"zerofp/models"

	"zerofp/admin/internal/svc"
	"zerofp/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductModifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductModifyLogic {
	return &ProductModifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductModifyLogic) ProductModify(req *types.ProductModifyRequest) (resp *types.BaseReply, err error) {
	resp = &types.BaseReply{
		Code: 0,
		Msg:  "success",
	}

	err = l.svcCtx.DB.Where("identity = ?", req.Identity).Updates(&models.Product{
		Name: req.Name,
		Desc: req.Desc,
	}).Error
	if err != nil {
		resp.Code = 500
		resp.Msg = "修改失败，请稍后重试"
		logx.Error("[DB ERROR] : ", err)
	}

	return
}
