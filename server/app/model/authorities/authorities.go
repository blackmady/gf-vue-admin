// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package authorities

import (
	"server/app/model/menus"

	"github.com/gogf/gf/os/gtime"
)

type Authorities struct {
	CreateAt      *gtime.Time    `orm:"create_at"            json:"CreateAt"`                     // 创建时间
	UpdateAt      *gtime.Time    `orm:"update_at"            json:"UpdateAt"`                     // 更新时间
	DeleteAt      *gtime.Time    `orm:"delete_at"            json:"DeleteAt"`                     // 删除时间
	AuthorityId   string         `r:"authorityId"  orm:"authority_id,primary" json:"authorityId"` // 角色ID
	AuthorityName string         `orm:"authority_name"       json:"authorityName"`                // 角色名
	ParentId      string         `orm:"parent_id"            json:"parentId"`                     // 父角色ID
	ResourcesId   string         `r:"dataAuthorityId" orm:"resources_id"         json:"-"`        // 资源权限ID
	DataAuthority []Authorities  `json:"dataAuthorityId"`
	Children      []*Authorities `json:"children"`
	Menus         []menus.Entity `json:"menus"`
}

// RecordNotFound 根据条件判断数据是否存在
// 有数据返回false
// 没数据 true
func RecordNotFound(where ...interface{}) bool {
	return Model.RecordNotFound(where...)
}

func (m *arModel) RecordNotFound(where ...interface{}) bool {
	r, err := m.M.FindOne(where...)
	if r == nil && err == nil {
		return true
	}
	return false
}
