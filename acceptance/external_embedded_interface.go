package acceptance

import (
	other "github.com/momchil-atanasov/gostub/acceptance/external"
	"github.com/momchil-atanasov/gostub/acceptance/external/external_dup"
)

//go:generate gostub ExternalEmbeddedInterfaceSupport

type ExternalEmbeddedInterfaceSupport interface {
	external.Runner
	Method(other.Runner) other.Runner
}
