package foo

type quxer interface {
	Do()
}

type Bar struct {
	qux quxer
}

// Want this method to only be callable by packages in this module
func New(q quxer) Bar {
	return Bar{qux: q}
}

func (b Bar) Do() {
	b.qux.Do()
}
