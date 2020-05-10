package tailer

import (
  "fmt"
  "github.com/hpcloud/tail"
  "error"
  "sync"
)

func Init (file string, wg *sync.WaitGroup) {
  defer wg.Done()
  t, err := tail.TailFile(file, tail.Config{Follow: true})
  error.Check(err)
  for line := range t.Lines {
      fmt.Println(line.Text)
  }
}
