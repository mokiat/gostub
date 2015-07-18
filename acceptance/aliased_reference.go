package acceptance

import custom "github.com/momchil-atanasov/gostub/acceptance/aliased"

//go:generate gostub AliasedReference

type AliasedReference interface {
	Aliased(custom.User) custom.User
	Slice([]custom.User) []custom.User
	Pointer(*custom.User) *custom.User
}
