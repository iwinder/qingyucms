package auth

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"github.com/iwinder/qingyucms/internal/qycms_blog/conf"
)

type CustomClaims struct {
	ID          uint64
	NickName    string
	AuthorityId uint64
	jwtV4.RegisteredClaims
}

// CreateToken generate token
func CreateToken(c CustomClaims, key string) (string, error) {
	claims := jwtV4.NewWithClaims(jwtV4.SigningMethodHS256, c)
	signedString, err := claims.SignedString([]byte(key))
	if err != nil {
		return "", errors.New("generate token failed" + err.Error())
	}

	return signedString, nil
}

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]bool)
	whiteList["/api.qycms_blog.admin.v1.QyBlogAdmin/Login"] = true
	whiteList["/api.qycms_user.v1.User/CreateUser"] = true

	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

func NewMiddleware(authConf *conf.Auth, logger log.Logger) http.ServerOption {
	return http.Middleware(
		recovery.Recovery(),
		tracing.Server(),
		logging.Server(logger),
		selector.Server(
			jwt.Server(
				func(token *jwtV4.Token) (interface{}, error) {
					return []byte(authConf.JwtSecret), nil
				},
				jwt.WithSigningMethod(jwtV4.SigningMethodHS256),
			),
		).Match(NewWhiteListMatcher()).Build(),
	)
}