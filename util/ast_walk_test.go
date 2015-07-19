package util_test

import (
	"go/ast"

	. "github.com/momchil-atanasov/gostub/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AST Walk", func() {
	Describe("FieldTypeReuseCount", func() {
		var field *ast.Field

		Context("when field is standard", func() {
			BeforeEach(func() {
				field = &ast.Field{
					Names: []*ast.Ident{
						ast.NewIdent("first"),
						ast.NewIdent("second"),
					},
				}
			})

			It("returns the correct count", func() {
				Ω(FieldTypeReuseCount(field)).Should(Equal(2))
			})
		})

		Context("when field is anonymous", func() {
			BeforeEach(func() {
				field = &ast.Field{}
			})

			It("returns 1", func() {
				Ω(FieldTypeReuseCount(field)).Should(Equal(1))
			})
		})
	})

	Describe("EachFieldInFieldList", func() {
		var fieldList *ast.FieldList
		var firstParam *ast.Field
		var secondParam *ast.Field

		BeforeEach(func() {
			firstParam = &ast.Field{
				Names: []*ast.Ident{
					ast.NewIdent("first"),
				},
			}
			secondParam = &ast.Field{
				Names: []*ast.Ident{
					ast.NewIdent("second"),
				},
			}
		})

		Context("when field list is standard", func() {
			BeforeEach(func() {
				fieldList = &ast.FieldList{
					List: []*ast.Field{
						firstParam,
						secondParam,
					},
				}
			})

			It("returns all fields", func() {
				fieldChan := EachFieldInFieldList(fieldList)
				Ω(<-fieldChan).Should(Equal(firstParam))
				Ω(<-fieldChan).Should(Equal(secondParam))
				Eventually(fieldChan).Should(BeClosed())
			})
		})

		Context("when field list is nil", func() {
			BeforeEach(func() {
				fieldList = nil
			})

			It("returns no fields", func() {
				fieldChan := EachFieldInFieldList(fieldList)
				Eventually(fieldChan).Should(BeClosed())
			})
		})

		Context("when List in field list is nil", func() {
			BeforeEach(func() {
				fieldList = &ast.FieldList{
					List: nil,
				}
			})

			It("returns no fields", func() {
				fieldChan := EachFieldInFieldList(fieldList)
				Eventually(fieldChan).Should(BeClosed())
			})
		})
	})
})
