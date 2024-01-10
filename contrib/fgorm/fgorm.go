package fgorm

import (
  "errors"
  "fmt"
  "sync"

  "dario.cat/mergo"
  "github.com/joomcode/errorx"
  "gorm.io/gorm"
)

var lock sync.RWMutex
var SupportDBTypes = []string{"mysql", "pgsql", "oracle", "mssql", "sqlite"}

func errrorwrap(err error) error {
  return errorx.DecorateMany("DBConfiger错误:", err)
}
func errorStackTrace(err error, message string, args ...interface{}) *errorx.Error {
  return errorx.EnhanceStackTrace(errrorwrap(err), message, args...)
}

type DBConfiger struct {
  DBList   []*SpecializedDB
  DBMap    map[string]*gorm.DB
  AutoConn bool
}

type Option func(dbc *DBConfiger)

func (dbc *DBConfiger) checkkey(aliasname string) bool {
  if _, ok := dbc.DBMap[aliasname]; ok {
    return ok
  }
  return false
}

func (dbc *DBConfiger) Dbmap(m map[string]any) (dbs *SpecializedDB) {
  dbs = &SpecializedDB{}
  err := errrorwrap(
    mergo.Map(
      dbs, m, mergo.WithOverride, mergo.WithOverrideEmptySlice,
      mergo.WithOverwriteWithEmptyValue,
    ),
  )
  if err != nil {
    fmt.Println("Dbmap函数错误", err)
    return nil
  }
  return dbs
}

func (dbc *DBConfiger) GetDBByName(aliasname string) (db *gorm.DB) {
  lock.RLock()
  defer lock.RUnlock()
  return dbc.DBMap[aliasname]
}

func (dbc *DBConfiger) MustGetDBByName(aliasname string) (db *gorm.DB) {
  lock.RLock()
  defer lock.RUnlock()
  db, ok := dbc.DBMap[aliasname]
  if !ok || db == nil {
    panic("db no init")
  }
  return db
}

func (dbc *DBConfiger) AddDB(dbs *SpecializedDB) *DBConfiger {
  dbc.DBList = append(dbc.DBList, dbs)
  dbc.autoConn()
  return dbc
}
func (dbc *DBConfiger) AddDBMap(m map[string]any) *DBConfiger {
  dbs := dbc.Dbmap(m)
  if dbs != nil {
    dbc.DBList = append(dbc.DBList, dbs)
  }
  dbc.autoConn()
  return dbc
}

func (dbc *DBConfiger) autoConn() {
  if dbc.AutoConn {
    dbc.Conn()
  }
}

func (dbc *DBConfiger) Conn() {
  for _, db := range dbc.DBList {
    if db.Disable {
      continue
    }
    if dbc.checkkey(db.AliasName) {
      fmt.Println("DBConfiger:检测到相同DB-aslias，跳过", db.AliasName, db)
      continue
    }

    switch db.Type {
    case "mysql":
      m := Mysql{GeneralDB: db.GeneralDB}
      dbc.DBMap[db.AliasName] = GormMysqlByConfig(m)
    // case "pgsql":
    //   m := config.Pgsql{GeneralDB:db.GeneralDB}
    //   dbc.DBMap[db.AliasName] = GormPgSqlByConfig(m)
    // case "oracle":
    //   return GormOracle()
    // case "mssql":
    //   return GormMssql()
    case "sqlite":
      m := Sqlite{GeneralDB: db.GeneralDB}
      dbc.DBMap[db.AliasName] = GormSqliteByConfig(m)
    }
  }
}

func NewDBConfiger(options ...Option) (dbConfiger *DBConfiger) {
  dbc := &DBConfiger{
    DBList:   make([]*SpecializedDB, 0),
    DBMap:    make(map[string]*gorm.DB),
    AutoConn: true,
  }

  for _, option := range options {
    option(dbc)
  }

  dbc.autoConn()

  return dbc
}

func WithDBList(list []*SpecializedDB) Option {
  return func(dbc *DBConfiger) {
    for _, db := range list {
      isSupported := false
      for _, dbType := range SupportDBTypes {
        if db.Type == dbType {
          isSupported = true
        }
      }
      if !isSupported {
        fmt.Println(errrorwrap(errors.New("不支持的数据库类型:" + db.Type)))
        db.Type = "mysql"
      }
    }
    dbc.DBList = list
  }
}
func WithDBMapList(list []map[string]any) Option {
  return func(dbc *DBConfiger) {
    for _, m := range list {
      dbs := dbc.Dbmap(m)
      if dbs != nil {
        dbc.DBList = append(dbc.DBList, dbs)
      }
    }
    for _, db := range dbc.DBList {
      isSupported := false
      for _, dbType := range SupportDBTypes {
        if db.Type == dbType {
          isSupported = true
        }
      }
      if !isSupported {
        fmt.Println(errrorwrap(errors.New("不支持的数据库类型:" + db.Type)))
        db.Type = "mysql"
      }
    }
  }
}

func WithAutoConn(b bool) Option {
  return func(dbc *DBConfiger) {
    dbc.AutoConn = b
  }
}
