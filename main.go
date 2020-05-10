package main

import (
    "fmt"
    "flag"
    "sync"
    "./pkg/config"
    "./pkg/print"
    "./pkg/reader"
)

var wg sync.WaitGroup

func main() {
    var c config.Conf
    configPath := flag.String("config", "./config/config.yaml", "path to config yaml file")
    flag.Parse()
    c.GetConf(*configPath)

    fmt.Println("Inputs detected:")
    print.PrintFiles(c.Inputs["files"])
    fmt.Println("Outputs detected:")
    print.PrintFiles(c.Outputs["files"])

    wg.Add(len(c.Inputs["files"]))
    for _, file := range c.Inputs["files"] {
      go reader.ReadFile(file, c.Outputs["files"], &wg)
    }

    wg.Wait()
}
