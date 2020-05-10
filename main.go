package main

import (
	"config"
	"flag"
	"fmt"
	"print"
	"tailer"
  "sync"
)

func main() {
	var c config.Conf
	configPath := flag.String("config", "./config/config.yaml", "path to config yaml file")
	flag.Parse()
	c.GetConf(*configPath)

	fmt.Println("Inputs detected:")
	print.PrintFiles(c.Inputs["files"])
	fmt.Println("Outputs detected:")
	print.PrintFiles(c.Outputs["files"])

  var wg sync.WaitGroup
  wg.Add(len(c.Inputs["files"]))
	for _, file := range c.Inputs["files"] {
		go tailer.Init(file, &wg)
	}

  wg.Wait()
}
