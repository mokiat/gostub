package util

import "go/ast"

func FieldTypeReuseCount(field *ast.Field) int {
	if len(field.Names) == 0 {
		return 1
	}
	return len(field.Names)
}

func EachFieldInFieldList(fieldList *ast.FieldList) <-chan *ast.Field {
	result := make(chan *ast.Field)
	go func() {
		if fieldList == nil {
			close(result)
			return
		}
		for _, field := range fieldList.List {
			result <- field
		}
		close(result)
	}()
	return result
}
