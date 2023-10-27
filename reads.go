package tdata

// returns you file body as []byte. If file cannot be read or
// it doesn't exits, then this function gives you empty byte array
//
// Deprecated: use Testdata.Read() instead
func Read(file string) []byte {
	return globalTestdata.Read(file)
}

// returns you file body as string. If file cannot be read or
// it doesn't exits, then this function gives you empty string
//
// Deprecated: use Testdata.ReadStr() instead
func ReadStr(file string) string {
	return globalTestdata.ReadStr(file)
}

// function load YAML file and fill the data into
// given out
//
// Deprecated: use Testdata.ReadYAML() instead
func ReadYAML(file string, out any) {
	globalTestdata.ReadYAML(file, out)
}

// function load JSON file and fill the data into
// given out
//
// Deprecated: use Testdata.ReadJSON() instead
func ReadJSON(file string, out any) {
	globalTestdata.ReadJSON(file, out)
}
