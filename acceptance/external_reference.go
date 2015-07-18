package acceptance

import "github.com/momchil-atanasov/gostub/acceptance/external"

//go:generate gostub ExternalReference

type ExternalReference interface {
	External(external.Address) external.Address
	Array([]external.Address) []external.Address
}
