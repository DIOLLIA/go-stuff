package main

type Folder struct {
	child []Cloneable
	name  string
}

func (f *Folder) print() {
	println("folder name:" + INDENTATION + f.name)
	for _, child := range f.child {
		child.print()
	}
}

func (f *Folder) clone() Cloneable {
	clonedFolder := &Folder{name: f.name + "_clone"}
	childFolders := []Cloneable{}
	for _, v := range f.child {
		cllone := v.clone()
		childFolders = append(childFolders, cllone)
	}

	clonedFolder.child = childFolders
	return clonedFolder
}
