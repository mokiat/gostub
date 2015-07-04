package util

import "go/ast"

func CreateField(name string, fieldType ast.Expr) *ast.Field {
	return &ast.Field{
		Names: []*ast.Ident{
			ast.NewIdent(name),
		},
		Type: fieldType,
	}
}

func FieldReuseCount(field *ast.Field) int {
	if len(field.Names) == 0 {
		return 1
	}
	return len(field.Names)
}

func CreateFuncType() *ast.FuncType {
	return &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{},
		},
		Results: &ast.FieldList{
			List: []*ast.Field{},
		},
	}
}

func FuncTypeParamCount(funcType *ast.FuncType) int {
	if funcType.Params == nil {
		return 0
	}
	return len(funcType.Params.List)
}

func FuncTypeResultCount(funcType *ast.FuncType) int {
	if funcType.Results == nil {
		return 0
	}
	return len(funcType.Results.List)
}
