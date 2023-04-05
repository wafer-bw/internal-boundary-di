package factory

import (
	"github.com/wafer-bw/internal-boundary-di/foo"
	"github.com/wafer-bw/internal-boundary-di/internal/baz"
)

type Factory struct {
	Details string
}

func (f Factory) NewBar() foo.Bar {
	qux := baz.New(f.Details)
	return foo.New(qux)
}
