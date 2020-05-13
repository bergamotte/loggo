package exit

import (
  "fmt"
  "os"
	"os/signal"
  "github.com/tarrynn/loggo/pid"
	"syscall"
)

func SetupExitListener(pidPath string) {
  pid.CreatePidFile(pidPath)

  c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\rTerminating...")
    pid.RemovePidFile(pidPath)
		os.Exit(0)
	}()
}
