package generator

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"os"
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

	err := target.Save(config.TargetFilePath)
	if err != nil {
		return err
	}

	fmt.Printf("Stub '%s' successfully created in '%s'.\n", config.TargetStructName, config.TargetFilePath)
	return nil
}

func newTarget(pkgName, stubName string) *genTarget {
	file := &ast.File{
		Name: ast.NewIdent(pkgName),
	}
	stubType := &ast.StructType{
		Fields: &ast.FieldList{},
	}
	file.Decls = append(file.Decls, &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ast.NewIdent(stubName),
				Type: stubType,
			},
		},
	})
	return &genTarget{
		structName:    stubName,
		astFile:       file,
		astStructType: stubType,
	}
}

type genTarget struct {
	structName    string
	astFile       *ast.File
	astStructType *ast.StructType
}

func (t *genTarget) Save(filePath string) error {
	osFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer osFile.Close()

	err = format.Node(osFile, token.NewFileSet(), t.astFile)
	if err != nil {
		return err
	}

	return nil
}
