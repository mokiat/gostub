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

func GetFieldsAsAnonymous(fields []*ast.Field) []*ast.Field {
	result := make([]*ast.Field, len(fields))
	for i, field := range fields {
		result[i] = &ast.Field{
			Type: field.Type,
		}
	}
	return result
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

func EachParamInFunc(funcType *ast.FuncType) <-chan *ast.Field {
	result := make(chan *ast.Field)
	go func() {
		if funcType.Params != nil {
			for _, param := range funcType.Params.List {
				result <- param
			}
		}
		close(result)
	}()
	return result
}

func FuncTypeResultCount(funcType *ast.FuncType) int {
	if funcType.Results == nil {
		return 0
	}
	return len(funcType.Results.List)
}

func EachResultInFunc(funcType *ast.FuncType) <-chan *ast.Field {
	result := make(chan *ast.Field)
	go func() {
		if funcType.Results != nil {
			for _, param := range funcType.Results.List {
				result <- param
			}
		}
		close(result)
	}()
	return result
}

func EachDeclarationInFile(file *ast.File) <-chan ast.Decl {
	result := make(chan ast.Decl)
	go func() {
		for _, decl := range file.Decls {
			result <- decl
		}
		close(result)
	}()
	return result
}

func EachGenericDeclarationInFile(file *ast.File) <-chan *ast.GenDecl {
	result := make(chan *ast.GenDecl)
	go func() {
		for decl := range EachDeclarationInFile(file) {
			if genDecl, ok := decl.(*ast.GenDecl); ok {
				result <- genDecl
			}
		}
		close(result)
	}()
	return result
}

func EachSpecificationInGenericDeclaration(decl *ast.GenDecl) <-chan ast.Spec {
	result := make(chan ast.Spec)
	go func() {
		for _, spec := range decl.Specs {
			result <- spec
		}
		close(result)
	}()
	return result
}

func EachTypeSpecificationInGenericDeclaration(decl *ast.GenDecl) <-chan *ast.TypeSpec {
	result := make(chan *ast.TypeSpec)
	go func() {
		for spec := range EachSpecificationInGenericDeclaration(decl) {
			if typeSpec, ok := spec.(*ast.TypeSpec); ok {
				result <- typeSpec
			}
		}
		close(result)
	}()
	return result
}

func EachTypeSpecificationInFile(file *ast.File) <-chan *ast.TypeSpec {
	result := make(chan *ast.TypeSpec)
	go func() {
		for decl := range EachGenericDeclarationInFile(file) {
			for spec := range EachTypeSpecificationInGenericDeclaration(decl) {
				result <- spec
			}
		}
		close(result)
	}()
	return result
}

func EachInterfaceDeclarationInFile(file *ast.File) <-chan *ast.TypeSpec {
	result := make(chan *ast.TypeSpec)
	go func() {
		for spec := range EachTypeSpecificationInFile(file) {
			if _, ok := spec.Type.(*ast.InterfaceType); ok {
				result <- spec
			}
		}
		close(result)
	}()
	return result
}

func EachMethodInInterfaceType(iFaceType *ast.InterfaceType) <-chan *ast.Field {
	result := make(chan *ast.Field)
	go func() {
		if iFaceType.Methods != nil {
			for _, method := range iFaceType.Methods.List {
				if _, ok := method.Type.(*ast.FuncType); ok {
					result <- method
				}
			}
		}
		close(result)
	}()
	return result
}

func EachSubInterfaceInInterfaceType(iFaceType *ast.InterfaceType) <-chan *ast.Field {
	result := make(chan *ast.Field)
	go func() {
		if iFaceType.Methods != nil {
			for _, method := range iFaceType.Methods.List {
				if _, ok := method.Type.(*ast.SelectorExpr); ok {
					result <- method
				}
			}
		}
		close(result)
	}()
	return result
}
