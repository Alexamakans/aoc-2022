package main

type file struct {
	parent entry
	name   string
	size   int
}

// compile-check conformity to entry interface
var _ entry = &file{}

func (f *file) getName() string {
	return f.name
}

func (f *file) getSize() int {
	return f.size
}

func (f *file) isDir() bool {
	return false
}

func (f *file) setParent(parent entry) {
	f.parent = parent
}

func (f *file) getParent() entry {
	return f.parent
}

func (f *file) getEntries() []entry {
	panic("getEntries() not implemented for file")
}

func (f *file) addChild(child entry) {
	panic("addChild(entry) not implemented for file")
}
