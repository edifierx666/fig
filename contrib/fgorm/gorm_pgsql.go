package fgorm

import (
  "github.com/edifierx666/fig/contrib/fgorm/internal"
  // "github.com/flipped-aurora/gin-vue-admin/server/initialize/internal"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

type Pgsql struct {
  GeneralDB `yaml:",inline" mapstructure:",squash"`
}

// Dsn 基于配置文件获取 dsn
// Author [SliverHorn](https://github.com/SliverHorn)
func (p *Pgsql) Dsn() string {
  return "host=" + p.Path + " user=" + p.Username + " password=" + p.Password + " dbname=" + p.Dbname + " port=" + p.Port + " " + p.Config
}

// LinkDsn 根据 dbname 生成 dsn
// Author [SliverHorn](https://github.com/SliverHorn)
func (p *Pgsql) LinkDsn(dbname string) string {
  return "host=" + p.Path + " user=" + p.Username + " password=" + p.Password + " dbname=" + dbname + " port=" + p.Port + " " + p.Config
}

func (m *Pgsql) GetLogMode() string {
  return m.LogMode
}

//
// // GormPgSql 初始化 Postgresql 数据库
// // Author [piexlmax](https://github.com/piexlmax)
// // Author [SliverHorn](https://github.com/SliverHorn)
// func GormPgSql() *gorm.DB {
// 	p := global.GVA_CONFIG.Pgsql
// 	if p.Dbname == "" {
// 		return nil
// 	}
// 	pgsqlConfig := postgres.Config{
// 		DSN:                  p.Dsn(), // DSN data source name
// 		PreferSimpleProtocol: false,
// 	}
// 	if Instance, err := gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(p.Prefix, p.Singular)); err != nil {
// 		return nil
// 	} else {
// 		sqlDB, _ := Instance.DB()
// 		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
// 		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
// 		return Instance
// 	}
// }

// GormPgSqlByConfig 初始化 Postgresql 数据库 通过参数
func GormPgSqlByConfig(p Pgsql) *gorm.DB {
  if p.Dbname == "" {
    return nil
  }
  pgsqlConfig := postgres.Config{
    DSN:                  p.Dsn(), // DSN data source name
    PreferSimpleProtocol: false,
  }
  if db, err := gorm.Open(
    postgres.New(pgsqlConfig), internal.Gorm.Config(p.Prefix, p.Singular),
  ); err != nil {
    panic(err)
  } else {
    sqlDB, _ := db.DB()
    sqlDB.SetMaxIdleConns(p.MaxIdleConns)
    sqlDB.SetMaxOpenConns(p.MaxOpenConns)
    return db
  }
}
