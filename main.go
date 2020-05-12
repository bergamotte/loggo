package main

import (
	"config"
	"daemon"
	"exit"
	"flag"
	"fmt"
	"os"
	"print"
  "sync"
	"tailer"
)

func main() {
	configPath := flag.String("config", "./config/config.yaml", "path to config yaml file")
	pidPath := flag.String("pid", "./tmp/loggo.pid", "path to pid file")
	dmon := flag.Bool("daemon", false, "run as daemon false|true")
	flag.Parse()

	if *dmon {
		daemon.Start(os.Args)
	} else {
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
}
