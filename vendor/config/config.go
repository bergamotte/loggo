package config

import (
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "error"
)

type Conf struct {
  Inputs map[string][]string `yaml:"input"`
  Outputs map[string][]string `yaml:"output"`
}

func (c *Conf) GetConf(path string) *Conf {
  yamlFile, err := ioutil.ReadFile(path)
  error.Check(err)

  err = yaml.Unmarshal(yamlFile, c)
  error.Check(err)

  return c
}
