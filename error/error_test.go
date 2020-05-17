package error_test

import (
  "errors"
  "github.com/tarrynn/loggo/error"
  "testing"
)

func TestCheck (t *testing.T) {
  e := errors.New("new error")
  defer func() {
    r := recover()
    if r == nil {
      t.Errorf("The code did not panic")
    }
  }()

  error.Check(e)
}

func TestCheck_2 (t *testing.T) {
  defer func() {
    r := recover()
    if r != nil {
      t.Errorf("The code panicked!")
    }
  }()

  error.Check(nil)
}
