package model

import "github.com/flipped-aurora/gva/library/global"

type Menu struct {
	global.Model
	Path      string `json:"path" gorm:"comment:路由path"`        // 路由path
	Name      string `json:"name" gorm:"comment:路由name"`        // 路由name
	Sort      int    `json:"sort" gorm:"comment:排序标记"`          // 排序标记
	Hidden    bool   `json:"hidden" gorm:"comment:是否在列表隐藏"`     // 是否在列表隐藏
	ParentId  string `json:"parentId" gorm:"comment:父菜单ID"`     // 父菜单ID
	Component string `json:"component" gorm:"comment:对应前端文件路径"` // 对应前端文件路径
	MenuLevel uint   `json:"-"`
	Meta      `json:"meta" gorm:"comment:附加属性"` // 附加属性

	Children    []Menu          `json:"children" gorm:"-"`
	Parameters  []MenuParameter `json:"parameters" gorm:"many2many:system_menus_parameters;foreignKey:ID;joinForeignKey:MenuID;References:MenuID;JoinReferences:ParameterID"`
	Authorities []Authority     `json:"authoritys" gorm:"many2many:system_authorities_menus;foreignKey:ID;joinForeignKey:MenuID;References:AuthorityId;JoinReferences:AuthorityId"`
}

type Meta struct {
	Title       string `json:"title" gorm:"comment:菜单名"`                // 菜单名
	Icon        string `json:"icon" gorm:"comment:菜单图标"`                // 菜单图标
	CloseTab    bool   `json:"closeTab" gorm:"comment:自动关闭tab"`         // 自动关闭tab
	KeepAlive   bool   `json:"keepAlive" gorm:"comment:是否缓存"`           // 是否缓存
	DefaultMenu bool   `json:"defaultMenu" gorm:"comment:是否是基础路由（开发中）"` // 是否是基础路由（开发中）
}

func (m *Menu) TableName() string {
	return "system_menus"
}
