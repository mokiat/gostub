package acceptance

import "github.com/momchil-atanasov/gostub/acceptance/external"

//go:generate gostub ExternalReference

type ExternalReference interface {
	External(external.Address) external.Address
	Slice([]external.Address) []external.Address
	Pointer(*external.Address) *external.Address
}
