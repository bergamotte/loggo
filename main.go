package main

import (
	"config"
	"exit"
	"flag"
	"fmt"
	"print"
	"tailer"
  "sync"
)

func main() {
	configPath := flag.String("config", "./config/config.yaml", "path to config yaml file")
	pidPath := flag.String("pid", "./tmp/loggo.pid", "path to pid file")
	flag.Parse()

	exit.SetupExitListener(*pidPath)

	var config config.Conf
	config.GetConf(*configPath)

	fmt.Println("Inputs detected:")
	print.PrintFiles(config.Inputs["files"])
	fmt.Println("Outputs detected:")
	print.PrintFiles(config.Outputs["files"])

  var wg sync.WaitGroup
  wg.Add(len(config.Inputs["files"]))
	for _, file := range config.Inputs["files"] {
		go tailer.Init(file, &wg, config.Outputs["files"])
	}

  wg.Wait()
}
