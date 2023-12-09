package simple

import "fmt"

type File struct {
	Name string
}

func (f *File) Close() {
	fmt.Println("close file")
}
func NewFile(name string) (*File, func()) {
	file := &File{Name: name}
	return file, func() {
		file.Close()
	}
}
