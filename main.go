package main

import (
	"github.com/tarrynn/loggo/conf"
	"github.com/tarrynn/loggo/daemon"
	"github.com/tarrynn/loggo/exit"
	"github.com/tarrynn/loggo/print"
	"github.com/tarrynn/loggo/tailer"
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
		print.PrintFiles(config.Inputs["files"])
		fmt.Println("Outputs detected:")
		print.PrintFiles(config.Outputs["files"])

	  var wg sync.WaitGroup
	  wg.Add(len(config.Inputs["files"]))
		for _, file := range config.Inputs["files"] {
			go tailer.Init(file, &wg, config.Outputs["files"], *full)
		}

	  wg.Wait()
	}
}
