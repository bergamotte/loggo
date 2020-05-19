package conf

import (
  "github.com/tarrynn/loggo/error"
  "gopkg.in/yaml.v2"
  "io/ioutil"
)

type Conf struct {
  Inputs map[string][]string `yaml:"input"`
  Outputs map[string][]string `yaml:"output"`
  Config map[string]int `yaml:config`
}

func (c *Conf) GetConf(path string) *Conf {
  yamlFile, err := ioutil.ReadFile(path)
  error.Check(err)

  err = yaml.Unmarshal(yamlFile, c)
  error.Check(err)

  return c
}
