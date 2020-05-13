package tailer

import (
  "error"
  "github.com/hpcloud/tail"
  "io"
  "sync"
  "writer"
)

func position (full bool) int {
  if full {
    return io.SeekStart
  } else {
    return io.SeekEnd
  }
}

func Init (file string, wg *sync.WaitGroup, dest []string, full bool) {
  defer wg.Done()

  seekInfo := &tail.SeekInfo{ Offset: 0, Whence: position(full) }

  t, err := tail.TailFile(file, tail.Config{Follow: true, Location: seekInfo})
  error.Check(err)
  for line := range t.Lines {
    for _, file := range dest {
      writer.WriteToFile(file, line.Text)
    }
  }
}
