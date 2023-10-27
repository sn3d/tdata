package tdata

import (
	"path/filepath"
)

// returns you absolute path to given path in temporary directory.
func Abs(path ...string) string {
	absPath := append([]string{Pwd()}, path...)
	return filepath.Join(absPath...)
}

// Print working dir. In this case it will print root of
// temp directory
//
// Deprecated: use Testdata.Pwd()
func Pwd() string {
	return globalTestdata.Pwd()
}

// compare content o files, function returns true if they're
// matching, otherwise it returns false.
//
// the a and b are relative paths to need to temporary testdata
// directory
//
// The false is returned also when any error occurs.
//
// The function  isn't optimal but it serves only for
// testing purposes
//
// Deprecated: use Testdata.CompareFiles() instead
func CompareFiles(a, b string) bool {
	return globalTestdata.CompareFiles(a, b)
}
