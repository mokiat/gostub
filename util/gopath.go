package util

import (
	"errors"
	"go/build"
	"os"

	"golang.org/x/tools/go/packages"
)

// DirToImport converts a directory path on the local machine to a
// Go import path (usually relative to the $GOPATH/src directory)
//
// For example,
//     /Users/user/workspace/Go/github.com/mokiat/gostub
// will be converted to
//     github.com/mokiat/gostub
// should GOPATH include the location
//     /Users/user/workspace/Go
func DirToImport(p string) (string, error) {
	pkgs, err := packages.Load(nil, p)
	if err != nil {
		return "", err
	}
	if len(pkgs) > 1 {
		return "", errors.New("no packages could be found")
	}
	return pkgs[0].PkgPath, nil
}

// ImportToDir converts an import location to a directory path on
// the local machine.
//
// For example,
//     github.com/mokiat/gostub
// will be converted to
//     /Users/user/workspace/Go/github.com/mokiat/gostub
// should GOPATH be equal to
//     /Users/user/workspace/Go
func ImportToDir(imp string) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	pkg, err := build.Import(imp, pwd, build.FindOnly)
	if err != nil {
		return "", err
	}
	return pkg.Dir, nil
}
