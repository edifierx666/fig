package fgorm

import (
  "fmt"

  "dario.cat/mergo"
  "gorm.io/gorm"
)

type DB struct {
  Instance *gorm.DB
  Inner    *SpecializedDB
}

type DBOption func(db *DB)

func NewDB(options ...DBOption) *DB {
  db := &DB{}

  for _, option := range options {
    option(db)
  }

  if db.Inner == nil {
    db.Inner = &SpecializedDB{
      Type:      "sqlite",
      AliasName: "default",
      GeneralDB: GeneralDB{
        Dbname: "db",
        Path:   "./",
      },
      Disable: false,
    }
  }

  db.Open()
  return db
}

func WithDBConfig(dbc *SpecializedDB) DBOption {
  return func(db *DB) {
    db.Inner = dbc
  }
}
func WithDBMap(m map[string]any) DBOption {
  return func(db *DB) {
    db.FromMap(m)
  }
}

func (db *DB) Open() *gorm.DB {
  if db.Inner.Disable {
    return nil
  }
  switch db.Inner.Type {
  case "mysql":
    m := Mysql{GeneralDB: db.Inner.GeneralDB}
    db.Instance = GormMysqlByConfig(m)
  // case "pgsql":
  //   m := config.Pgsql{GeneralDB:Instance.GeneralDB}
  //   dbc.DBMap[Instance.AliasName] = GormPgSqlByConfig(m)
  // case "oracle":
  //   return GormOracle()
  // case "mssql":
  //   return GormMssql()
  case "sqlite":
    m := Sqlite{GeneralDB: db.Inner.GeneralDB}
    db.Instance = GormSqliteByConfig(m)
  }
  return db.Instance
}

func (db *DB) FromMap(m map[string]any) *DB {
  err := errrorwrap(
    mergo.Map(
      db.Inner, m, mergo.WithOverride, mergo.WithOverrideEmptySlice,
      mergo.WithOverwriteWithEmptyValue,
    ),
  )
  if err != nil {
    fmt.Println("Dbmap函数错误", err)
    return nil
  }
  return db
}
