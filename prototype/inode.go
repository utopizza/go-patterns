package prototype

type Inode interface {
	Print(indentation string)
	Clone() Inode
}
