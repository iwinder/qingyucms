package service

import (
	"context"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/admin/v1"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
)

// CreateApi 创建用户
func (s *BlogAdminUserService) CreateQyAdminApi(ctx context.Context, in *v1.CreateQyAdminApiRequest) (*v1.CreateQyAdminApiReply, error) {
	objDO := &biz.ApiDO{
		ApiGroup:    in.ApiGroup,
		Identifier:  in.Identifier,
		Method:      in.Method,
		Path:        in.Path,
		Description: in.Description,
	}
	obj, err := s.ac.Create(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.CreateQyAdminApiReply{Id: obj.ID}, nil
}

// UpdateApi 更新用户
func (s *BlogAdminUserService) UpdateQyAdminApi(ctx context.Context, in *v1.UpdateQyAdminApiRequest) (*v1.UpdateQyAdminApiReply, error) {
	objDO := &biz.ApiDO{
		ObjectMeta:  metaV1.ObjectMeta{ID: in.Id},
		ApiGroup:    in.ApiGroup,
		Identifier:  in.Identifier,
		Method:      in.Method,
		Path:        in.Path,
		Description: in.Description,
	}
	obj, err := s.ac.Update(ctx, objDO)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateQyAdminApiReply{Id: obj.ID}, nil
}

// DeleteApi 根据ID删除用户
func (s *BlogAdminUserService) DeleteQyAdminApi(ctx context.Context, in *v1.DeleteQyAdminApiRequest) (*v1.DeleteQyAdminApiReply, error) {
	err := s.ac.Delete(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteQyAdminApiReply{}, nil
}

// DeleteApis 根据ID批量删除用户
func (s *BlogAdminUserService) DeleteQyAdminApis(ctx context.Context, in *v1.DeleteQyAdminApisRequest) (*v1.DeleteQyAdminApisReply, error) {
	err := s.ac.DeleteList(ctx, in.Ids)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteQyAdminApisReply{}, nil
}

// GetApi 根据ID获取用户信息
func (s *BlogAdminUserService) GetQyAdminApi(ctx context.Context, in *v1.GetQyAdminApiRequest) (*v1.GetQyAdminApiReply, error) {
	obj, err := s.ac.FindOneByID(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	u := bizToApiResponse(obj)
	return &v1.GetQyAdminApiReply{Info: &u}, nil
}

// ListApi 获取用户列表
func (s *BlogAdminUserService) ListQyAdminApi(ctx context.Context, in *v1.ListQyAdminApiRequest) (*v1.ListQyAdminApiReply, error) {
	opts := biz.ApiDOListOption{}
	opts.ListOptions.Pages = 0
	opts.ListOptions.Current = -1
	opts.ListOptions.PageSize = 20
	if in.PageInfo != nil {
		opts.ListOptions.Pages = int64(in.PageInfo.Pages)
		opts.ListOptions.Current = int64(in.PageInfo.Page)
		opts.ListOptions.PageSize = int64(in.PageInfo.Size)
	}

	opts.ListOptions.Init()
	objList, err := s.ac.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	pageInfo := &v1.APIPageInfo{
		Page:      uint64(objList.Pages),
		Size:      uint64(objList.PageSize),
		Total:     uint64(objList.TotalCount),
		Pages:     uint64(objList.Pages),
		FirstFlag: objList.FirstFlag,
		LastFlag:  objList.LastFlag,
	}
	objs := make([]*v1.ApiInfoResponse, 0, len(objList.Items))
	for _, item := range objList.Items {
		titem := bizToApiResponse(item)
		objs = append(objs, &titem)
	}
	return &v1.ListQyAdminApiReply{PageInfo: pageInfo, Items: objs}, nil
}
func bizToApiResponse(obj *biz.ApiDO) v1.ApiInfoResponse {
	objInfoRsp := v1.ApiInfoResponse{
		Id:          obj.ID,
		ApiGroup:    obj.ApiGroup,
		Identifier:  obj.Identifier,
		Method:      obj.Method,
		Path:        obj.Path,
		Description: obj.Description,
	}
	return objInfoRsp
}
