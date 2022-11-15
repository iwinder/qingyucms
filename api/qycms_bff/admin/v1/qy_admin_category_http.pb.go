// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.1
// - protoc             v3.21.1
// source: api/qycms_bff/admin/v1/qy_admin_category.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationQyAdminCategoryCreateQyAdminCategory = "/api.qycms_bff.admin.v1.QyAdminCategory/CreateQyAdminCategory"
const OperationQyAdminCategoryDeleteQyAdminCategory = "/api.qycms_bff.admin.v1.QyAdminCategory/DeleteQyAdminCategory"
const OperationQyAdminCategoryListQyAdminCategory = "/api.qycms_bff.admin.v1.QyAdminCategory/ListQyAdminCategory"
const OperationQyAdminCategoryUpdateQyAdminCategory = "/api.qycms_bff.admin.v1.QyAdminCategory/UpdateQyAdminCategory"

type QyAdminCategoryHTTPServer interface {
	CreateQyAdminCategory(context.Context, *CreateQyAdminCategoryRequest) (*CreateQyAdminCategoryReply, error)
	DeleteQyAdminCategory(context.Context, *DeleteQyAdminCategoryRequest) (*DeleteQyAdminCategoryReply, error)
	ListQyAdminCategory(context.Context, *ListQyAdminCategoryRequest) (*ListQyAdminCategoryReply, error)
	UpdateQyAdminCategory(context.Context, *UpdateQyAdminCategoryRequest) (*UpdateQyAdminCategoryReply, error)
}

func RegisterQyAdminCategoryHTTPServer(s *http.Server, srv QyAdminCategoryHTTPServer) {
	r := s.Route("/")
	r.POST("/api/admin/v1/category", _QyAdminCategory_CreateQyAdminCategory0_HTTP_Handler(srv))
	r.PUT("/api/admin/v1/category/{id}", _QyAdminCategory_UpdateQyAdminCategory0_HTTP_Handler(srv))
	r.DELETE("/api/admin/v1/category", _QyAdminCategory_DeleteQyAdminCategory0_HTTP_Handler(srv))
	r.GET("/api/admin/v1/category", _QyAdminCategory_ListQyAdminCategory0_HTTP_Handler(srv))
}

func _QyAdminCategory_CreateQyAdminCategory0_HTTP_Handler(srv QyAdminCategoryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateQyAdminCategoryRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminCategoryCreateQyAdminCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateQyAdminCategory(ctx, req.(*CreateQyAdminCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateQyAdminCategoryReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminCategory_UpdateQyAdminCategory0_HTTP_Handler(srv QyAdminCategoryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateQyAdminCategoryRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminCategoryUpdateQyAdminCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateQyAdminCategory(ctx, req.(*UpdateQyAdminCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateQyAdminCategoryReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminCategory_DeleteQyAdminCategory0_HTTP_Handler(srv QyAdminCategoryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteQyAdminCategoryRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminCategoryDeleteQyAdminCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteQyAdminCategory(ctx, req.(*DeleteQyAdminCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteQyAdminCategoryReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminCategory_ListQyAdminCategory0_HTTP_Handler(srv QyAdminCategoryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListQyAdminCategoryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminCategoryListQyAdminCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListQyAdminCategory(ctx, req.(*ListQyAdminCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListQyAdminCategoryReply)
		return ctx.Result(200, reply)
	}
}

type QyAdminCategoryHTTPClient interface {
	CreateQyAdminCategory(ctx context.Context, req *CreateQyAdminCategoryRequest, opts ...http.CallOption) (rsp *CreateQyAdminCategoryReply, err error)
	DeleteQyAdminCategory(ctx context.Context, req *DeleteQyAdminCategoryRequest, opts ...http.CallOption) (rsp *DeleteQyAdminCategoryReply, err error)
	ListQyAdminCategory(ctx context.Context, req *ListQyAdminCategoryRequest, opts ...http.CallOption) (rsp *ListQyAdminCategoryReply, err error)
	UpdateQyAdminCategory(ctx context.Context, req *UpdateQyAdminCategoryRequest, opts ...http.CallOption) (rsp *UpdateQyAdminCategoryReply, err error)
}

type QyAdminCategoryHTTPClientImpl struct {
	cc *http.Client
}

func NewQyAdminCategoryHTTPClient(client *http.Client) QyAdminCategoryHTTPClient {
	return &QyAdminCategoryHTTPClientImpl{client}
}

func (c *QyAdminCategoryHTTPClientImpl) CreateQyAdminCategory(ctx context.Context, in *CreateQyAdminCategoryRequest, opts ...http.CallOption) (*CreateQyAdminCategoryReply, error) {
	var out CreateQyAdminCategoryReply
	pattern := "/api/admin/v1/category"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQyAdminCategoryCreateQyAdminCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminCategoryHTTPClientImpl) DeleteQyAdminCategory(ctx context.Context, in *DeleteQyAdminCategoryRequest, opts ...http.CallOption) (*DeleteQyAdminCategoryReply, error) {
	var out DeleteQyAdminCategoryReply
	pattern := "/api/admin/v1/category"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQyAdminCategoryDeleteQyAdminCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminCategoryHTTPClientImpl) ListQyAdminCategory(ctx context.Context, in *ListQyAdminCategoryRequest, opts ...http.CallOption) (*ListQyAdminCategoryReply, error) {
	var out ListQyAdminCategoryReply
	pattern := "/api/admin/v1/category"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationQyAdminCategoryListQyAdminCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminCategoryHTTPClientImpl) UpdateQyAdminCategory(ctx context.Context, in *UpdateQyAdminCategoryRequest, opts ...http.CallOption) (*UpdateQyAdminCategoryReply, error) {
	var out UpdateQyAdminCategoryReply
	pattern := "/api/admin/v1/category/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQyAdminCategoryUpdateQyAdminCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
