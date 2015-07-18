package acceptance

import "github.com/momchil-atanasov/gostub/acceptance/external/external_dup"

//go:generate gostub ExternalEmbeddedInterfaceSupport

type ExternalEmbeddedInterfaceSupport interface {
	external.Runner
	Method(int) int
}
