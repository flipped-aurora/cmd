package postgres

import "github.com/flipped-aurora/gva/library/global"

var Postgres = new(_postgres)

type _postgres struct{}

func (p *_postgres) CreateDatabase() error {
	return global.Db.Exec("CREATE DATABASE "+global.Config.Gorm.Dsn.Sources[0].DbName).Error
}
