package v1

import (
	"encoding/json"
	"gitee.com/windcoder/qingyucms/internal/pkg/ qygo-common/utils/idUtil"
	metaV1 "gitee.com/windcoder/qingyucms/internal/pkg/qycms-common/meta/v1"
	"gorm.io/gorm"
)

type Menu struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	Name              string `json:"name" gorm:"name"`
	Path              string `json:"path" gorm:"path"`
	TargetId          uint64 `json:"targetId" gorm:"target_id"`
	TargetType        string `json:"targetType" gorm:"target_type"`
	blanked           bool   `json:"blanked" gorm:"blanked"`
	parentId          uint64 `json:"parentId" gorm:"column:parent_id"`
}

func (m *Menu) TableName() string {
	return "qy_sys_menu"
}

func (m *Menu) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (m *Menu) AfterCreate(tx *gorm.DB) (err error) {
	m.InstanceID = idUtil.GetInstanceID(m.ID, "user-")
	return tx.Save(m).Error
}

func (u *Menu) BeforeUpdate(tx *gorm.DB) (err error) {
	//u.Password, err = auth.Encrypt(u.Password + u.Salt)
	u.ExtendShadow = u.Extend.String()
	return err
}

func (m *Menu) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(m.ExtendShadow), &m.Extend); err != nil {
		return err
	}
	return nil
}
