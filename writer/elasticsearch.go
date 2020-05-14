package writer

import (
  "encoding/json"
  "fmt"
  "github.com/elastic/go-elasticsearch/v7"
  "github.com/tarrynn/loggo/error"
  "strings"
  "time"
)

func jsonEscape(i string) string {
	b, err := json.Marshal(i)
	error.Check(err)
	s := string(b)
	return s[1:len(s)-1]
}

func CreateIndex(path string) {
  cfg := elasticsearch.Config{
    Addresses: []string{
      path,
    },
  }
  es, _ := elasticsearch.NewClient(cfg)
  _, err := es.Indices.Create(
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
	//fmt.Println(res, err)
	if err != nil {
		fmt.Println(err)
	}
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
		fmt.Println("Error getting response: ", err)
	}

  //fmt.Println(res, " - message was: ", msg)
	defer res.Body.Close()
}
