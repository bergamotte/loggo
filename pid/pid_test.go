package pid_test

import (
  "os"
  "github.com/tarrynn/loggo/pid"
  "testing"
)

func TestPidFile(t *testing.T) {
  pid.CreatePidFile("test.txt")
  _, err := os.Stat("test.txt")
  if err != nil {
    t.Errorf("Didn't create pid file")
  }

  pid.RemovePidFile("test.txt")
  _, err2 := os.Stat("test.txt")
  if err2 == nil {
    t.Errorf("Didn't delete pid file")
  }
}
