// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dictionary_details

import "github.com/gogf/gf/os/gtime"

// DictionaryDetails
type DictionaryDetails struct {
	Id           uint        `orm:"id,primary"    json:"ID"`              // 自增ID
	CreateAt     *gtime.Time `orm:"create_at"     json:"CreatedAt"`       // 创建时间
	UpdateAt     *gtime.Time `orm:"update_at"     json:"UpdatedAt"`       // 更新时间
	DeleteAt     *gtime.Time `orm:"delete_at"     json:"DeletedAt"`       // 删除时间
	Label        string      `orm:"label"         json:"label"`           // 展示值
	Value        int         `orm:"value"         json:"value"`           // 字典值
	Status       int         `orm:"status"        json:"status"`          // 启用状态
	Sort         int         `orm:"sort"          json:"sort"`            // 排序标记
	DictionaryId int         `orm:"dictionary_id" json:"sysDictionaryID"` // 关联标记
}
