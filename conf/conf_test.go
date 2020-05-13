package conf_test

import (
  "github.com/tarrynn/loggo/conf"
  //"os"
  "testing"
)

func TestGetConf(t *testing.T) {
  testCases := []struct {
    Id int
    Name string
    Expected string
    File string
  }{
    {
      Id: 1,
      Name: "when no config file is present",
      Expected: "panic",
      File: "test.yml",
    },
    {
      Id: 2,
      Name: "when config file is present and returns proper values",
      Expected: "no panic",
      File: "conf_test/config.yaml",
    },
  }

  for _, tc := range testCases {
    tc := tc // capture range variable
    t.Run(tc.Name, func(t *testing.T) {
        t.Parallel()
        defer func() {
          r := recover()
          if r == nil && tc.Expected == "panic" {
            t.Errorf("The code did not panic")
          }

          if r != nil && tc.Expected == "no panic" {
            t.Errorf("The code panicked")
          }
        }()

        var config conf.Conf
        config.GetConf(tc.File)

        // only applies to the second test case where we check the actual config values read from the file
        if tc.Id == 2 {
          if len(config.Inputs["files"]) != 2 || len(config.Outputs["files"]) != 2 {
            t.Errorf("Config file wasn't read properly")
          }
        }
    })
  }
}
