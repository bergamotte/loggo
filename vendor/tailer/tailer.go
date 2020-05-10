package tailer

import (
  "error"
  "github.com/hpcloud/tail"
  "sync"
  "writer"
)

func Init (file string, wg *sync.WaitGroup, dest []string) {
  defer wg.Done()
  t, err := tail.TailFile(file, tail.Config{Follow: true})
  error.Check(err)
  for line := range t.Lines {
    for _, file := range dest {
      writer.WriteToFile(file, line.Text)
    }
  }
}
