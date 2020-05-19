package writer

import (
  "encoding/json"
  "github.com/tarrynn/loggo/error"
)

func jsonEscape(i string) string {
	b, err := json.Marshal(i)
	error.Check(err)
	s := string(b)
	return s[1:len(s)-1]
}

func Write(dest map[string][]string, hostname string, source string, channel <-chan string) {
  for line := range channel {
    for key, values := range dest {
      if key == "files" {
        for _, out := range values {
          WriteToFile(out, hostname, source, line)
        }
      }

      if key == "elasticsearch" {
        elastic := NewElasticConn(values)
        WriteToElastic(elastic, hostname, source, line)
      }

      if key == "redis" {
        for _, path := range values {
          WriteToRedis(path, hostname, source, line)
        }
      }
    }
  }
}
