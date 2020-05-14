package writer

import (
  "github.com/tarrynn/loggo/error"
  "os"
  "strings"
  "sync"
  "time"
)

var mu sync.Mutex

func WriteToFile(filename string, hostname string, log string, msg string) {
  mu.Lock()
  defer mu.Unlock()

  f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  error.Check(err)
  defer f.Close()

  logname := strings.Split(log, "/")
  f.WriteString("[" + time.Now().Format("2006-01-02 15:04:05") + "] [" + hostname + "] " + logname[len(logname)-1] + ": " + msg + "\n")
}
