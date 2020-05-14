package writer

func Write(dest map[string][]string, hostname string, source string, line string) {
  for key, values := range dest {
    if key == "files" {
      for _, out := range values {
        WriteToFile(out, hostname, source, line)
      }
    }

    if key == "elasticsearch" {
      for _, path := range values {
        WriteToElastic(path, hostname, source, line)
      }
    }
  }
}
