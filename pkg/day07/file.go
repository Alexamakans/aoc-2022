package main

type file struct {
	parent entry
	name   string
	size   int
}

// compile-check conformity to entry interface
var _ entry = file{}

func (f file) getName() string {
	return f.name
}

func (f file) getSize() int {
	return f.size
}

func (f file) isDir() bool {
	return false
}

func (f file) getParent() entry {
	return f.parent
}

func (f file) getEntries() []entry {
	return nil
}
