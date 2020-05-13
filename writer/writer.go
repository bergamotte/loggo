package writer

import (
  "fmt"
  "github.com/elastic/go-elasticsearch/v7"
  "github.com/tarrynn/loggo/error"
  "os"
  "sync"
  "strings"
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

func CreateIndex(path string) {
  cfg := elasticsearch.Config{
    Addresses: []string{
      path,
    },
  }
  es, _ := elasticsearch.NewClient(cfg)
	res, err := es.Indices.Create("logs")
	fmt.Println(res, err)
	if err != nil {
		fmt.Println(err)
	}
}

func WriteToElastic(path string, msg string) {
  cfg := elasticsearch.Config{
    Addresses: []string{
      path,
    },
  }
  es, _ := elasticsearch.NewClient(cfg)

  res, err := es.Index(
		"logs",
		strings.NewReader(`{
		  "user": "tarrynn",
		  "message": "`+ msg +`"
		}`),
		es.Index.WithPretty(),
	)

	if err != nil {
		fmt.Println("Error getting response: %s", err)
	}

  fmt.Println(res)
	defer res.Body.Close()
}
