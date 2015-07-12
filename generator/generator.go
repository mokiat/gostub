package generator

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

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

	iface, err := findTypeDeclaration(config.SourceInterfaceName, config.SourcePackageLocation)
	if err != nil {
		return err
	}

	iFaceType, isIFace := iface.Type.(*ast.InterfaceType)
	if !isIFace {
		return errors.New(fmt.Sprintf("Type '%s' in '%s' is not interface!", config.SourceInterfaceName, config.SourcePackageLocation))
	}

	err = generateIFace(iFaceType, model)
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

func generateIFace(iFaceType *ast.InterfaceType, model *GeneratorModel) error {
	for method := range util.EachMethodInInterfaceType(iFaceType) {
		funcType := method.Type.(*ast.FuncType)
		source := &MethodConfig{
			MethodName:    method.Names[0].String(),
			MethodParams:  getNormalizedParams(funcType),
			MethodResults: getNormalizedResults(funcType),
		}
		err := model.AddMethod(source)
		if err != nil {
			return err
		}
	}
	return nil
}

func getNormalizedParams(funcType *ast.FuncType) []*ast.Field {
	normalizedParams := []*ast.Field{}
	paramIndex := 1
	for param := range util.EachParamInFunc(funcType) {
		count := util.FieldReuseCount(param)
		for i := 0; i < count; i++ {
			normalizedParam := util.CreateField(fmt.Sprintf("arg%d", paramIndex), param.Type)
			normalizedParams = append(normalizedParams, normalizedParam)
			paramIndex++
		}
	}
	return normalizedParams
}

func getNormalizedResults(funcType *ast.FuncType) []*ast.Field {
	normalizedResults := []*ast.Field{}
	resultIndex := 1
	for result := range util.EachResultInFunc(funcType) {
		count := util.FieldReuseCount(result)
		for i := 0; i < count; i++ {
			normalizedResult := util.CreateField(fmt.Sprintf("result%d", resultIndex), result.Type)
			normalizedResults = append(normalizedResults, normalizedResult)
			resultIndex++
		}
	}
	return normalizedResults
}

func findTypeDeclaration(name, location string) (*ast.TypeSpec, error) {
	sourcePath, err := util.ImportToDir(location)
	if err != nil {
		return nil, err
	}

	pkgs, err := parser.ParseDir(token.NewFileSet(), sourcePath, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for spec := range util.EachInterfaceDeclarationInFile(file) {
				if spec.Name.String() == name {
					return spec, nil
				}
			}
		}
	}

	return nil, errors.New(fmt.Sprintf("Could not find interface '%s'.", name))
}
