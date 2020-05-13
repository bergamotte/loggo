package daemon

import (
	"fmt"
	"os"
	"os/exec"
)

func Start(args []string) {
		i := 0
		for ; i < len(args); i++ {
			if args[i] == "-daemon=true" {
				args[i] = "-daemon=false"
				break
			}
		}
		cmd := exec.Command(os.Args[0], args...)
		cmd.Start()
		fmt.Println("[PID]", cmd.Process.Pid)
		os.Exit(0)
}
