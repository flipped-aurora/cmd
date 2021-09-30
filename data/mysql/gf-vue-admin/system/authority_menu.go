package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gva/library/global"
	"github.com/gookit/color"
)

var AuthorityMenu = new(authorityMenu)

type authorityMenu struct{}

// Init authority_menu 视图创建
func (a *authorityMenu) Init() error {
	if global.Db.Raw("select * from authority_menu").RowsAffected > 0 {
		color.Danger.Println("\n[Mysql] --> authority_menu 视图已存在!")
		return nil
	}
	if err := global.Db.Exec("CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `authority_menu` AS select `menus`.`id` AS `id`, `menus`.`created_at` AS `created_at`, `menus`.`updated_at` AS `updated_at`, `menus`.`deleted_at` AS `deleted_at`, `menus`.`menu_level` AS `menu_level`, `menus`.`parent_id` AS `parent_id`, `menus`.`path` AS `path`, `menus`.`name` AS `name`, `menus`.`hidden` AS `hidden`, `menus`.`component` AS `component`, `menus`.`title` AS `title`, `menus`.`icon` AS `icon`, `menus`.`sort` AS `sort`, `authorities_menus`.`authority_id` AS `authority_id`, `authorities_menus`.`menu_id` AS `menu_id`, `menus`.`keep_alive` AS `keep_alive`, `menus`.`default_menu` AS `default_menu`from (`authorities_menus` join `menus` on ((`authorities_menus`.`menu_id` = `menus`.`id`)));").Error; err != nil {
		return err
	}
	return nil
}

// TableName 定义表名
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *authorityMenu) TableName() string {
	var entity system.AuthorityMenu
	return entity.TableName()
}
