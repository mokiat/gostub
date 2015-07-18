package acceptance

import "github.com/momchil-atanasov/gostub/acceptance/external"

//go:generate gostub ExternalReference

type ExternalReference interface {
	External(external.Address) external.Address
	Array([3]external.Address) [3]external.Address
	Slice([]external.Address) []external.Address
	Pointer(*external.Address) *external.Address
}
