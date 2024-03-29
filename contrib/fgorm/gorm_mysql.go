package fgorm

import (
  "github.com/edifierx666/fig/contrib/fgorm/internal"
  _ "github.com/go-sql-driver/mysql"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

type Mysql struct {
  GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *Mysql) Dsn() string {
  return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func (m *Mysql) GetLogMode() string {
  return m.LogMode
}

// // GormMysql 初始化Mysql数据库
// // Author [piexlmax](https://github.com/piexlmax)
// // Author [SliverHorn](https://github.com/SliverHorn)
// func GormMysql() *gorm.DB {
//   m := global.GVA_CONFIG.Mysql
//   if m.Dbname == "" {
//     return nil
//   }
//   mysqlConfig := mysql.Config{
//     DSN:                       m.Dsn(), // DSN data source name
//     DefaultStringSize:         191,     // string 类型字段的默认长度
//     SkipInitializeWithVersion: false,   // 根据版本自动配置
//   }
//   if Instance, err := gorm.Open(
//     mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular),
//   ); err != nil {
//     return nil
//   } else {
//     Instance.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
//     sqlDB, _ := Instance.DB()
//     sqlDB.SetMaxIdleConns(m.MaxIdleConns)
//     sqlDB.SetMaxOpenConns(m.MaxOpenConns)
//     return Instance
//   }
// }

// GormMysqlByConfig 初始化Mysql数据库用过传入配置
func GormMysqlByConfig(m Mysql) *gorm.DB {
  if m.Dbname == "" {
    return nil
  }
  mysqlConfig := mysql.Config{
    DSN:                       m.Dsn(), // DSN data source name
    DefaultStringSize:         191,     // string 类型字段的默认长度
    SkipInitializeWithVersion: false,   // 根据版本自动配置
  }
  if db, err := gorm.Open(
    mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular),
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
