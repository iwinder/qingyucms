package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/app/qycms_comments/internal/biz"

	"github.com/iwinder/qingyucms/app/qycms_comments/internal/data/po"
)

type CommentAgentRepo struct {
	data *Data
	log  *log.Helper
}

// NewCommentIndexRepo .
func NewCommentAgentRepo(data *Data, logger log.Logger) biz.CommentAgentRepo {
	return &CommentAgentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *CommentAgentRepo) Save(ctx context.Context, g *biz.CommentAgentDO) (*po.CommentAgentPO, error) {
	newData := &po.CommentAgentPO{
		ObjectMeta: g.ObjectMeta,
		ObjId:      g.ObjId,
		ObjType:    g.ObjType,
		MemberId:   g.MemberId,
		Count:      g.Count,
		RootCount:  g.RootCount,
		AllCount:   g.AllCount,
		State:      g.State,
		Attrs:      g.Attrs,
	}
	err := r.data.db.Create(newData).Error
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func (r *CommentAgentRepo) Update(ctx context.Context, g *biz.CommentAgentDO) (*po.CommentAgentPO, error) {
	newData := &po.CommentAgentPO{
		ObjId:     g.ObjId,
		ObjType:   g.ObjType,
		MemberId:  g.MemberId,
		Count:     g.Count,
		RootCount: g.RootCount,
		AllCount:  g.AllCount,
		State:     g.State,
		Attrs:     g.Attrs,
	}
	tData := &po.CommentAgentPO{}
	tData.ID = g.ID
	err := r.data.db.Model(&tData).Updates(&newData).Error
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func (r *CommentAgentRepo) FindByID(cxt context.Context, id uint64) (*po.CommentAgentPO, error) {
	tData := &po.CommentAgentPO{}
	err := r.data.db.Where("id = ?", id).First(&tData).Error
	if err != nil {
		return nil, err
	}

	return tData, nil
}
