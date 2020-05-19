package main

import (
	"github.com/tarrynn/loggo/conf"
	"github.com/tarrynn/loggo/exit"
	"github.com/tarrynn/loggo/tailer"
	"github.com/tarrynn/loggo/writer"
	"flag"
	"fmt"
	"os"
  "sync"
)

func main() {
	configPath := flag.String("config", "./config/config.yaml", "path to config yaml file")
	pidPath := flag.String("pid", "./tmp/loggo.pid", "path to pid file")
	full := flag.Bool("full", false, "start straight at the end or full index of the input false|true")
	flag.Parse()

	exit.SetupExitListener(*pidPath)

	var config conf.Conf
	config.GetConf(*configPath)

	fmt.Println("Inputs detected:")
	for key, value := range config.Inputs {
			for _, path := range value {
				fmt.Println(key, " ->", path)
			}
	}
	fmt.Println("Outputs detected:")
	for key, value := range config.Outputs {
		for _, path := range value {
			fmt.Println(key, " ->", path)
		}

		if key == "elasticsearch" {
			elastic := writer.NewElasticConn(value)
			writer.CreateIndex(elastic)
		}

		if key == "redis" {
			writer.NewRedisConn(value[0])
		}
	}

	if len(config.Outputs) != 0 && len(config.Inputs) != 0  {
		var wg sync.WaitGroup
	  wg.Add(len(config.Inputs["files"]))
		for _, file := range config.Inputs["files"] {
			go tailer.Init(file, &wg, config.Outputs, *full)
		}

	  wg.Wait()
	} else {
		fmt.Println("No inputs/outputs detected, terminating.")
		os.Exit(0)
	}
}
