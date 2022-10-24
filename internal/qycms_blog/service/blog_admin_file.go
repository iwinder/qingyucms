package service

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/admin/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
)

func (s *BlogAdminUserService) CreateQyAdminFileLibType(ctx context.Context, in *v1.CreateQyAdminFileLibTypeRequest) (*v1.UpdateQyAdminFileLibTypeReply, error) {
	dataDO := &biz.FileLibTypeDO{
		Name:       in.Name,
		Identifier: int(in.Identifier),
		Ftype:      in.Type,
	}
	dataDO.StatusFlag = int(in.StatusFlag)
	ret, err := s.fit.Save(ctx, dataDO)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateQyAdminFileLibTypeReply{Id: ret.ID}, err
}

func (s *BlogAdminUserService) UpdateQyAdminFileLibType(ctx context.Context, in *v1.UpdateQyAdminFileLibTypeRequest) (*v1.UpdateQyAdminFileLibTypeReply, error) {
	dataDO := &biz.FileLibTypeDO{
		Name:       in.Name,
		Identifier: int(in.Identifier),
		Ftype:      in.Type,
	}
	dataDO.ID = in.Id
	dataDO.StatusFlag = int(in.StatusFlag)
	ret, err := s.fit.Update(ctx, dataDO)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateQyAdminFileLibTypeReply{Id: ret.ID}, err
}
func (s *BlogAdminUserService) DeleteQyAdminFileLibType(ctx context.Context, in *v1.DeleteQyAdminFileLibTypeRequest) (*v1.DeleteQyAdminFileLibTypeReply, error) {
	err := s.fit.DeleteList(ctx, in.Ids)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteQyAdminFileLibTypeReply{}, nil
}
func (s *BlogAdminUserService) ListQyAdminFileLibType(ctx context.Context, in *v1.ListQyAdminFileLibTypeRequest) (*v1.ListQyAdminFileLibTypeReply, error) {
	opts := biz.FileLibTypeDOListOption{}
	opts.ListOptions.Pages = 0
	opts.ListOptions.Current = -1
	opts.ListOptions.PageSize = 20
	if in.Current > 0 {
		opts.ListOptions.Pages = in.Pages
		opts.ListOptions.Current = in.Current
		opts.ListOptions.PageSize = in.PageSize
	}
	opts.Name = in.Name
	opts.StatusFlag = int(in.StatusFlag)
	opts.ListOptions.Init()
	objList, err := s.fit.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	pageInfo := &v1.FilePageInfo{
		Current:   objList.Current,
		PageSize:  objList.PageSize,
		Total:     objList.TotalCount,
		Pages:     objList.Pages,
		FirstFlag: objList.FirstFlag,
		LastFlag:  objList.LastFlag,
	}
	objs := make([]*v1.FileLibTypeResponse, 0, len(objList.Items))
	for _, item := range objList.Items {
		titem := bizToFileLibTypeResponse(item)
		objs = append(objs, &titem)
	}
	return &v1.ListQyAdminFileLibTypeReply{PageInfo: pageInfo, Items: objs}, nil
}

func (s *BlogAdminUserService) CreateQyAdminFileLibConfig(ctx context.Context, in *v1.CreateQyAdminFileLibConfigRequest) (*v1.CreateQyAdminFileLibConfigReply, error) {
	dataDO := &biz.FileLibConfigDO{
		AccessKey: in.AccessKey,
		SecretKey: in.SecretKey,
		Bucket:    in.Bucket,
		Prefix:    in.Prefix,
		Domain:    in.Domain,
		Endpoint:  in.Endpoint,
		TypeId:    in.TypeId,
	}
	if in.Id > 0 {
		dataDO.ID = in.Id
	}

	ret, err := s.fic.SaveOrUpdate(ctx, dataDO)
	if err != nil {
		return nil, err
	}
	return &v1.CreateQyAdminFileLibConfigReply{Id: ret.ID}, nil
}

func (s *BlogAdminUserService) GetQyAdminFileLibConfig(ctx context.Context, in *v1.GetQyAdminFileLibConfigRequest) (*v1.GetQyAdminFileLibConfigReply, error) {
	data, err := s.fic.FindByTypeId(ctx, in.TypeId)
	if err != nil {
		return nil, err
	}
	return &v1.GetQyAdminFileLibConfigReply{
		Id:        data.ID,
		AccessKey: data.AccessKey,
		SecretKey: data.SecretKey,
		Bucket:    data.Bucket,
		Prefix:    data.Prefix,
		Domain:    data.Domain,
		Endpoint:  data.Endpoint,
		TypeId:    data.TypeId,
	}, nil
}

func bizToFileLibTypeResponse(obj *biz.FileLibTypeDO) v1.FileLibTypeResponse {
	objInfoRsp := v1.FileLibTypeResponse{
		Id:         obj.ID,
		Name:       obj.Name,
		Identifier: int32(obj.Identifier),
		Type:       obj.Ftype,
		StatusFlag: int32(obj.StatusFlag),
	}
	return objInfoRsp
}
