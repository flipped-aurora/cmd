package model

type AuthorityMenu struct {
	Menu
	MenuId      string          `json:"menuId" gorm:"comment:菜单ID"`
	AuthorityId string          `json:"-" gorm:"comment:角色ID"`
	Children    []AuthorityMenu `json:"children" gorm:"-"`
	Parameters  []MenuParameter `json:"parameters" gorm:"foreignKey:SysBaseMenuID;references:MenuId"`
}

func (a *AuthorityMenu) TableName() string {
	return "authority_menu"
}
