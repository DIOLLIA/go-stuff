package main

type File struct {
	name string
}

func (f *File) print() {
	println("file name:" + INDENTATION + f.name)
}

func (f *File) clone() Cloneable {
	return &File{name: f.name + "_clone"}
}
