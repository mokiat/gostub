package generator

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/momchil-atanasov/gostub/util"
)

func NewLocator() *Locator {
	return &Locator{
		cache: make(map[string][]TypeDiscovery),
	}
}

type Locator struct {
	cache map[string][]TypeDiscovery
}

type TypeDiscovery struct {
	Location string
	File     *ast.File
	Spec     *ast.TypeSpec
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

func (l *Locator) FindTypeDeclarationInLocation(name string, location string) (TypeDiscovery, bool, error) {
	discoveries, err := l.discoverTypes(location)
	if err != nil {
		return TypeDiscovery{}, false, err
	}
	for _, discovery := range discoveries {
		if discovery.Spec.Name.String() == name {
			return discovery, true, nil
		}
	}
	return TypeDiscovery{}, false, nil
}

func (l *Locator) discoverTypes(location string) ([]TypeDiscovery, error) {
	discoveries, found := l.cache[location]
	if found {
		return discoveries, nil
	}

	sourcePath, err := util.ImportToDir(location)
	if err != nil {
		return nil, err
	}

	pkgs, err := parser.ParseDir(token.NewFileSet(), sourcePath, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	discoveries = make([]TypeDiscovery, 0)
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for spec := range util.EachTypeSpecificationInFile(file) {
				discoveries = append(discoveries, TypeDiscovery{
					Location: location,
					File:     file,
					Spec:     spec,
				})
			}
		}
	}
	l.cache[location] = discoveries
	return discoveries, nil
}
