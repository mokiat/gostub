package acceptance

import . "github.com/momchil-atanasov/gostub/acceptance/embedded"

//go:generate gostub EmbeddedReference

type EmbeddedReference interface {
	Embedded(Resource) Resource
	Pointer(*Resource) *Resource
}
