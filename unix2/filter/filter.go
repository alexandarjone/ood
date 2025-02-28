package filter

import (
	"reflect"
	"unix/file"
)

type Filter interface {
	IsMatched(file.File) bool
}

type AggregateFilter struct {
	filters map[reflect.Type]Filter
}

func (a AggregateFilter) IsMatched(f file.File) bool {
	for _, filter := range a.filters {
		if !filter.IsMatched(f) {
			return false
		}
	}
	return true
}

// TODO: handle duplicate filter
func (a *AggregateFilter) AddFilter(filter Filter) {
	a.filters[reflect.TypeOf(filter)] = filter
}

// TODO: implement remove filter function
func NewAggregateFilter(filters []Filter) Filter {
	filterMap := make(map[reflect.Type]Filter)
	for _, filter := range filters {
		filterMap[reflect.TypeOf(filter)] = filter
	}
	return &AggregateFilter{filters: filterMap}
}

type SizeFilter struct {
	maximumSize int
}

func (s SizeFilter) IsMatched(f file.File) bool {
	return !f.IsDirectory() && f.GetSize() < s.maximumSize
}

func NewSizeFilter(maximumSize int) Filter {
	return &SizeFilter{maximumSize: maximumSize}
}

type ExtensionFilter struct {
	extensionSet map[string]bool
}

func (e ExtensionFilter) IsMatched(f file.File) bool {
	return !f.IsDirectory() && e.extensionSet[f.GetExtension()]
}

func NewExtensionFiler(extensions []string) Filter {
	filter := &ExtensionFilter{extensionSet: make(map[string]bool)}
	for _, extension := range extensions {
		filter.extensionSet[extension] = true
	}
	return filter
}