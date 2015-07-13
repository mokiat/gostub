package generator

import (
	"errors"
	"fmt"
	"go/ast"

	"github.com/momchil-atanasov/gostub/util"
)

// Config is used to pass a rather large configuration to the
// Generate method.
type Config struct {

	// SourcePackageLocation specifies the location
	// (e.g. "github.com/momchil-atanasov/gostub") where the interface
	// to be stubbed is located.
	SourcePackageLocation string

	// SourceInterfaceName specifies the name of the interface to be stubbed
	SourceInterfaceName string

	// TargetFilePath specifies the file in which the stub will be saved.
	TargetFilePath string

	// TargetPackageName specifies the name of the package in which the
	// stub will be saved. Ideally, this should equal the last segment of
	// the TargetPackageLocation (e.g. "gostub_stubs")
	TargetPackageName string

	// TargetStructName specifies the name of the stub structure
	// that will implement the interface
	TargetStructName string
}

func Generate(config Config) error {
	model := NewGeneratorModel(config.TargetPackageName, config.TargetStructName)
	locator := NewLocator()
	resolver := NewResolver(model, locator)

	discovery, found, err := locator.FindTypeDeclarationInLocation(config.SourceInterfaceName, config.SourcePackageLocation)
	if err != nil {
		return err
	}
	if !found {
		return errors.New(fmt.Sprintf("Could not find interface '%s'.", config.SourceInterfaceName))
	}
	iFaceType, isIFace := discovery.Spec.Type.(*ast.InterfaceType)
	if !isIFace {
		return errors.New(fmt.Sprintf("Type '%s' in '%s' is not interface!", config.SourceInterfaceName, config.SourcePackageLocation))
	}

	resolver.SetContext(discovery.File, config.SourcePackageLocation)
	stubGen := newGenerator(model, resolver)

	err = stubGen.generateIFace(iFaceType)
	if err != nil {
		return err
	}

	err = model.Save(config.TargetFilePath)
	if err != nil {
		return err
	}

	fmt.Printf("Stub '%s' successfully created in '%s'.\n", config.TargetStructName, config.TargetFilePath)
	return nil
}

func newGenerator(model *GeneratorModel, resolver *Resolver) *stubGenerator {
	return &stubGenerator{
		model:    model,
		resolver: resolver,
	}
}

type stubGenerator struct {
	model    *GeneratorModel
	resolver *Resolver
}

func (g *stubGenerator) generateIFace(iFaceType *ast.InterfaceType) error {
	for method := range util.EachMethodInInterfaceType(iFaceType) {
		funcType := method.Type.(*ast.FuncType)
		source := &MethodConfig{
			MethodName:    method.Names[0].String(),
			MethodParams:  g.getNormalizedParams(funcType),
			MethodResults: g.getNormalizedResults(funcType),
		}
		err := g.model.AddMethod(source)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *stubGenerator) getNormalizedParams(funcType *ast.FuncType) []*ast.Field {
	normalizedParams := []*ast.Field{}
	paramIndex := 1
	for param := range util.EachParamInFunc(funcType) {
		count := util.FieldReuseCount(param)
		for i := 0; i < count; i++ {
			fieldName := fmt.Sprintf("arg%d", paramIndex)
			fieldType, _ := g.resolver.ResolveType(param.Type)
			normalizedParam := util.CreateField(fieldName, fieldType)
			normalizedParams = append(normalizedParams, normalizedParam)
			paramIndex++
		}
	}
	return normalizedParams
}

func (g *stubGenerator) getNormalizedResults(funcType *ast.FuncType) []*ast.Field {
	normalizedResults := []*ast.Field{}
	resultIndex := 1
	for result := range util.EachResultInFunc(funcType) {
		count := util.FieldReuseCount(result)
		for i := 0; i < count; i++ {
			fieldName := fmt.Sprintf("result%d", resultIndex)
			fieldType, _ := g.resolver.ResolveType(result.Type)
			normalizedResult := util.CreateField(fieldName, fieldType)
			normalizedResults = append(normalizedResults, normalizedResult)
			resultIndex++
		}
	}
	return normalizedResults
}
