package main

func main() {
	file1 := File{name: "file.txt"}
	file2 := File{name: "table.csv"}
	file3 := File{name: "photo.jpg"}

	folder1 := Folder{name: "FolderTXT_CSV"}
	folder2 := Folder{name: "FolderPIC"}

	folder1.add(&file1)
	folder1.add(&file2)

	folder2.add(&file3)
	folder1.add(&folder2)

	folder1.search("mkv")
}
