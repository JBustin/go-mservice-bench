package mocks

type fs struct {
	Calls              map[string]int
	StubRead           func(filepath string) ([]byte, error)
	StubReadLineByLine func(filepath string, walker func(line string)) error
}

func NewFs() fs {
	f := fs{}
	f.Reset()
	return f
}

func (f *fs) Reset() {
	f.Calls = map[string]int{
		"Read":           0,
		"ReadLineByLine": 0,
	}
	f.Calls = map[string]int{}
	f.StubRead = func(filepath string) ([]byte, error) { return []byte{}, nil }
	f.StubReadLineByLine = func(filepath string, walker func(line string)) error { return nil }
}

func (f *fs) SetStubRead(fn func(filepath string) ([]byte, error)) {
	f.StubRead = fn
}

func (f fs) Read(filepath string) ([]byte, error) {
	f.Calls["Read"]++
	return f.StubRead(filepath)
}

func (f *fs) SetStubReadLineByLine(fn func(filepath string, walker func(line string)) error) {
	f.StubReadLineByLine = fn
}

func (f fs) ReadLineByLine(filepath string, walker func(line string)) error {
	f.Calls["ReadLineByLine"]++
	return f.StubReadLineByLine(filepath, walker)
}
