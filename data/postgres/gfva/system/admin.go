//+build postgres

package data

import (
	"github.com/flipped-aurora/gva/library/global"
	model "github.com/flipped-aurora/gva/model/gfva/system"
	"time"

	"github.com/gookit/color"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var Admin = new(admin)

type admin struct{}

// Init admins 表数据初始化
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *admin) Init() error {
	admins := []model.Admin{
		{Model: global.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Uuid: uuid.NewV4().String(), Username: "admin", Password: "123456", Nickname: I18nHash["SuperAdmin"], Avatar: "http://qmplusimg.henrongyi.top/gva_header.jpg", AuthorityId: "888"},
		{Model: global.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Uuid: uuid.NewV4().String(), Username: "a303176530", Password: "123456", Nickname: I18nHash["OtherUser"], Avatar: "http://qmplusimg.henrongyi.top/1572075907logo.png", AuthorityId: "9528"},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		for i := range admins {
			_ = admins[i].EncryptedPassword()
		}
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.Admin{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> admins 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

// TableName 定义表名
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *admin) TableName() string {
	return "admins"
}
