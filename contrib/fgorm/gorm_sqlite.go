package fgorm

import (
  "path/filepath"

  "github.com/edifierx666/fig/contrib/fgorm/internal"
  // "github.com/flipped-aurora/gin-vue-admin/server/initialize/internal"
  "github.com/glebarez/sqlite"
  "gorm.io/gorm"
)

type Sqlite struct {
  GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (s *Sqlite) Dsn() string {
  return filepath.Join(s.Path, s.Dbname+".db")
}

func (s *Sqlite) GetLogMode() string {
  return s.LogMode
}

// GormSqlite 初始化Sqlite数据库
func GormSqlite() *gorm.DB {
  s := &Sqlite{}
  if db, err := gorm.Open(sqlite.Open(s.Dsn()), internal.Gorm.Config("", false)); err != nil {
    panic(err)
  } else {
    sqlDB, _ := db.DB()
    sqlDB.SetMaxIdleConns(s.MaxIdleConns)
    sqlDB.SetMaxOpenConns(s.MaxOpenConns)
    return db
  }
}

// GormSqliteByConfig 初始化Sqlite数据库用过传入配置
func GormSqliteByConfig(s Sqlite) *gorm.DB {
  if s.Dbname == "" {
    return nil
  }

  if db, err := gorm.Open(
    sqlite.Open(s.Dsn()), internal.Gorm.Config(s.Prefix, s.Singular),
  ); err != nil {
    panic(err)
  } else {
    sqlDB, _ := db.DB()
    sqlDB.SetMaxIdleConns(s.MaxIdleConns)
    sqlDB.SetMaxOpenConns(s.MaxOpenConns)
    return db
  }
}
