// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CategoryInfo is the golang structure of table category_info for DAO operations like Where/Data.
type CategoryInfo struct {
	g.Meta      `orm:"table:category_info, do:true"`
	Id          interface{} //
	ParentId    interface{} // 父级id
	Name        interface{} // 分类名称
	PicUrl      interface{} // 分类图标
	Level       interface{} // 等级 默认1级分类
	Sort        interface{} // 排序值
	Status      interface{} // 状态 1:启用 0:禁用
	Description interface{} // 分类描述
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	DeletedAt   *gtime.Time //
}
