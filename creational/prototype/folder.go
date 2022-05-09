package prototype

import "fmt"

type Folder struct {
	Name     string
	Children []Inode
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, inode := range f.Children {
		inode.Print(indentation + indentation)
	}
}

func (f *Folder) Clone() Inode {
	cloneFolder := &Folder{Name: f.Name + "_clone"}
	var cloneChildren []Inode
	for _, inode := range f.Children {
		cloneInode := inode.Clone()
		cloneChildren = append(cloneChildren, cloneInode)
	}
	cloneFolder.Children = cloneChildren
	return cloneFolder
}
