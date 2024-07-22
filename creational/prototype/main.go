package main

func main() {

	file1 := &File{"file1"}
	file2 := &File{"file2"}
	file3 := &File{"file3"}

	folder1 := &Folder{name: "folder1", child: []Cloneable{file1}}

	folder2 := &Folder{name: "folder2", child: []Cloneable{folder1, file2, file3}}
	folder3 := &Folder{name: "folder3", child: []Cloneable{folder2}}

	//file1.print()
	//file2.print()
	//file3.print()
	//println("print folder 1 details")
	//folder1.print()
	//println("-------------------")
	//println("print folder 2 details")
	//folder2.print()
	//println("-------------------")

	//fileCopy1 := file1.clone()
	//
	//fileCopy1.print()
	//
	//foldCopy2 := folder2.clone()
	//foldCopy2.print()
	//

	foldCopy3 := folder3.clone()
	foldCopy3.print()

}
