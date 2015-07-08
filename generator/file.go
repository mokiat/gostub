package generator

import (
	"fmt"
	"go/ast"
	"go/token"
)

func NewFileModel() *FileModel {
	return &FileModel{
		importToAlias:        make(map[string]string),
		aliasToImport:        make(map[string]string),
		structs:              make([]*StructModel, 0),
		functionDeclarations: make([]*ast.FuncDecl, 0),
	}
}

type FileModel struct {
	filePackageName      string
	importToAlias        map[string]string
	aliasToImport        map[string]string
	aliasCounter         int
	structs              []*StructModel
	functionDeclarations []*ast.FuncDecl
}

func (m *FileModel) SetPackage(name string) {
	m.filePackageName = name
}

// AddImport assures that the specified package name in the specified
// location will be added as an import.
// This function returns the alias to be used in selector expressions.
// If the specified location is already added, then just the alias for
// that package is returned.
func (m *FileModel) AddImport(pkgName, location string) string {
	alias, locationAlreadyRegistered := m.importToAlias[location]
	if locationAlreadyRegistered {
		return alias
	}

	_, aliasAlreadyRegistered := m.aliasToImport[pkgName]
	if aliasAlreadyRegistered {
		alias = m.allocateUniqueAlias()
	} else {
		alias = pkgName
	}

	m.importToAlias[location] = alias
	m.aliasToImport[alias] = location
	return alias
}

func (m *FileModel) allocateUniqueAlias() string {
	m.aliasCounter++
	return fmt.Sprintf("alias%d", m.aliasCounter)
}

func (m *FileModel) AddStructure(declaration *StructModel) {
	m.structs = append(m.structs, declaration)
}

func (m *FileModel) AddFunctionDeclaration(declaration *ast.FuncDecl) {
	m.functionDeclarations = append(m.functionDeclarations, declaration)
}

func (m *FileModel) BuildASTFile() *ast.File {
	file := &ast.File{
		Name: ast.NewIdent(m.filePackageName),
	}

	if len(m.aliasToImport) > 0 {
		importDeclaration := &ast.GenDecl{
			Tok:    token.IMPORT,
			Lparen: token.Pos(1),
			Specs:  []ast.Spec{},
		}
		for alias, location := range m.aliasToImport {
			importDeclaration.Specs = append(importDeclaration.Specs, &ast.ImportSpec{
				Name: ast.NewIdent(alias),
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf("\"%s\"", location),
				},
			})
		}
		file.Decls = append(file.Decls, importDeclaration)
	}

	for _, structDeclaration := range m.structs {
		file.Decls = append(file.Decls, structDeclaration.BuildASTStructDeclaration())
	}

	for _, functionDeclaration := range m.functionDeclarations {
		file.Decls = append(file.Decls, functionDeclaration)
	}

	return file
}
