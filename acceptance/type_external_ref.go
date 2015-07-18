package acceptance

import "github.com/momchil-atanasov/gostub/acceptance/external"

//go:generate gostub ExternalRefSupport

type ExternalRefSupport interface {
	Method(external.Address) external.Address
}
