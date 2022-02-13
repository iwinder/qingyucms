package user

import (
	"gitee.com/windcoder/qingyucms/internal/pkg/qy-common/core"
	metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qy-common/meta/v1"
	log "gitee.com/windcoder/qingyucms/internal/pkg/qy-log"
	"github.com/gin-gonic/gin"
)

func (u *UserController) Delete(c *gin.Context) {
	log.L(c).Info("delete user function called.")

	if err := u.srv.Users().Delete(c, c.Param("username"), metav1.DeleteOptions{
		Unscoped: false,
	}); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, nil)
}
