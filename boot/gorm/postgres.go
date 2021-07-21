//+build postgres

package boot

import (
	"github.com/flipped-aurora/gva/library/global"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var DbResolver = new(_postgres)

type _postgres struct {
	dsn string
}

// GetSources 获取主库的 gorm.Dialector 切片对象
// Author [SliverHorn](https://github.com/SliverHorn)
func (p *_postgres) GetSources() (directories []gorm.Dialector) {
	length := len(global.GinVueAdminConfig.Gorm.Dsn.Sources)
	directories = make([]gorm.Dialector, 0, length)
	for i := 0; i < length; i++ {
		if !global.GinVueAdminConfig.Gorm.Dsn.Sources[i].IsEmpty() {
			dsn := global.GinVueAdminConfig.Gorm.Dsn.Sources[i].GetDsn(global.GinVueAdminConfig.Gorm.Config)
			if i == 0 {
				p.dsn = dsn
			}
			directories = append(directories, postgres.Open(dsn))
		} else {
			continue
		}
	}
	return directories
}

// GetReplicas 获取从库库的 gorm.Dialector 切片对象
// Author [SliverHorn](https://github.com/SliverHorn)
func (p *_postgres) GetReplicas() (directories []gorm.Dialector) {
	length := len(global.GinVueAdminConfig.Gorm.Dsn.Replicas)
	directories = make([]gorm.Dialector, 0, length)
	for i := 0; i < length; i++ {
		if !global.GinVueAdminConfig.Gorm.Dsn.Replicas[i].IsEmpty() {
			dsn := global.GinVueAdminConfig.Gorm.Dsn.Replicas[i].GetDsn(global.GinVueAdminConfig.Gorm.Config)
			directories = append(directories, postgres.Open(dsn))
		} else {
			continue
		}
	}
	return directories
}

// GetResolver 通过主库与从库的链接组装 gorm.Plugin
// Author [SliverHorn](https://github.com/SliverHorn)
func (p *_postgres) GetResolver() gorm.Plugin {
	sources := p.GetSources()
	resolver := dbresolver.Register(dbresolver.Config{
		Sources:  sources,
		Replicas: p.GetReplicas(),
		Policy:   dbresolver.RandomPolicy{}, // sources/replicas 负载均衡策略
	})
	resolver.SetMaxIdleConns(global.GinVueAdminConfig.Gorm.GetMaxOpenConnes())
	resolver.SetMaxOpenConns(global.GinVueAdminConfig.Gorm.GetMaxOpenConnes())
	resolver.SetConnMaxIdleTime(global.GinVueAdminConfig.Gorm.GetConnMaxIdleTime())
	resolver.SetConnMaxLifetime(global.GinVueAdminConfig.Gorm.GetConnMaxLifetime())
	return resolver
}

// GetGormDialector 获取数据库的 gorm.Dialector
// Author [SliverHorn](https://github.com/SliverHorn)
func (p *_postgres) GetGormDialector() gorm.Dialector {
	return postgres.New(postgres.Config{
		DSN:                  p.dsn,
		PreferSimpleProtocol: false,
	})
}

func (p *_postgres) GetConfigPath() string {
	return "config/config.postgres.yaml"
}
