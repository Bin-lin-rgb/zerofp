package logic

import (
	"context"
	"zerofp/models"
	"zerofp/pkg"

	"zerofp/admin/internal/svc"
	"zerofp/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductListLogic) ProductList(req *types.ProductListRequst) (resp *types.ProductListReply, err error) {
	list := make([]*types.ProductListBaisc, 0)
	req.Size = pkg.If(req.Size == 0, 20, req.Size).(int)
	req.Page = pkg.If(req.Page == 0, 0, (req.Page-1)*req.Size).(int)
	resp = new(types.ProductListReply)
	var count int64

	err = models.ProductList(req.Name).Count(&count).Offset(req.Page).Limit(req.Size).Find(&list).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return
	}
	for _, v := range list {
		v.CreatedAt = pkg.RFC3339ToNormalTime(v.CreatedAt)
	}

	resp.Count = count
	resp.List = list

	return
}
