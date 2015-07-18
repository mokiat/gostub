package acceptance

import . "github.com/momchil-atanasov/gostub/acceptance/external/external_dup"

//go:generate gostub EmbeddedEmbeddedInterfaceSupport

type EmbeddedEmbeddedInterfaceSupport interface {
	Runner
	Method(int) int
}
