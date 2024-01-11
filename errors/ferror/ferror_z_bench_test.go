// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ferror_test

import (
  "errors"
  "testing"

  "github.com/edifierx666/fig/errors/fcode"
  "github.com/edifierx666/fig/errors/ferror"
)

var (
  // base error for benchmark testing of Wrap* functions.
  baseError = errors.New("test")
)

func Benchmark_New(b *testing.B) {
  for i := 0; i < b.N; i++ {
    ferror.New("test")
  }
}

func Benchmark_Newf(b *testing.B) {
  for i := 0; i < b.N; i++ {
    ferror.Newf("%s", "test")
  }
}

func Benchmark_Wrap(b *testing.B) {
  for i := 0; i < b.N; i++ {
    ferror.Wrap(baseError, "test")
  }
}

func Benchmark_Wrapf(b *testing.B) {
  for i := 0; i < b.N; i++ {
    ferror.Wrapf(baseError, "%s", "test")
  }
}

func Benchmark_NewSkip(b *testing.B) {
  for i := 0; i < b.N; i++ {
    ferror.NewSkip(1, "test")
  }
}

func Benchmark_NewSkipf(b *testing.B) {
  for i := 0; i < b.N; i++ {
    ferror.NewSkipf(1, "%s", "test")
  }
}

func Benchmark_NewCode(b *testing.B) {
  for i := 0; i < b.N; i++ {
    ferror.NewCode(fcode.New(500, "", nil), "test")
  }
}

func Benchmark_NewCodef(b *testing.B) {
  for i := 0; i < b.N; i++ {
    ferror.NewCodef(fcode.New(500, "", nil), "%s", "test")
  }
}

func Benchmark_NewCodeSkip(b *testing.B) {
  for i := 0; i < b.N; i++ {
    ferror.NewCodeSkip(fcode.New(1, "", nil), 500, "test")
  }
}

func Benchmark_NewCodeSkipf(b *testing.B) {
  for i := 0; i < b.N; i++ {
    ferror.NewCodeSkipf(fcode.New(1, "", nil), 500, "%s", "test")
  }
}

func Benchmark_WrapCode(b *testing.B) {
  for i := 0; i < b.N; i++ {
    ferror.WrapCode(fcode.New(500, "", nil), baseError, "test")
  }
}

func Benchmark_WrapCodef(b *testing.B) {
  for i := 0; i < b.N; i++ {
    ferror.WrapCodef(fcode.New(500, "", nil), baseError, "test")
  }
}
