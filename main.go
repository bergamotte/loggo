package main

import (
    "os"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "./pkg/print"
)

type conf struct {
    Inputs map[string][]string `yaml:"input"`
    Outputs map[string][]string `yaml:"output"`
}

func (c *conf) getConf(path string) *conf {

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

func main() {
    var c conf
    configPath := os.Args[1]
    c.getConf(configPath)

    print.PrintFiles(c.Inputs["files"])
    print.PrintFiles(c.Outputs["files"])
}
