// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ferror

import (
  "fmt"
  "strings"

  "github.com/edifierx666/fig/errors/fcode"
)

// NewCode creates and returns an error that has error code and given text.
func NewCode(code fcode.Code, text ...string) error {
  return &Error{
    stack: callers(),
    text:  strings.Join(text, commaSeparatorSpace),
    code:  code,
  }
}

// NewCodef returns an error that has error code and formats as the given format and args.
func NewCodef(code fcode.Code, format string, args ...interface{}) error {
  return &Error{
    stack: callers(),
    text:  fmt.Sprintf(format, args...),
    code:  code,
  }
}

// NewCodeSkip creates and returns an error which has error code and is formatted from given text.
// The parameter `skip` specifies the stack callers skipped amount.
func NewCodeSkip(code fcode.Code, skip int, text ...string) error {
  return &Error{
    stack: callers(skip),
    text:  strings.Join(text, commaSeparatorSpace),
    code:  code,
  }
}

// NewCodeSkipf returns an error that has error code and formats as the given format and args.
// The parameter `skip` specifies the stack callers skipped amount.
func NewCodeSkipf(code fcode.Code, skip int, format string, args ...interface{}) error {
  return &Error{
    stack: callers(skip),
    text:  fmt.Sprintf(format, args...),
    code:  code,
  }
}

// WrapCode wraps error with code and text.
// It returns nil if given err is nil.
func WrapCode(code fcode.Code, err error, text ...string) error {
  if err == nil {
    return nil
  }
  return &Error{
    error: err,
    stack: callers(),
    text:  strings.Join(text, commaSeparatorSpace),
    code:  code,
  }
}

// WrapCodef wraps error with code and format specifier.
// It returns nil if given `err` is nil.
func WrapCodef(code fcode.Code, err error, format string, args ...interface{}) error {
  if err == nil {
    return nil
  }
  return &Error{
    error: err,
    stack: callers(),
    text:  fmt.Sprintf(format, args...),
    code:  code,
  }
}

// WrapCodeSkip wraps error with code and text.
// It returns nil if given err is nil.
// The parameter `skip` specifies the stack callers skipped amount.
func WrapCodeSkip(code fcode.Code, skip int, err error, text ...string) error {
  if err == nil {
    return nil
  }
  return &Error{
    error: err,
    stack: callers(skip),
    text:  strings.Join(text, commaSeparatorSpace),
    code:  code,
  }
}

// WrapCodeSkipf wraps error with code and text that is formatted with given format and args.
// It returns nil if given err is nil.
// The parameter `skip` specifies the stack callers skipped amount.
func WrapCodeSkipf(
  code fcode.Code, skip int, err error, format string, args ...interface{},
) error {
  if err == nil {
    return nil
  }
  return &Error{
    error: err,
    stack: callers(skip),
    text:  fmt.Sprintf(format, args...),
    code:  code,
  }
}

// Code returns the error code of current error.
// It returns `CodeNil` if it has no error code neither it does not implement interface Code.
func Code(err error) fcode.Code {
  if err == nil {
    return fcode.CodeNil
  }
  if e, ok := err.(ICode); ok {
    return e.Code()
  }
  if e, ok := err.(IUnwrap); ok {
    return Code(e.Unwrap())
  }
  return fcode.CodeNil
}

// HasCode checks and reports whether `err` has `code` in its chaining errors.
func HasCode(err error, code fcode.Code) bool {
  if err == nil {
    return false
  }
  if e, ok := err.(ICode); ok {
    return code == e.Code()
  }
  if e, ok := err.(IUnwrap); ok {
    return HasCode(e.Unwrap(), code)
  }
  return false
}
