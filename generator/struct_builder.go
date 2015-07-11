package generator

import (
	"go/ast"
	"go/token"
)

func NewStructBuilder() *StructBuilder {
	return &StructBuilder{
		fields: make([]*ast.Field, 0),
	}
}

type StructBuilder struct {
	name   string
	fields []*ast.Field
}

func (m *StructBuilder) SetName(name string) {
	m.name = name
}

func (m *StructBuilder) AddField(field *ast.Field) {
	m.fields = append(m.fields, field)
}

func (m *StructBuilder) Build() *ast.GenDecl {
	return &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ast.NewIdent(m.name),
				Type: &ast.StructType{
					Fields: &ast.FieldList{
						List: m.fields,
					},
				},
			},
		},
	}
}
