package resolution

import (
	"go/ast"
	"strings"

	"github.com/momchil-atanasov/gostub/util"
)

func NewSingleLocationContext(location string) *LocatorContext {
	return &LocatorContext{
		imports: []importEntry{
			importEntry{
				Alias:    ".",
				Location: location,
			},
		},
	}
}

func NewASTFileLocatorContext(astFile *ast.File, location string) *LocatorContext {
	imports := []importEntry{
		importEntry{
			Alias:    ".",
			Location: location,
		},
	}
	for decl := range util.EachGenericDeclarationInFile(astFile) {
		for spec := range util.EachSpecificationInGenericDeclaration(decl) {
			if importSpec, ok := spec.(*ast.ImportSpec); ok {
				imp := importEntry{}
				if importSpec.Name != nil {
					imp.Alias = importSpec.Name.String()
				}
				imp.Location = strings.Trim(importSpec.Path.Value, "\"")
				imports = append(imports, imp)
			}
		}
	}
	return &LocatorContext{
		imports: imports,
	}
}

type LocatorContext struct {
	imports []importEntry
}

type importEntry struct {
	Alias    string
	Location string
}

func (c *LocatorContext) CandidateLocations(alias string) []string {
	if alias == "." {
		return c.LocalLocations()
	}
	if location, found := c.AliasedLocation(alias); found {
		return []string{location}
	}
	return c.NonLocalLocations()
}

func (c *LocatorContext) LocalLocations() []string {
	result := []string{}
	for _, imp := range c.imports {
		if imp.Alias == "." {
			result = append(result, imp.Location)
		}
	}
	return result
}

func (c *LocatorContext) NonLocalLocations() []string {
	result := []string{}
	for _, imp := range c.imports {
		if imp.Alias != "." {
			result = append(result, imp.Location)
		}
	}
	return result
}

func (c *LocatorContext) AliasedLocation(alias string) (string, bool) {
	for _, imp := range c.imports {
		if imp.Alias == alias {
			return imp.Location, true
		}
	}
	return "", false
}
