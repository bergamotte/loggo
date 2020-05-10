package config

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

type Conf struct {
    Inputs map[string][]string `yaml:"input"`
    Outputs map[string][]string `yaml:"output"`
}

func (c *Conf) GetConf(path string) *Conf {

    yamlFile, err := ioutil.ReadFile(path)
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return c
}
