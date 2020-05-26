package writer

import (
  "fmt"
  "github.com/elastic/go-elasticsearch/v6"
  "github.com/tarrynn/loggo/error"
  "strings"
  "time"
)

func NewElasticConn(path []string) elasticsearch.Client {
  defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered from ", r)
    }
  }()

  elastic, err := elasticsearch.NewClient(elasticsearch.Config{
    Addresses: path,
  })

  error.Check(err)

  return *elastic
}

func CreateIndex(es elasticsearch.Client) {
  defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered from ", r)
    }
  }()

  _, err := es.Indices.Create(
		"logs",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "settings": {
		    "number_of_shards": 1
		  },
		  "mappings": {
        "_doc": {
          "properties": {
  		      "log": { "type": "text" },
            "message": { "type": "text" },
            "hostname": { "type": "text" },
            "timestamp": { "type": "date", "format": "yyyy-MM-dd HH:mm:ss.SSS" }
  		    }
        }
		  }
		}`)),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func DeleteOldEntries(es elasticsearch.Client) {
  defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered from ", r)
    }
  }()

  res, err := es.DeleteByQuery(
		[]string{"logs"},
		strings.NewReader(`{
		  "query": {
		    "range": {
		      "timestamp": {
		        "lt": "now-5d",
            "format": "yyyy-MM-dd HH:mm:ss.SSS"
		      }
		    }
		  }
		}`),
	)

  if err != nil {
		fmt.Println(err)
	}

  defer res.Body.Close()
}

func WriteToElastic(es elasticsearch.Client, hostname string, log string, msg string) {
  defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered from ", r)
    }
  }()

  logname := strings.Split(log, "/")

  res, err := es.Index(
		"logs",
		strings.NewReader(`{
		  "log": "`+ logname[len(logname)-1] +`",
      "hostname": "`+ hostname +`",
      "timestamp": "`+ time.Now().Format("2006-01-02 15:04:05.000") +`",
		  "message": "`+ jsonEscape(msg) +`"
		}`),
		es.Index.WithPretty(),
	)

	if err != nil {
		fmt.Println("Error getting response: ", err)
	}

	defer res.Body.Close()
}
