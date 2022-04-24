package user

import (
	"gitee.com/windcoder/qingyucms/internal/pkg/qycms-common/core"
	metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qycms-common/meta/v1"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qygo-log"
	"github.com/gin-gonic/gin"
)

func (u *UserController) Get(c *gin.Context) {
	log.L(c).Info("user create function called.")
	user, err := u.srv.Users().Get(c, c.Param("username"), metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, user)
}

func (u *UserController) GetWithUnscoped(c *gin.Context) {
	log.L(c).Info("user create function called.")
	user, err := u.srv.Users().Get(c, c.Param("username"), metav1.GetOptions{Unscoped: true})
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, user)
}
