package prototype

import (
	"fmt"
	"testing"
)

func TestInode(t *testing.T) {
	file1 := &File{Name: "File1"}
	file2 := &File{Name: "File2"}
	file3 := &File{Name: "File3"}

	folder1 := &Folder{
		Children: []Inode{file1},
		Name:     "Folder1",
	}

	folder2 := &Folder{
		Children: []Inode{folder1, file2, file3},
		Name:     "Folder2",
	}
	fmt.Println("Printing hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("Printing hierarchy for clone Folder")
	cloneFolder.Print("  ")
}
