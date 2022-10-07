package biz

import (
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewGreeterUsecase,
	NewUserRoleUsecase, NewRoleMenusUsecase, NewRoleApiUsecase,
	NewRoleUsecase, NewUserUsecase,
	NewApiUsecase, NewMenusAdminUsecase, NewApiGroupUsecase, NewCasbinRuleUsecase,
	NewArticleUsecase, NewArticleContentUsecase,
	NewCommentAgentUsecase, NewCommentIndexUsecase, NewCommentContentUsecase,
)
