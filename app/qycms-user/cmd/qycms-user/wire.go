// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/iwinder/qingyucms/app/qycms-user/internal/biz"
	"github.com/iwinder/qingyucms/app/qycms-user/internal/conf"
	"github.com/iwinder/qingyucms/app/qycms-user/internal/data"
	"github.com/iwinder/qingyucms/app/qycms-user/internal/server"
	"github.com/iwinder/qingyucms/app/qycms-user/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
