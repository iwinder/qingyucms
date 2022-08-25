package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/iwinder/qingyucms/api/qycms_blog/admin/v1"
	userV1 "github.com/iwinder/qingyucms/api/qycms_user/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/auth"
	"github.com/iwinder/qingyucms/internal/qycms_blog/conf"
	"github.com/iwinder/qingyucms/internal/qycms_blog/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, authConf *conf.Auth, greeter *service.BlogAdminService, user *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		auth.NewMiddleware(authConf, logger),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterQyBlogAdminHTTPServer(srv, greeter)
	userV1.RegisterUserHTTPServer(srv, user)
	return srv
}
