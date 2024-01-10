package fgorm

import (
  "fmt"
  "sync"

  "github.com/pkg/errors"
  "gorm.io/gorm"
)

var lock sync.RWMutex
var SupportDBTypes = []string{"mysql", "pgsql", "oracle", "mssql", "sqlite"}

func errrorwrap(err error) error {
  return errors.WithStack(errors.Wrap(err, "DBConfiger错误:"))
}

type DBList struct {
  DBList []*DB
  DBMap  map[string]*gorm.DB
}

type DBListOption func(dbc *DBList)

func (dbc *DBList) checkkey(aliasname string) bool {
  if _, ok := dbc.DBMap[aliasname]; ok {
    return ok
  }
  return false
}

func (dbc *DBList) Dbmap(m map[string]any) (dbs *DB) {
  dbs = &DB{}
  return dbs.FromMap(m)
}

func (dbc *DBList) GetDBByName(aliasname string) (db *gorm.DB) {
  lock.RLock()
  defer lock.RUnlock()
  return dbc.DBMap[aliasname]
}

func (dbc *DBList) RemoveDBByName(aliasname string) {
  lock.RLock()
  defer lock.RUnlock()
  // 关闭数据库连接
  db, _ := dbc.DBMap[aliasname].DB()
  _ = db.Close()
  // 删除map
  delete(dbc.DBMap, aliasname)
  var filterSlice []*DB
  for _, db := range dbc.DBList {
    if db.Inner.AliasName != aliasname {
      filterSlice = append(filterSlice, db)
    }
  }

  dbc.DBList = filterSlice
}

func (dbc *DBList) MustGetDBByName(aliasname string) (db *gorm.DB) {
  lock.RLock()
  defer lock.RUnlock()
  db, ok := dbc.DBMap[aliasname]
  if !ok || db == nil {
    panic("Instance no init")
  }
  return db
}

func (dbc *DBList) AddDB(dbs *DB) *DBList {
  dbc.DBList = append(dbc.DBList, dbs)
  dbc.Open()
  return dbc
}

func (dbc *DBList) AddDBMap(m map[string]any) *DBList {
  dbs := dbc.Dbmap(m)
  if dbs != nil {
    dbc.DBList = append(dbc.DBList, dbs)
  }
  dbc.Open()
  return dbc
}

func (dbc *DBList) Open() {
  for _, db := range dbc.DBList {
    if dbc.checkkey(db.Inner.AliasName) {
      fmt.Println("DBList:检测到相同DB-aslias，跳过", db.Inner.AliasName, db)
      continue
    }
    dbc.DBMap[db.Inner.AliasName] = db.Open()
  }
}

func NewDBConfiger(options ...DBListOption) (dbConfiger *DBList) {
  dbc := &DBList{
    DBList: make([]*DB, 0),
    DBMap:  make(map[string]*gorm.DB),
  }

  for _, option := range options {
    option(dbc)
  }

  dbc.Open()

  return dbc
}

func WithDBList(list []*DB) DBListOption {
  return func(dbc *DBList) {
    for _, db := range list {
      isSupported := false
      for _, dbType := range SupportDBTypes {
        if db.Inner.Type == dbType {
          isSupported = true
        }
      }
      if !isSupported {
        fmt.Println(errrorwrap(errors.New("不支持的数据库类型:" + db.Inner.Type)))
        db.Inner.Type = "mysql"
      }
    }
    dbc.DBList = list
  }
}

func WithDBMapList(list []map[string]any) DBListOption {
  return func(dbc *DBList) {
    for _, m := range list {
      dbs := dbc.Dbmap(m)
      if dbs != nil {
        dbc.DBList = append(dbc.DBList, dbs)
      }
    }
    for _, db := range dbc.DBList {
      isSupported := false
      for _, dbType := range SupportDBTypes {
        if db.Inner.Type == dbType {
          isSupported = true
        }
      }
      if !isSupported {
        fmt.Println(errrorwrap(errors.New("不支持的数据库类型:" + db.Inner.Type)))
        db.Inner.Type = "mysql"
      }
    }
  }
}
