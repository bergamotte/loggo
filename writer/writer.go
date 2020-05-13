package writer

import (
  "encoding/json"
  "fmt"
  "github.com/elastic/go-elasticsearch/v7"
  "github.com/tarrynn/loggo/error"
  "os"
  "sync"
  "strings"
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
  f.WriteString("[" + hostname + "] " + logname[len(logname)-1] + ": " + msg + "\n")
}

func CreateIndex(path string) {
  cfg := elasticsearch.Config{
    Addresses: []string{
      path,
    },
  }
  es, _ := elasticsearch.NewClient(cfg)
  res, err := es.Indices.Create(
		"logs",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "settings": {
		    "number_of_shards": 1
		  },
		  "mappings": {
		    "properties": {
		      "log": { "type": "text" },
          "message": { "type": "text" },
          "hostname": { "type": "text" },
          "timestamp": { "type": "date", "format": "yyyy-MM-dd HH:mm:ss" }
		    }
		  }
		}`)),
	)
	fmt.Println(res, err)
	if err != nil {
		fmt.Println(err)
	}
}

func jsonEscape(i string) string {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	s := string(b)
	return s[1:len(s)-1]
}

func WriteToElastic(path string, hostname string, log string, msg string) {
  cfg := elasticsearch.Config{
    Addresses: []string{
      path,
    },
  }
  es, _ := elasticsearch.NewClient(cfg)

  logname := strings.Split(log, "/")

  res, err := es.Index(
		"logs",
		strings.NewReader(`{
		  "log": "`+ logname[len(logname)-1] +`",
      "hostname": "`+ hostname +`",
      "timestamp": "`+ time.Now().Format("2006-01-02 15:04:05") +`",
		  "message": "`+ jsonEscape(msg) +`"
		}`),
		es.Index.WithPretty(),
	)

	if err != nil {
		fmt.Println("Error getting response: %s", err)
	}

  //fmt.Println(res, " - message was: ", msg)
	defer res.Body.Close()
}
