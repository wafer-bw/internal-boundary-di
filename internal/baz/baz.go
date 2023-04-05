package baz

type Qux struct {
	details string
}

func New(details string) Qux {
	return Qux{details: details}
}

func (q Qux) Do() {

}
