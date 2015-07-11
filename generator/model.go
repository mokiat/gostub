package generator

import (
	"go/ast"
	"go/format"
	"go/token"
	"os"
)

func NewGeneratorModel(pkgName, stubName string) *GeneratorModel {
	fileBuilder := NewFileBuilder()
	fileBuilder.SetPackage(pkgName)

	structBuilder := NewStructBuilder()
	structBuilder.SetName(stubName)

	return &GeneratorModel{
		fileBuilder:   fileBuilder,
		structBuilder: structBuilder,
		structName:    stubName,
	}
}

type GeneratorModel struct {
	fileBuilder   *FileBuilder
	structBuilder *StructBuilder
	structName    string
}

func (t *GeneratorModel) AddMethod(source *genSource) error {
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

func (t *GeneratorModel) createMethodStubField(source *genSource) {
	builder := NewMethodStubFieldBuilder()
	builder.SetFieldName(source.StubMethodName())
	builder.SetParams(source.MethodParams)
	builder.SetResults(source.MethodResults)
	t.structBuilder.AddField(builder.Build())
}

func (t *GeneratorModel) createMutexField(source *genSource) {
	builder := NewMethodMutexFieldBuilder()
	builder.SetFieldName(source.MutexName())
	builder.SetMutexType(t.resolveMutexType())
	t.structBuilder.AddField(builder.Build())
}

func (t *GeneratorModel) createArgsForCallField(source *genSource) {
	builder := NewMethodArgsFieldBuilder()
	builder.SetFieldName(source.ArgsForCallName())
	builder.SetParams(source.MethodParams)
	t.structBuilder.AddField(builder.Build())
}

func (t *GeneratorModel) createReturnsField(source *genSource) {
	builder := NewReturnsFieldBuilder()
	builder.SetFieldName(source.ReturnsName())
	builder.SetResults(source.MethodResults)
	t.structBuilder.AddField(builder.Build())
}

func (t *GeneratorModel) createStubMethod(source *genSource) {
	methodBuilder := t.createMethodBuilder(source, source.MethodName)
	builder := NewStubMethodBuilder(methodBuilder)
	builder.SetMutexFieldSelector(source.MutexFieldSelector())
	builder.SetArgsFieldSelector(source.ArgsFieldSelector())
	builder.SetReturnsFieldSelector(source.ReturnsFieldSelector())
	builder.SetStubFieldSelector(source.StubFieldSelector())
	builder.SetParams(source.MethodParams)
	builder.SetResults(source.MethodResults)
	t.fileBuilder.AddFunctionDeclaration(builder.Build())
}

func (t *GeneratorModel) createCallCountMethod(source *genSource) {
	methodBuilder := t.createMethodBuilder(source, source.CallCountMethodName())
	builder := NewCountMethodBuilder(methodBuilder)
	builder.SetMutexFieldSelector(source.MutexFieldSelector())
	builder.SetArgsFieldSelector(source.ArgsFieldSelector())
	t.fileBuilder.AddFunctionDeclaration(builder.Build())
}

func (t *GeneratorModel) createArgsForCallMethod(source *genSource) {
	methodBuilder := t.createMethodBuilder(source, source.ArgsForCallMethodName())
	builder := NewArgsMethodBuilder(methodBuilder)
	builder.SetMutexFieldSelector(source.MutexFieldSelector())
	builder.SetArgsFieldSelector(source.ArgsFieldSelector())
	builder.SetParams(source.MethodParams)
	t.fileBuilder.AddFunctionDeclaration(builder.Build())
}

func (t *GeneratorModel) createReturnsMethod(source *genSource) {
	methodBuilder := t.createMethodBuilder(source, source.ReturnsMethodName())
	builder := NewReturnsMethodBuilder(methodBuilder)
	builder.SetMutexFieldSelector(source.MutexFieldSelector())
	builder.SetReturnsFieldSelector(source.ReturnsFieldSelector())
	builder.SetResults(source.MethodResults)
	t.fileBuilder.AddFunctionDeclaration(builder.Build())
}

func (t *GeneratorModel) createMethodBuilder(source *genSource, name string) *MethodBuilder {
	builder := NewMethodBuilder()
	builder.SetName(name)
	builder.SetReceiver(source.StructSelfName(), t.structName)
	return builder
}

func (t *GeneratorModel) resolveMutexType() ast.Expr {
	alias := t.fileBuilder.AddImport("sync", "sync")
	return &ast.SelectorExpr{
		X:   ast.NewIdent(alias),
		Sel: ast.NewIdent("RWMutex"),
	}
}

func (t *GeneratorModel) Save(filePath string) error {
	osFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer osFile.Close()

	t.fileBuilder.AddGeneralDeclaration(t.structBuilder.Build())
	astFile := t.fileBuilder.Build()
	err = format.Node(osFile, token.NewFileSet(), astFile)
	if err != nil {
		return err
	}

	return nil
}
