package main

type directory struct {
	parent  entry
	name    string
	entries []entry
}

// compile-check conformity to entry interface
var _ entry = &directory{}

func (d *directory) getName() string {
	return d.name
}

func (d *directory) isDir() bool {
	return true
}

func (d *directory) getSize() int {
	size := 0
	for _, child := range d.entries {
		size += child.getSize()
	}
	return size
}

func (d *directory) setParent(parent entry) {
	d.parent = parent
}

func (d *directory) getParent() entry {
	return d.parent
}

func (d *directory) getEntries() []entry {
	return d.entries
}

func (d *directory) addChild(child entry) {
	child.setParent(d)
	d.entries = append(d.entries, child)
}
