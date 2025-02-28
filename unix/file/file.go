package file

// File represents a file or a directory in a Unix system
type File struct {
	name string
	extension string
	size int64
	isDir bool
	children []*File
}

func NewFile(
	name string,
	extension string,
	size int64,
) *File {
	return &File{
		name: name,
		extension: extension,
		size: size,
		isDir: false,
		children: nil,
	}
}

func NewFolder(
	name string,
	children []*File,
) *File {
	return &File{
		name: name,
		isDir: true,
		children: children,
	}
}

func (f File) IsDirectory() bool {return f.isDir}
func (f File) GetSize() int64 {return f.size}
func (f File) ListOfSubDirector() []*File {return f.children}
func (f File) GetExtension() string {return f.extension}
func (f File) GetFileName() string {return f.name + f.extension}
