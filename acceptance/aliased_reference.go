package acceptance

import custom "github.com/momchil-atanasov/gostub/acceptance/aliased"

//go:generate gostub AliasedReference

type AliasedReference interface {
	Aliased(custom.User) custom.User
	Array([3]custom.User) [3]custom.User
	Slice([]custom.User) []custom.User
}
