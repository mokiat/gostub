package generator

import "go/ast"

func NewMethodModel() *MethodModel {
	return &MethodModel{}
}

type MethodModel struct {
	name         string
	funcType     *ast.FuncType
	receiverName string
	receiverType string
	statements   []ast.Stmt
}

func (m *MethodModel) SetName(name string) {
	m.name = name
}

func (m *MethodModel) SetReceiver(name, recType string) {
	m.receiverName = name
	m.receiverType = recType
}

func (m *MethodModel) SetType(funcType *ast.FuncType) {
	m.funcType = funcType
}

func (m *MethodModel) AddStatement(statement ast.Stmt) {
	m.statements = append(m.statements, statement)
}

func (m *MethodModel) BuildASTFuncDecl() *ast.FuncDecl {
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
