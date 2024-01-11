// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ferror

import (
  "github.com/edifierx666/fig/errors/fcode"
)

// Code returns the error code.
// It returns CodeNil if it has no error code.
func (err *Error) Code() fcode.Code {
  if err == nil {
    return fcode.CodeNil
  }
  if err.code == fcode.CodeNil {
    return Code(err.Unwrap())
  }
  return err.code
}

// SetCode updates the internal code with given code.
func (err *Error) SetCode(code fcode.Code) {
  if err == nil {
    return
  }
  err.code = code
}
