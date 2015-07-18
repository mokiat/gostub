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
	imports []importEntry
}

func (r *Resolver) SetContext(astFile *ast.File, fileLocation string) {
	r.imports = []importEntry{}
	r.imports = append(r.imports, importEntry{
		Alias:    ".",
		Location: fileLocation,
	})
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
		return r.resolveIdent(t)
	case *ast.SelectorExpr:
		return r.resolveSelectorExpr(t)
	case *ast.ArrayType:
		return r.resolveArrayType(t)
	case *ast.MapType:
		return r.resolveMapType(t)
	case *ast.ChanType:
		return r.resolveChanType(t)
	case *ast.StarExpr:
		return r.resolveStarType(t)
	}
	return astType, nil
}

func (r *Resolver) resolveIdent(ident *ast.Ident) (ast.Expr, error) {
	if r.isBuiltIn(ident.String()) {
		return ident, nil
	}
	locations := r.findPotentialLocations(".")
	discovery, found, err := r.locator.FindTypeDeclarationInLocations(ident.String(), locations)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New(fmt.Sprintf("Type '%s' not found.", ident.String()))
	}
	al := r.model.AddImport("", discovery.Location)
	return &ast.SelectorExpr{
		X:   ast.NewIdent(al),
		Sel: ast.NewIdent(ident.String()),
	}, nil
}

func (r *Resolver) resolveSelectorExpr(expr *ast.SelectorExpr) (ast.Expr, error) {
	if alias, ok := expr.X.(*ast.Ident); ok {
		locations := r.findPotentialLocations(alias.String())
		discovery, found, err := r.locator.FindTypeDeclarationInLocations(expr.Sel.String(), locations)
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, errors.New(fmt.Sprintf("Type '%s' not found.", expr.Sel.String()))
		}
		al := r.model.AddImport("", discovery.Location)
		return &ast.SelectorExpr{
			X:   ast.NewIdent(al),
			Sel: expr.Sel,
		}, nil
	}
	return expr, nil
}

func (r *Resolver) resolveArrayType(astType *ast.ArrayType) (ast.Expr, error) {
	var err error
	astType.Elt, err = r.ResolveType(astType.Elt)
	return astType, err
}

func (r *Resolver) resolveMapType(astType *ast.MapType) (ast.Expr, error) {
	var err error
	astType.Key, err = r.ResolveType(astType.Key)
	if err != nil {
		return nil, err
	}
	astType.Value, err = r.ResolveType(astType.Value)
	if err != nil {
		return nil, err
	}
	return astType, nil
}

func (r *Resolver) resolveChanType(astType *ast.ChanType) (ast.Expr, error) {
	var err error
	astType.Value, err = r.ResolveType(astType.Value)
	return astType, err
}

func (r *Resolver) resolveStarType(astType *ast.StarExpr) (ast.Expr, error) {
	var err error
	astType.X, err = r.ResolveType(astType.X)
	return astType, err
}

func (r *Resolver) isBuiltIn(name string) bool {
	// Either builtin or private (which is not supported either way)
	return strings.ToLower(name) == name
}

func (r *Resolver) findPotentialLocations(alias string) []string {
	if alias == "." {
		result := []string{}
		for _, imp := range r.imports {
			if imp.Alias == "." {
				result = append(result, imp.Location)
			}
		}
		return result
	}
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
