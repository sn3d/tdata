package tdata

import (
	"encoding/json"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

// returns you file body as []byte. If file cannot be read or
// it doesn't exits, then this function gives you empty byte array
func Read(file string) []byte {
	fileParts := strings.Split(file, "/")
	content, err := os.ReadFile(Abs(fileParts...))
	if err != nil {
		return []byte{}
	} else {
		return content
	}
}

// returns you file body as string. If file cannot be read or
// it doesn't exits, then this function gives you empty string
func ReadStr(file string) string {
	return string(Read(file))
}

// function load YAML file and fill the data into
// given out
func ReadYAML(file string, out any) {
	data := Read(file)
	yaml.Unmarshal(data, out)
}

// function load YAML file and fill the data into
// given out
func ReadJSON(file string, out any) {
	data := Read(file)
	json.Unmarshal(data, out)
}
