package acceptance

import . "github.com/momchil-atanasov/gostub/acceptance/embedded"

//go:generate gostub EmbeddedRefSupport

type EmbeddedRefSupport interface {
	Method(Resource) Resource
}
