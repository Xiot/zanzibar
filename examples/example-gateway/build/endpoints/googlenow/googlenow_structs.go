// Code generated by zanzibar
// @generated

package googlenow

import (
	"os"
	"path/filepath"
	"runtime"
)

func getDirName() string {
	_, file, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(file)

	// Strip _obj dirs generated by test -cover ...
	if filepath.Base(dirname) == "_obj" {
		dirname = filepath.Dir(dirname)
	}

	// Strip _obj_test in go test -cover
	if filepath.Base(dirname) == "_obj_test" {
		dirname = filepath.Dir(dirname)
	}

	// go test -cover does weird folder stuff
	if filepath.Base(dirname) == "_test" {
		dirname = filepath.Dir(dirname)
	}

	// if filepath then we are done, otherwise its go package name
	if filepath.IsAbs(dirname) {
		return dirname
	}

	// If dirname is not absolute then its a package name...
	return filepath.Join(os.Getenv("GOPATH"), "src", dirname)
}

// AddCredentialsHTTPRequest is the http body type for endpoint addCredentials.
type AddCredentialsHTTPRequest struct {
	AuthCode string
}
