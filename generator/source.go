package generator

import (
	"go/ast"
	"strings"
)

type genSource struct {
	MethodName string
	MethodType *ast.FuncType
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

func (s *genSource) MutexName() string {
	return toPrivate(s.MethodName + "Mutex")
}

func (s *genSource) ArgsForCallName() string {
	return toPrivate(s.MethodName + "ArgsForCall")
}

func toPrivate(name string) string {
	return strings.ToLower(name[0:1]) + name[1:]
}
