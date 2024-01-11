// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ferror_test

import (
  "errors"
  "fmt"

  "github.com/edifierx666/fig/errors/fcode"
  "github.com/edifierx666/fig/errors/ferror"
)

func ExampleNewCode() {
  err := ferror.NewCode(fcode.New(10000, "", nil), "My Error")
  fmt.Println(err.Error())
  fmt.Println(ferror.Code(err))

  // Output:
  // My Error
  // 10000
}

func ExampleNewCodef() {
  err := ferror.NewCodef(fcode.New(10000, "", nil), "It's %s", "My Error")
  fmt.Println(err.Error())
  fmt.Println(ferror.Code(err).Code())

  // Output:
  // It's My Error
  // 10000
}

func ExampleWrapCode() {
  err1 := errors.New("permission denied")
  err2 := ferror.WrapCode(fcode.New(10000, "", nil), err1, "Custom Error")
  fmt.Println(err2.Error())
  fmt.Println(ferror.Code(err2).Code())

  // Output:
  // Custom Error: permission denied
  // 10000
}

func ExampleWrapCodef() {
  err1 := errors.New("permission denied")
  err2 := ferror.WrapCodef(fcode.New(10000, "", nil), err1, "It's %s", "Custom Error")
  fmt.Println(err2.Error())
  fmt.Println(ferror.Code(err2).Code())

  // Output:
  // It's Custom Error: permission denied
  // 10000
}

func ExampleEqual() {
  err1 := errors.New("permission denied")
  err2 := ferror.New("permission denied")
  err3 := ferror.NewCode(fcode.CodeNotAuthorized, "permission denied")
  fmt.Println(ferror.Equal(err1, err2))
  fmt.Println(ferror.Equal(err2, err3))

  // Output:
  // true
  // false
}

func ExampleIs() {
  err1 := errors.New("permission denied")
  err2 := ferror.Wrap(err1, "operation failed")
  fmt.Println(ferror.Is(err1, err1))
  fmt.Println(ferror.Is(err2, err2))
  fmt.Println(ferror.Is(err2, err1))
  fmt.Println(ferror.Is(err1, err2))

  // Output:
  // false
  // true
  // true
  // false
}
