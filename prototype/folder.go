package prototype

import "fmt"

type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *Folder) close() Inode {
	cloneFolder := &Folder{
		name: f.name + "_close",
	}
	var tempChildren []Inode
	for _, i := range f.children {
		cp := i.close()
		tempChildren = append(tempChildren, cp)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
