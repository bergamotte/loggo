package main

import (
	"github.com/tarrynn/loggo/conf"
	"github.com/tarrynn/loggo/daemon"
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
	dmon := flag.Bool("daemon", false, "run as daemon false|true")
	full := flag.Bool("full", false, "start straight at the end or full index of the input false|true")
	flag.Parse()

	if *dmon {
		daemon.Start(os.Args)
	} else {
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
					if key == "elasticsearch" {
						 writer.CreateIndex(path)
					}
				}
		}

	  var wg sync.WaitGroup
	  wg.Add(len(config.Inputs["files"]))
		for _, file := range config.Inputs["files"] {
			go tailer.Init(file, &wg, config.Outputs, *full)
		}

	  wg.Wait()
	}
}
