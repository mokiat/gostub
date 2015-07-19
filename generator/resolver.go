package generator

import (
	"go/ast"
	"strings"

	"github.com/momchil-atanasov/gostub/resolution"
	"github.com/momchil-atanasov/gostub/util"
)

func NewResolver(model *GeneratorModel, locator *resolution.Locator) *Resolver {
	return &Resolver{
		model:   model,
		locator: locator,
	}
}

type importEntry struct {
	Alias    string
	Location string
}

type Resolver struct {
	model   *GeneratorModel
	locator *resolution.Locator
	context *resolution.LocatorContext
}

func (r *Resolver) SetLocatorContext(context *resolution.LocatorContext) {
	r.context = context
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
	case *ast.FuncType:
		return r.resolveFuncType(t)
	case *ast.StructType:
		return r.resolveStructType(t)
	case *ast.InterfaceType:
		return r.resolveInterfaceType(t)
	}
	return astType, nil
}

func (r *Resolver) resolveIdent(ident *ast.Ident) (ast.Expr, error) {
	if r.isBuiltIn(ident.String()) {
		return ident, nil
	}
	discovery, err := r.locator.FindIdentType(r.context, ident)
	if err != nil {
		return nil, err
	}
	al := r.model.AddImport("", discovery.Location)
	return &ast.SelectorExpr{
		X:   ast.NewIdent(al),
		Sel: ast.NewIdent(ident.String()),
	}, nil
}

func (r *Resolver) resolveSelectorExpr(expr *ast.SelectorExpr) (ast.Expr, error) {
	if _, ok := expr.X.(*ast.Ident); ok {
		discovery, err := r.locator.FindSelectorType(r.context, expr)
		if err != nil {
			return nil, err
		}
		al := r.model.AddImport("", discovery.Location)
		return &ast.SelectorExpr{
			X:   ast.NewIdent(al),
			Sel: expr.Sel,
		}, nil
	}
	// TODO: Maybe return an error
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

func (r *Resolver) resolveFuncType(astType *ast.FuncType) (ast.Expr, error) {
	var err error
	for param := range util.EachParamInFunc(astType) {
		param.Type, err = r.ResolveType(param.Type)
		if err != nil {
			return nil, err
		}
	}
	for result := range util.EachResultInFunc(astType) {
		result.Type, err = r.ResolveType(result.Type)
		if err != nil {
			return nil, err
		}
	}
	return astType, nil
}

func (r *Resolver) resolveStructType(astType *ast.StructType) (ast.Expr, error) {
	var err error
	for field := range util.EachFieldInStruct(astType) {
		field.Type, err = r.ResolveType(field.Type)
		if err != nil {
			return nil, err
		}
	}
	return astType, nil
}

func (r *Resolver) resolveInterfaceType(astType *ast.InterfaceType) (ast.Expr, error) {
	var err error
	for field := range util.EachFieldInInterface(astType) {
		field.Type, err = r.ResolveType(field.Type)
		if err != nil {
			return nil, err
		}
	}
	return astType, nil
}

func (r *Resolver) isBuiltIn(name string) bool {
	// Either builtin or private (which is not supported either way)
	return strings.ToLower(name) == name
}
