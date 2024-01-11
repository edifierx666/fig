// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package fcode_test

import (
  "testing"

  "github.com/edifierx666/fig/errors/fcode"
  "github.com/gogf/gf/v2/test/gtest"
)

func Test_Case(t *testing.T) {
  gtest.C(
    t, func(t *gtest.T) {
      t.Assert(fcode.CodeNil.String(), "-1")
      t.Assert(fcode.CodeInternalError.String(), "50:Internal Error")
    },
  )
}

func Test_Nil(t *testing.T) {
  gtest.C(
    t, func(t *gtest.T) {
      c := fcode.New(1, "custom error", "detailed description")
      t.Assert(c.Code(), 1)
      t.Assert(c.Message(), "custom error")
      t.Assert(c.Detail(), "detailed description")
    },
  )
}

func Test_WithCode(t *testing.T) {
  gtest.C(
    t, func(t *gtest.T) {
      c := fcode.WithCode(fcode.CodeInternalError, "CodeInternalError")
      t.Assert(c.Code(), fcode.CodeInternalError.Code())
      t.Assert(c.Detail(), "CodeInternalError")
    },
  )
}
