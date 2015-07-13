package generator

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/momchil-atanasov/gostub/util"
)

func NewLocator() *Locator {
	return &Locator{}
}

type Locator struct {
}

type TypeDiscovery struct {
	Location string
	Package  *ast.Package
	File     *ast.File
	Spec     *ast.TypeSpec
}

func (l *Locator) FindTypeDeclarationInLocation(name string, location string) (TypeDiscovery, bool, error) {
	sourcePath, err := util.ImportToDir(location)
	if err != nil {
		return TypeDiscovery{}, false, err
	}

	pkgs, err := parser.ParseDir(token.NewFileSet(), sourcePath, nil, parser.AllErrors)
	if err != nil {
		return TypeDiscovery{}, false, err
	}

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for spec := range util.EachTypeSpecificationInFile(file) {
				if spec.Name.String() == name {
					discovery := TypeDiscovery{
						Location: location,
						Package:  pkg,
						File:     file,
						Spec:     spec,
					}
					return discovery, true, nil
				}
			}
		}
	}
	return TypeDiscovery{}, false, nil
}

func (l *Locator) FindTypeDeclarationInLocations(name string, candidateLocations []string) (TypeDiscovery, bool, error) {
	for _, location := range candidateLocations {
		discovery, found, err := l.FindTypeDeclarationInLocation(name, location)
		if err != nil {
			return TypeDiscovery{}, false, err
		}
		if found {
			return discovery, true, nil
		}
	}
	return TypeDiscovery{}, false, nil
}
