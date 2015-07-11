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
	target := newTarget(config.TargetPackageName, config.TargetStructName)

	iface, err := findTypeDeclaration(config.SourceInterfaceName, config.SourcePackageLocation)
	if err != nil {
		return err
	}

	iFaceType, isIFace := iface.Type.(*ast.InterfaceType)
	if !isIFace {
		return errors.New(fmt.Sprintf("Type '%s' in '%s' is not interface!", config.SourceInterfaceName, config.SourcePackageLocation))
	}

	err = generateIFace(iFaceType, target)
	if err != nil {
		return err
	}

	err = target.Save(config.TargetFilePath)
	if err != nil {
		return err
	}

	fmt.Printf("Stub '%s' successfully created in '%s'.\n", config.TargetStructName, config.TargetFilePath)
	return nil
}

func generateIFace(iFaceType *ast.InterfaceType, target *genTarget) error {
	for method := range util.EachMethodInInterfaceType(iFaceType) {
		funcType := method.Type.(*ast.FuncType)
		source := &genSource{
			MethodName:    method.Names[0].String(),
			MethodParams:  getParams(funcType),
			MethodResults: getResults(funcType),
		}
		err := target.GenerateMethod(source)
		if err != nil {
			return err
		}
	}
	return nil
}

func getParams(funcType *ast.FuncType) []*ast.Field {
	params := []*ast.Field{}
	paramIndex := 1
	for param := range util.EachParamInFunc(funcType) {
		param.Names = []*ast.Ident{
			ast.NewIdent(fmt.Sprintf("arg%d", paramIndex)),
		}
		params = append(params, param)
		paramIndex++
	}
	return params
}

func getResults(funcType *ast.FuncType) []*ast.Field {
	results := []*ast.Field{}
	resultIndex := 1
	for result := range util.EachResultInFunc(funcType) {
		result.Names = []*ast.Ident{
			ast.NewIdent(fmt.Sprintf("result%d", resultIndex)),
		}
		results = append(results, result)
		resultIndex++
	}
	return results
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
