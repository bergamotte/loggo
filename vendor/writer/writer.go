package writer

import (
  "error"
  "os"
  "sync"
)

var mu sync.Mutex

func WriteToFile(filename string, msg string) {
  mu.Lock()
  defer mu.Unlock()

  f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  error.Check(err)
  defer f.Close()

  f.WriteString(msg + "\n")
}
