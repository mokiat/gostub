package acceptance

import "github.com/momchil-atanasov/gostub/acceptance/external/external_dup"

//go:generate gostub PointerSupport

type PointerSupport interface {
	Method(*external.Address) *external.Address
}
