package generator

import (
	"go/ast"
	"strings"
)

type genSource struct {
	MethodName    string
	MethodParams  []*ast.Field
	MethodResults []*ast.Field
}

func (s *genSource) HasParams() bool {
	return len(s.MethodParams) > 0
}

func (s *genSource) HasResults() bool {
	return len(s.MethodResults) > 0
}

func (s *genSource) StructSelfName() string {
	return "stub"
}

func (s *genSource) StubMethodName() string {
	return s.MethodName + "Stub"
}

func (s *genSource) CallCountMethodName() string {
	return s.MethodName + "CallCount"
}

func (s *genSource) ArgsForCallMethodName() string {
	return s.MethodName + "ArgsForCall"
}

func (s *genSource) ReturnsMethodName() string {
	return s.MethodName + "Returns"
}

func (s *genSource) MutexName() string {
	return toPrivate(s.MethodName + "Mutex")
}

func (s *genSource) ArgsForCallName() string {
	return toPrivate(s.MethodName + "ArgsForCall")
}

func (s *genSource) ReturnsName() string {
	return toPrivate(s.MethodName + "Returns")
}

func toPrivate(name string) string {
	return strings.ToLower(name[0:1]) + name[1:]
}