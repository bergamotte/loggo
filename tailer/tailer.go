package tailer

import (
  "github.com/hpcloud/tail"
  "github.com/tarrynn/loggo/error"
  "github.com/tarrynn/loggo/writer"
  "io"
  "sync"
)

func position (full bool) int {
  if full {
    return io.SeekStart
  } else {
    return io.SeekEnd
  }
}

func Init (file string, wg *sync.WaitGroup, dest map[string][]string, full bool) {
  defer wg.Done()

  seekInfo := &tail.SeekInfo{ Offset: 0, Whence: position(full) }

  t, err := tail.TailFile(file, tail.Config{Follow: true, Location: seekInfo})
  error.Check(err)
  for line := range t.Lines {
    for key, values := range dest {
      if key == "files" {
        for _, file := range values {
          writer.WriteToFile(file, line.Text)
        }
      }

      if key == "elasticsearch" {
        for _, path := range values {
          writer.WriteToElastic(path, line.Text)
        }
      }
    }

  }
}
