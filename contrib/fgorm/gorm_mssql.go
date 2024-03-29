package fgorm

import (
  "github.com/edifierx666/fig/contrib/fgorm/internal"
  "gorm.io/driver/sqlserver"
  "gorm.io/gorm"
)

type Mssql struct {
  GeneralDB `yaml:",inline" mapstructure:",squash"`
}

// dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
func (m *Mssql) Dsn() string {
  return "sqlserver://" + m.Username + ":" + m.Password + "@" + m.Path + ":" + m.Port + "?database=" + m.Dbname + "&encrypt=disable"
}

func (m *Mssql) GetLogMode() string {
  return m.LogMode
}

// GormMssql 初始化Mssql数据库
// Author [LouisZhang](191180776@qq.com)
// func GormMssql() *gorm.DB {
// 	m := global.GVA_CONFIG.Mssql
// 	if m.Dbname == "" {
// 		return nil
// 	}
// 	mssqlConfig := sqlserver.Config{
// 		DSN:               m.Dsn(), // DSN data source name
// 		DefaultStringSize: 191,     // string 类型字段的默认长度
// 	}
// 	if Instance, err := gorm.Open(sqlserver.New(mssqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
// 		return nil
// 	} else {
// 		Instance.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
// 		sqlDB, _ := Instance.DB()
// 		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
// 		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
// 		return Instance
// 	}
// }

// GormMssqlByConfig 初始化Mysql数据库用过传入配置
func GormMssqlByConfig(m Mssql) *gorm.DB {
  if m.Dbname == "" {
    return nil
  }
  mssqlConfig := sqlserver.Config{
    DSN:               m.Dsn(), // DSN data source name
    DefaultStringSize: 191,     // string 类型字段的默认长度
  }
  if db, err := gorm.Open(
    sqlserver.New(mssqlConfig), internal.Gorm.Config(m.Prefix, m.Singular),
  ); err != nil {
    panic(err)
  } else {
    db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
    sqlDB, _ := db.DB()
    sqlDB.SetMaxIdleConns(m.MaxIdleConns)
    sqlDB.SetMaxOpenConns(m.MaxOpenConns)
    return db
  }
}
