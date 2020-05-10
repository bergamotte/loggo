package reader

import (
  "bufio"
  "log"
  "os"
  "sync"
)

func ReadFile(path string, channel *chan string, wg *sync.WaitGroup) {
  defer wg.Done()
  readFile, err := os.Open(path)

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
    *channel <- path + ": " + fileScanner.Text()
	}

	readFile.Close()
}
