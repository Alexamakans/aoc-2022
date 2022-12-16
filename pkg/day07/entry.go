package main

type entry interface {
	getName() string
	getParent() entry
	getSize() int
	getEntries() []entry
	isDir() bool
}
