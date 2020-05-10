package writer

import (
    "os"
    "sync"
    "time"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func WriteToFile(path string, channel <-chan string, wg *sync.WaitGroup) {
  defer wg.Done()
  f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  check(err)
  defer f.Close()

  for a := 0; a < 6; a++{
		m, more := <-channel
		f.WriteString("[" + time.Now().Format("2006-01-02 15:04:05.000000") + "] " + m + "\n")
		if more == false {
      f.Sync()
			return
		}
	}
}
