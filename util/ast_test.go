package util_test

import (
	"go/ast"

	. "github.com/momchil-atanasov/gostub/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AST", func() {

	Describe("CreateField", func() {
		var field *ast.Field
		var fieldType ast.Expr

		BeforeEach(func() {
			fieldType = ast.NewIdent("string")
			field = CreateField("Name", fieldType)
		})

		It("has correct name", func() {
			Ω(field.Names).ShouldNot(BeNil())
			Ω(field.Names).Should(HaveLen(1))
			Ω(field.Names[0].String()).Should(Equal("Name"))
		})

		It("has correct type", func() {
			Ω(field.Type).Should(Equal(fieldType))
		})
	})

	Describe("FieldReuseCount", func() {
		var anonymousField *ast.Field
		var field *ast.Field

		BeforeEach(func() {
			anonymousField = &ast.Field{}
			field = &ast.Field{
				Names: []*ast.Ident{
					ast.NewIdent("first"),
					ast.NewIdent("second"),
				},
			}
		})

		It("returns 1 for anonymous fields", func() {
			Ω(FieldReuseCount(anonymousField)).Should(Equal(1))
		})

		It("returns the correct count for a reused field", func() {
			Ω(FieldReuseCount(field)).Should(Equal(2))
		})
	})

	Describe("CreateFuncType", func() {
		var funcType *ast.FuncType

		BeforeEach(func() {
			funcType = CreateFuncType()
		})

		It("is not nil", func() {
			Ω(funcType).ShouldNot(BeNil())
		})

		It("has zero params", func() {
			Ω(funcType.Params).ShouldNot(BeNil())
			Ω(funcType.Params.List).ShouldNot(BeNil())
			Ω(funcType.Params.List).Should(HaveLen(0))
		})

		It("has zero results", func() {
			Ω(funcType.Results).ShouldNot(BeNil())
			Ω(funcType.Results.List).ShouldNot(BeNil())
			Ω(funcType.Results.List).Should(HaveLen(0))
		})
	})

	Describe("FuncTypeParamCount", func() {
		var funcType *ast.FuncType
		var emptyFuncType *ast.FuncType

		BeforeEach(func() {
			funcType = &ast.FuncType{
				Params: &ast.FieldList{
					List: []*ast.Field{
						&ast.Field{},
						&ast.Field{},
					},
				},
			}
			emptyFuncType = &ast.FuncType{}
		})

		It("returns zero for empty func types", func() {
			Ω(FuncTypeParamCount(emptyFuncType)).Should(Equal(0))
		})

		It("return correct param count for non-empty func types", func() {
			Ω(FuncTypeParamCount(funcType)).Should(Equal(2))
		})
	})

	Describe("FuncTypeResultCount", func() {
		var funcType *ast.FuncType
		var emptyFuncType *ast.FuncType

		BeforeEach(func() {
			funcType = &ast.FuncType{
				Results: &ast.FieldList{
					List: []*ast.Field{
						&ast.Field{},
						&ast.Field{},
					},
				},
			}
			emptyFuncType = &ast.FuncType{}
		})

		It("returns zero for empty func types", func() {
			Ω(FuncTypeResultCount(emptyFuncType)).Should(Equal(0))
		})

		It("return correct result count for non-empty func types", func() {
			Ω(FuncTypeResultCount(funcType)).Should(Equal(2))
		})
	})
})
