package tailer

import (
  "github.com/hpcloud/tail"
  "github.com/tarrynn/loggo/error"
  "github.com/tarrynn/loggo/writer"
  "io"
  "os"
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

  channel := make(chan string, 200)

  hostname, err := os.Hostname()
  error.Check(err)

  seekInfo := &tail.SeekInfo{ Offset: 0, Whence: position(full) }
  t, err := tail.TailFile(file, tail.Config{Follow: true, Location: seekInfo})
  error.Check(err)

  go writer.Write(dest, hostname, file, channel)

  for line := range t.Lines {
    for {
      if len(channel) != cap(channel) {
        channel <- line.Text
        break
      }
    }
  }

  close(channel)
}
