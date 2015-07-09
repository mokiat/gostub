package generator

import (
	"go/ast"
	"go/format"
	"go/token"
	"os"
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
	t.createMethodStubField(source)
	t.createMutexField(source)
	t.createArgsForCallField(source)
	if source.HasResults() {
		t.createReturnsField(source)
	}
	t.createStubMethod(source)
	t.createCallCountMethod(source)
	if source.HasParams() {
		t.createArgsForCallMethod(source)
	}
	if source.HasResults() {
		t.createReturnsMethod(source)
	}
	return nil
}

func (t *genTarget) createMethodStubField(source *genSource) {
	builder := NewMethodStubFieldBuilder()
	builder.SetFieldName(source.StubMethodName())
	builder.SetParams(source.MethodParams)
	builder.SetResults(source.MethodResults)
	t.structModel.AddField(builder.Build())
}

func (t *genTarget) createMutexField(source *genSource) {
	builder := NewMethodMutexFieldBuilder()
	builder.SetFieldName(source.MutexName())
	builder.SetMutexType(t.resolveMutexType())
	t.structModel.AddField(builder.Build())
}

func (t *genTarget) createArgsForCallField(source *genSource) {
	builder := NewMethodArgsFieldBuilder()
	builder.SetFieldName(source.ArgsForCallName())
	builder.SetParams(source.MethodParams)
	t.structModel.AddField(builder.Build())
}

func (t *genTarget) createReturnsField(source *genSource) {
	builder := NewReturnsFieldBuilder()
	builder.SetFieldName(source.ReturnsName())
	builder.SetResults(source.MethodResults)
	t.structModel.AddField(builder.Build())
}

func (t *genTarget) createStubMethod(source *genSource) {
	methodBuilder := t.createMethodBuilder(source, source.MethodName)
	builder := NewStubMethodBuilder(methodBuilder)
	builder.SetMutexFieldSelector(source.MutexFieldSelector())
	builder.SetArgsFieldSelector(source.ArgsFieldSelector())
	builder.SetReturnsFieldSelector(source.ReturnsFieldSelector())
	builder.SetStubFieldSelector(source.StubFieldSelector())
	builder.SetParams(source.MethodParams)
	builder.SetResults(source.MethodResults)
	t.fileModel.AddFunctionDeclaration(builder.Build())
}

func (t *genTarget) createCallCountMethod(source *genSource) {
	methodBuilder := t.createMethodBuilder(source, source.CallCountMethodName())
	builder := NewCountMethodBuilder(methodBuilder)
	builder.SetMutexFieldSelector(source.MutexFieldSelector())
	builder.SetArgsFieldSelector(source.ArgsFieldSelector())
	t.fileModel.AddFunctionDeclaration(builder.Build())
}

func (t *genTarget) createArgsForCallMethod(source *genSource) {
	methodBuilder := t.createMethodBuilder(source, source.ArgsForCallMethodName())
	builder := NewArgsMethodBuilder(methodBuilder)
	builder.SetMutexFieldSelector(source.MutexFieldSelector())
	builder.SetArgsFieldSelector(source.ArgsFieldSelector())
	builder.SetParams(source.MethodParams)
	t.fileModel.AddFunctionDeclaration(builder.Build())
}

func (t *genTarget) createReturnsMethod(source *genSource) {
	methodBuilder := t.createMethodBuilder(source, source.ReturnsMethodName())
	builder := NewReturnsMethodBuilder(methodBuilder)
	builder.SetMutexFieldSelector(source.MutexFieldSelector())
	builder.SetReturnsFieldSelector(source.ReturnsFieldSelector())
	builder.SetResults(source.MethodResults)
	t.fileModel.AddFunctionDeclaration(builder.Build())
}

func (t *genTarget) createMethodBuilder(source *genSource, name string) *MethodBuilder {
	builder := NewMethodBuilder()
	builder.SetName(name)
	builder.SetReceiver(source.StructSelfName(), t.structName)
	return builder
}

func (t *genTarget) resolveMutexType() ast.Expr {
	alias := t.fileModel.AddImport("sync", "sync")
	return &ast.SelectorExpr{
		X:   ast.NewIdent(alias),
		Sel: ast.NewIdent("RWMutex"),
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
