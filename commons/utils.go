package commons

import (
	"os"
	"encoding/json"
)

// read json file and parse into struct with standard go library
func ReadJSON(filename string, v any) error {
  file, err := os.Open(filename)
  defer file.Close()

  if err!= nil {
    return err
  }

  json.NewDecoder(file).Decode(v)

  return nil
}

