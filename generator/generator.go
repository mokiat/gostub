package generator

import (
	"errors"
	"fmt"
	"go/ast"

	"github.com/momchil-atanasov/gostub/resolution"
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
	locator := resolution.NewLocator()

	// Do an initial search only with what we have as input
	context := resolution.NewSingleLocationContext(config.SourcePackageLocation)
	discovery, err := locator.FindIdentType(context, ast.NewIdent(config.SourceInterfaceName))
	if err != nil {
		return err
	}

	model := NewGeneratorModel(config.TargetPackageName, config.TargetStructName)
	stubGen := newGenerator(model, locator)
	err = stubGen.ProcessInterface(discovery)
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

func newGenerator(model *GeneratorModel, locator *resolution.Locator) *stubGenerator {
	return &stubGenerator{
		model:    model,
		locator:  locator,
		resolver: NewResolver(model, locator),
	}
}

type stubGenerator struct {
	model    *GeneratorModel
	locator  *resolution.Locator
	resolver *Resolver
}

func (g *stubGenerator) ProcessInterface(discovery resolution.TypeDiscovery) error {
	context := resolution.NewASTFileLocatorContext(discovery.File, discovery.Location)
	g.resolver.SetLocatorContext(context)
	iFaceType, isIFace := discovery.Spec.Type.(*ast.InterfaceType)
	if !isIFace {
		return errors.New(fmt.Sprintf("Type '%s' in '%s' is not interface!", discovery.Spec.Name.String(), discovery.Location))
	}
	for method := range util.EachMethodInInterfaceType(iFaceType) {
		funcType := method.Type.(*ast.FuncType)
		normalizedParams, err := g.getNormalizedParams(funcType)
		if err != nil {
			return err
		}
		normalizedResults, err := g.getNormalizedResults(funcType)
		if err != nil {
			return err
		}
		source := &MethodConfig{
			MethodName:    method.Names[0].String(),
			MethodParams:  normalizedParams,
			MethodResults: normalizedResults,
		}
		err = g.model.AddMethod(source)
		if err != nil {
			return err
		}
	}
	for subIFaceType := range util.EachSubInterfaceInInterfaceType(iFaceType) {
		switch t := subIFaceType.Type.(type) {
		case *ast.Ident:
			discovery, err := g.locator.FindIdentType(context, t)
			if err != nil {
				return err
			}
			err = g.ProcessInterface(discovery)
			if err != nil {
				return err
			}
		case *ast.SelectorExpr:
			discovery, err := g.locator.FindSelectorType(context, t)
			if err != nil {
				return err
			}
			err = g.ProcessInterface(discovery)
			if err != nil {
				return err
			}
		default:
			return errors.New("Unknown statement in interface declaration.")
		}
	}
	return nil
}

func (g *stubGenerator) getNormalizedParams(funcType *ast.FuncType) ([]*ast.Field, error) {
	normalizedParams := []*ast.Field{}
	paramIndex := 1
	for param := range util.EachFieldInFieldList(funcType.Params) {
		count := util.FieldTypeReuseCount(param)
		for i := 0; i < count; i++ {
			fieldName := fmt.Sprintf("arg%d", paramIndex)
			fieldType, err := g.resolver.ResolveType(param.Type)
			if err != nil {
				return nil, err
			}
			normalizedParam := util.CreateField(fieldName, fieldType)
			normalizedParams = append(normalizedParams, normalizedParam)
			paramIndex++
		}
	}
	return normalizedParams, nil
}

func (g *stubGenerator) getNormalizedResults(funcType *ast.FuncType) ([]*ast.Field, error) {
	normalizedResults := []*ast.Field{}
	resultIndex := 1
	for result := range util.EachFieldInFieldList(funcType.Results) {
		count := util.FieldTypeReuseCount(result)
		for i := 0; i < count; i++ {
			fieldName := fmt.Sprintf("result%d", resultIndex)
			fieldType, err := g.resolver.ResolveType(result.Type)
			if err != nil {
				return nil, err
			}
			normalizedResult := util.CreateField(fieldName, fieldType)
			normalizedResults = append(normalizedResults, normalizedResult)
			resultIndex++
		}
	}
	return normalizedResults, nil
}
