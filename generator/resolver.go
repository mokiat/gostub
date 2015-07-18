package generator

import (
	"errors"
	"fmt"
	"go/ast"
	"strings"

	"github.com/momchil-atanasov/gostub/util"
)

func NewResolver(model *GeneratorModel, locator *Locator) *Resolver {
	return &Resolver{
		model:   model,
		locator: locator,
		imports: make([]importEntry, 0),
	}
}

type importEntry struct {
	Alias    string
	Location string
}

type Resolver struct {
	model   *GeneratorModel
	locator *Locator
	astFile *ast.File
	imports []importEntry
}

func (r *Resolver) SetContext(astFile *ast.File, fileLocation string) {
	r.imports = []importEntry{}
	for decl := range util.EachGenericDeclarationInFile(astFile) {
		for spec := range util.EachSpecificationInGenericDeclaration(decl) {
			if importSpec, ok := spec.(*ast.ImportSpec); ok {
				imp := importEntry{}
				if importSpec.Name != nil {
					imp.Alias = importSpec.Name.String()
				}
				imp.Location = strings.Trim(importSpec.Path.Value, "\"")
				r.imports = append(r.imports, imp)
			}
		}
	}
}

func (r *Resolver) ResolveType(astType ast.Expr) (ast.Expr, error) {
	switch t := astType.(type) {
	case *ast.Ident:
		if r.isBuiltIn(t.String()) {
			return t, nil
		}
		locations := r.findPotentialLocations(".")
		discovery, found, err := r.locator.FindTypeDeclarationInLocations(t.String(), locations)
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, errors.New(fmt.Sprintf("Type '%s' not found.", t.String()))
		}
		al := r.model.AddImport("", discovery.Location)
		return &ast.SelectorExpr{
			X:   ast.NewIdent(al),
			Sel: ast.NewIdent(t.String()),
		}, nil
	case *ast.SelectorExpr:
		if alias, ok := t.X.(*ast.Ident); ok {
			locations := r.findPotentialLocations(alias.String())
			discovery, found, err := r.locator.FindTypeDeclarationInLocations(t.Sel.String(), locations)
			if err != nil {
				return nil, err
			}
			if !found {
				return nil, errors.New(fmt.Sprintf("Type '%s' not found.", t.Sel.String()))
			}
			al := r.model.AddImport("", discovery.Location)
			return &ast.SelectorExpr{
				X:   ast.NewIdent(al),
				Sel: t.Sel,
			}, nil
		}
		return astType, nil
	case *ast.ArrayType:
		return r.resolveArrayType(t)
	}
	return astType, nil
}

func (r *Resolver) resolveArrayType(astType *ast.ArrayType) (ast.Expr, error) {
	var err error
	astType.Elt, err = r.ResolveType(astType.Elt)
	return astType, err
}

func (r *Resolver) isBuiltIn(name string) bool {
	// Either builtin or private (which is not supported either way)
	return strings.ToLower(name) == name
}

func (r *Resolver) findPotentialLocations(alias string) []string {
	for _, imp := range r.imports {
		if imp.Alias == alias {
			return []string{imp.Location}
		}
	}

	result := []string{}
	for _, imp := range r.imports {
		result = append(result, imp.Location)
	}
	return result
}
