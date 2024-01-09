package fig

import (
  "testing"

  "github.com/edifierx666/fig/contrib/fviper"
  "github.com/gookit/goutil/dump"
)

func TestA1(t *testing.T) {
  var m map[any]any
  fviper.New(fviper.WithResult(&m))
  dump.P(m)
}
