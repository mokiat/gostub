package acceptance

import "github.com/momchil-atanasov/gostub/acceptance/mismatch"

//go:generate gostub MismatchedRefSupport

type MismatchedRefSupport interface {
	Method(wrong.Job) wrong.Job
}
