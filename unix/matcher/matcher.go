package matcher

import (
	"unix/file"
)

// Match is an interface for checking if file matches the constraints
type Matcher interface {
	// IsMatched returns if the file matches the constraints
	IsMatched(file file.File) bool
}

type SizeFilter struct {
	maximumSize int64
}

func NewSizeFilter(maximumSize int64) *SizeFilter {
	return &SizeFilter{maximumSize: maximumSize}
}

// IsMatched checks if the file's size pass the filter's condition
func (f SizeFilter) IsMatched(file file.File) bool {
	return !file.IsDirectory() && file.GetSize() <= f.maximumSize
}

type ExtensionFilter struct {
	extensionSet map[string]bool
}

func NewExtensionFilter(extensions []string) *ExtensionFilter {
	extensionSet := make(map[string]bool)
	for _, extension := range extensions {
		extensionSet[extension] = true
	}
	return &ExtensionFilter{extensionSet: extensionSet}
}

// IsMatched checks if the file's  extension pass the filter's condition
func (f ExtensionFilter) IsMatched(file file.File) bool {
	return !file.IsDirectory() && f.extensionSet[file.GetExtension()]
}