//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/iwinder/qingyucms/app/qycms_user/internal/biz"
	"github.com/iwinder/qingyucms/app/qycms_user/internal/conf"
	"github.com/iwinder/qingyucms/app/qycms_user/internal/data/db"
	"github.com/iwinder/qingyucms/app/qycms_user/internal/server"
	"github.com/iwinder/qingyucms/app/qycms_user/internal/service"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Qycms, *conf.Data, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, db.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
