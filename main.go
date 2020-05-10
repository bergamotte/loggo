package main

import (
    "flag"
    "./pkg/config"
    "./pkg/print"
)

func main() {
    var c config.Conf
    configPath := flag.String("config", "./config/config.yaml", "path to config yaml file")
    flag.Parse()
    c.GetConf(*configPath)

    print.PrintFiles(c.Inputs["files"])
    print.PrintFiles(c.Outputs["files"])
}
