package fig

import (
  "fmt"
  "testing"

  "github.com/edifierx666/fig/contrib/fgorm"
  "github.com/edifierx666/fig/contrib/fviper"
  "github.com/edifierx666/fig/errors/ferror"
  "github.com/gookit/goutil/dump"
)

func TestA1(t *testing.T) {
  var m map[any]any
  fviper.New(fviper.WithResult(&m))
  dump.P(m)
}
func TestA2(t *testing.T) {
  m := []map[string]any{
    {
      "type":      "sqlite",
      "aliasName": "a1",
      "path":      "./",
      "dbname":    "db",
    },
  }
  dbConfiger := fgorm.NewDBConfiger(fgorm.WithDBMapList(m))
  db := dbConfiger.GetDBByName("a1")
  db.AutoMigrate(fgorm.GeneralDB{})
}
func TestA3(t *testing.T) {
  db := fgorm.NewDB()
  db.Instance.AutoMigrate(&fgorm.GeneralDB{})

}
func TestA4(t *testing.T) {
  err := ferror.New("错误")
  fmt.Printf("err: %+v", err)
  fmt.Println("--------------------")
  fmt.Println(err)
}
