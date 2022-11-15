// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.1
// - protoc             v3.21.1
// source: api/qycms_bff/admin/v1/qy_admin_tags.proto

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

const OperationQyAdminTagsCreateQyAdminTags = "/api.qycms_bff.admin.v1.QyAdminTags/CreateQyAdminTags"
const OperationQyAdminTagsDeleteQyAdminTags = "/api.qycms_bff.admin.v1.QyAdminTags/DeleteQyAdminTags"
const OperationQyAdminTagsListQyAdminTags = "/api.qycms_bff.admin.v1.QyAdminTags/ListQyAdminTags"
const OperationQyAdminTagsUpdateQyAdminTags = "/api.qycms_bff.admin.v1.QyAdminTags/UpdateQyAdminTags"

type QyAdminTagsHTTPServer interface {
	CreateQyAdminTags(context.Context, *CreateQyAdminTagsRequest) (*CreateQyAdminTagsReply, error)
	DeleteQyAdminTags(context.Context, *DeleteQyAdminTagsRequest) (*DeleteQyAdminTagsReply, error)
	ListQyAdminTags(context.Context, *ListQyAdminTagsRequest) (*ListQyAdminTagsReply, error)
	UpdateQyAdminTags(context.Context, *UpdateQyAdminTagsRequest) (*UpdateQyAdminTagsReply, error)
}

func RegisterQyAdminTagsHTTPServer(s *http.Server, srv QyAdminTagsHTTPServer) {
	r := s.Route("/")
	r.POST("/api/admin/v1/tags", _QyAdminTags_CreateQyAdminTags0_HTTP_Handler(srv))
	r.PUT("/api/admin/v1/tags/{id}", _QyAdminTags_UpdateQyAdminTags0_HTTP_Handler(srv))
	r.DELETE("/api/admin/v1/tags", _QyAdminTags_DeleteQyAdminTags0_HTTP_Handler(srv))
	r.GET("/api/admin/v1/tags", _QyAdminTags_ListQyAdminTags0_HTTP_Handler(srv))
}

func _QyAdminTags_CreateQyAdminTags0_HTTP_Handler(srv QyAdminTagsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateQyAdminTagsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminTagsCreateQyAdminTags)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateQyAdminTags(ctx, req.(*CreateQyAdminTagsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateQyAdminTagsReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminTags_UpdateQyAdminTags0_HTTP_Handler(srv QyAdminTagsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateQyAdminTagsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminTagsUpdateQyAdminTags)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateQyAdminTags(ctx, req.(*UpdateQyAdminTagsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateQyAdminTagsReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminTags_DeleteQyAdminTags0_HTTP_Handler(srv QyAdminTagsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteQyAdminTagsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminTagsDeleteQyAdminTags)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteQyAdminTags(ctx, req.(*DeleteQyAdminTagsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteQyAdminTagsReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminTags_ListQyAdminTags0_HTTP_Handler(srv QyAdminTagsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListQyAdminTagsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminTagsListQyAdminTags)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListQyAdminTags(ctx, req.(*ListQyAdminTagsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListQyAdminTagsReply)
		return ctx.Result(200, reply)
	}
}

type QyAdminTagsHTTPClient interface {
	CreateQyAdminTags(ctx context.Context, req *CreateQyAdminTagsRequest, opts ...http.CallOption) (rsp *CreateQyAdminTagsReply, err error)
	DeleteQyAdminTags(ctx context.Context, req *DeleteQyAdminTagsRequest, opts ...http.CallOption) (rsp *DeleteQyAdminTagsReply, err error)
	ListQyAdminTags(ctx context.Context, req *ListQyAdminTagsRequest, opts ...http.CallOption) (rsp *ListQyAdminTagsReply, err error)
	UpdateQyAdminTags(ctx context.Context, req *UpdateQyAdminTagsRequest, opts ...http.CallOption) (rsp *UpdateQyAdminTagsReply, err error)
}

type QyAdminTagsHTTPClientImpl struct {
	cc *http.Client
}

func NewQyAdminTagsHTTPClient(client *http.Client) QyAdminTagsHTTPClient {
	return &QyAdminTagsHTTPClientImpl{client}
}

func (c *QyAdminTagsHTTPClientImpl) CreateQyAdminTags(ctx context.Context, in *CreateQyAdminTagsRequest, opts ...http.CallOption) (*CreateQyAdminTagsReply, error) {
	var out CreateQyAdminTagsReply
	pattern := "/api/admin/v1/tags"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQyAdminTagsCreateQyAdminTags))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminTagsHTTPClientImpl) DeleteQyAdminTags(ctx context.Context, in *DeleteQyAdminTagsRequest, opts ...http.CallOption) (*DeleteQyAdminTagsReply, error) {
	var out DeleteQyAdminTagsReply
	pattern := "/api/admin/v1/tags"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQyAdminTagsDeleteQyAdminTags))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminTagsHTTPClientImpl) ListQyAdminTags(ctx context.Context, in *ListQyAdminTagsRequest, opts ...http.CallOption) (*ListQyAdminTagsReply, error) {
	var out ListQyAdminTagsReply
	pattern := "/api/admin/v1/tags"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationQyAdminTagsListQyAdminTags))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminTagsHTTPClientImpl) UpdateQyAdminTags(ctx context.Context, in *UpdateQyAdminTagsRequest, opts ...http.CallOption) (*UpdateQyAdminTagsReply, error) {
	var out UpdateQyAdminTagsReply
	pattern := "/api/admin/v1/tags/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQyAdminTagsUpdateQyAdminTags))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
