package generator

import (
	"go/ast"
	"go/token"
)

func NewStructModel() *StructModel {
	return &StructModel{}
}

type StructModel struct {
	name   string
	fields []*ast.Field
}

func (m *StructModel) SetName(name string) {
	m.name = name
}

func (m *StructModel) AddField(field *ast.Field) {
	m.fields = append(m.fields, field)
}

func (m *StructModel) BuildASTStructDeclaration() *ast.GenDecl {
	structType := &ast.StructType{
		Fields: &ast.FieldList{},
	}
	for _, field := range m.fields {
		structType.Fields.List = append(structType.Fields.List, field)
	}
	return &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ast.NewIdent(m.name),
				Type: structType,
			},
		},
	}
}
