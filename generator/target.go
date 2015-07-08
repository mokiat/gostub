package generator

import (
	"go/ast"
	"go/format"
	"go/token"
	"os"

	"github.com/momchil-atanasov/gostub/util"
)

func newTarget(pkgName, stubName string) *genTarget {
	fileModel := NewFileModel()
	fileModel.SetPackage(pkgName)

	structModel := NewStructModel()
	structModel.SetName(stubName)
	fileModel.AddStructure(structModel)

	return &genTarget{
		fileModel:   fileModel,
		structModel: structModel,
		structName:  stubName,
	}
}

type genTarget struct {
	fileModel   *FileModel
	structModel *StructModel
	structName  string
}

func (t *genTarget) GenerateMethod(source *genSource) error {
	t.createStubField(source)
	t.createMutexField(source)
	t.createArgsForCallField(source)
	t.createMethod(source)
	t.createCallCountMethod(source)
	return nil
}

func (t *genTarget) createStubField(source *genSource) {
	field := util.CreateField(source.StubMethodName(), source.MethodType)
	t.structModel.AddField(field)
}

func (t *genTarget) createMutexField(source *genSource) {
	alias := t.fileModel.AddImport("sync", "sync")
	mutexType := &ast.SelectorExpr{
		X:   ast.NewIdent(alias),
		Sel: ast.NewIdent("RWMutex"),
	}
	field := util.CreateField(source.MutexName(), mutexType)
	t.structModel.AddField(field)
}

func (t *genTarget) createArgsForCallField(source *genSource) {
	argsForCallType := &ast.ArrayType{
		Elt: &ast.StructType{
			Fields: &ast.FieldList{},
		},
	}
	field := util.CreateField(source.ArgsForCallName(), argsForCallType)
	t.structModel.AddField(field)
}

func (t *genTarget) createMethod(source *genSource) {
	method := NewMethodModel()
	method.SetName(source.MethodName)
	method.SetReceiver(source.StructSelfName(), t.structName)
	method.SetType(source.MethodType)
	method.AddStatement(t.createMutexStatement(source, "Lock"))
	method.AddStatement(t.createMutexDeferStatement(source, "Unlock"))
	method.AddStatement(&ast.AssignStmt{
		Lhs: []ast.Expr{
			&ast.SelectorExpr{
				X:   ast.NewIdent(source.StructSelfName()),
				Sel: ast.NewIdent(source.ArgsForCallName()),
			},
		},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: ast.NewIdent("append"),
				Args: []ast.Expr{
					&ast.SelectorExpr{
						X:   ast.NewIdent(source.StructSelfName()),
						Sel: ast.NewIdent(source.ArgsForCallName()),
					},
					&ast.CompositeLit{
						Type: &ast.StructType{
							Fields: &ast.FieldList{},
						},
						Elts: []ast.Expr{},
					},
				},
			},
		},
	})
	method.AddStatement(&ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X: &ast.SelectorExpr{
				X:   ast.NewIdent(source.StructSelfName()),
				Sel: ast.NewIdent(source.StubMethodName()),
			},
			Op: token.NEQ,
			Y:  ast.NewIdent("nil"),
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent(source.StructSelfName()),
							Sel: ast.NewIdent(source.StubMethodName()),
						},
					},
				},
			},
		},
	})
	t.fileModel.AddMethod(method)
}

func (t *genTarget) createCallCountMethod(source *genSource) {
	method := NewMethodModel()
	method.SetName(source.CallCountMethodName())
	method.SetReceiver(source.StructSelfName(), t.structName)
	method.SetType(&ast.FuncType{
		Params: &ast.FieldList{},
		Results: &ast.FieldList{
			List: []*ast.Field{
				&ast.Field{
					Type: ast.NewIdent("int"),
				},
			},
		},
	})
	method.AddStatement(t.createMutexStatement(source, "RLock"))
	method.AddStatement(t.createMutexDeferStatement(source, "RUnlock"))
	method.AddStatement(&ast.ReturnStmt{
		Results: []ast.Expr{
			&ast.CallExpr{
				Fun: ast.NewIdent("len"),
				Args: []ast.Expr{
					&ast.SelectorExpr{
						X:   ast.NewIdent(source.StructSelfName()),
						Sel: ast.NewIdent(source.ArgsForCallName()),
					},
				},
			},
		},
	})
	t.fileModel.AddMethod(method)
}

func (t *genTarget) createMutexStatement(source *genSource, action string) ast.Stmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X: &ast.SelectorExpr{
					X:   ast.NewIdent(source.StructSelfName()),
					Sel: ast.NewIdent(source.MutexName()),
				},
				Sel: ast.NewIdent(action),
			},
		},
	}
}

func (t *genTarget) createMutexDeferStatement(source *genSource, action string) ast.Stmt {
	return &ast.DeferStmt{
		Call: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X: &ast.SelectorExpr{
					X:   ast.NewIdent(source.StructSelfName()),
					Sel: ast.NewIdent(source.MutexName()),
				},
				Sel: ast.NewIdent(action),
			},
		},
	}
}

func (t *genTarget) Save(filePath string) error {
	osFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer osFile.Close()

	astFile := t.fileModel.BuildASTFile()
	err = format.Node(osFile, token.NewFileSet(), astFile)
	if err != nil {
		return err
	}

	return nil
}
