package reader

import (
  "bufio"
  "os"
  "sync"
  "time"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func ReadFile(path string, dest []string, wg *sync.WaitGroup) {
  defer wg.Done()

  readFile, err := os.Open(path)
  check(err)
  defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
    for _, file := range dest {
      f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
      check(err)
      f.WriteString("[" + time.Now().Format("2006-01-02 15:04:05.000000") + "] " + path + ": " + fileScanner.Text() + "\n")
    }
	}
}
