package prototype

import "fmt"

type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) close() Inode {
	return &File{
		name: f.name + "_close",
	}
}
