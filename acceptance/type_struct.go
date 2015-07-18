package acceptance

import "github.com/momchil-atanasov/gostub/acceptance/external/external_dup"

//go:generate gostub StructSupport

type StructSupport interface {
	Method(struct {
		Input external.Address
	}) struct {
		Output external.Address
	}
}
