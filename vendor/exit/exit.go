package exit

import (
  "fmt"
  "os"
	"os/signal"
  "pid"
	"syscall"
)

func SetupExitListener(pidPath string) {
  pid.CreatePidFile(pidPath)

  c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Terminating...")
    pid.RemovePidFile(pidPath)
		os.Exit(0)
	}()
}
