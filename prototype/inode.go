package prototype

type Inode interface {
	print(string)
	close() Inode
}
