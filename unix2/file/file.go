package file

type File interface {
	IsDirectory() bool
	GetSize() int
	ListOfSubDirectory() []File
	GetExtension() string
	GetName() string
}

type file struct {
	name string
	extension string
	size int
	isDirectory bool
	children []File
}

func (f file) GetName() string {
	return f.name
}

func (f file) IsDirectory() bool {
	return f.isDirectory
}

func (f file) GetSize() int {
	return f.size
}

func (f file) GetExtension() string {
	return f.extension
}

func (f file) ListOfSubDirectory() []File {
	return f.children
}

func NewFile(name, extension string, size int) File {
	return &file{name: name, extension: extension, size: size}
}

func NewFolder(name string, children []File) File {
	return &file{
		name: name,
		isDirectory: true,
		children: children,
	}
}