package acceptance

import alias "github.com/momchil-atanasov/gostub/acceptance/external/external_dup"

//go:generate gostub AliasedEmbeddedInterfaceSupport

type AliasedEmbeddedInterfaceSupport interface {
	alias.Runner
	Method(int) int
}
