package f

import (
  "github.com/gogf/gf/v2/container/gvar"
  "github.com/gookit/goutil/dump"
)

import (
  "io"
)

func Dump(vs ...any) {
  dump.P(vs...)
}

func DumpTo(out io.Writer, value interface{}, options ...dump.OptionFunc) {
  dump.NewDumper(out, 3).WithOptions(options...).Dump(value)
}

// NewVar returns a gvar.Var.
func NewVar(i interface{}, safe ...bool) *gvar.Var {
  return gvar.New(i, safe...)
}
