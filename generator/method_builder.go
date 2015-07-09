package generator

import "go/ast"

func NewMethodBuilder() *MethodBuilder {
	return &MethodBuilder{}
}

type MethodBuilder struct {
	name         string
	funcType     *ast.FuncType
	receiverName string
	receiverType string
	statements   []ast.Stmt
}

func (m *MethodBuilder) SetName(name string) {
	m.name = name
}

func (m *MethodBuilder) SetReceiver(name, recType string) {
	m.receiverName = name
	m.receiverType = recType
}

func (m *MethodBuilder) SetType(funcType *ast.FuncType) {
	m.funcType = funcType
}

func (m *MethodBuilder) AddStatement(statement ast.Stmt) {
	m.statements = append(m.statements, statement)
}

func (m *MethodBuilder) Build() *ast.FuncDecl {
	body := &ast.BlockStmt{
		List: []ast.Stmt{},
	}
	for _, statement := range m.statements {
		body.List = append(body.List, statement)
	}
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				&ast.Field{
					Names: []*ast.Ident{
						ast.NewIdent(m.receiverName),
					},
					Type: &ast.StarExpr{
						X: ast.NewIdent(m.receiverType),
					},
				},
			},
		},
		Name: ast.NewIdent(m.name),
		Type: m.funcType,
		Body: body,
	}
}
